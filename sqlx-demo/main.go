package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	type User struct {
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Email     string
	}
	var schema = `
CREATE TABLE IF NOT EXISTS users (
    first_name varchar(255),
    last_name varchar(255),
    email varchar(255)
);
`
	db.MustExec(schema)

	tx := db.MustBegin()
	p1 := &User{"Taro", "Tanaka", "t-tanaka@test.com"}
	tx.NamedExec("INSERT INTO users (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", p1)
	tx.Commit()

	users := []User{}
	db.Select(&users, "SELECT * FROM users ORDER BY email ASC")
	fmt.Println(users)

	db.MustExec(`DROP TABLE users;`)
}
