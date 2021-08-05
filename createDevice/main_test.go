package main

import (
	"AWS_challenge1/createDevice/data"
	"errors"
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
	}{
		{"ok",data.Input{},nil,nil,nil,data.Output{"device added successfully"}},
		{"marshalMethodErr",data.Input{},errors.New(""),nil,errors.New("server error"),data.Output{}},
		{"putItemErr",data.Input{},nil,errors.New(""),errors.New("server error"),data.Output{}},
	}
	_ = tests

}
