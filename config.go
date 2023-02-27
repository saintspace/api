package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func getConfigurationParameter(paramName string, ssmSession *ssm.SSM, isSecret bool) (string, error) {
	parameter, err := ssmSession.GetParameter(&ssm.GetParameterInput{
		Name:           &paramName,
		WithDecryption: &isSecret,
	})
	if err != nil {
		return "", fmt.Errorf("error while retrieving %s parameter: %s", paramName, err.Error())
	}
	value := *parameter.Parameter.Value
	return value, nil
}

type AppConfig struct {
	EmailSubscriptionsTableName string
	WorkerTasksTopicArn         string
}

var appConfig AppConfig = AppConfig{
	EmailSubscriptionsTableName: "",
}

func initConfig() error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ssmSession := ssm.New(sess)
	var err error
	var value string
	// Get EmailSubscriptionsTableName parameter
	value, err = getConfigurationParameter(
		"email-subscriptions-table-name",
		ssmSession,
		false,
	)
	if err != nil {
		return err
	}
	appConfig.EmailSubscriptionsTableName = value
	// Get WorkerTasksTopicArn parameter
	value, err = getConfigurationParameter(
		"worker-tasks-topic-arn",
		ssmSession,
		false,
	)
	if err != nil {
		return err
	}
	appConfig.WorkerTasksTopicArn = value
	return nil
}
