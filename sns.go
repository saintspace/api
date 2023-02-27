package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/google/uuid"
)

func publishTask(message string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)
	uniqueMessageId := uuid.New().String()
	_, err := svc.Publish(&sns.PublishInput{
		Message:                &message,
		TopicArn:               &appConfig.WorkerTasksTopicArn,
		MessageGroupId:         &uniqueMessageId,
		MessageDeduplicationId: &uniqueMessageId,
	})
	return err
}
