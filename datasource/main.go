package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// This file should only be run using `go run` command
// i.e. one-off command/script
//
// To execute the script, run: `go run datasource/main.go`

func main() {
	db, err := sql.Open("sqlite3", "file:datasource/quran.sqlite")
	if err != nil {
		log.Fatalf("failed to open sqlite db: %v", err)
	}

	// show db version
	var version string
	db.QueryRow(`SELECT sqlite_version()`).Scan(&version)
	fmt.Printf("running sqlite version %s\n", version)

	// show tables
	tablerows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table' ORDER BY name`)
	if err != nil {
		log.Fatalf("failed to query tables: %v", err)
	}
	defer tablerows.Close()

	var tables []string
	for tablerows.Next() {
		var t string
		err = tablerows.Scan(&t)
		if err != nil {
			log.Fatalf("failed to parse table: %v", err)
		}
		tables = append(tables, t)
	}
	fmt.Printf("tables: %#v\n", tables)
}
