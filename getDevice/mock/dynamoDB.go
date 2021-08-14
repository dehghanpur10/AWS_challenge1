package mock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dynamoMock struct {
	dynamodbiface.DynamoDBAPI
	ErrReturn    error
	resultReturn *dynamodb.GetItemOutput
}

func (s dynamoMock) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return s.resultReturn, s.ErrReturn
}

func NewMockDynamo(result *dynamodb.GetItemOutput, err error) *dynamoMock {
	return &dynamoMock{
		ErrReturn:    err,
		resultReturn: result,
	}
}
