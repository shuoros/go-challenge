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

func Test_SuccessResponseMustHaveOkTrue(t *testing.T) {
	resp, _ := handlers.SuccessResponse(200, "OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(resp.Body), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["ok"] != true {
		t.Errorf("got %q, wanted %t", jsonMap["ok"], true)
	}
}

func Test_ErrorResponseMustHaveOkFalse(t *testing.T) {
	resp, _ := handlers.ErrorResponse(409, "Not OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(resp.Body), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["ok"] != false {
		t.Errorf("got %q, wanted %t", jsonMap["ok"], false)
	}
}

func Test_SuccessResponseMustHaveMessageBasedOnItsGivenStatusCode(t *testing.T) {
	resp, _ := handlers.SuccessResponse(200, "OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(resp.Body), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["message"] != "OK" {
		t.Errorf("got %q, wanted %q", jsonMap["message"], "OK")
	}
}

func Test_ErrorResponseMustHaveMessageBasedOnItsGivenStatusCode(t *testing.T) {
	resp, _ := handlers.ErrorResponse(409, "Not OK")
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(resp.Body), &jsonMap); err != nil {
		t.Error("Response body is not a valid JSON")
	}
	if jsonMap["message"] != "Conflict" {
		t.Errorf("got %q, wanted %q", jsonMap["message"], "Conflict")
	}
}
