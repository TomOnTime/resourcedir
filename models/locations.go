package models

//go:generate stringer -type=Region

type Location struct {
	Id    int    `db:"location_id"`
	Name  string `db:"location_name"`
	Group Region `db:"state_province_country"`
}

type Region int

const (
	USA    Region = 1
	Canada        = 2
	Other         = 3
)
