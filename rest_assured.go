package golibtest

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// RestAssured represent rest assertion
type RestAssured struct {
	t                *testing.T
	requestBuilder   *RequestBuilder
	request          *http.Request
	responseRecorder *httptest.ResponseRecorder
}

// NewRestAssured return a new instance of RestAssured
func NewRestAssured(t *testing.T) *RestAssured {
	return &RestAssured{
		t: t,
	}
}

// When create an new request builder
func (r *RestAssured) When() *RequestBuilder {
	r.requestBuilder = NewRequestBuilder(r)
	return r.requestBuilder
}

// Status assert http status code
func (r *RestAssured) Status(expected int) *RestAssured {
	require.Equal(r.t, expected, r.responseRecorder.Code)
	return r
}

// Header assert http header
func (r *RestAssured) Header(key string, expected interface{}) *RestAssured {
	require.EqualValues(r.t, expected, r.responseRecorder.Header().Get(key))
	return r
}

// Body assert json body in response
// Read https://github.com/tidwall/gjson for more information about path syntax
func (r *RestAssured) Body(key string, expected interface{}) *RestAssured {
	NewJsonAssured(r.t, r.responseRecorder.Body.String()).Has(key, expected)
	return r
}
