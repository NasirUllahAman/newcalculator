package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "scott:tiger@tcp(127.0.0.1)/employees"
	log.Println("using", dsn)

	db, _ := sql.Open("mysql", dsn)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("successfully connected using native password")
	}

	dsn = dsn + "?allowNativePasswords=false"
	log.Println("using", dsn)

	db, _ = sql.Open("mysql", dsn)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
