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
	t        *testing.T
	resp     *http.Response
	respBody string
}

func NewRestAssured(t *testing.T, resp *http.Response) *RestAssured {
	respBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	return &RestAssured{
		t:        t,
		resp:     resp,
		respBody: string(respBody),
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
	v := gjson.Get(r.respBody, key)
	assert.EqualValues(r.t, expected, v.Value(), fmt.Sprintf("Expected value of key %v is %v, but got: %v", key, expected, v.Value()))
	return r
}
