package test

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/shuoros/go-challenge/pkg/device"
	"github.com/shuoros/go-challenge/pkg/handlers"
	"net/http"
	"testing"
)

const tableName = "VeryImportantTable"

// A fakeDynamoDB instance for mocking test that emulates real DynamoDB
type FakeDynamoDBAPI struct {
	dynamodbiface.DynamoDBAPI
}

// CreateTestCase struct that contains all requested and expected values for unit testing
type CreateTestCase struct {
	Name               string
	Request            events.APIGatewayProxyRequest
	Device             device.Device
	ExpectedStatusCode int
}

// GetTestCase struct that contains all requested and expected values for unit testing
type GetTestCase struct {
	Name                   string
	InputId                events.APIGatewayProxyRequest
	InputIdString          string
	DatabaseOutput         dynamodb.GetItemOutput
	Error                  error
	ExpectedBody           string
	ExpectedStatusCode     int
	ExpectedDatabaseOutput dynamodb.GetItemOutput
}

/*
 a mocked version of DynamoDB's PutItem function.
 in testing state, instead of calling real DynamoDB's PutItem, we try to emulate it.
 insertItemToDatabase function of addDevice.go calls this function in Testing state.
*/
func (fd *FakeDynamoDBAPI) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return new(dynamodb.PutItemOutput), nil
}

/*
 a mocked version of DynamoDB's GetItem function.
 in testing state, instead of calling real DynamoDB's GetItem, we try to emulate it.
 Get function of getDeviceById.go calls this function in Testing state.
*/
func (fd *FakeDynamoDBAPI) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	output := new(dynamodb.GetItemOutput)
	id := input.Key["id"].S

	if *id == "id_test" {
		output.SetItem(
			map[string]*dynamodb.AttributeValue{
				"id":          &dynamodb.AttributeValue{S: aws.String("id_test")},
				"deviceModel": &dynamodb.AttributeValue{S: aws.String("deviceModel_test")},
				"name":        &dynamodb.AttributeValue{S: aws.String("name_test")},
				"note":        &dynamodb.AttributeValue{S: aws.String("note_test")},
				"serial":      &dynamodb.AttributeValue{S: aws.String("serial_test")},
			},
		)
	}

	return output, nil
}

func Test_CreateDevice(t *testing.T) {
	dynaClient := &FakeDynamoDBAPI{}
	testCases := []CreateTestCase{
		{
			Name:               "** Testing empty body input **",
			Request:            events.APIGatewayProxyRequest{Body: ""},
			ExpectedStatusCode: 422,
		},
		{
			Name:               "** Testing wrong json format **",
			Request:            events.APIGatewayProxyRequest{Body: "{{{}"},
			ExpectedStatusCode: 422,
		},
		{
			Name:               "** Testing json with missing field {id} **",
			Request:            events.APIGatewayProxyRequest{Body: "{\"id\":\"\" , \"deviceModel\":\"testDeviceModel\" , \"name\":\"testName\" , \"note\":\"testNote\" , \"serial\":\"testSerial\" }"},
			ExpectedStatusCode: 422,
		},
		{
			Name:               "** Testing json with missing field {deviceModel, note} **",
			Request:            events.APIGatewayProxyRequest{Body: "{\"id\":\"1\" , \"deviceModel\":\"\" , \"name\":\"testName\" , \"note\":\"\" , \"serial\":\"testSerial\" }"},
			ExpectedStatusCode: 422,
		},

		{
			Name:               "** Testing json with missing field {serial, name, deviceModel} **",
			Request:            events.APIGatewayProxyRequest{Body: "{\"id\":\"1\" , \"deviceModel\":\"\" , \"name\":\"\" , \"note\":\"testNote\" , \"serial\":\"\" }"},
			ExpectedStatusCode: 422,
		},

		{
			Name:               "** Testing valid json with all fields **",
			Request:            events.APIGatewayProxyRequest{Body: "{\"id\":\"1\" , \"deviceModel\":\"testDeviceModel\" , \"name\":\"testName\" , \"note\":\"testNote\" , \"serial\":\"testSerial\"}"},
			ExpectedStatusCode: 201,
		},
	}

	for _, test := range testCases {

		// calls AddDevice function.
		resp, _ := handlers.AddDevice(test.Request, tableName, dynaClient)

		if resp.StatusCode != test.ExpectedStatusCode {
			t.Errorf("%s \n \t<expected error-code: %d> <resulted error-code: %d>", test.Name, test.ExpectedStatusCode, resp.StatusCode)
		}

	}

}

func Test_GetDevice(t *testing.T) {
	dynaClient := &FakeDynamoDBAPI{}
	testCases := []GetTestCase{
		{
			Name: "** Testing empty input id **",
			InputId: events.APIGatewayProxyRequest{PathParameters: map[string]string{
				"id": ""}},
			ExpectedStatusCode: 404,
		},
		{
			Name: "** Testing valid input id **",
			InputId: events.APIGatewayProxyRequest{PathParameters: map[string]string{
				"id": "id_test"}},
			ExpectedStatusCode: 200,
		},
	}

	for _, test := range testCases {

		// calls GetDevice function.
		resp, _ := handlers.GetDevice(test.InputId, tableName, dynaClient)

		if resp.StatusCode != test.ExpectedStatusCode {
			t.Errorf("%s \n \t<expected error-code: %d> <resulted error-code: %d>", test.Name, test.ExpectedStatusCode, resp.StatusCode)
		}
	}

}

func Test_UnhandledMethodMustReturnAnErrorResponseWith405Status(t *testing.T) {
	resp, _ := handlers.UnhandledMethod()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func Test_CatchExceptionMustCatchErrorCodeAndReturnValidResponseWithRightStatusCode(t *testing.T) {
	resp, _ := handlers.CatchException(errors.New("200-OK"))
	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %q, wanted %q", resp.StatusCode, http.StatusOK)
	}
}

func Test_CatchExceptionMustCatchErrorCodeAndReturnValidResponseWithRightMessage(t *testing.T) {
	resp, _ := handlers.CatchException(errors.New("200-OK"))
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(resp.Body), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["message"] != "OK" {
		t.Errorf("got %q, wanted %q", resp.StatusCode, http.StatusOK)
	}
}
