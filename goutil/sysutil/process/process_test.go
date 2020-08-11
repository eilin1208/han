package process_test

import (
	"os"
	"testing"

	"github.com/hanlin1123/go-boxs/goutil/sysutil/process"
	"github.com/stretchr/testify/assert"
)

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, process.Exists(pid))
}
