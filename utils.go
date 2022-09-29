package golibtest

import (
	"context"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func WaitUntil(cb func() bool, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for {
		if ctx.Err() != nil || cb() {
			break
		}
	}
}

func WaitUntilT(t testing.TB, cb func() bool, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for {
		if ctx.Err() != nil {
			assert.FailNowf(t, "Condition not matched",
				"Condition not matched with timeout [%s], err [%v]", timeout.String(), ctx.Err())
			break
		}
		if cb() {
			break
		}
	}
}
