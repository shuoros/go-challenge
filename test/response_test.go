package test

import (
	"encoding/json"
	"github.com/shuoros/go-challenge/pkg/handlers"
	"testing"
)

func Test_CreateResponseBodyMustGenerateAJSONableString(t *testing.T) {
	responseInString := handlers.CreateResponseBody(true, 200, "OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(responseInString), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
}

func Test_ErrorResponseMustHaveOkFalse(t *testing.T) {
	responseInString := handlers.CreateResponseBody(false, 200, "OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(responseInString), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["ok"] != false {
		t.Error("Error response must have ok false")
	}
}

func Test_SuccessResponseMustHaveOkTrue(t *testing.T) {
	responseInString := handlers.CreateResponseBody(false, 200, "OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(responseInString), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["ok"] != false {
		t.Error("Error response must have ok false")
	}
}

func Test_SuccessResponseMustHaveMessageBasedOnItsGivenStatusCode(t *testing.T) {
	responseInString := handlers.CreateResponseBody(true, 200, "OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(responseInString), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["message"] != "OK" {
		t.Error("Error response must have ok false")
	}
}

func Test_ErrorResponseMustHaveMessageBasedOnItsGivenStatusCode(t *testing.T) {
	responseInString := handlers.CreateResponseBody(false, 409, "Not OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(responseInString), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["message"] != "Conflict" {
		t.Error("Error response must have ok false")
	}
}

func Test_CreateResponseMustHaveGivenData(t *testing.T) {
	responseInString := handlers.CreateResponseBody(true, 200, "OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(responseInString), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["data"] == nil {
		t.Error("Response must have data")
	}
}
