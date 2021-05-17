package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := log.New()
	log.FieldLogger(logger).Info("field logger")
	log.WithField("component", "handler").Info("starting handler invocation...")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello AWS Lambda and Netlify",
	}, nil
}

func main() {
	log.WithField("component", "main").Info("starting function call...")
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
