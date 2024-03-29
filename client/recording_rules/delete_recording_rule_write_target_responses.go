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

// DeleteRecordingRuleWriteTargetReader is a Reader for the DeleteRecordingRuleWriteTarget structure.
type DeleteRecordingRuleWriteTargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecordingRuleWriteTargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRecordingRuleWriteTargetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRecordingRuleWriteTargetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRecordingRuleWriteTargetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRecordingRuleWriteTargetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRecordingRuleWriteTargetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /recording-rules/writer] deleteRecordingRuleWriteTarget", response, response.Code())
	}
}

// NewDeleteRecordingRuleWriteTargetOK creates a DeleteRecordingRuleWriteTargetOK with default headers values
func NewDeleteRecordingRuleWriteTargetOK() *DeleteRecordingRuleWriteTargetOK {
	return &DeleteRecordingRuleWriteTargetOK{}
}

/*
DeleteRecordingRuleWriteTargetOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteRecordingRuleWriteTargetOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete recording rule write target Ok response has a 2xx status code
func (o *DeleteRecordingRuleWriteTargetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete recording rule write target Ok response has a 3xx status code
func (o *DeleteRecordingRuleWriteTargetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule write target Ok response has a 4xx status code
func (o *DeleteRecordingRuleWriteTargetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete recording rule write target Ok response has a 5xx status code
func (o *DeleteRecordingRuleWriteTargetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recording rule write target Ok response a status code equal to that given
func (o *DeleteRecordingRuleWriteTargetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete recording rule write target Ok response
func (o *DeleteRecordingRuleWriteTargetOK) Code() int {
	return 200
}

func (o *DeleteRecordingRuleWriteTargetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetOk %s", 200, payload)
}

func (o *DeleteRecordingRuleWriteTargetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetOk %s", 200, payload)
}

func (o *DeleteRecordingRuleWriteTargetOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleWriteTargetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingRuleWriteTargetUnauthorized creates a DeleteRecordingRuleWriteTargetUnauthorized with default headers values
func NewDeleteRecordingRuleWriteTargetUnauthorized() *DeleteRecordingRuleWriteTargetUnauthorized {
	return &DeleteRecordingRuleWriteTargetUnauthorized{}
}

/*
DeleteRecordingRuleWriteTargetUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteRecordingRuleWriteTargetUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete recording rule write target unauthorized response has a 2xx status code
func (o *DeleteRecordingRuleWriteTargetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recording rule write target unauthorized response has a 3xx status code
func (o *DeleteRecordingRuleWriteTargetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule write target unauthorized response has a 4xx status code
func (o *DeleteRecordingRuleWriteTargetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete recording rule write target unauthorized response has a 5xx status code
func (o *DeleteRecordingRuleWriteTargetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recording rule write target unauthorized response a status code equal to that given
func (o *DeleteRecordingRuleWriteTargetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete recording rule write target unauthorized response
func (o *DeleteRecordingRuleWriteTargetUnauthorized) Code() int {
	return 401
}

func (o *DeleteRecordingRuleWriteTargetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetUnauthorized %s", 401, payload)
}

func (o *DeleteRecordingRuleWriteTargetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetUnauthorized %s", 401, payload)
}

func (o *DeleteRecordingRuleWriteTargetUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleWriteTargetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingRuleWriteTargetForbidden creates a DeleteRecordingRuleWriteTargetForbidden with default headers values
func NewDeleteRecordingRuleWriteTargetForbidden() *DeleteRecordingRuleWriteTargetForbidden {
	return &DeleteRecordingRuleWriteTargetForbidden{}
}

/*
DeleteRecordingRuleWriteTargetForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteRecordingRuleWriteTargetForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete recording rule write target forbidden response has a 2xx status code
func (o *DeleteRecordingRuleWriteTargetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recording rule write target forbidden response has a 3xx status code
func (o *DeleteRecordingRuleWriteTargetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule write target forbidden response has a 4xx status code
func (o *DeleteRecordingRuleWriteTargetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete recording rule write target forbidden response has a 5xx status code
func (o *DeleteRecordingRuleWriteTargetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recording rule write target forbidden response a status code equal to that given
func (o *DeleteRecordingRuleWriteTargetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete recording rule write target forbidden response
func (o *DeleteRecordingRuleWriteTargetForbidden) Code() int {
	return 403
}

func (o *DeleteRecordingRuleWriteTargetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetForbidden %s", 403, payload)
}

func (o *DeleteRecordingRuleWriteTargetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetForbidden %s", 403, payload)
}

func (o *DeleteRecordingRuleWriteTargetForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleWriteTargetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingRuleWriteTargetNotFound creates a DeleteRecordingRuleWriteTargetNotFound with default headers values
func NewDeleteRecordingRuleWriteTargetNotFound() *DeleteRecordingRuleWriteTargetNotFound {
	return &DeleteRecordingRuleWriteTargetNotFound{}
}

/*
DeleteRecordingRuleWriteTargetNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteRecordingRuleWriteTargetNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete recording rule write target not found response has a 2xx status code
func (o *DeleteRecordingRuleWriteTargetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recording rule write target not found response has a 3xx status code
func (o *DeleteRecordingRuleWriteTargetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule write target not found response has a 4xx status code
func (o *DeleteRecordingRuleWriteTargetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete recording rule write target not found response has a 5xx status code
func (o *DeleteRecordingRuleWriteTargetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recording rule write target not found response a status code equal to that given
func (o *DeleteRecordingRuleWriteTargetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete recording rule write target not found response
func (o *DeleteRecordingRuleWriteTargetNotFound) Code() int {
	return 404
}

func (o *DeleteRecordingRuleWriteTargetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetNotFound %s", 404, payload)
}

func (o *DeleteRecordingRuleWriteTargetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetNotFound %s", 404, payload)
}

func (o *DeleteRecordingRuleWriteTargetNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleWriteTargetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingRuleWriteTargetInternalServerError creates a DeleteRecordingRuleWriteTargetInternalServerError with default headers values
func NewDeleteRecordingRuleWriteTargetInternalServerError() *DeleteRecordingRuleWriteTargetInternalServerError {
	return &DeleteRecordingRuleWriteTargetInternalServerError{}
}

/*
DeleteRecordingRuleWriteTargetInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteRecordingRuleWriteTargetInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete recording rule write target internal server error response has a 2xx status code
func (o *DeleteRecordingRuleWriteTargetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete recording rule write target internal server error response has a 3xx status code
func (o *DeleteRecordingRuleWriteTargetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recording rule write target internal server error response has a 4xx status code
func (o *DeleteRecordingRuleWriteTargetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete recording rule write target internal server error response has a 5xx status code
func (o *DeleteRecordingRuleWriteTargetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete recording rule write target internal server error response a status code equal to that given
func (o *DeleteRecordingRuleWriteTargetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete recording rule write target internal server error response
func (o *DeleteRecordingRuleWriteTargetInternalServerError) Code() int {
	return 500
}

func (o *DeleteRecordingRuleWriteTargetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetInternalServerError %s", 500, payload)
}

func (o *DeleteRecordingRuleWriteTargetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /recording-rules/writer][%d] deleteRecordingRuleWriteTargetInternalServerError %s", 500, payload)
}

func (o *DeleteRecordingRuleWriteTargetInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteRecordingRuleWriteTargetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
