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

func SuccessResponse(status int, data interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       CreateResponseBody(true, status, data),
	}

	return &resp, nil
}

func ErrorResponse(status int, data interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       CreateResponseBody(false, status, data),
	}

	return &resp, nil
}

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
