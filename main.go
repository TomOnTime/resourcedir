package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/TomOnTime/velma/db"
	auth "github.com/abbot/go-http-auth"
	"github.com/go-sql-driver/mysql"
)

var flagDbUser = flag.String("dbuser", "remotedebug", "help")
var flagDbPass = flag.String("dbpass", os.Getenv("RD_DBPASS"), "help")
var flagDbServer = flag.String("dbserver", "jj.whatexit.org", "help")
var flagDbPort = flag.String("dbport", "3306", "help")
var flagDbNet = flag.String("dbnet", "tcp", "help")
var flagDbDatabase = flag.String("dbdatabase", "polylistBetaDB", "help")
var flagRealm = flag.String("realm", "example.com", "help")
var flagPort = flag.String("port", ":8008", "help")
var flagConf = flag.String("flagfile", "", "Location of flagfile")

func Secret(user, realm string) string {
	fmt.Printf("User=%#v\n", user)
	return data.GetPasswordHash(user)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><body><h1>Hello main.</h1></body></html>")
}

var data db.Db

func readConf() {
	loc := *flagConf
	if loc == "" {
		return
	}
	b, err := ioutil.ReadFile(loc)
	if err != nil {
		log.Fatal("readConf could not open file: ", err)
	}
	for i, line := range strings.Split(string(b), "\n") {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		sp := strings.SplitN(line, "=", 2)
		if len(sp) != 2 {
			log.Fatalf("expected = in %v:%v", loc, i+1)
		}
		k := strings.TrimSpace(sp[0])
		v := strings.TrimSpace(sp[1])
		err := flag.Set(k, v)
		if err != nil {
			log.Fatalf("unknown key %#v in %v:%v", k, loc, i+1)
		}
	}
}

func main() {
	flag.Parse()
	readConf()

	var err error
	var addr string

	// Connect to the database.
	if *flagDbNet == "unix" {
		addr = ""
	} else {
		addr = net.JoinHostPort(*flagDbServer, *flagDbPort)
	}
	cst := mysql.Config{
		User:   *flagDbUser,
		Passwd: *flagDbPass,
		Net:    *flagDbNet,
		Addr:   addr,
		DBName: *flagDbDatabase,
	}
	connection := cst.FormatDSN()
	fmt.Println("CONNECT STRING:", connection)
	data, err = db.New("mysql", connection)
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}

	// Set up authentication.
	authenticator := auth.NewBasicAuthenticator(*flagRealm, Secret)

	// Routes.
	http.HandleFunc("/admin/loc", authenticator.Wrap(handleLocation))
	http.HandleFunc("/", handleMain)
	http.Handle("/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("/tmp/public"))))
	log.Fatal(http.ListenAndServe(*flagPort, nil))
}
