package db

import (
	"log"
	"os"

	"github.com/anthdm/superkit/db"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// use signel pattern
var dbInstance *gorm.DB

// get the db
func Get() *gorm.DB {
	return dbInstance
}

func InitDB() error {
	config := db.Config{
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		Host:     os.Getenv("DB_HOST"),
	}
	dbinst, err := db.NewSQL(config)
	if err != nil {
		return err
	}
	// Based on the driver create the corresponding DB instance.
	// By default, the SuperKit boilerplate comes with a pre-configured
	// ORM called Gorm. https://gorm.io.
	//
	// You can change this to any other DB interaction tool
	// of your liking. EG:
	// - uptrace bun -> https://bun.uptrace.dev
	// - SQLC -> https://github.com/sqlc-dev/sqlc
	// - gojet -> https://github.com/go-jet/jet
	switch config.Driver {
	case db.DriverSqlite3:
		dbInstance, err = gorm.Open(sqlite.New(sqlite.Config{
			Conn: dbinst,
		}))
	case db.DriverMysql:
		// ...
	default:
		log.Fatal("invalid driver:", config.Driver)
	}
	return err
}
