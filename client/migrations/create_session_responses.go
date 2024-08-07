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

// CreateSessionReader is a Reader for the CreateSession structure.
type CreateSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloudmigration/migration] createSession", response, response.Code())
	}
}

// NewCreateSessionOK creates a CreateSessionOK with default headers values
func NewCreateSessionOK() *CreateSessionOK {
	return &CreateSessionOK{}
}

/*
CreateSessionOK describes a response with status code 200, with default header values.

(empty)
*/
type CreateSessionOK struct {
	Payload *models.CloudMigrationSessionResponseDTO
}

// IsSuccess returns true when this create session Ok response has a 2xx status code
func (o *CreateSessionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create session Ok response has a 3xx status code
func (o *CreateSessionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create session Ok response has a 4xx status code
func (o *CreateSessionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create session Ok response has a 5xx status code
func (o *CreateSessionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create session Ok response a status code equal to that given
func (o *CreateSessionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create session Ok response
func (o *CreateSessionOK) Code() int {
	return 200
}

func (o *CreateSessionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionOk %s", 200, payload)
}

func (o *CreateSessionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionOk %s", 200, payload)
}

func (o *CreateSessionOK) GetPayload() *models.CloudMigrationSessionResponseDTO {
	return o.Payload
}

func (o *CreateSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudMigrationSessionResponseDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSessionBadRequest creates a CreateSessionBadRequest with default headers values
func NewCreateSessionBadRequest() *CreateSessionBadRequest {
	return &CreateSessionBadRequest{}
}

/*
CreateSessionBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CreateSessionBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create session bad request response has a 2xx status code
func (o *CreateSessionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create session bad request response has a 3xx status code
func (o *CreateSessionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create session bad request response has a 4xx status code
func (o *CreateSessionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create session bad request response has a 5xx status code
func (o *CreateSessionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create session bad request response a status code equal to that given
func (o *CreateSessionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create session bad request response
func (o *CreateSessionBadRequest) Code() int {
	return 400
}

func (o *CreateSessionBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionBadRequest %s", 400, payload)
}

func (o *CreateSessionBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionBadRequest %s", 400, payload)
}

func (o *CreateSessionBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSessionUnauthorized creates a CreateSessionUnauthorized with default headers values
func NewCreateSessionUnauthorized() *CreateSessionUnauthorized {
	return &CreateSessionUnauthorized{}
}

/*
CreateSessionUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreateSessionUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create session unauthorized response has a 2xx status code
func (o *CreateSessionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create session unauthorized response has a 3xx status code
func (o *CreateSessionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create session unauthorized response has a 4xx status code
func (o *CreateSessionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create session unauthorized response has a 5xx status code
func (o *CreateSessionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create session unauthorized response a status code equal to that given
func (o *CreateSessionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create session unauthorized response
func (o *CreateSessionUnauthorized) Code() int {
	return 401
}

func (o *CreateSessionUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionUnauthorized %s", 401, payload)
}

func (o *CreateSessionUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionUnauthorized %s", 401, payload)
}

func (o *CreateSessionUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSessionForbidden creates a CreateSessionForbidden with default headers values
func NewCreateSessionForbidden() *CreateSessionForbidden {
	return &CreateSessionForbidden{}
}

/*
CreateSessionForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateSessionForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create session forbidden response has a 2xx status code
func (o *CreateSessionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create session forbidden response has a 3xx status code
func (o *CreateSessionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create session forbidden response has a 4xx status code
func (o *CreateSessionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create session forbidden response has a 5xx status code
func (o *CreateSessionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create session forbidden response a status code equal to that given
func (o *CreateSessionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create session forbidden response
func (o *CreateSessionForbidden) Code() int {
	return 403
}

func (o *CreateSessionForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionForbidden %s", 403, payload)
}

func (o *CreateSessionForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionForbidden %s", 403, payload)
}

func (o *CreateSessionForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSessionInternalServerError creates a CreateSessionInternalServerError with default headers values
func NewCreateSessionInternalServerError() *CreateSessionInternalServerError {
	return &CreateSessionInternalServerError{}
}

/*
CreateSessionInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateSessionInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create session internal server error response has a 2xx status code
func (o *CreateSessionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create session internal server error response has a 3xx status code
func (o *CreateSessionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create session internal server error response has a 4xx status code
func (o *CreateSessionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create session internal server error response has a 5xx status code
func (o *CreateSessionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create session internal server error response a status code equal to that given
func (o *CreateSessionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create session internal server error response
func (o *CreateSessionInternalServerError) Code() int {
	return 500
}

func (o *CreateSessionInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionInternalServerError %s", 500, payload)
}

func (o *CreateSessionInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration][%d] createSessionInternalServerError %s", 500, payload)
}

func (o *CreateSessionInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
