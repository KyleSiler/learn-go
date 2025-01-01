package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/tursodatabase/go-libsql"
)

func main() {
	log.Println("We in")
	// db, err := sql.Open("libsql", "libsql://192.168.0.66:9080")
	db, err := sql.Open("libsql", "libsql://oasis.local:9080")
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxIdleTime(9 * time.Second)

	defer db.Close()

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Query returned")

	defer rows.Close()

	log.Println("Key - Value")
	var key int
	var value string
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d - %s", key, value)
	}

	log.Println("Done")
}
