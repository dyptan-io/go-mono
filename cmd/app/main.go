package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/dyptan-io/go-mono/internal/service/app"
)

func main() {
	cfg := ReadConfig()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := app.Run(context.Background(), cfg.BindAddr, log); err != nil {
		log.Error("fatal error occurred", "error", err)
		os.Exit(1)
	}
}
