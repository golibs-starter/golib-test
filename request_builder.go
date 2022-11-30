package golibtest

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// RequestBuilder represent http request builder for test
type RequestBuilder struct {
	method string
	url    string
	header map[string]string
	body   io.Reader
	ra     *RestAssured
}

// NewRequestBuilder return a new instance of RequestBuilder
func NewRequestBuilder(restAssured *RestAssured) *RequestBuilder {
	return &RequestBuilder{
		header: make(map[string]string),
		ra:     restAssured,
	}
}

// Method set http method
func (r *RequestBuilder) Method(method string) *RequestBuilder {
	r.method = method
	return r
}

// URL set url method
func (r *RequestBuilder) URL(url string) *RequestBuilder {
	r.url = url
	return r
}

func (r *RequestBuilder) Header(key string, value string) *RequestBuilder {
	r.header[key] = value
	return r
}

// Body set request body
func (r *RequestBuilder) Body(body string) *RequestBuilder {
	r.body = bytes.NewReader([]byte(body))
	return r
}

// Get set get method and url to request builder
func (r *RequestBuilder) Get(url string) *RequestBuilder {
	r.Method(http.MethodGet).URL(url)
	return r
}

// Post set post method and url to request builder
func (r *RequestBuilder) Post(url string) *RequestBuilder {
	r.Method(http.MethodPost).URL(url)
	return r
}

// Put set post method and url to request builder
func (r *RequestBuilder) Put(url string) *RequestBuilder {
	r.Method(http.MethodPut).URL(url)
	return r
}

// Patch set patch method and url to request builder
func (r *RequestBuilder) Patch(url string) *RequestBuilder {
	r.Method(http.MethodPatch).URL(url)
	return r
}

// Delete set delete method and url to request builder
func (r *RequestBuilder) Delete(url string) *RequestBuilder {
	r.Method(http.MethodDelete).URL(url)
	return r
}

// BearerToken set bearer token to request builder
func (r *RequestBuilder) BearerToken(token string) *RequestBuilder {
	r.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	return r
}

// BasicAuth set basic auth credentials to request builder
func (r *RequestBuilder) BasicAuth(username string, password string) *RequestBuilder {
	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	r.Header("Authorization", fmt.Sprintf("Basic %s", token))
	return r
}

// Build create request from request builder
func (r *RequestBuilder) Build() *http.Request {
	request, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		panic(fmt.Sprintf("could not create request: %v", err))
	}
	for key, value := range r.header {
		request.Header.Add(key, value)
	}
	return request
}

// Then perform request
func (r *RequestBuilder) Then() *RestAssured {
	r.Header("Content-Type", "application/json")
	r.ra.request = r.Build()
	r.ra.responseRecorder = httptest.NewRecorder()
	httpServer.Handler.ServeHTTP(r.ra.responseRecorder, r.ra.request)
	return r.ra
}
