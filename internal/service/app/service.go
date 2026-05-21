// Package app contains the wiring and request handlers for the app service.
package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/dyptan-io/go-mono/internal/pkg/server"
)

// Version is the semantic version of the service.
const Version = "0.0.0"

// Run starts the service and blocks until ctx is cancelled or an OS signal
// triggers graceful shutdown.
func Run(ctx context.Context, cfg Config, log *slog.Logger) error {
	router := http.NewServeMux()

	srv := server.New(&http.Server{
		Addr:    cfg.BindAddr,
		Handler: router,
	}, log)

	return srv.Serve(ctx)
}
