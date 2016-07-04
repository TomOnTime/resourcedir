package db

import (
	"fmt"
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

// UpdateLocation updates a location or creates a new one if Id == 0.
func (d *dataAccess) UpdateLocation(loc *models.Location) error {
	fmt.Printf("WOULD update loc: %#v\n", loc)
	return nil
}
