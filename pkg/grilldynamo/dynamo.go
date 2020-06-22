package grilldynamo

import (
	"context"

	"bitbucket.org/swigy/grill/internal/canned"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Dynamo struct {
	dynamo *canned.DynamoDB
}

func Start() (*Dynamo, error) {
	dynamo, err := canned.NewDynamoDB(context.TODO())
	if err != nil {
		return nil, err
	}

	return &Dynamo{
		dynamo: dynamo,
	}, nil
}

func (gd *Dynamo) Client() dynamodbiface.DynamoDBAPI {
	return gd.dynamo.Client
}

func (gd *Dynamo) Host() string {
	return gd.dynamo.Host
}

func (gd *Dynamo) Port() string {
	return gd.dynamo.Port
}

func (gd *Dynamo) Region() string {
	return gd.dynamo.Region
}

func (gd *Dynamo) AccessKey() string {
	return gd.dynamo.AccessKey
}

func (gd *Dynamo) SecretKey() string {
	return gd.dynamo.SecretKey
}

func (gd *Dynamo) Stop() error {
	return gd.dynamo.Container.Terminate(context.TODO())
}
