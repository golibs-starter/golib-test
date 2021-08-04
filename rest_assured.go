package golibtest

import (
	"fmt"
	assert "github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"testing"
)

type RestAssured struct {
	t    *testing.T
	resp *http.Response
}

func NewRestAssured(t *testing.T, resp *http.Response) *RestAssured {
	return &RestAssured{
		t:    t,
		resp: resp,
	}
}

func (r *RestAssured) Status(httpStatusCode int) *RestAssured {
	assert.Equal(r.t, httpStatusCode, r.resp.StatusCode)
	return r
}

func (r *RestAssured) Header(key string, expected interface{}) *RestAssured {
	assert.EqualValues(r.t, expected, r.resp.Header.Get(key))
	return r
}

func (r *RestAssured) Body(key string, expected interface{}) *RestAssured {
	byteBody, err := ioutil.ReadAll(r.resp.Body)
	assert.NoError(r.t, err)
	v := gjson.Get(string(byteBody), key)
	assert.EqualValues(r.t, expected, v.Value(), fmt.Sprintf("Expected value of key %v is %v, but got: %v", key, expected, v.Value()))
	return r
}
