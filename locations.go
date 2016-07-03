package main

import (
	"fmt"
	"net/http"

	"github.com/abbot/go-http-auth"
)

func handleLocation(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)

	all := data.GetAllLocations()
	fmt.Fprintf(w, "<table border=1>\n")
	for _, row := range all {
		fmt.Fprintf(w, "<tr> <td>%d</td><td>%s</td><td>%s</td> </tr>\n",
			row.Id, row.Name, row.Group)
	}
	fmt.Fprintf(w, "</table>\n")
}
