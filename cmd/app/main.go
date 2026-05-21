package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/dyptan-io/mono-go/internal/platform/server"
)

func main() {
	config := readConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	router := http.NewServeMux()

	srv := server.New(&http.Server{
		Addr:    config.BindAddr,
		Handler: router,
	}, logger)

	if err := srv.Serve(context.Background()); err != nil {
		logger.Error("fatal error occurred", "error", err)
		os.Exit(1)
	}
}
