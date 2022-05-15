package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"time"
)

type ResponseBody struct {
	Timestamp string `json:"timestamp"`
	Ok        bool   `json:"ok"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Data      string `json:"data"`
}

func successResponse(status int, data interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       string(createResponseBody(true, status, data)),
	}

	return &resp, nil
}

func errorResponse(status int, data interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       string(createResponseBody(false, status, data)),
	}

	return &resp, nil
}

func createResponseBody(ok bool, status int, data interface{}) []byte {
	dataInString, _ := json.Marshal(data)
	rb := ResponseBody{
		Timestamp: time.Now().Format(time.RFC1123Z),
		Ok:        ok,
		Status:    status,
		Message:   http.StatusText(status),
		Data:      string(dataInString),
	}

	rbInString, _ := json.Marshal(rb)
	return rbInString
}
