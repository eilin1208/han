package strutil_test

import (
	"testing"

	"github.com/hanlin1123/go-boxs/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	s, err := strutil.RandomString(3)

	assert.NoError(t, err)
	assert.True(t, len(s) > 3)
}
