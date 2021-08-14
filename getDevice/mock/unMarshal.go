package mock

import (
	"AWS_challenge1/getDevice/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type UnMarshalType func(m map[string]*dynamodb.AttributeValue, out interface{}) error

func UnMarshalMock(err error) UnMarshalType {
	unmarshal := func(m map[string]*dynamodb.AttributeValue, out interface{}) error {
		out.(*model.Output).Name = aws.StringValue(m["name"].S)
		return err
	}
	return unmarshal
}
