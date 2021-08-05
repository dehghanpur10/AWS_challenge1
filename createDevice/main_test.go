package main

import (
	"AWS_challenge1/createDevice/data"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	err := os.Setenv("TABLE_NAME", "dummy")
	assert.NoError(t, err)
	err = os.Setenv("AWS_REGION", "dummy")
	assert.NoError(t, err)

	tests := []struct {
		name           string
		input          data.Input
		marshalErr     error
		putItemErr     error
		expectedErr    error
		expectedOutput data.Output
	}{{}}
	_ = tests

}
