package golibtest

import (
	"go.uber.org/fx"
	"net/http"
)

var httpServer *http.Server

func EnableWebTestUtil() fx.Option {
	return fx.Invoke(func(s *http.Server) {
		httpServer = s
	})
}
