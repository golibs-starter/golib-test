package golibtest

import (
	"net/http"
)

var handler httpHandler

// httpHandler represent ServeHTTP function in http package
type httpHandler func(w http.ResponseWriter, r *http.Request)

// RegisterHttpHandler register http handler for test
func RegisterHttpHandler(httpHandler httpHandler) {
	handler = httpHandler
}
