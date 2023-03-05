package main

import (
	"fmt"
	"log"

	"api/app"
	"api/app/config"
	"api/app/handler"
	"api/app/router"
	"api/app/services/aws/dynamodb"
	"api/app/services/aws/sns"
	"api/app/services/datastore"
	"api/app/services/email"
	"api/app/services/taskpub"

	"github.com/aws/aws-sdk-go/aws/session"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var application *app.App

func init() {

	// Retrieve application parameters

	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	configService := config.New(awsSession)
	err := configService.InitializeParameters()
	if err != nil {
		err = fmt.Errorf("error while initializing app config => %v", err.Error())
		log.Panic(err.Error())
	}

	// Build application dependencies

	dynamoDbService := dynamodb.New(awsSession, configService)
	snsService := sns.New(awsSession, configService)
	datastoreService := datastore.New(dynamoDbService)
	taskPublisherService := taskpub.New(snsService, configService)
	emailService := email.New(datastoreService, taskPublisherService)
	routeHandler := handler.New(emailService)
	apiRouter := router.New(routeHandler)
	ginLambdaAdapter := ginadapter.New(apiRouter.GetRouter())
	application = app.New(ginLambdaAdapter)
}

func main() {
	application.Start()
}
