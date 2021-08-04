package golibtest

import (
	"fmt"
	assert "github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"net/http/httptest"
	"testing"
)

type RestAssured struct {
	t *testing.T
	w *httptest.ResponseRecorder
}

func NewRestAssured(t *testing.T, w *httptest.ResponseRecorder) *RestAssured {
	return &RestAssured{
		t: t,
		w: w,
	}
}

func (r *RestAssured) Status(httpStatusCode int) *RestAssured {
	assert.Equal(r.t, httpStatusCode, r.w.Code)
	return r
}

func (r *RestAssured) Header(key string, expected interface{}) *RestAssured {
	assert.EqualValues(r.t, expected, r.w.Header().Get(key))
	return r
}

func (r *RestAssured) Body(key string, expected interface{}) *RestAssured {
	v := gjson.Get(r.w.Body.String(), key)
	assert.EqualValues(r.t, expected, v.Value(), fmt.Sprintf("Expected value of key %v is %v, but got: %v", key, expected, v.Value()))
	return r
}
