package main

import (
	"AWS_challenge1/createDevice/serviceHandler"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

var dynamoDB *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Println(err)
		return
	}
	dynamoDB = dynamodb.New(awsSession)
}
func main() {
	core := serviceHandler.NewCore(dynamoDB, dynamodbattribute.MarshalMap)
	lambda.Start(core.Handler)
}
