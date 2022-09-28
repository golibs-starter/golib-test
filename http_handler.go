package golibtest

import (
	"context"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"net/http"
)

var handler httpHandler

// httpHandler represent ServeHTTP function in http package
type httpHandler func(w http.ResponseWriter, r *http.Request)

// RegisterHttpHandler register http handler for test
func RegisterHttpHandler(httpHandler httpHandler) {
	handler = httpHandler
}

func RegisterHttpHandlerOnStartOpt() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, httpServer *http.Server) {
		lc.Append(
			fx.Hook{
				OnStart: func(ctx context.Context) error {
					RegisterHttpHandler(httpServer.Handler.ServeHTTP)
					log.Info("Register http handler on application start successful")
					return nil
				},
			},
		)
	})
}
