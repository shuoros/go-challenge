package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/shuoros/go-challenge/pkg/handlers"
	"os"
)

const tableName = "VeryImportantTable"

var dynaClient dynamodbiface.DynamoDBAPI

func main() {
	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return
	}

	dynaClient = dynamodb.New(awsSession)

	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetDevice(req, tableName, dynaClient)
	case "POST":
		return handlers.AddDevice(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}
}
