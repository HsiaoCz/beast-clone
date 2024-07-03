package db

import (
	"os"

	"github.com/anthdm/superkit/db"
	"github.com/anthdm/superkit/kit"

	_ "github.com/mattn/go-sqlite3"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

// TODO this name need change
// but now use this
var Query *bun.DB

func InitDB() error {
	config := db.Config{
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		Host:     os.Getenv("DB_HOST"),
	}
	db, err := db.NewSQL(config)
	if err != nil {
		return err
	}
	Query = bun.NewDB(db, sqlitedialect.New())
	if kit.IsDevelopment() {
		Query.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	return nil
}
