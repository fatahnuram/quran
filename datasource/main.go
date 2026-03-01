package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

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

const (
	DB_PATH           = "datasource/quran.sqlite"
	OUTPUT_SURAH_PATH = "surah.csv"
	OUTPUT_AYAT_PATH  = "ayat.csv"
)

// This script export current sqlite database
// into CSV format
func main() {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", DB_PATH))
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

	var surahs []SurahT
	for surahrows.Next() {
		var s SurahT
		err = surahrows.Scan(&s.No, &s.NameAR, &s.NameID, &s.AyatCount)
		if err != nil {
			log.Fatalf("failed to parse surah: %v", err)
		}
		surahs = append(surahs, s)
	}

	// show ayat
	ayatrows, err := db.Query(`SELECT surah, ayat, arab, terjemahan FROM table_ayat`)
	if err != nil {
		log.Fatalf("failed to query ayat: %v", err)
	}
	defer ayatrows.Close()

	var ayats []AyatT
	for ayatrows.Next() {
		var s AyatT
		err = ayatrows.Scan(&s.SurahNo, &s.AyatNo, &s.AR, &s.ID)
		if err != nil {
			log.Fatalf("failed to parse ayat: %v", err)
		}
		ayats = append(ayats, s)
	}

	// export surah to csv
	var surahrecords [][]string
	for _, surah := range surahs {
		rec := make([]string, 4)
		rec[0] = surah.No
		rec[1] = surah.NameAR
		rec[2] = surah.NameID
		rec[3] = surah.AyatCount

		surahrecords = append(surahrecords, rec)
	}

	surahfile, err := os.Create(OUTPUT_SURAH_PATH)
	if err != nil {
		log.Fatalf("failed to create csv file %s: %v", OUTPUT_SURAH_PATH, err)
	}
	defer surahfile.Close()

	ws := csv.NewWriter(surahfile)
	err = ws.WriteAll(surahrecords)
	if err != nil {
		log.Fatalf("failed to write csv file %s: %v", OUTPUT_SURAH_PATH, err)
	}
}
