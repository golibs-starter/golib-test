package golibtest

import (
	"net/http"
	"net/http/httptest"
)

var handler httpHandler

// httpHandler represent ServeHTTP function in http package
type httpHandler func(w *httptest.ResponseRecorder, r *http.Request)

// RegisterHttpHandler register http handler for test
func RegisterHttpHandler(httpHandler httpHandler) {
	handler = httpHandler
}
