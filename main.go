package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/TomOnTime/resourcedir/db"
	"github.com/go-sql-driver/mysql"

	auth "github.com/abbot/go-http-auth"
)

var flagDbUser = flag.String("dbuser", "remotedebug", "help")
var flagDbPass = flag.String("dbpass", os.Getenv("RD_DBPASS"), "help")
var flagDbServer = flag.String("dbserver", "jj.whatexit.org", "help")
var flagDbDatabase = flag.String("dbdatabase", "polylistBetaDB", "help")
var flagRealm = flag.String("realm", "example.com", "help")

func Secret(user, realm string) string {
	fmt.Printf("User=%#v\n", user)
	return data.GetPasswordHash(user)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><body><h1>Hello main.</h1></body></html>")
}

var data db.Db

func main() {
	var err error

	// Connect to the database.
	cst := mysql.Config{
		User:   *flagDbUser,
		Passwd: *flagDbPass,
		Net:    "tcp",
		Addr:   net.JoinHostPort(*flagDbServer, "3306"),
		DBName: *flagDbDatabase,
	}
	connection := cst.FormatDSN()
	fmt.Println("CONNECT STRING:", connection)
	data, err = db.New("mysql", connection)
	if err != nil {
		log.Fatal("Could not connect to database: %v", err)
	}

	// Set up authentication.
	authenticator := auth.NewBasicAuthenticator(*flagRealm, Secret)

	// Routes.
	http.HandleFunc("/admin/loc", authenticator.Wrap(handleLocation))
	http.HandleFunc("/", handleMain)
	http.Handle("/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("/tmp/public"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
