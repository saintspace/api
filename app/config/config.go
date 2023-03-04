package config

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Config struct {
	systemManager      *ssm.SSM
	standardParameters map[string]string
	secretParameters   map[string]string
}

const (
	taskSnsTopicArnParameterName             = "worker-tasks-topic-arn"
	emailSubscriptionsTableNameParameterName = "email-subscriptions-table-name"
)

func New(awsSession *session.Session) *Config {
	return &Config{
		systemManager: ssm.New(awsSession),
		standardParameters: map[string]string{
			taskSnsTopicArnParameterName:             "",
			emailSubscriptionsTableNameParameterName: "",
		},
		secretParameters: map[string]string{},
	}
}

func (s *Config) initStandardParameters() error {
	paramNames := []*string{}
	for paramName, _ := range s.standardParameters {
		paramNames = append(paramNames, aws.String(paramName))
	}
	paramOutput, err := s.systemManager.GetParameters(&ssm.GetParametersInput{
		Names:          paramNames,
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		return fmt.Errorf("error while retrieving standard SSM parameters => %v", err.Error())
	}
	for _, param := range paramOutput.Parameters {
		s.standardParameters[*param.Name] = *param.Value
	}
	return nil
}

func (s *Config) InitializeParameters() error {
	return s.initStandardParameters()
}

func (s *Config) TaskSnsTopicArn() string {
	return s.standardParameters[taskSnsTopicArnParameterName]
}

func (s *Config) EmailSubscriptionsTableName() string {
	return s.standardParameters[emailSubscriptionsTableNameParameterName]
}
