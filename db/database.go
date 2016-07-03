package db

import "github.com/jmoiron/sqlx"

type dataAccess struct {
	db *sqlx.DB
}

func New(driver, connection string) (Db, error) {
	var err error
	var d *sqlx.DB

	d, err = sqlx.Connect(driver, connection)
	if err != nil {
		return nil, err
	}

	err = d.Ping()
	return &dataAccess{
		db: d,
	}, err
}
