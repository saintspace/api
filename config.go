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
	value, err = getConfigurationParameter(
		"email-subscriptions-table-name",
		ssmSession,
		false,
	)
	if err != nil {
		return err
	}
	appConfig.EmailSubscriptionsTableName = value
	return nil
}
