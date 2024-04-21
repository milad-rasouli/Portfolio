package db

import (
	"errors"
	"log"

	"crawshaw.io/sqlite/sqlitex"
)

var unableToEstablishDatabase = errors.New("the Database cannot be established")

func New(cfg Config) (dbpool *sqlitex.Pool, err error) {

	if cfg.IsSqlite == true {
		dbpool, err := sqlitex.Open("file:./data/locked.sqlite?cache=shared", 0, 10)
		if err != nil {
			log.Fatal(err)
		}
		return dbpool, err
	}
	return nil, unableToEstablishDatabase
}
