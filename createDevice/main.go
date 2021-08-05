package main

import (
	"AWS_challenge1/createDevice/data"
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
	marshalMap func(in interface{}) (map[string]*dynamodb.AttributeValue, error)
}

//Handler is a lambda for handle post request from api Getway
func (d *Core) Handler(ctx context.Context, entity data.Input) (data.Output, error) {
	av, err := d.marshalMap(entity)
	if err != nil {
		return data.Output{}, errors.New("server error")
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	_, err = d.db.PutItem(input)
	if err != nil {
		return data.Output{}, errors.New("server error")
	}

	return data.Output{Message: "device added successfully"}, nil
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
		marshalMap: dynamodbattribute.MarshalMap,
	}
	lambda.Start(core.Handler)
}
