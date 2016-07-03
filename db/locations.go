package db

import (
	"log"

	"github.com/TomOnTime/resourcedir/models"
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
