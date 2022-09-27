package golibtest

import (
	"context"
	"go.uber.org/fx"
	"net/http"
	"net/http/httptest"
)

const (
	RequestContextKey        = "r"
	ResponseWriterContextKey = "w"
)

var handler httpHandler

// httpHandler represent ServeHTTP function in http package
type httpHandler func(w http.ResponseWriter, r *http.Request)

// RegisterHttpHandler register http handler for test
func RegisterHttpHandler(httpHandler httpHandler) {
	handler = httpHandler
}

func FxHttpHandler(app *fx.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctxWithResponseWriter := context.WithValue(ctx, ResponseWriterContextKey, w)
		ctxWithRequest := context.WithValue(ctxWithResponseWriter, RequestContextKey, r)
		_ = app.Start(ctxWithRequest)
	}
}

func HttpTestServerStartupOpt() fx.Option {
	return fx.Invoke(AppendHttpTestServerStartup)
}

func AppendHttpTestServerStartup(lc fx.Lifecycle, httpServer *http.Server) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				r := ctx.Value(RequestContextKey).(*http.Request)
				w := ctx.Value(ResponseWriterContextKey).(*httptest.ResponseRecorder)
				httpServer.Handler.ServeHTTP(w, r)
				return nil
			},
		},
	)
}
