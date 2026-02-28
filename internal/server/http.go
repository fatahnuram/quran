package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

func ServeHttp(port int) error {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	slog.Info(fmt.Sprintf("starting http server on port %d ..", port))
	return http.ListenAndServe(addr, nil)
}
