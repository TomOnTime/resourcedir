package db

import "github.com/TomOnTime/resourcedir/models"

type Db interface {
	GetPasswordHash(user string) string
	GetAllLocations() []models.Location
	UpdateLocation(*models.Location) error
}
