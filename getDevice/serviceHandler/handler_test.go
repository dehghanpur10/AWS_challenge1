package serviceHandler

import (
	"AWS_challenge1/getDevice/mock"
	"AWS_challenge1/getDevice/model"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	err := os.Setenv("TABLE_NAME", "dummy")
	assert.NoError(t, err)
	item := map[string]*dynamodb.AttributeValue{"name": &dynamodb.AttributeValue{S: aws.String("mohammad")}}
	tests := []struct {
		name           string
		getItemErr     error
		unMarshalErr   error
		item           map[string]*dynamodb.AttributeValue
		expectedOutput model.Output
		expectedErr    error
	}{
		{name: "ok", item: item, expectedOutput: model.Output{Name: "mohammad"}},
		{name: "getItemErr", getItemErr: errors.New(""), expectedErr: errors.New("server error")},
		{name: "not found", expectedErr: errors.New("device not found")},
		{name: "unMarshalErr", item: item, unMarshalErr: errors.New(""), expectedErr: errors.New("server error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dyMock := mock.NewMockDynamo(test.item, test.getItemErr)
			unmarshalMock := mock.UnMarshalMock(test.unMarshalErr)
			core := NewCore(dyMock, unMarshalType(unmarshalMock))

			output, err := core.Handler(context.TODO(), model.Input{})

			if err == nil {
				assert.Nil(t, test.expectedErr)
			} else {
				assert.EqualError(t, err, test.expectedErr.Error())
			}
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}
