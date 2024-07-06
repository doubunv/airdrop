package concurrency

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"runtime"
)

func Go(ctx context.Context, f func(context.Context)) {
	go func() {
		defer func() {
			if rerr := recover(); rerr != nil {
				buf := make([]byte, 64<<10) //nolint:gomnd
				n := runtime.Stack(buf, false)
				buf = buf[:n]
				logx.Errorf(" unhandle error: %+v:\n%s\n", rerr, buf)
			}
		}()
		f(ctx)
	}()
}
