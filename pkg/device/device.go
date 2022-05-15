package device

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"strings"
)

var (
	ErrorFailedToPutDevice       = "500-Failed to create device!"
	ErrorFailedToFetchDevice     = "500-Failed to fetch device!"
	ErrorFailedToMarshalDevice   = "500-Failed to marshal device!"
	ErrorFailedToUnmarshalDevice = "500-Failed to unmarshal device!"
	ErrorInvalidDeviceData       = "422-Please provide valid device data!"
	ErrorDeviceAlreadyExists     = "409-Device with given id is already exists!"
	ErrorDeviceNotFound          = "404-Device with given id is not found!"
)

type Device struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Model  string `json:"deviceModel"`
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

	// if some fields are missing, report it as an error
	errorMessage, errorFlag := ValidateFields(d)
	if errorFlag == true {
		return nil, errors.New(errorMessage)
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

	_, errorFlag := ValidateFields(*item)
	if errorFlag == true {
		return nil, errors.New(ErrorDeviceNotFound)
	}

	return item, nil
}

func ValidateFields(d Device) (string, bool) {
	errorMessage := "422-Following fields are not provided: "
	var errorFlag bool = false

	if len(d.ID) == 0 {
		errorMessage += "id, "
		errorFlag = true
	}

	if len(d.Model) == 0 {
		errorMessage += "deviceModel, "
		errorFlag = true
	}

	if len(d.Name) == 0 {
		errorMessage += "name, "
		errorFlag = true
	}

	if len(d.Note) == 0 {
		errorMessage += "note, "
		errorFlag = true
	}

	if len(d.Serial) == 0 {
		errorMessage += "serial, "
		errorFlag = true
	}

	return strings.TrimSuffix(errorMessage, ", "), errorFlag
}
