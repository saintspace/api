package app

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	lambdaHttpRouterAdapter iLambdaHttpRouterAdatper
}

type iLambdaHttpRouterAdatper interface {
	ProxyWithContext(
		ctx context.Context,
		req events.APIGatewayProxyRequest,
	) (events.APIGatewayProxyResponse, error)
}

func New(lambdaHttpRouterAdapter iLambdaHttpRouterAdatper) *App {
	return &App{
		lambdaHttpRouterAdapter: lambdaHttpRouterAdapter,
	}
}

func (s *App) Start() {
	lambda.Start(
		func(
			ctx context.Context,
			req events.APIGatewayProxyRequest,
		) (
			events.APIGatewayProxyResponse,
			error,
		) {
			return s.lambdaHttpRouterAdapter.ProxyWithContext(ctx, req)
		},
	)
}
