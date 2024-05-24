package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewMysqlClient() *sql.DB {
	// Connection string
	dsn := "root@tcp(127.0.0.1:3306)/akshara"

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to create connection", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database!")

	return db
}
