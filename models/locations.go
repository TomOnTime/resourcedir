package models

//go:generate stringer -type=Region

type Location struct {
	Id      int    `db:"location_id"`
	Name    string `db:"location_name"`
	GroupId Region `db:"state_province_country"`
}

type Region int

const (
	USA    Region = 1
	Canada Region = 2
	Other  Region = 3
)
