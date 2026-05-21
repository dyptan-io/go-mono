package server

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Listener is a network listener (e.g. *http.Server).
type Listener interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

// Server wraps a Listener with graceful shutdown on OS signals.
type Server struct {
	srv Listener
	log *slog.Logger
}

// New returns a Server with a 5-second shutdown timeout.
func New(srv Listener, log *slog.Logger) Server {
	return Server{
		srv: srv,
		log: log,
	}
}

// Serve starts the listener and blocks until ctx is cancelled or an OS signal
// (SIGINT/SIGTERM) is received, then shuts down gracefully.
func (s Server) Serve(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		s.log.Info("server started")

		errCh <- s.srv.ListenAndServe()
	}()

	sigCtx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	select {
	case err := <-errCh:
		return err
	case <-sigCtx.Done():
		s.log.Info("shutting down server")

		// parent ctx is already cancelled here; a fresh context is required.
		shutCtx, shutCancel := context.WithTimeout(context.Background(), time.Second)
		defer shutCancel()

		return s.srv.Shutdown(shutCtx) //nolint:contextcheck
	}
}
