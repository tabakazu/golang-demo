package datastore

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewConnection() *sqlx.DB {
	db, err := sqlx.Connect("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err)
	}
	return db
}
