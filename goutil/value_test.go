package goutil_test

import (
	"testing"

	"github.com/hanlin1123/go-boxs/goutil"
	"github.com/stretchr/testify/assert"
)

func TestValue_Val(t *testing.T) {
	v := goutil.Value{V: 23}

	assert.Equal(t, 23, v.Val())
	assert.Equal(t, 23, v.Int())
	assert.Equal(t, int64(23), v.Int64())
	assert.Equal(t, float64(23), v.Float64())
	assert.Equal(t, "23", v.String())
	assert.False(t, v.IsEmpty())
	assert.False(t, v.Bool())

	v.V = []string{"a", "b"}
	assert.Equal(t, []string{"a", "b"}, v.Val())
	assert.Equal(t, []string{"a", "b"}, v.Strings())

	v.Reset()
	assert.Nil(t, v.V)
	assert.True(t, v.IsEmpty())
}
