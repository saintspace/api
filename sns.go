package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"fmt"
)

func publishTask(message string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	_, err := svc.Publish(&sns.PublishInput{
		Message:  &message,
		TopicArn: &appConfig.WorkerTasksTopicArn,
	})
	return fmt.Errorf("error while publishing task to SNS => %v", err.Error())
}
