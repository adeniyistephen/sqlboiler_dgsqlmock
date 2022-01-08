package main

//go:generate sqlboiler --wipe mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"mysql/example/user"
)

func main() {

	// TODO: read from config
	HOSTNAME := "localhost"
	PORT := "3306"
	USERNAME := "root"
	PASSWORD := "password_here"
	DATABASE := "user"

	connection := USERNAME + ":" + PASSWORD + "@tcp(" + HOSTNAME + ":" + PORT + ")/" + DATABASE + "?parseTime=true"
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	user.InsertUser(db)
}