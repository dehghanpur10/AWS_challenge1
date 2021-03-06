//Package mock for mocking dependency
package mock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dynamoMock struct {
	dynamodbiface.DynamoDBAPI
	ErrReturn error
}

func (s dynamoMock) PutItem(i *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, s.ErrReturn
}
//NewMockDynamo is func that return mock dynamoDB
func NewMockDynamo(err error) *dynamoMock {
	return &dynamoMock{ErrReturn: err}
}
