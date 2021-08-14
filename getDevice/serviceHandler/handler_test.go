package serviceHandler

import (
	"AWS_challenge1/getDevice/mock"
	"AWS_challenge1/getDevice/model"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
		entity         model.Input
		getItemErr     error
		resultGetItem  *dynamodb.GetItemOutput
		expectedOutput model.Output
		expectedErr    error
	}{
		{name: "ok", entity: model.Input{Id: "1"}, resultGetItem: &dynamodb.GetItemOutput{Item: map[string]*dynamodb.AttributeValue{"name": &dynamodb.AttributeValue{S: aws.String("mohammad")}}}, expectedOutput: model.Output{Name: "mohammad"}},
		{name: "getItemErr", entity: model.Input{Id: "1"}, getItemErr: errors.New(""), expectedErr: errors.New("server error"), expectedOutput: model.Output{}},
		{name: "resultItemErr", entity: model.Input{Id: "1"}, resultGetItem: &dynamodb.GetItemOutput{}, expectedOutput: model.Output{}, expectedErr: errors.New("device not found")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dyMock := mock.NewMockDynamo(test.resultGetItem, test.getItemErr)
			core := NewCore(dyMock, dynamodbattribute.UnmarshalMap)

			output, err := core.Handler(context.TODO(), test.entity)

			if err == nil {
				assert.Nil(t, test.expectedErr)
			} else {
				assert.EqualError(t, err, test.expectedErr.Error())
			}
			assert.Equal(t, test.expectedOutput.Name, output.Name)
		})
	}
}


