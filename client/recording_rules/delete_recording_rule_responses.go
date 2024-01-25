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

// DeleteRecordingRuleReader is a Reader for the DeleteRecordingRule structure.
type DeleteRecordingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecordingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRecordingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRecordingRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRecordingRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRecordingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRecordingRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /recording-rules/{recordingRuleID}] deleteRecordingRule", response, response.Code())
	}
}

// NewDeleteRecordingRuleOK creates a DeleteRecordingRuleOK with default headers values
func NewDeleteRecordingRuleOK() *DeleteRecordingRuleOK {
	return &DeleteRecordingRuleOK{}
}

/*
DeleteRecordingRuleOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteRecordingRuleOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete recording rule Ok response has a 2xx status code
func (o *DeleteRecordingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete recording rule Ok response has a 3xx status code
func (o *DeleteRecordingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule Ok response has a 4xx status code
func (o *DeleteRecordingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete recording rule Ok response has a 5xx status code
func (o *DeleteRecordingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recording rule Ok response a status code equal to that given
func (o *DeleteRecordingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete recording rule Ok response
func (o *DeleteRecordingRuleOK) Code() int {
	return 200
}

func (o *DeleteRecordingRuleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleOk %s", 200, payload)
}

func (o *DeleteRecordingRuleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleOk %s", 200, payload)
}

func (o *DeleteRecordingRuleOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingRuleUnauthorized creates a DeleteRecordingRuleUnauthorized with default headers values
func NewDeleteRecordingRuleUnauthorized() *DeleteRecordingRuleUnauthorized {
	return &DeleteRecordingRuleUnauthorized{}
}

/*
DeleteRecordingRuleUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteRecordingRuleUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete recording rule unauthorized response has a 2xx status code
func (o *DeleteRecordingRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recording rule unauthorized response has a 3xx status code
func (o *DeleteRecordingRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule unauthorized response has a 4xx status code
func (o *DeleteRecordingRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete recording rule unauthorized response has a 5xx status code
func (o *DeleteRecordingRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recording rule unauthorized response a status code equal to that given
func (o *DeleteRecordingRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete recording rule unauthorized response
func (o *DeleteRecordingRuleUnauthorized) Code() int {
	return 401
}

func (o *DeleteRecordingRuleUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleUnauthorized %s", 401, payload)
}

func (o *DeleteRecordingRuleUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleUnauthorized %s", 401, payload)
}

func (o *DeleteRecordingRuleUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingRuleForbidden creates a DeleteRecordingRuleForbidden with default headers values
func NewDeleteRecordingRuleForbidden() *DeleteRecordingRuleForbidden {
	return &DeleteRecordingRuleForbidden{}
}

/*
DeleteRecordingRuleForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteRecordingRuleForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete recording rule forbidden response has a 2xx status code
func (o *DeleteRecordingRuleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recording rule forbidden response has a 3xx status code
func (o *DeleteRecordingRuleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule forbidden response has a 4xx status code
func (o *DeleteRecordingRuleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete recording rule forbidden response has a 5xx status code
func (o *DeleteRecordingRuleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recording rule forbidden response a status code equal to that given
func (o *DeleteRecordingRuleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete recording rule forbidden response
func (o *DeleteRecordingRuleForbidden) Code() int {
	return 403
}

func (o *DeleteRecordingRuleForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleForbidden %s", 403, payload)
}

func (o *DeleteRecordingRuleForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleForbidden %s", 403, payload)
}

func (o *DeleteRecordingRuleForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingRuleNotFound creates a DeleteRecordingRuleNotFound with default headers values
func NewDeleteRecordingRuleNotFound() *DeleteRecordingRuleNotFound {
	return &DeleteRecordingRuleNotFound{}
}

/*
DeleteRecordingRuleNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteRecordingRuleNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete recording rule not found response has a 2xx status code
func (o *DeleteRecordingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recording rule not found response has a 3xx status code
func (o *DeleteRecordingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule not found response has a 4xx status code
func (o *DeleteRecordingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete recording rule not found response has a 5xx status code
func (o *DeleteRecordingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recording rule not found response a status code equal to that given
func (o *DeleteRecordingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete recording rule not found response
func (o *DeleteRecordingRuleNotFound) Code() int {
	return 404
}

func (o *DeleteRecordingRuleNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleNotFound %s", 404, payload)
}

func (o *DeleteRecordingRuleNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleNotFound %s", 404, payload)
}

func (o *DeleteRecordingRuleNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingRuleInternalServerError creates a DeleteRecordingRuleInternalServerError with default headers values
func NewDeleteRecordingRuleInternalServerError() *DeleteRecordingRuleInternalServerError {
	return &DeleteRecordingRuleInternalServerError{}
}

/*
DeleteRecordingRuleInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteRecordingRuleInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete recording rule internal server error response has a 2xx status code
func (o *DeleteRecordingRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recording rule internal server error response has a 3xx status code
func (o *DeleteRecordingRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule internal server error response has a 4xx status code
func (o *DeleteRecordingRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete recording rule internal server error response has a 5xx status code
func (o *DeleteRecordingRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete recording rule internal server error response a status code equal to that given
func (o *DeleteRecordingRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete recording rule internal server error response
func (o *DeleteRecordingRuleInternalServerError) Code() int {
	return 500
}

func (o *DeleteRecordingRuleInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleInternalServerError %s", 500, payload)
}

func (o *DeleteRecordingRuleInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/{recordingRuleID}][%d] deleteRecordingRuleInternalServerError %s", 500, payload)
}

func (o *DeleteRecordingRuleInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
