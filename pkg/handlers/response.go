package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"time"
)

/*
	ResponseBody is a custom representation for the APIGatewayProxyResponse body.
*/
type ResponseBody struct {
	Timestamp string `json:"timestamp"`
	Ok        bool   `json:"ok"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Data      string `json:"data"`
}

/*
	SuccessResponse creates a custom success response body (ok = true) for the APIGatewayProxyResponse based on the
	given status code and message.
*/
func SuccessResponse(status int, data interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       CreateResponseBody(true, status, data),
	}

	return &resp, nil
}

/*
	ErrorResponse creates a custom success response body (ok = false) for the APIGatewayProxyResponse based on the
	given status code and message.
*/
func ErrorResponse(status int, data interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       CreateResponseBody(false, status, data),
	}

	return &resp, nil
}

/*
	CreateResponseBody creates a custom response body based on the given status code and message.
*/
func CreateResponseBody(ok bool, status int, data interface{}) string {
	dataInString, _ := json.Marshal(data)
	rb := ResponseBody{
		Timestamp: time.Now().Format(time.RFC1123Z),
		Ok:        ok,
		Status:    status,
		Message:   http.StatusText(status),
		Data:      string(dataInString),
	}

	rbInString, _ := json.Marshal(rb)
	return string(rbInString)
}
