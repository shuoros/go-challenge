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

// AddDevice /**
// * @api {post} /api/devices
// * @apiDescription Add a new device in the database
// * @apiParamExample {json} Request-Example:
// * {
// * 	"name": "id1",
// * 	"description": "model1",
// * 	"type": "Sensor",
// * 	"location": "Testing a sensor.",
// * 	"serial": "A020000102"
// * }
// * @apiSuccessExample {json} Success-Response:
// * HTTP/1.1 201 OK
// * {
// * 	"timestamp": "Sun, 15 May 2022 20:31:55 +0000",
// * 	"ok": true,
// * 	"status": 201,
// * 	"message": "Created",
// * 	"data": "{
// * 		"id": "id1",
// * 		"deviceModel": "model1",
// * 		"name": "Sensor",
// * 		"note": "Testing a sensor.",
// * 		"serial": "A020000102"}"
// * }
// */
func AddDevice(req events.APIGatewayProxyRequest, table string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	result, err := device.CreateDevice(req, table, dynaClient)
	if err != nil {
		return CatchException(err)
	}

	return SuccessResponse(http.StatusCreated, result)
}

// GetDevice /**
// * @api {post} /api/devices{id}
// * @apiDescription Get a device from the database by its id
// * @pathParam {String} [id] Unique identifier of the device
// * @apiSuccessExample {json} Success-Response:
// * HTTP/1.1 200 OK
// * {
// * 	"timestamp": "Sun, 15 May 2022 20:35:03 +0000",
// * 	"ok": true,
// * 	"status": 200,
// * 	"message": "OK",
// * 	"data": "{
// * 		"id": "id1",
// * 		"deviceModel": "model1",
// * 		"name": "Sensor",
// * 		"note": "Testing a sensor.",
// * 		"serial": "A020000102"}"
// * }
// */
func GetDevice(req events.APIGatewayProxyRequest, table string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	deviceId := req.PathParameters["id"]

	result, err := device.FetchDevice(deviceId, table, dynaClient)
	if err != nil {
		return CatchException(err)
	}

	return SuccessResponse(http.StatusOK, result)
}

/*
	return a http 405 error
*/
func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return ErrorResponse(http.StatusMethodNotAllowed, "Method Not Allowed")
}

/*
	Catches exceptions and extract status code and message and return a http error response with the status code
	and message.
*/
func CatchException(err error) (*events.APIGatewayProxyResponse, error) {
	arr := strings.Split(err.Error(), "-")
	if len(arr) == 2 {
		statusCode, _ := strconv.Atoi(arr[0])
		return ErrorResponse(statusCode, aws.String(arr[1]))
	} else {
		return ErrorResponse(http.StatusInternalServerError, aws.String(err.Error()))
	}
}
