package cmd

import (
	"fmt"
	"hot-coffee-sql/internal/router"
	"hot-coffee-sql/pkg/config"
	"log/slog"
	"net/http"
)

func Init() {
	// Handel config
	cfg := config.LoadConfig()

	// Start server
	mux := router.NewRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HttpPort),
		Handler: mux,
	}

	slog.Info(
		"starting http server",
		slog.String("Env", cfg.Environment),
		slog.String("addr", fmt.Sprintf("http://127.0.0.1:%d", cfg.HttpPort)),
		slog.String("dir", cfg.DatabaseURI),
	)
	slog.Debug(fmt.Sprint(srv.ListenAndServe()))
}
