package sigtrap

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var ErrSignalReceived = errors.New("signal received")

func Wait(ctx context.Context) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-sigCh:
		return fmt.Errorf("%w: %s", ErrSignalReceived, s)
	case <-ctx.Done():
		return ctx.Err()
	}
}
