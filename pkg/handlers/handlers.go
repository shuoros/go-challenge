package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/shuoros/go-challenge/pkg/device"
	"net/http"
	"strconv"
	"strings"
)

func AddDevice(req events.APIGatewayProxyRequest, table string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	result, err := device.CreateDevice(req, table, dynaClient)
	if err != nil {
		return CatchException(err)
	}

	return SuccessResponse(http.StatusCreated, result)
}

func GetDevice(req events.APIGatewayProxyRequest, table string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	deviceId := req.PathParameters["id"]

	result, err := device.FetchDevice(deviceId, table, dynaClient)
	if err != nil {
		return CatchException(err)
	}

	return SuccessResponse(http.StatusOK, result)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return ErrorResponse(http.StatusMethodNotAllowed, "Method Not Allowed")
}

func CatchException(err error) (*events.APIGatewayProxyResponse, error) {
	arr := strings.Split(err.Error(), "-")
	if len(arr) == 2 {
		statusCode, _ := strconv.Atoi(arr[0])
		return ErrorResponse(statusCode, aws.String(arr[1]))
	} else {
		return ErrorResponse(http.StatusInternalServerError, aws.String(err.Error()))
	}
}
