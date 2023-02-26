package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createEmailSubscription(email string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	// Check if the email address is already subscribed
	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(appConfig.EmailSubscriptionsTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
	}
	result, err := svc.GetItem(getInput)
	if err != nil {
		return fmt.Errorf("error while checking for email subscription => %v", err.Error())
	}
	if result.Item != nil { // Do nothing if the email address is already subscribed
		return nil
	}
	// Proceed to add the email subscription to the database
	timestamp := time.Now().Unix()
	subscriptionToken, err := GenerateEmailSubscriptionToken(email)
	if err != nil {
		return fmt.Errorf("error while generating email subscription token => %v", err.Error())
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String("my-email-subscriptions-table"),
		Item: map[string]*dynamodb.AttributeValue{
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
				BOOL: aws.Bool(false),
			},
		},
	}

	// Add the item to the table using the PutItem operation
	_, err = svc.PutItem(putInput)
	if err != nil {
		return err
	}
	return nil
}
