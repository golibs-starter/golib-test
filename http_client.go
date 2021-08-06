package golibtest

import (
	assert "github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

type HttpClient struct {
	t      *testing.T
	client *http.Client
}

func NewHttpClient(t *testing.T, client *http.Client) *HttpClient {
	return &HttpClient{
		t:      t,
		client: client,
	}
}

func NewDefaultHttpClient(t *testing.T) *HttpClient {
	return &HttpClient{
		t:      t,
		client: &http.Client{},
	}
}

func (h *HttpClient) Do(request *http.Request) *http.Response {
	resp, err := h.client.Do(request)
	assert.NoError(h.t, err)
	return resp
}
