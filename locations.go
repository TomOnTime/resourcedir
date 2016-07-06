package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"

	"github.com/TomOnTime/velma/models"
	"github.com/abbot/go-http-auth"
)

func handleLocation(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	locs := data.GetAllLocations()
	messages := []string{}

	r.ParseForm()
	//fmt.Println(r.Form)
	if len(r.Form) != 0 {
		fmt.Println("Received Form")
		m, err := updateLocationsFromForm(r.Form, locs)
		if err != nil {
			messages = append(messages, fmt.Sprintf("ERROR SAVING: %v", err))
		}
		messages = append(messages, m...)
		// Now that the data is updated, we must re-read it.
		locs = data.GetAllLocations()
	}
	fmt.Printf("DONE\n")

	renderLocation(w, locs, messages)
}

func updateLocationsFromForm(updates url.Values, locs []models.Location) ([]string, error) {
	messages := []string{}
	messages = append(messages, "Saved")
	var finalerror error

	for k, v := range updates {
		// Get the key.
		if !strings.HasPrefix(k, "name_") {
			messages = append(messages, fmt.Sprintf("Invalid index %#v", k))
			continue
		}
		id, err := strconv.Atoi(k[5:])
		if err != nil {
			messages = append(messages, fmt.Sprintf("Invalid index %#v", k))
			continue
		}
		if len(v) != 1 {
			messages = append(messages, fmt.Sprintf("Multiple data on Index %#v %#v", k, v))
			continue
		}
		// Get the value.
		name := strings.TrimSpace(v[0])
		//fmt.Printf("Processing %d:%#v\n", id, name)
		if name == "" { // Skip if blank.
			continue
		}
		// Get the groupid
		parts := strings.SplitN(v[0], "-", 2)
		groupId := models.Other
		if len(parts) == 2 {
			groupId = parseRegion(parts[0])
		}

		// Update:
		changed := false
		loc := findLocById(locs, id)
		if loc == nil {
			loc = &models.Location{Name: name, GroupId: groupId}
			messages = append(messages, fmt.Sprintf("Created location %d:%#v:%s", id, name, groupId))
			changed = true
		} else {
			if loc.Name != name {
				loc.Name = name
				changed = true
				messages = append(messages, fmt.Sprintf("Updated location %d:%#v:%s", id, name, groupId))
			}
			if loc.GroupId != groupId {
				loc.GroupId = groupId
				changed = true
			}
		}

		if changed {
			err = data.UpdateLocation(loc)
			if err != nil {
				messages = append(messages, fmt.Sprintf("DB error: %v", err))
				finalerror = fmt.Errorf("There were errors")
			}
			break
		}
	}

	return messages, finalerror
}

func parseRegion(r string) models.Region {
	switch r {
	case "USA":
		return models.USA
	case "CAN":
		return models.Canada
	default:
		return models.Other
	}
}

func maxLoc(locs []models.Location) int {
	maxLocId := 0
	if len(locs) > 0 {
		for _, l := range locs {
			if maxLocId < l.Id {
				maxLocId = l.Id
			}
		}
	}
	return maxLocId
}

func findLocById(locs []models.Location, target int) *models.Location {
	for _, l := range locs {
		if l.Id == target {
			return &l
		}
	}
	return nil
}

func renderLocation(w http.ResponseWriter, locs []models.Location, messages []string) {

	maxLocId := maxLoc(locs)
	data := struct {
		Locations []models.Location
		Messages  []string
		NewIds    []int
	}{
		Locations: locs,
		Messages:  messages,
		NewIds:    []int{maxLocId + 1, maxLocId + 2},
	}

	loctemplate := `<html><body>
{{range .Messages}}
<p style="background-color: yellow;">Note: {{.}}</p>
{{end}}
<form method="post">
<p><input type="submit" value="Save"></p>
<table border=1>
<tr><th>Id</th><th>Location Name</th><th>Region (updated on save)</th></tr>

{{- range .NewIds -}}
<tr><td>NEW</td><td><input type="text" name="name_{{.}}"></td><td>Add new location</td></tr>
{{end}}

{{- range .Locations -}}
<tr><td>{{.Id}}</td><td><input type="text" name="name_{{.Id}}" value="{{.Name}}"></td><td>{{.GroupId}}</td></tr>
{{end}}

</table>
</form>
<p><input type="submit" value="Save"></p>
</body></html>
`

	tmpl, err := template.New("location").Parse(loctemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}

}
