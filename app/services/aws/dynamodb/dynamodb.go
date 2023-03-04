package dynamodb

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDB is responsible for interfacing with AWS DynamoDB Service
type DynamoDB struct {
	svc    *dynamodb.DynamoDB
	config iConfig
}

func New(awsSession *session.Session, config iConfig) *DynamoDB {
	dynamoDBSession := dynamodb.New(awsSession)
	return &DynamoDB{
		svc:    dynamoDBSession,
		config: config,
	}
}

type iConfig interface {
	EmailSubscriptionsTableName() string
}

func (s *DynamoDB) itemExists(
	tableName string,
	key map[string]*dynamodb.AttributeValue,
) (bool, error) {
	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}
	result, err := s.svc.GetItem(getInput)
	if err != nil {
		return false, err
	}
	if result.Item != nil {
		return true, nil
	}
	return false, nil
}

func (s *DynamoDB) putItem(
	tableName string,
	item map[string]*dynamodb.AttributeValue,
) error {
	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}
	_, err := s.svc.PutItem(putInput)
	return err
}

func (s *DynamoDB) EmailSubscriptionItemExists(email string) (bool, error) {
	tableName := s.config.EmailSubscriptionsTableName()
	key := map[string]*dynamodb.AttributeValue{
		"email": {
			S: aws.String(email),
		},
	}
	return s.itemExists(tableName, key)
}

func (s *DynamoDB) CreateEmailSubscriptionItem(
	email string,
	subscriptionToken string,
	isVerified bool,
) error {
	timestamp := time.Now().Unix()
	tableName := s.config.EmailSubscriptionsTableName()
	item := map[string]*dynamodb.AttributeValue{
		"email": {
			S: aws.String(email),
		},
		"creation_date": {
			N: aws.String(fmt.Sprintf("%d", timestamp)),
		},
		"subscription_token": {
			S: aws.String(subscriptionToken),
		},
		"email_verified": {
			BOOL: aws.Bool(isVerified),
		},
	}
	return s.putItem(tableName, item)
}
