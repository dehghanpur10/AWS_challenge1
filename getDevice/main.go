package main

import (
	"AWS_challenge1/getDevice/serviceHandler"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

var dynamoDB *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	accessToken := os.Getenv("ACCESS_TOKEN")
	secretKey := os.Getenv("SECRET_KEY")
	credential := credentials.NewStaticCredentials(accessToken, secretKey, "")
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credential,
	},

	)
	if err != nil {
		log.Println(err)
		return
	}
	dynamoDB = dynamodb.New(awsSession)
}
func main() {
	core := serviceHandler.NewCore(dynamoDB, dynamodbattribute.UnmarshalMap)
	lambda.Start(core.Handler)
}
