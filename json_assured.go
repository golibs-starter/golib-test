package golibtest

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"testing"
)

// JsonAssured ...
type JsonAssured struct {
	t       *testing.T
	message string
}

// NewJsonAssured ...
func NewJsonAssured(t *testing.T, message string) *JsonAssured {
	return &JsonAssured{t: t, message: message}
}

func (k *JsonAssured) Get(key string) interface{} {
	return gjson.Get(k.message, key).Value()
}

// Has ...
func (k *JsonAssured) Has(key string, expected interface{}) *JsonAssured {
	v := gjson.Get(k.message, key)
	require.EqualValues(k.t, expected, v.Value(), fmt.Sprintf("Expected value of key %v is %v, but got: %v", key, expected, v.Value()))
	return k
}
