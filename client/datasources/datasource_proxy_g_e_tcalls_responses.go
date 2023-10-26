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

// DatasourceProxyGETcallsReader is a Reader for the DatasourceProxyGETcalls structure.
type DatasourceProxyGETcallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DatasourceProxyGETcallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDatasourceProxyGETcallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDatasourceProxyGETcallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDatasourceProxyGETcallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDatasourceProxyGETcallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDatasourceProxyGETcallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDatasourceProxyGETcallsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /datasources/proxy/{id}/{datasource_proxy_route}] datasourceProxyGETcalls", response, response.Code())
	}
}

// NewDatasourceProxyGETcallsOK creates a DatasourceProxyGETcallsOK with default headers values
func NewDatasourceProxyGETcallsOK() *DatasourceProxyGETcallsOK {
	return &DatasourceProxyGETcallsOK{}
}

/*
DatasourceProxyGETcallsOK describes a response with status code 200, with default header values.

(empty)
*/
type DatasourceProxyGETcallsOK struct {
}

// IsSuccess returns true when this datasource proxy g e tcalls Ok response has a 2xx status code
func (o *DatasourceProxyGETcallsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this datasource proxy g e tcalls Ok response has a 3xx status code
func (o *DatasourceProxyGETcallsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e tcalls Ok response has a 4xx status code
func (o *DatasourceProxyGETcallsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this datasource proxy g e tcalls Ok response has a 5xx status code
func (o *DatasourceProxyGETcallsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e tcalls Ok response a status code equal to that given
func (o *DatasourceProxyGETcallsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the datasource proxy g e tcalls Ok response
func (o *DatasourceProxyGETcallsOK) Code() int {
	return 200
}

func (o *DatasourceProxyGETcallsOK) Error() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsOk ", 200)
}

func (o *DatasourceProxyGETcallsOK) String() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsOk ", 200)
}

func (o *DatasourceProxyGETcallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDatasourceProxyGETcallsBadRequest creates a DatasourceProxyGETcallsBadRequest with default headers values
func NewDatasourceProxyGETcallsBadRequest() *DatasourceProxyGETcallsBadRequest {
	return &DatasourceProxyGETcallsBadRequest{}
}

/*
DatasourceProxyGETcallsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DatasourceProxyGETcallsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e tcalls bad request response has a 2xx status code
func (o *DatasourceProxyGETcallsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e tcalls bad request response has a 3xx status code
func (o *DatasourceProxyGETcallsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e tcalls bad request response has a 4xx status code
func (o *DatasourceProxyGETcallsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy g e tcalls bad request response has a 5xx status code
func (o *DatasourceProxyGETcallsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e tcalls bad request response a status code equal to that given
func (o *DatasourceProxyGETcallsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the datasource proxy g e tcalls bad request response
func (o *DatasourceProxyGETcallsBadRequest) Code() int {
	return 400
}

func (o *DatasourceProxyGETcallsBadRequest) Error() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsBadRequest  %+v", 400, o.Payload)
}

func (o *DatasourceProxyGETcallsBadRequest) String() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsBadRequest  %+v", 400, o.Payload)
}

func (o *DatasourceProxyGETcallsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETcallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyGETcallsUnauthorized creates a DatasourceProxyGETcallsUnauthorized with default headers values
func NewDatasourceProxyGETcallsUnauthorized() *DatasourceProxyGETcallsUnauthorized {
	return &DatasourceProxyGETcallsUnauthorized{}
}

/*
DatasourceProxyGETcallsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DatasourceProxyGETcallsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e tcalls unauthorized response has a 2xx status code
func (o *DatasourceProxyGETcallsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e tcalls unauthorized response has a 3xx status code
func (o *DatasourceProxyGETcallsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e tcalls unauthorized response has a 4xx status code
func (o *DatasourceProxyGETcallsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy g e tcalls unauthorized response has a 5xx status code
func (o *DatasourceProxyGETcallsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e tcalls unauthorized response a status code equal to that given
func (o *DatasourceProxyGETcallsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the datasource proxy g e tcalls unauthorized response
func (o *DatasourceProxyGETcallsUnauthorized) Code() int {
	return 401
}

func (o *DatasourceProxyGETcallsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsUnauthorized  %+v", 401, o.Payload)
}

func (o *DatasourceProxyGETcallsUnauthorized) String() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsUnauthorized  %+v", 401, o.Payload)
}

func (o *DatasourceProxyGETcallsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETcallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyGETcallsForbidden creates a DatasourceProxyGETcallsForbidden with default headers values
func NewDatasourceProxyGETcallsForbidden() *DatasourceProxyGETcallsForbidden {
	return &DatasourceProxyGETcallsForbidden{}
}

/*
DatasourceProxyGETcallsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DatasourceProxyGETcallsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e tcalls forbidden response has a 2xx status code
func (o *DatasourceProxyGETcallsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e tcalls forbidden response has a 3xx status code
func (o *DatasourceProxyGETcallsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e tcalls forbidden response has a 4xx status code
func (o *DatasourceProxyGETcallsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy g e tcalls forbidden response has a 5xx status code
func (o *DatasourceProxyGETcallsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e tcalls forbidden response a status code equal to that given
func (o *DatasourceProxyGETcallsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the datasource proxy g e tcalls forbidden response
func (o *DatasourceProxyGETcallsForbidden) Code() int {
	return 403
}

func (o *DatasourceProxyGETcallsForbidden) Error() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsForbidden  %+v", 403, o.Payload)
}

func (o *DatasourceProxyGETcallsForbidden) String() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsForbidden  %+v", 403, o.Payload)
}

func (o *DatasourceProxyGETcallsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETcallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyGETcallsNotFound creates a DatasourceProxyGETcallsNotFound with default headers values
func NewDatasourceProxyGETcallsNotFound() *DatasourceProxyGETcallsNotFound {
	return &DatasourceProxyGETcallsNotFound{}
}

/*
DatasourceProxyGETcallsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DatasourceProxyGETcallsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e tcalls not found response has a 2xx status code
func (o *DatasourceProxyGETcallsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e tcalls not found response has a 3xx status code
func (o *DatasourceProxyGETcallsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e tcalls not found response has a 4xx status code
func (o *DatasourceProxyGETcallsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy g e tcalls not found response has a 5xx status code
func (o *DatasourceProxyGETcallsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e tcalls not found response a status code equal to that given
func (o *DatasourceProxyGETcallsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the datasource proxy g e tcalls not found response
func (o *DatasourceProxyGETcallsNotFound) Code() int {
	return 404
}

func (o *DatasourceProxyGETcallsNotFound) Error() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsNotFound  %+v", 404, o.Payload)
}

func (o *DatasourceProxyGETcallsNotFound) String() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsNotFound  %+v", 404, o.Payload)
}

func (o *DatasourceProxyGETcallsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETcallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyGETcallsInternalServerError creates a DatasourceProxyGETcallsInternalServerError with default headers values
func NewDatasourceProxyGETcallsInternalServerError() *DatasourceProxyGETcallsInternalServerError {
	return &DatasourceProxyGETcallsInternalServerError{}
}

/*
DatasourceProxyGETcallsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DatasourceProxyGETcallsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e tcalls internal server error response has a 2xx status code
func (o *DatasourceProxyGETcallsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e tcalls internal server error response has a 3xx status code
func (o *DatasourceProxyGETcallsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e tcalls internal server error response has a 4xx status code
func (o *DatasourceProxyGETcallsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this datasource proxy g e tcalls internal server error response has a 5xx status code
func (o *DatasourceProxyGETcallsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this datasource proxy g e tcalls internal server error response a status code equal to that given
func (o *DatasourceProxyGETcallsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the datasource proxy g e tcalls internal server error response
func (o *DatasourceProxyGETcallsInternalServerError) Code() int {
	return 500
}

func (o *DatasourceProxyGETcallsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsInternalServerError  %+v", 500, o.Payload)
}

func (o *DatasourceProxyGETcallsInternalServerError) String() string {
	return fmt.Sprintf("[GET /datasources/proxy/{id}/{datasource_proxy_route}][%d] datasourceProxyGETcallsInternalServerError  %+v", 500, o.Payload)
}

func (o *DatasourceProxyGETcallsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETcallsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
