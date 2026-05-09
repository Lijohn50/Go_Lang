package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:admin21@tcp(127.0.0.1:3306)/goStart"

	db, err := sql.Open("mysql", dsn)
	if err != nil {

		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("Database connection successful!")
}
