package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/shuoros/go-challenge/pkg/device"
	"net/http"
)

func AddDevice(req events.APIGatewayProxyRequest, table string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	result, err := device.CreateDevice(req, table, dynaClient)
	if err != nil {
		return response(http.StatusBadRequest, aws.String(err.Error()))
	}

	return response(http.StatusCreated, result)
}

func GetDevice(req events.APIGatewayProxyRequest, table string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	deviceId := req.QueryStringParameters["id"]

	result, err := device.FetchDevice(deviceId, table, dynaClient)
	if err != nil {
		return response(http.StatusBadRequest, aws.String(err.Error()))
	}

	return response(http.StatusOK, result)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return response(http.StatusMethodNotAllowed, "Method Not Allowed")
}
