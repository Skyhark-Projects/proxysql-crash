package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func iterate(db *sql.DB) {
	for i := 0; i < 1000; i++ {
		_, err := db.Exec("INSERT INTO test VALUES (NULL, 'TEST')")
		if err != nil {
			panic(err)
		}

		fmt.Println("Passed testes", i)
	}
}

func main() {
	// Connect to mysql
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:6033)/mcm")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)

	// Run async inserts
	go iterate(db)
	go iterate(db)
	go iterate(db)
	go iterate(db)
	iterate(db)
}
