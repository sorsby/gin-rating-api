package dynamo

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createTestTable() (client *dynamodb.DynamoDB, name string, err error) {
	config := aws.NewConfig().
		WithEndpoint("http://localhost:4569/").
		WithRegion("eu-west-2").
		WithCredentials(credentials.NewStaticCredentials("fake", "fake", ""))
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: *config,
	})
	if err != nil {
		return
	}
	name = fmt.Sprintf("test_gin_rating_%d", time.Now().UnixNano())
	client = dynamodb.New(sess)
	_, err = client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(name),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("pk"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("sk"), AttributeType: aws.String("S")},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("pk"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("sk"), KeyType: aws.String("RANGE")},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("gsi1"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String("sk"), KeyType: aws.String("HASH")},
					{AttributeName: aws.String("pk"), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
				},
			},
		},
	})
	if err != nil {
		return
	}
	return
}

func dropTable(client *dynamodb.DynamoDB, name string) error {
	_, err := client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(name),
	})
	return err
}
