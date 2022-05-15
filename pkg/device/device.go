package device

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorFailedToPutDevice       = "Failed to put device"
	ErrorFailedToFetchDevice     = "failed to fetch device"
	ErrorFailedToMarshalDevice   = "failed to marshal device"
	ErrorFailedToUnmarshalDevice = "failed to unmarshal device"
	ErrorInvalidDeviceData       = "invalid device data"
	ErrorDeviceAlreadyExists     = "device already exists"
)

type Device struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	Serial string `json:"serial"`
	Note   string `json:"note"`
}

func CreateDevice(req events.APIGatewayProxyRequest, table string, dynaClient dynamodbiface.DynamoDBAPI) (
	*Device, error) {
	// Validate request
	var d Device
	if err := json.Unmarshal([]byte(req.Body), &d); err != nil {
		return nil, errors.New(ErrorInvalidDeviceData)
	}

	// Check if device already exists
	possibleExistingDevice, _ := FetchDevice(d.ID, table, dynaClient)
	if possibleExistingDevice != nil && len(possibleExistingDevice.ID) != 0 {
		return nil, errors.New(ErrorDeviceAlreadyExists)
	}

	// Create device
	item, err := dynamodbattribute.MarshalMap(d)
	if err != nil {
		return nil, errors.New(ErrorFailedToMarshalDevice)
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(table),
	}
	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToPutDevice)
	}

	return &d, nil
}

func FetchDevice(deviceId string, table string, dynaClient dynamodbiface.DynamoDBAPI) (*Device, error) {
	// query the device
	input := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(deviceId),
			},
		},
	}

	// fetch the device
	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchDevice)
	}

	// unmarshal the device
	item := new(Device)
	if err := dynamodbattribute.UnmarshalMap(result.Item, item); err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalDevice)
	}

	return item, nil
}
