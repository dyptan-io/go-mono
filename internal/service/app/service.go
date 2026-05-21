// Package app contains the wiring and request handlers for the app service.
package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/dyptan-io/go-mono/internal/pkg/server"
)

// Run starts the service and blocks until ctx is cancelled or an OS signal
// triggers graceful shutdown.
func Run(ctx context.Context, addr string, log *slog.Logger) error {
	router := http.NewServeMux()

	srv := server.New(&http.Server{
		Addr:    addr,
		Handler: router,
	}, log)

	return srv.Serve(ctx)
}
