package golibtest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type ResponseOpt func(w http.ResponseWriter)

func WithHeader(key string, value string) ResponseOpt {
	return func(w http.ResponseWriter) {
		w.Header().Add(key, value)
	}
}

func NewHttpTestServer(httpStatus int, responseBody string, responseOpts ...ResponseOpt) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		if len(responseOpts) > 0 {
			for _, option := range responseOpts {
				option(w)
			}
		}
		_, _ = fmt.Fprint(w, responseBody)
	}))
}
