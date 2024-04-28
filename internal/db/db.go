package db

import (
	"errors"
	"log"

	"crawshaw.io/sqlite/sqlitex"
)

var unableToEstablishDatabase = errors.New("the Database cannot be established")

func New(cfg Config) (*sqlitex.Pool, error) {
	var (
		dbPool *sqlitex.Pool
		err    error
	)
	if cfg.IsSqlite == true {
		dbPool, err = sqlitex.Open("file:./data/locked.sqlite?cache=shared", 0, 10)
		if err != nil {
			log.Fatal(err)
		}
		return dbPool, err
	}
	return dbPool, unableToEstablishDatabase
}
