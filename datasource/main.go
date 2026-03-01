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

type SurahT struct {
	No        string `json:"no"`
	NameAR    string `json:"name_ar"`
	NameID    string `json:"name_id"`
	AyatCount string `json:"ayat_count"`
}

type AyatT struct {
	SurahNo string `json:"surah_no"`
	AyatNo  string `json:"ayat_no"`
	AR      string `json:"ar"`
	ID      string `json:"id"`
}

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

	// show surah
	surahrows, err := db.Query(`SELECT surah, ayat, terjemahan, jumlah_ayat FROM table_surah`)
	if err != nil {
		log.Fatalf("failed to query surah: %v", err)
	}
	defer surahrows.Close()

	var surah []SurahT
	for surahrows.Next() {
		var s SurahT
		err = surahrows.Scan(&s.No, &s.NameAR, &s.NameID, &s.AyatCount)
		if err != nil {
			log.Fatalf("failed to parse surah: %v", err)
		}
		surah = append(surah, s)
	}
	fmt.Printf("surah: %#v\n", surah)

	// show surah
	ayatrows, err := db.Query(`SELECT surah, ayat, arab, terjemahan FROM table_ayat`)
	if err != nil {
		log.Fatalf("failed to query ayat: %v", err)
	}
	defer ayatrows.Close()

	var ayat []AyatT
	for ayatrows.Next() {
		var s AyatT
		err = ayatrows.Scan(&s.SurahNo, &s.AyatNo, &s.AR, &s.ID)
		if err != nil {
			log.Fatalf("failed to parse ayat: %v", err)
		}
		ayat = append(ayat, s)
	}
	fmt.Printf("ayat: %#v\n", ayat)
}
