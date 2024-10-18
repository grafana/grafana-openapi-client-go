// Code generated by go-swagger; DO NOT EDIT.

package migrations

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

// CancelSnapshotReader is a Reader for the CancelSnapshot structure.
type CancelSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelSnapshotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel] cancelSnapshot", response, response.Code())
	}
}

// NewCancelSnapshotOK creates a CancelSnapshotOK with default headers values
func NewCancelSnapshotOK() *CancelSnapshotOK {
	return &CancelSnapshotOK{}
}

/*
CancelSnapshotOK describes a response with status code 200, with default header values.

(empty)
*/
type CancelSnapshotOK struct {
}

// IsSuccess returns true when this cancel snapshot Ok response has a 2xx status code
func (o *CancelSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel snapshot Ok response has a 3xx status code
func (o *CancelSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel snapshot Ok response has a 4xx status code
func (o *CancelSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel snapshot Ok response has a 5xx status code
func (o *CancelSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel snapshot Ok response a status code equal to that given
func (o *CancelSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cancel snapshot Ok response
func (o *CancelSnapshotOK) Code() int {
	return 200
}

func (o *CancelSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotOk", 200)
}

func (o *CancelSnapshotOK) String() string {
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotOk", 200)
}

func (o *CancelSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelSnapshotBadRequest creates a CancelSnapshotBadRequest with default headers values
func NewCancelSnapshotBadRequest() *CancelSnapshotBadRequest {
	return &CancelSnapshotBadRequest{}
}

/*
CancelSnapshotBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CancelSnapshotBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this cancel snapshot bad request response has a 2xx status code
func (o *CancelSnapshotBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel snapshot bad request response has a 3xx status code
func (o *CancelSnapshotBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel snapshot bad request response has a 4xx status code
func (o *CancelSnapshotBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel snapshot bad request response has a 5xx status code
func (o *CancelSnapshotBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel snapshot bad request response a status code equal to that given
func (o *CancelSnapshotBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cancel snapshot bad request response
func (o *CancelSnapshotBadRequest) Code() int {
	return 400
}

func (o *CancelSnapshotBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotBadRequest %s", 400, payload)
}

func (o *CancelSnapshotBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotBadRequest %s", 400, payload)
}

func (o *CancelSnapshotBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CancelSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSnapshotUnauthorized creates a CancelSnapshotUnauthorized with default headers values
func NewCancelSnapshotUnauthorized() *CancelSnapshotUnauthorized {
	return &CancelSnapshotUnauthorized{}
}

/*
CancelSnapshotUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CancelSnapshotUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this cancel snapshot unauthorized response has a 2xx status code
func (o *CancelSnapshotUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel snapshot unauthorized response has a 3xx status code
func (o *CancelSnapshotUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel snapshot unauthorized response has a 4xx status code
func (o *CancelSnapshotUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel snapshot unauthorized response has a 5xx status code
func (o *CancelSnapshotUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel snapshot unauthorized response a status code equal to that given
func (o *CancelSnapshotUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cancel snapshot unauthorized response
func (o *CancelSnapshotUnauthorized) Code() int {
	return 401
}

func (o *CancelSnapshotUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotUnauthorized %s", 401, payload)
}

func (o *CancelSnapshotUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotUnauthorized %s", 401, payload)
}

func (o *CancelSnapshotUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CancelSnapshotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSnapshotForbidden creates a CancelSnapshotForbidden with default headers values
func NewCancelSnapshotForbidden() *CancelSnapshotForbidden {
	return &CancelSnapshotForbidden{}
}

/*
CancelSnapshotForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CancelSnapshotForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this cancel snapshot forbidden response has a 2xx status code
func (o *CancelSnapshotForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel snapshot forbidden response has a 3xx status code
func (o *CancelSnapshotForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel snapshot forbidden response has a 4xx status code
func (o *CancelSnapshotForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel snapshot forbidden response has a 5xx status code
func (o *CancelSnapshotForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel snapshot forbidden response a status code equal to that given
func (o *CancelSnapshotForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cancel snapshot forbidden response
func (o *CancelSnapshotForbidden) Code() int {
	return 403
}

func (o *CancelSnapshotForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotForbidden %s", 403, payload)
}

func (o *CancelSnapshotForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotForbidden %s", 403, payload)
}

func (o *CancelSnapshotForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CancelSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSnapshotInternalServerError creates a CancelSnapshotInternalServerError with default headers values
func NewCancelSnapshotInternalServerError() *CancelSnapshotInternalServerError {
	return &CancelSnapshotInternalServerError{}
}

/*
CancelSnapshotInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CancelSnapshotInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this cancel snapshot internal server error response has a 2xx status code
func (o *CancelSnapshotInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel snapshot internal server error response has a 3xx status code
func (o *CancelSnapshotInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel snapshot internal server error response has a 4xx status code
func (o *CancelSnapshotInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel snapshot internal server error response has a 5xx status code
func (o *CancelSnapshotInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel snapshot internal server error response a status code equal to that given
func (o *CancelSnapshotInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cancel snapshot internal server error response
func (o *CancelSnapshotInternalServerError) Code() int {
	return 500
}

func (o *CancelSnapshotInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotInternalServerError %s", 500, payload)
}

func (o *CancelSnapshotInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/snapshot/{snapshotUid}/cancel][%d] cancelSnapshotInternalServerError %s", 500, payload)
}

func (o *CancelSnapshotInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CancelSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}