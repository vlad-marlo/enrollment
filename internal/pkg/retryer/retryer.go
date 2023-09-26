package retryer

import (
	"context"
	"go.uber.org/zap"
	"time"
)

// TryWithAttempts tries to get non-error result of calling function f with delay.
func TryWithAttempts(f func() error, attempts uint, delay time.Duration) (err error) {
	err = f()
	if err == nil {
		return nil
	}

	for i := uint(1); i < attempts; i++ {
		if err = f(); err == nil {
			return nil
		}
		zap.L().Warn("got error in attempter", zap.Uint("attempts", i+1), zap.NamedError("error", err))
		time.Sleep(delay)
	}
	return err
}

// TryWithAttemptsCtx is helper function that calls TryWithAttempts with function f transformed to closure that does not
// require ctx as necessary argument.
func TryWithAttemptsCtx(ctx context.Context, f func(context.Context) error, attempts uint, delay time.Duration) (err error) {
	return TryWithAttempts(func() error { return f(ctx) }, attempts, delay)
}
