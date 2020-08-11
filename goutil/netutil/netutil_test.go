package netutil_test

import (
	"testing"

	"github.com/hanlin1123/go-boxs/goutil/netutil"
	"github.com/stretchr/testify/assert"
)

func TestInternalIP(t *testing.T) {
	assert.NotEmpty(t, netutil.InternalIP())
}
