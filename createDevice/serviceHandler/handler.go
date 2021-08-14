package serviceHandler

import (
	"AWS_challenge1/createDevice/model"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"log"
	"os"
)

//marshalType is type for marshal function
type marshalType func(in interface{}) (map[string]*dynamodb.AttributeValue, error)

//CoreInterface is interface for core
type CoreInterface interface {
	Handler(ctx context.Context, entity model.Input) (model.Output, error)
}

//NewCore is function for create new core for handler lambada
func NewCore(db dynamodbiface.DynamoDBAPI, marshal marshalType) CoreInterface {
	return &Core{
		db:         db,
		marshalMap: marshal,
	}
}

//Core is struct for handle request, dynamoDB client and marshalMap are dependency injection
type Core struct {
	db         dynamodbiface.DynamoDBAPI
	marshalMap marshalType
}

//Handler is a lambda for handle post request from api Getway
func (d *Core) Handler(ctx context.Context, entity model.Input) (model.Output, error) {
	device, err := d.marshalMap(entity)
	if err != nil {
		log.Println(err)
		return model.Output{}, errors.New("server error")
	}
	input := &dynamodb.PutItemInput{
		Item:      device,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	_, err = d.db.PutItem(input)
	if err != nil {
		log.Println(err)

		return model.Output{}, errors.New("server error")
	}
	return model.Output{Message: "device added successfully"}, nil
}
