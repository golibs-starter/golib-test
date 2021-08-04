package golibtest

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

type Request struct {
	headers map[string]string
	body    io.Reader
}

func NewRequestBuilder() *Request {
	return &Request{
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func (r *Request) WithHeader(key string, value string) *Request {
	r.headers[key] = value
	return r
}

func (r *Request) WithAuthorizationHeader(value string) *Request {
	r.headers["Authorization"] = value
	return r
}

func (r *Request) WithBody(body io.Reader) *Request {
	r.body = body
	return r
}

func (r *Request) WithBodyString(body string) *Request {
	r.body = bytes.NewReader([]byte(body))
	return r
}

func (r *Request) Request(t *testing.T, method string, url string) *http.Request {
	req, err := http.NewRequest(method, url, r.body)
	assert.NoError(t, err)
	for k, val := range r.headers {
		req.Header.Add(k, val)
	}
	return req
}

func (r *Request) Get(t *testing.T, url string) *http.Request {
	return r.Request(t, http.MethodGet, url)
}

func (r *Request) Post(t *testing.T, url string) *http.Request {
	return r.Request(t, http.MethodPost, url)
}

func (r *Request) Put(t *testing.T, url string) *http.Request {
	return r.Request(t, http.MethodPut, url)
}

func (r *Request) Patch(t *testing.T, url string) *http.Request {
	return r.Request(t, http.MethodPatch, url)
}

func (r *Request) Delete(t *testing.T, url string) *http.Request {
	return r.Request(t, http.MethodDelete, url)
}
