package golibtest

import (
	assert "github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

type HttpClient struct {
	t       *testing.T
	client  *http.Client
	request *http.Request
}

func NewHttpClient(t *testing.T, client *http.Client) *HttpClient {
	return &HttpClient{
		t:      t,
		client: client,
	}
}

func NewDefaultHttpClient(t *testing.T, request *http.Request) *HttpClient {
	return &HttpClient{
		t:       t,
		client:  &http.Client{},
		request: request,
	}
}

func (h *HttpClient) WithRequest(request *http.Request) *HttpClient {
	h.request = request
	return h
}

func (h *HttpClient) Do() *http.Response {
	resp, err := h.client.Do(h.request)
	assert.NoError(h.t, err)
	return resp
}
