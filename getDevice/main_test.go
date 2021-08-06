package main

import (
	"AWS_challenge1/getDevice/data"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
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
		entity         data.Input
		getItemErr     error
		resultGetItem  *dynamodb.GetItemOutput
		expectedOutput data.Output
		expectedErr    error
	}{
		{name: "ok", entity: data.Input{Id: "1"}, resultGetItem: &dynamodb.GetItemOutput{Item: map[string]*dynamodb.AttributeValue{"name": &dynamodb.AttributeValue{S: aws.String("mohammad")}}}, expectedOutput: data.Output{Name: "mohammad"}},
		{name: "getItemErr", entity: data.Input{Id: "1"}, getItemErr: errors.New(""), expectedErr: errors.New("server error"), expectedOutput: data.Output{}},
		{name:"resultItemErr",entity: data.Input{Id: "1"},resultGetItem: &dynamodb.GetItemOutput{},expectedOutput: data.Output{},expectedErr: errors.New("device not found")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dyMock := newMockDynamo(test.resultGetItem, test.getItemErr)

			core := Core{
				db:           dyMock,
				unMarshalMap: dynamodbattribute.UnmarshalMap,
			}

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

type dynamoMock struct {
	dynamodbiface.DynamoDBAPI
	ErrReturn    error
	resultReturn *dynamodb.GetItemOutput
}

func (s dynamoMock) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return s.resultReturn, s.ErrReturn
}

func newMockDynamo(result *dynamodb.GetItemOutput, err error) dynamoMock {
	return dynamoMock{ErrReturn: err, resultReturn: result}
}
