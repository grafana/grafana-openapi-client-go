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

// GetSessionListReader is a Reader for the GetSessionList structure.
type GetSessionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSessionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSessionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSessionListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSessionListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSessionListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloudmigration/migration] getSessionList", response, response.Code())
	}
}

// NewGetSessionListOK creates a GetSessionListOK with default headers values
func NewGetSessionListOK() *GetSessionListOK {
	return &GetSessionListOK{}
}

/*
GetSessionListOK describes a response with status code 200, with default header values.

(empty)
*/
type GetSessionListOK struct {
	Payload *models.CloudMigrationSessionListResponseDTO
}

// IsSuccess returns true when this get session list Ok response has a 2xx status code
func (o *GetSessionListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get session list Ok response has a 3xx status code
func (o *GetSessionListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session list Ok response has a 4xx status code
func (o *GetSessionListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get session list Ok response has a 5xx status code
func (o *GetSessionListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get session list Ok response a status code equal to that given
func (o *GetSessionListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get session list Ok response
func (o *GetSessionListOK) Code() int {
	return 200
}

func (o *GetSessionListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getSessionListOk %s", 200, payload)
}

func (o *GetSessionListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getSessionListOk %s", 200, payload)
}

func (o *GetSessionListOK) GetPayload() *models.CloudMigrationSessionListResponseDTO {
	return o.Payload
}

func (o *GetSessionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudMigrationSessionListResponseDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionListUnauthorized creates a GetSessionListUnauthorized with default headers values
func NewGetSessionListUnauthorized() *GetSessionListUnauthorized {
	return &GetSessionListUnauthorized{}
}

/*
GetSessionListUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetSessionListUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get session list unauthorized response has a 2xx status code
func (o *GetSessionListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session list unauthorized response has a 3xx status code
func (o *GetSessionListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session list unauthorized response has a 4xx status code
func (o *GetSessionListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get session list unauthorized response has a 5xx status code
func (o *GetSessionListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get session list unauthorized response a status code equal to that given
func (o *GetSessionListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get session list unauthorized response
func (o *GetSessionListUnauthorized) Code() int {
	return 401
}

func (o *GetSessionListUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getSessionListUnauthorized %s", 401, payload)
}

func (o *GetSessionListUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getSessionListUnauthorized %s", 401, payload)
}

func (o *GetSessionListUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSessionListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionListForbidden creates a GetSessionListForbidden with default headers values
func NewGetSessionListForbidden() *GetSessionListForbidden {
	return &GetSessionListForbidden{}
}

/*
GetSessionListForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetSessionListForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get session list forbidden response has a 2xx status code
func (o *GetSessionListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session list forbidden response has a 3xx status code
func (o *GetSessionListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session list forbidden response has a 4xx status code
func (o *GetSessionListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get session list forbidden response has a 5xx status code
func (o *GetSessionListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get session list forbidden response a status code equal to that given
func (o *GetSessionListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get session list forbidden response
func (o *GetSessionListForbidden) Code() int {
	return 403
}

func (o *GetSessionListForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getSessionListForbidden %s", 403, payload)
}

func (o *GetSessionListForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getSessionListForbidden %s", 403, payload)
}

func (o *GetSessionListForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSessionListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionListInternalServerError creates a GetSessionListInternalServerError with default headers values
func NewGetSessionListInternalServerError() *GetSessionListInternalServerError {
	return &GetSessionListInternalServerError{}
}

/*
GetSessionListInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetSessionListInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get session list internal server error response has a 2xx status code
func (o *GetSessionListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session list internal server error response has a 3xx status code
func (o *GetSessionListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session list internal server error response has a 4xx status code
func (o *GetSessionListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get session list internal server error response has a 5xx status code
func (o *GetSessionListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get session list internal server error response a status code equal to that given
func (o *GetSessionListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get session list internal server error response
func (o *GetSessionListInternalServerError) Code() int {
	return 500
}

func (o *GetSessionListInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getSessionListInternalServerError %s", 500, payload)
}

func (o *GetSessionListInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getSessionListInternalServerError %s", 500, payload)
}

func (o *GetSessionListInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSessionListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
