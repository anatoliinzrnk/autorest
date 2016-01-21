package autorest

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"testing"
)

func TestNewErrorWithError_AssignsPackageType(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", nil, "message")

	if e.PackageType() != "packageType" {
		t.Errorf("autorest: Error failed to set package type -- expected %v, received %v", "packageType", e.PackageType())
	}
}

func TestNewErrorWithError_AssignsMethod(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", nil, "message")

	if e.Method() != "method" {
		t.Errorf("autorest: Error failed to set method -- expected %v, received %v", "method", e.Method())
	}
}

func TestNewErrorWithError_AssignsMessage(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", nil, "message")

	if e.Message() != "message" {
		t.Errorf("autorest: Error failed to set message -- expected %v, received %v", "message", e.Message())
	}
}

func TestNewErrorWithError_AssignsUndefinedStatusCodeIfRespNil(t *testing.T) {
	e := NewErrorWithError(nil, "packageType", "method", nil, "message")
	if e.StatusCode() != UndefinedStatusCode {
		t.Errorf("autorest: Error failed to set status code -- expected %v, received %v", UndefinedStatusCode, e.StatusCode())
	}
}

func TestNewErrorWithError_AssignsStatusCode(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", &http.Response{
		StatusCode: http.StatusBadRequest,
		Status:     http.StatusText(http.StatusBadRequest)}, "message")

	if e.StatusCode() != http.StatusBadRequest {
		t.Errorf("autorest: Error failed to set status code -- expected %v, received %v", http.StatusBadRequest, e.StatusCode())
	}
}

func TestNewErrorWithError_AcceptsArgs(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", nil, "message %s", "arg")

	if matched, _ := regexp.MatchString(`.*arg.*`, e.Message()); !matched {
		t.Errorf("autorest: Error failed to apply message arguments -- expected %v, received %v",
			`.*arg.*`, e.Message())
	}
}

func TestNewErrorWithError_AssignsError(t *testing.T) {
	err := fmt.Errorf("original")
	e := NewErrorWithError(err, "packageType", "method", nil, "message")

	if e.Original() != err {
		t.Errorf("autorest: Error failed to set error -- expected %v, received %v", err, e.Original())
	}
}

func TestNewErrorWithResponse_ContainsStatusCode(t *testing.T) {
	e := NewErrorWithResponse("packageType", "method", &http.Response{
		StatusCode: http.StatusBadRequest,
		Status:     http.StatusText(http.StatusBadRequest)}, "message")

	if e.StatusCode() != http.StatusBadRequest {
		t.Errorf("autorest: Error failed to set status code -- expected %v, received %v", http.StatusBadRequest, e.StatusCode())
	}
}

func TestNewErrorWithResponse_nilResponse_ReportsUndefinedStatusCode(t *testing.T) {
	e := NewErrorWithResponse("packageType", "method", nil, "message")

	if e.StatusCode() != UndefinedStatusCode {
		t.Errorf("autorest: Error failed to set status code -- expected %v, received %v", UndefinedStatusCode, e.StatusCode())
	}
}

func TestNewErrorForwards(t *testing.T) {
	e1 := NewError("packageType", "method", "message %s", "arg")
	e2 := NewErrorWithError(nil, "packageType", "method", nil, "message %s", "arg")

	if !reflect.DeepEqual(e1, e2) {
		t.Error("autorest: NewError did not return an error equivelent to NewErrorWithError")
	}
}

func TestErrorError(t *testing.T) {
	err := fmt.Errorf("original")
	e := NewErrorWithError(err, "packageType", "method", nil, "message")

	if matched, _ := regexp.MatchString(`.*original.*`, e.Error()); !matched {
		t.Errorf("autorest: Error#Error failed to return original error message -- expected %v, received %v",
			`.*original.*`, e.Error())
	}
}

func TestErrorStringConstainsPackageType(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", nil, "message")

	if matched, _ := regexp.MatchString(`.*packageType.*`, e.String()); !matched {
		t.Errorf("autorest: Error#String failed to include PackageType -- expected %v, received %v",
			`.*packageType.*`, e.String())
	}
}

func TestErrorStringConstainsMethod(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", nil, "message")

	if matched, _ := regexp.MatchString(`.*method.*`, e.String()); !matched {
		t.Errorf("autorest: Error#String failed to include Method -- expected %v, received %v",
			`.*method.*`, e.String())
	}
}

func TestErrorStringConstainsMessage(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", nil, "message")

	if matched, _ := regexp.MatchString(`.*message.*`, e.String()); !matched {
		t.Errorf("autorest: Error#String failed to include Message -- expected %v, received %v",
			`.*message.*`, e.String())
	}
}

func TestErrorStringConstainsStatusCode(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", &http.Response{
		StatusCode: http.StatusBadRequest,
		Status:     http.StatusText(http.StatusBadRequest)}, "message")

	if matched, _ := regexp.MatchString(`.*400.*`, e.String()); !matched {
		t.Errorf("autorest: Error#String failed to include Status Code -- expected %v, received %v",
			`.*400.*`, e.String())
	}
}

func TestErrorStringConstainsOriginal(t *testing.T) {
	e := NewErrorWithError(fmt.Errorf("original"), "packageType", "method", nil, "message")

	if matched, _ := regexp.MatchString(`.*original.*`, e.String()); !matched {
		t.Errorf("autorest: Error#String failed to include Original error -- expected %v, received %v",
			`.*original.*`, e.String())
	}
}

func TestErrorStringSkipsOriginal(t *testing.T) {
	e := NewError("packageType", "method", "message")

	if matched, _ := regexp.MatchString(`.*Original.*`, e.String()); matched {
		t.Errorf("autorest: Error#String included missing Original error -- unexpected %v, received %v",
			`.*Original.*`, e.String())
	}
}
