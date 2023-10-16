// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetDataSourceIDByNameReader is a Reader for the GetDataSourceIDByName structure.
type GetDataSourceIDByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataSourceIDByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataSourceIDByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDataSourceIDByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDataSourceIDByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDataSourceIDByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDataSourceIDByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /datasources/id/{name}] getDataSourceIdByName", response, response.Code())
	}
}

// NewGetDataSourceIDByNameOK creates a GetDataSourceIDByNameOK with default headers values
func NewGetDataSourceIDByNameOK() *GetDataSourceIDByNameOK {
	return &GetDataSourceIDByNameOK{}
}

/*
GetDataSourceIDByNameOK describes a response with status code 200, with default header values.

(empty)
*/
type GetDataSourceIDByNameOK struct {
	Payload *models.GetDataSourceIDByNameOKBody
}

// IsSuccess returns true when this get data source Id by name o k response has a 2xx status code
func (o *GetDataSourceIDByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get data source Id by name o k response has a 3xx status code
func (o *GetDataSourceIDByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source Id by name o k response has a 4xx status code
func (o *GetDataSourceIDByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data source Id by name o k response has a 5xx status code
func (o *GetDataSourceIDByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source Id by name o k response a status code equal to that given
func (o *GetDataSourceIDByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get data source Id by name o k response
func (o *GetDataSourceIDByNameOK) Code() int {
	return 200
}

func (o *GetDataSourceIDByNameOK) Error() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameOK  %+v", 200, o.Payload)
}

func (o *GetDataSourceIDByNameOK) String() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameOK  %+v", 200, o.Payload)
}

func (o *GetDataSourceIDByNameOK) GetPayload() *models.GetDataSourceIDByNameOKBody {
	return o.Payload
}

func (o *GetDataSourceIDByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDataSourceIDByNameOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceIDByNameUnauthorized creates a GetDataSourceIDByNameUnauthorized with default headers values
func NewGetDataSourceIDByNameUnauthorized() *GetDataSourceIDByNameUnauthorized {
	return &GetDataSourceIDByNameUnauthorized{}
}

/*
GetDataSourceIDByNameUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetDataSourceIDByNameUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get data source Id by name unauthorized response has a 2xx status code
func (o *GetDataSourceIDByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data source Id by name unauthorized response has a 3xx status code
func (o *GetDataSourceIDByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source Id by name unauthorized response has a 4xx status code
func (o *GetDataSourceIDByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data source Id by name unauthorized response has a 5xx status code
func (o *GetDataSourceIDByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source Id by name unauthorized response a status code equal to that given
func (o *GetDataSourceIDByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get data source Id by name unauthorized response
func (o *GetDataSourceIDByNameUnauthorized) Code() int {
	return 401
}

func (o *GetDataSourceIDByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDataSourceIDByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDataSourceIDByNameUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDataSourceIDByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceIDByNameForbidden creates a GetDataSourceIDByNameForbidden with default headers values
func NewGetDataSourceIDByNameForbidden() *GetDataSourceIDByNameForbidden {
	return &GetDataSourceIDByNameForbidden{}
}

/*
GetDataSourceIDByNameForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetDataSourceIDByNameForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get data source Id by name forbidden response has a 2xx status code
func (o *GetDataSourceIDByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data source Id by name forbidden response has a 3xx status code
func (o *GetDataSourceIDByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source Id by name forbidden response has a 4xx status code
func (o *GetDataSourceIDByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data source Id by name forbidden response has a 5xx status code
func (o *GetDataSourceIDByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source Id by name forbidden response a status code equal to that given
func (o *GetDataSourceIDByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get data source Id by name forbidden response
func (o *GetDataSourceIDByNameForbidden) Code() int {
	return 403
}

func (o *GetDataSourceIDByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetDataSourceIDByNameForbidden) String() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameForbidden  %+v", 403, o.Payload)
}

func (o *GetDataSourceIDByNameForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDataSourceIDByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceIDByNameNotFound creates a GetDataSourceIDByNameNotFound with default headers values
func NewGetDataSourceIDByNameNotFound() *GetDataSourceIDByNameNotFound {
	return &GetDataSourceIDByNameNotFound{}
}

/*
GetDataSourceIDByNameNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetDataSourceIDByNameNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get data source Id by name not found response has a 2xx status code
func (o *GetDataSourceIDByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data source Id by name not found response has a 3xx status code
func (o *GetDataSourceIDByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source Id by name not found response has a 4xx status code
func (o *GetDataSourceIDByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data source Id by name not found response has a 5xx status code
func (o *GetDataSourceIDByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source Id by name not found response a status code equal to that given
func (o *GetDataSourceIDByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get data source Id by name not found response
func (o *GetDataSourceIDByNameNotFound) Code() int {
	return 404
}

func (o *GetDataSourceIDByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetDataSourceIDByNameNotFound) String() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetDataSourceIDByNameNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDataSourceIDByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceIDByNameInternalServerError creates a GetDataSourceIDByNameInternalServerError with default headers values
func NewGetDataSourceIDByNameInternalServerError() *GetDataSourceIDByNameInternalServerError {
	return &GetDataSourceIDByNameInternalServerError{}
}

/*
GetDataSourceIDByNameInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetDataSourceIDByNameInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get data source Id by name internal server error response has a 2xx status code
func (o *GetDataSourceIDByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data source Id by name internal server error response has a 3xx status code
func (o *GetDataSourceIDByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source Id by name internal server error response has a 4xx status code
func (o *GetDataSourceIDByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data source Id by name internal server error response has a 5xx status code
func (o *GetDataSourceIDByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get data source Id by name internal server error response a status code equal to that given
func (o *GetDataSourceIDByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get data source Id by name internal server error response
func (o *GetDataSourceIDByNameInternalServerError) Code() int {
	return 500
}

func (o *GetDataSourceIDByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDataSourceIDByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /datasources/id/{name}][%d] getDataSourceIdByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDataSourceIDByNameInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDataSourceIDByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}