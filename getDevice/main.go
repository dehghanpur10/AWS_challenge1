package main

import (
	"AWS_challenge1/getDevice/data"
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"os"
)

//Core is struct for handle request, dynamoDB client and marshalMap are dependency injection
type Core struct {
	db         dynamodbiface.DynamoDBAPI
	unMarshalMap func (m map[string]*dynamodb.AttributeValue, out interface{}) error
}

//Handler is a lambda for handle post request from api Getway
func (d *Core) Handler(ctx context.Context, entity data.Input) (data.Output, error) {
	key := map[string]*dynamodb.AttributeValue{
		"Id": {
			S: aws.String(entity.Id),
		},
	}
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key:       key,
	}
	result, err := d.db.GetItem(getItemInput)
	if err != nil {
		return data.Output{}, errors.New("server error")
	}
	if result.Item == nil {
		return data.Output{}, errors.New("device not found")
	}
	var device data.Output
	err = d.unMarshalMap(result.Item, &device)
	if err != nil {
		return data.Output{}, errors.New("server error")
	}
	return device, nil

}
func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return
	}
	dynaClient := dynamodb.New(awsSession)
	core := Core{
		db:         dynaClient,
		unMarshalMap: dynamodbattribute.UnmarshalMap,
	}
	lambda.Start(core.Handler)
}
