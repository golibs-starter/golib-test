package golibtest

import (
	"net/http"
	"testing"
)

func httpHandlerTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status": "OK"}`))
}

func TestRestAssuredWithHttpHandler(t *testing.T) {
	RegisterHttpHandler(httpHandlerTest)
	NewRestAssured(t).
		When().
		Get("/").
		Then().
		Status(http.StatusOK).
		Body("status", "OK")
}
