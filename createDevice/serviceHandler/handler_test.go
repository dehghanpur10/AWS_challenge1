package serviceHandler

import (
	"AWS_challenge1/createDevice/mock"
	"AWS_challenge1/createDevice/model"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
			marshalMock := func(in interface{}) (map[string]*dynamodb.AttributeValue, error) {
				return nil, test.marshalErr
			}
			core := Core{
				db:         dyMock,
				marshalMap: marshalMock,
			}

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
