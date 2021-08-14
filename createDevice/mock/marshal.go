package mock

import "github.com/aws/aws-sdk-go/service/dynamodb"

type Marshal func(in interface{}) (map[string]*dynamodb.AttributeValue, error)

//MarshalMock is function for mocking marshal method
func MarshalMock(err error) Marshal {
	return func(in interface{}) (map[string]*dynamodb.AttributeValue, error) {
		return nil, err
	}
}
