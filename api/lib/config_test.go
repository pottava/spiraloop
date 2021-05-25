package lib

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	os.Setenv("ACCESS_LOG", "false")

	Set()

	assert.Equal(t, false, Config.AccessLog, "Expected an error when AccessLog is not false")
}
