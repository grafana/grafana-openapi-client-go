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

// GetMigrationListReader is a Reader for the GetMigrationList structure.
type GetMigrationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMigrationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMigrationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMigrationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMigrationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMigrationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloudmigration/migration] getMigrationList", response, response.Code())
	}
}

// NewGetMigrationListOK creates a GetMigrationListOK with default headers values
func NewGetMigrationListOK() *GetMigrationListOK {
	return &GetMigrationListOK{}
}

/*
GetMigrationListOK describes a response with status code 200, with default header values.

(empty)
*/
type GetMigrationListOK struct {
	Payload *models.CloudMigrationListResponse
}

// IsSuccess returns true when this get migration list Ok response has a 2xx status code
func (o *GetMigrationListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get migration list Ok response has a 3xx status code
func (o *GetMigrationListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration list Ok response has a 4xx status code
func (o *GetMigrationListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get migration list Ok response has a 5xx status code
func (o *GetMigrationListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration list Ok response a status code equal to that given
func (o *GetMigrationListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get migration list Ok response
func (o *GetMigrationListOK) Code() int {
	return 200
}

func (o *GetMigrationListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getMigrationListOk %s", 200, payload)
}

func (o *GetMigrationListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getMigrationListOk %s", 200, payload)
}

func (o *GetMigrationListOK) GetPayload() *models.CloudMigrationListResponse {
	return o.Payload
}

func (o *GetMigrationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudMigrationListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationListUnauthorized creates a GetMigrationListUnauthorized with default headers values
func NewGetMigrationListUnauthorized() *GetMigrationListUnauthorized {
	return &GetMigrationListUnauthorized{}
}

/*
GetMigrationListUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetMigrationListUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get migration list unauthorized response has a 2xx status code
func (o *GetMigrationListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration list unauthorized response has a 3xx status code
func (o *GetMigrationListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration list unauthorized response has a 4xx status code
func (o *GetMigrationListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration list unauthorized response has a 5xx status code
func (o *GetMigrationListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration list unauthorized response a status code equal to that given
func (o *GetMigrationListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get migration list unauthorized response
func (o *GetMigrationListUnauthorized) Code() int {
	return 401
}

func (o *GetMigrationListUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getMigrationListUnauthorized %s", 401, payload)
}

func (o *GetMigrationListUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getMigrationListUnauthorized %s", 401, payload)
}

func (o *GetMigrationListUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetMigrationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationListForbidden creates a GetMigrationListForbidden with default headers values
func NewGetMigrationListForbidden() *GetMigrationListForbidden {
	return &GetMigrationListForbidden{}
}

/*
GetMigrationListForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetMigrationListForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get migration list forbidden response has a 2xx status code
func (o *GetMigrationListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration list forbidden response has a 3xx status code
func (o *GetMigrationListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration list forbidden response has a 4xx status code
func (o *GetMigrationListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration list forbidden response has a 5xx status code
func (o *GetMigrationListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration list forbidden response a status code equal to that given
func (o *GetMigrationListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get migration list forbidden response
func (o *GetMigrationListForbidden) Code() int {
	return 403
}

func (o *GetMigrationListForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getMigrationListForbidden %s", 403, payload)
}

func (o *GetMigrationListForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getMigrationListForbidden %s", 403, payload)
}

func (o *GetMigrationListForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetMigrationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationListInternalServerError creates a GetMigrationListInternalServerError with default headers values
func NewGetMigrationListInternalServerError() *GetMigrationListInternalServerError {
	return &GetMigrationListInternalServerError{}
}

/*
GetMigrationListInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetMigrationListInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get migration list internal server error response has a 2xx status code
func (o *GetMigrationListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration list internal server error response has a 3xx status code
func (o *GetMigrationListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration list internal server error response has a 4xx status code
func (o *GetMigrationListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get migration list internal server error response has a 5xx status code
func (o *GetMigrationListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get migration list internal server error response a status code equal to that given
func (o *GetMigrationListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get migration list internal server error response
func (o *GetMigrationListInternalServerError) Code() int {
	return 500
}

func (o *GetMigrationListInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getMigrationListInternalServerError %s", 500, payload)
}

func (o *GetMigrationListInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration][%d] getMigrationListInternalServerError %s", 500, payload)
}

func (o *GetMigrationListInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetMigrationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
