package async

import (
	"context"
	"log/slog"
	"time"
)

type asyncJobFn func(context.Context) error

// Schedule runs provided function at a configured intervals.
func Schedule(ctx context.Context, interval time.Duration, fn asyncJobFn, log *slog.Logger) {
	go func(ctx context.Context) {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := fn(ctx); err != nil {
					log.Error(err.Error())
				}
			}
		}
	}(ctx)
}
