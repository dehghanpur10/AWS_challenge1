package serviceHandler

import (
	"AWS_challenge1/createDevice/mock"
	"AWS_challenge1/createDevice/model"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name           string
		input          model.Input
		marshalErr     error
		putItemErr     error
		expectedErr    error
		expectedOutput model.Output
	}{
		{name: "ok", expectedOutput: model.Output{Message: "device added successfully"}},
		{name: "marshalMethodErr", marshalErr: errors.New(""), expectedErr: errors.New("server error")},
		{name: "putItemErr", putItemErr: errors.New(""), expectedErr: errors.New("server error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dyMock := mock.NewMockDynamo(test.putItemErr)
			marshalMock := mock.MarshalMock(test.marshalErr)
			core := NewCore(dyMock, marshalType(marshalMock))

			output, err := core.Handler(context.TODO(), test.input)

			if err == nil {
				assert.Nil(t, test.expectedErr)
			} else {
				assert.EqualError(t, err, test.expectedErr.Error())
			}
			assert.Equal(t, test.expectedOutput.Message, output.Message)
		})
	}

}
