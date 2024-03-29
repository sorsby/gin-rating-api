package dynamo

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sorsby/gin-rating-api/data"
)

const gsiIndexName = "gsi1"

// Manager manages interactions with the dynamodb table.
type Manager struct {
	db        *dynamodb.DynamoDB
	tableName string
	now       func() time.Time
}

// NewManager creates a new manager.
func NewManager(tableName string) *Manager {
	return &Manager{
		db:        dynamodb.New(session.New()),
		tableName: tableName,
		now:       func() time.Time { return time.Now() },
	}
}

// CreateGin attempts to create a unique gin and insert to dynamo, returns false if gin already exists.
func (mgr *Manager) CreateGin(in data.CreateGinInput) error {
	var writeItems []*dynamodb.TransactWriteItem

	ngi := newListGinItem(in, mgr.now())
	ginJSON, err := dynamodbattribute.MarshalMap(ngi)
	if err != nil {
		return fmt.Errorf("dynamo.CreateGin: error marshalling gin item: %w", err)
	}
	gp := dynamodb.Put{
		TableName: aws.String(mgr.tableName),
		Item:      ginJSON,
	}
	writeItems = append(writeItems, &dynamodb.TransactWriteItem{Put: &gp})

	nngi := newNamedGinItem(in, mgr.now())
	namedGinJSON, err := dynamodbattribute.MarshalMap(nngi)
	if err != nil {
		return fmt.Errorf("dynamo.CreateGin: error marshalling named gin item: %w", err)
	}
	ngp := dynamodb.Put{
		TableName: aws.String(mgr.tableName),
		Item:      namedGinJSON,
	}
	writeItems = append(writeItems, &dynamodb.TransactWriteItem{Put: &ngp})
	_, err = mgr.db.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		TransactItems: writeItems,
	})
	return err
}

// ListGins lists all the gins in the database.
func (mgr *Manager) ListGins() (o data.ListGinOutput, err error) {
	out, err := mgr.db.Query(&dynamodb.QueryInput{
		TableName:              aws.String(mgr.tableName),
		KeyConditionExpression: aws.String("pk=:pk_value AND begins_with (sk, :sk_value)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk_value": &dynamodb.AttributeValue{S: aws.String(ginFilter)},
			":sk_value": &dynamodb.AttributeValue{S: aws.String("gin_")},
		},
	})
	o.GinItems = make([]data.GinItem, len(out.Items))
	err = dynamodbattribute.UnmarshalListOfMaps(out.Items, &o.GinItems)
	return o, err
}

// GetGin gets a single gin by name if it exists.
func (mgr *Manager) GetGin(name string) (gi data.GinItem, found bool, err error) {
	out, err := mgr.db.Query(&dynamodb.QueryInput{
		TableName:              aws.String(mgr.tableName),
		KeyConditionExpression: aws.String("pk=:pk_value"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk_value": &dynamodb.AttributeValue{S: aws.String(name)},
		},
	})
	if err != nil {
		return
	}
	if len(out.Items) != 1 {
		return
	}
	err = dynamodbattribute.UnmarshalMap(out.Items[0], &gi)
	return gi, true, err
}
