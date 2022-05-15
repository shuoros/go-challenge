package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func response(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	bodyInString, _ := json.Marshal(body)

	resp := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       string(bodyInString),
	}

	return &resp, nil
}
