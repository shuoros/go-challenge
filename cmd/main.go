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

/*
	This is the entry point for the lambda function.
*/
func main() {
	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return
	}

	dynaClient = dynamodb.New(awsSession)

	lambda.Start(handler)
}

/*
	This function handles the incoming request based on its http method and then calls the appropriate handler.
	If the given request's http type is GET, it will call the Get handler and if it is POST, it will call the
	Post handler otherwise it will call the default handler which is the UnhandledMethod handler.
*/
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
