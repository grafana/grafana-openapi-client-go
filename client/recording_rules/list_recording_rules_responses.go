// Code generated by go-swagger; DO NOT EDIT.

package recording_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// ListRecordingRulesReader is a Reader for the ListRecordingRules structure.
type ListRecordingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRecordingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRecordingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRecordingRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRecordingRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRecordingRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRecordingRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /recording-rules] listRecordingRules", response, response.Code())
	}
}

// NewListRecordingRulesOK creates a ListRecordingRulesOK with default headers values
func NewListRecordingRulesOK() *ListRecordingRulesOK {
	return &ListRecordingRulesOK{}
}

/*
ListRecordingRulesOK describes a response with status code 200, with default header values.

(empty)
*/
type ListRecordingRulesOK struct {
	Payload []*models.RecordingRuleJSON
}

// IsSuccess returns true when this list recording rules Ok response has a 2xx status code
func (o *ListRecordingRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list recording rules Ok response has a 3xx status code
func (o *ListRecordingRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list recording rules Ok response has a 4xx status code
func (o *ListRecordingRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list recording rules Ok response has a 5xx status code
func (o *ListRecordingRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list recording rules Ok response a status code equal to that given
func (o *ListRecordingRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list recording rules Ok response
func (o *ListRecordingRulesOK) Code() int {
	return 200
}

func (o *ListRecordingRulesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesOk %s", 200, payload)
}

func (o *ListRecordingRulesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesOk %s", 200, payload)
}

func (o *ListRecordingRulesOK) GetPayload() []*models.RecordingRuleJSON {
	return o.Payload
}

func (o *ListRecordingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRecordingRulesUnauthorized creates a ListRecordingRulesUnauthorized with default headers values
func NewListRecordingRulesUnauthorized() *ListRecordingRulesUnauthorized {
	return &ListRecordingRulesUnauthorized{}
}

/*
ListRecordingRulesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type ListRecordingRulesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this list recording rules unauthorized response has a 2xx status code
func (o *ListRecordingRulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list recording rules unauthorized response has a 3xx status code
func (o *ListRecordingRulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list recording rules unauthorized response has a 4xx status code
func (o *ListRecordingRulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list recording rules unauthorized response has a 5xx status code
func (o *ListRecordingRulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list recording rules unauthorized response a status code equal to that given
func (o *ListRecordingRulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list recording rules unauthorized response
func (o *ListRecordingRulesUnauthorized) Code() int {
	return 401
}

func (o *ListRecordingRulesUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesUnauthorized %s", 401, payload)
}

func (o *ListRecordingRulesUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesUnauthorized %s", 401, payload)
}

func (o *ListRecordingRulesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListRecordingRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRecordingRulesForbidden creates a ListRecordingRulesForbidden with default headers values
func NewListRecordingRulesForbidden() *ListRecordingRulesForbidden {
	return &ListRecordingRulesForbidden{}
}

/*
ListRecordingRulesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ListRecordingRulesForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this list recording rules forbidden response has a 2xx status code
func (o *ListRecordingRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list recording rules forbidden response has a 3xx status code
func (o *ListRecordingRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list recording rules forbidden response has a 4xx status code
func (o *ListRecordingRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list recording rules forbidden response has a 5xx status code
func (o *ListRecordingRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list recording rules forbidden response a status code equal to that given
func (o *ListRecordingRulesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list recording rules forbidden response
func (o *ListRecordingRulesForbidden) Code() int {
	return 403
}

func (o *ListRecordingRulesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesForbidden %s", 403, payload)
}

func (o *ListRecordingRulesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesForbidden %s", 403, payload)
}

func (o *ListRecordingRulesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListRecordingRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRecordingRulesNotFound creates a ListRecordingRulesNotFound with default headers values
func NewListRecordingRulesNotFound() *ListRecordingRulesNotFound {
	return &ListRecordingRulesNotFound{}
}

/*
ListRecordingRulesNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type ListRecordingRulesNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this list recording rules not found response has a 2xx status code
func (o *ListRecordingRulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list recording rules not found response has a 3xx status code
func (o *ListRecordingRulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list recording rules not found response has a 4xx status code
func (o *ListRecordingRulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list recording rules not found response has a 5xx status code
func (o *ListRecordingRulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list recording rules not found response a status code equal to that given
func (o *ListRecordingRulesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list recording rules not found response
func (o *ListRecordingRulesNotFound) Code() int {
	return 404
}

func (o *ListRecordingRulesNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesNotFound %s", 404, payload)
}

func (o *ListRecordingRulesNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesNotFound %s", 404, payload)
}

func (o *ListRecordingRulesNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListRecordingRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRecordingRulesInternalServerError creates a ListRecordingRulesInternalServerError with default headers values
func NewListRecordingRulesInternalServerError() *ListRecordingRulesInternalServerError {
	return &ListRecordingRulesInternalServerError{}
}

/*
ListRecordingRulesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ListRecordingRulesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this list recording rules internal server error response has a 2xx status code
func (o *ListRecordingRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list recording rules internal server error response has a 3xx status code
func (o *ListRecordingRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list recording rules internal server error response has a 4xx status code
func (o *ListRecordingRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list recording rules internal server error response has a 5xx status code
func (o *ListRecordingRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list recording rules internal server error response a status code equal to that given
func (o *ListRecordingRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list recording rules internal server error response
func (o *ListRecordingRulesInternalServerError) Code() int {
	return 500
}

func (o *ListRecordingRulesInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesInternalServerError %s", 500, payload)
}

func (o *ListRecordingRulesInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /recording-rules][%d] listRecordingRulesInternalServerError %s", 500, payload)
}

func (o *ListRecordingRulesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListRecordingRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
