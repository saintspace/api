package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	err := initConfig()
	if err != nil {
		err = fmt.Errorf("error while initializing app config => %v", err.Error())
		log.Panic(err.Error())
	}
	log.Printf("Initializing lambda runtime with Gin router")
	ginLambda = ginadapter.New(getRouter())
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
