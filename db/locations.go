package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TomOnTime/velma/models"
)

func (d *dataAccess) GetAllLocations() []models.Location {

	ll := []models.Location{}
	err := d.db.Select(
		&ll,
		`SELECT location_id, location_name, state_province_country FROM locations ORDER BY state_province_country, location_name`,
	)
	if err != nil {
		log.Fatalf("GetLocationTable: %v", err)
	}

	return ll
}

// UpdateLocation updates a location or creates a new one if Id == 0.
func (d *dataAccess) UpdateLocation(loc *models.Location) error {
	var err error
	var result sql.Result

	if loc.Id == 0 {
		fmt.Printf("INSERT loc: %#v\n", loc)
		query := `
	      INSERT INTO locations
	      (location_name, state_province_country) VALUES (?, ?)
	  `
		result, err = d.db.Exec(query, loc.Name, loc.GroupId)
		fmt.Printf("RESULT=%#v\n", result)
	} else {
		fmt.Printf("UPDATE loc: %#v\n", loc)
		query := `
		    UPDATE locations
				SET
				    location_name = ?,
						state_province_country = ?
			  WHERE location_id = ?
		`
		result, err = d.db.Exec(query, loc.Name, loc.GroupId, loc.Id)
		fmt.Printf("RESULT=%#v\n", result)
	}

	return err
}
