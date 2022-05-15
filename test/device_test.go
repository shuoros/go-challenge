package test

import (
	"github.com/shuoros/go-challenge/pkg/device"
	"testing"
)

func Test_ValidateFieldsMustReturnFalseIfDetectMissingFieldsInDeviceObject(t *testing.T) {
	d := device.Device{}
	_, err := device.ValidateFields(d)
	if err != true {
		t.Error("validate fields must detect missing fields in device object")
	}
}

func Test_ValidateFieldsMustDetectAllFieldsMissingInDeviceObject(t *testing.T) {
	d := device.Device{}
	msg, _ := device.ValidateFields(d)
	expect := "422-Following fields are not provided: id, deviceModel, name, note, serial"
	if msg != expect {
		t.Errorf("got %q, wanted %q", msg, expect)
	}
}

func Test_ValidateFieldsMustDetectSomeFieldsMissingInDeviceObject(t *testing.T) {
	d := device.Device{
		ID:   "1",
		Name: "1",
		Note: "1",
	}
	msg, _ := device.ValidateFields(d)
	expect := "422-Following fields are not provided: deviceModel, serial"
	if msg != expect {
		t.Errorf("got %q, wanted %q", msg, expect)
	}
}

func Test_ValidateFieldsMustDetectOneFieldsMissingInDeviceObject(t *testing.T) {
	d := device.Device{
		ID:     "1",
		Name:   "1",
		Note:   "1",
		Serial: "1",
	}
	msg, _ := device.ValidateFields(d)
	expect := "422-Following fields are not provided: deviceModel"
	if msg != expect {
		t.Errorf("got %q, wanted %q", msg, expect)
	}
}
