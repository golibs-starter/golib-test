package golibtest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func httpHandlerTest(w *httptest.ResponseRecorder, r *http.Request) {
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
