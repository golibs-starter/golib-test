package golibtest

import (
	"bytes"
	assert "github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type Request struct {
	t       *testing.T
	headers map[string]string
	body    io.Reader
}

func NewRequestBuilder(t *testing.T) *Request {
	return &Request{
		t: t,
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func (r *Request) WithHeader(key string, value string) *Request {
	r.headers[key] = value
	return r
}

func (r *Request) WithContentType(value string) *Request {
	return r.WithHeader("Content-Type", value)
}

func (r *Request) WithAuthorization(value string) *Request {
	r.headers["Authorization"] = value
	return r
}

func (r *Request) WithBearerAuthorization(value string) *Request {
	return r.WithAuthorization("Bearer " + value)
}

func (r *Request) WithBasicAuthorization(value string) *Request {
	return r.WithAuthorization("Basic " + value)
}

func (r *Request) WithBody(body io.Reader) *Request {
	r.body = body
	return r
}

func (r *Request) WithBodyString(body string) *Request {
	r.body = bytes.NewReader([]byte(body))
	return r
}

func (r *Request) Request(method string, url string) *http.Request {
	req, err := http.NewRequest(method, url, r.body)
	assert.NoError(r.t, err)
	for k, val := range r.headers {
		req.Header.Add(k, val)
	}
	return req
}

func (r *Request) Get(url string) *http.Request {
	return r.Request(http.MethodGet, url)
}

func (r *Request) Post(url string) *http.Request {
	return r.Request(http.MethodPost, url)
}

func (r *Request) Put(url string) *http.Request {
	return r.Request(http.MethodPut, url)
}

func (r *Request) Patch(url string) *http.Request {
	return r.Request(http.MethodPatch, url)
}

func (r *Request) Delete(url string) *http.Request {
	return r.Request(http.MethodDelete, url)
}
