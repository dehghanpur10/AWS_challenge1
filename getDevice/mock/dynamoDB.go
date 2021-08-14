//Package mock for mocking dependency
package mock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dynamoMock struct {
	dynamodbiface.DynamoDBAPI
	ErrReturn error
	Item      map[string]*dynamodb.AttributeValue
}

func (s dynamoMock) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{
		Item: s.Item,
	}, s.ErrReturn
}

//NewMockDynamo is func that return mock dynamoDB
func NewMockDynamo(item map[string]*dynamodb.AttributeValue, err error) *dynamoMock {
	return &dynamoMock{
		ErrReturn: err,
		Item:      item,
	}
}
