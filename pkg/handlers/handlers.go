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
		return errorResponse(http.StatusBadRequest, aws.String(err.Error()))
	}

	return successResponse(http.StatusCreated, result)
}

func GetDevice(req events.APIGatewayProxyRequest, table string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	deviceId := req.PathParameters["id"]

	result, err := device.FetchDevice(deviceId, table, dynaClient)
	if err != nil {
		return errorResponse(http.StatusBadRequest, aws.String(err.Error()))
	}

	return successResponse(http.StatusOK, result)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return errorResponse(http.StatusMethodNotAllowed, "Method Not Allowed")
}
