package golibtest

import (
	"context"
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
