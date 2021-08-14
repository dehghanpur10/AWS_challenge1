package serviceHandler

import (
	"AWS_challenge1/getDevice/model"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"log"
	"os"
)
//unMarshalType is type for unMarshal function
type unMarshalType func(m map[string]*dynamodb.AttributeValue, out interface{}) error


//NewCore is function for create new core for handler lambada
func NewCore(db dynamodbiface.DynamoDBAPI, unMarshal unMarshalType) *Core {
	return &Core{
		db:           db,
		unMarshalMap: unMarshal,
	}
}

//Core is struct for handle request, dynamoDB client and marshalMap are dependency injection
type Core struct {
	db           dynamodbiface.DynamoDBAPI
	unMarshalMap unMarshalType
}

//Handler is a lambda for handle post request from api Getway
func (d *Core) Handler(ctx context.Context, entity model.Input) (model.Output, error) {
	key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(entity.Id),
		},
	}
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key:       key,
	}
	result, err := d.db.GetItem(getItemInput)
	if err != nil {
		log.Println(err)
		return model.Output{}, errors.New("server error")
	}
	if result.Item == nil {
		log.Println(err)
		return model.Output{}, errors.New("device not found")
	}
	var device model.Output
	err = d.unMarshalMap(result.Item, &device)
	if err != nil {
		log.Println(err)
		return model.Output{}, errors.New("server error")
	}
	return device, nil

}

