package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

// ResponseError
// A general http response error that follows a similar layout to shopline's response errors
type ResponseError struct {
	Status  int
	Code    string
	Message string
	Errors  []string
}

type RateLimitError struct {
	ResponseError
	RetryAfter int
}

// GetStatus returns http response status
func (e ResponseError) GetStatus() int {
	return e.Status
}

// GetMessage returns response error message
func (e ResponseError) GetMessage() string {
	return e.Message
}

// GetErrors returns response errors
func (e ResponseError) GetErrors() []string {
	return e.Errors
}

func (e ResponseError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	sort.Strings(e.Errors)
	s := strings.Join(e.Errors, ", ")

	if s != "" {
		return s
	}

	return "Unknown Error"
}

func CheckHttpResponseError(resp *http.Response) error {
	// 200 <= StatusCode < 300
	if http.StatusOK <= resp.StatusCode && resp.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	shoplineError := struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Errors  interface{} `json:"errors"`
	}{}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// empty body, this probably means shopline returned an error with no response body
	if len(bodyBytes) > 0 {
		err := json.Unmarshal(bodyBytes, &shoplineError)
		if err != nil {
			return ResponseDecodingError{
				Body:    bodyBytes,
				Message: err.Error(),
				Status:  resp.StatusCode,
			}
		}
	}

	// Create the response error from the shopline error.
	responseError := ResponseError{
		Status:  resp.StatusCode,
		Code:    shoplineError.Code,
		Message: shoplineError.Message,
	}

	if shoplineError.Errors == nil {
		return wrapSpecificError(resp, responseError)
	}

	switch reflect.TypeOf(shoplineError.Errors).Kind() {
	case reflect.String:
		// Single string, use as message
		responseError.Message = shoplineError.Errors.(string)
	case reflect.Slice:
		// An array, parse each entry as a string and join them on the message
		// json always serializes JSON arrays into []interface{}
		for _, elem := range shoplineError.Errors.([]interface{}) {
			responseError.Errors = append(responseError.Errors, fmt.Sprint(elem))
		}
		responseError.Message = strings.Join(responseError.Errors, ", ")
	case reflect.Map:
		// A map, parse each error for each key in the map.
		// json always serializes into map[string]interface{} for objects
		for k, v := range shoplineError.Errors.(map[string]interface{}) {
			switch reflect.TypeOf(v).Kind() {
			// Check to make sure the interface is a slice
			// json always serializes JSON arrays into []interface{}
			case reflect.Slice:
				for _, elem := range v.([]interface{}) {
					// If the primary message of the response error is not set, use
					// any message.
					if responseError.Message == "" {
						responseError.Message = fmt.Sprintf("%v: %v", k, elem)
					}
					topicAndElem := fmt.Sprintf("%v: %v", k, elem)
					responseError.Errors = append(responseError.Errors, topicAndElem)
				}
			case reflect.String:
				elem := v.(string)
				if responseError.Message == "" {
					responseError.Message = fmt.Sprintf("%v: %v", k, elem)
				}
				topicAndElem := fmt.Sprintf("%v: %v", k, elem)
				responseError.Errors = append(responseError.Errors, topicAndElem)
			}
		}
	}

	return responseError
}

func wrapSpecificError(r *http.Response, err ResponseError) error {

	if err.Status == http.StatusNotAcceptable {
		err.Message = http.StatusText(err.Status)
	}

	return err
}
