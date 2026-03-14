package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fatahnuram/quran/internal/data"
)

func ListSurat() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Quran)
	})
}

func GetSurat() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		ayats := data.Quran[0].Ayats
		json.NewEncoder(w).Encode(ayats)
	})
}

func GetAyat() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		ayats := data.Quran[0].Ayats
		json.NewEncoder(w).Encode(ayats[1])
	})
}

func initRoutes() *http.ServeMux {
	m := http.NewServeMux()
	m.Handle("GET /quran", ListSurat())
	m.Handle("GET /quran/{sid}", GetSurat())
	m.Handle("GET /quran/{sid}/{aid}", GetAyat())
	return m
}

func ServeHttp(port int) error {
	mux := initRoutes()
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	slog.Info(fmt.Sprintf("starting http server on port %d ..", port))
	return http.ListenAndServe(addr, mux)
}
