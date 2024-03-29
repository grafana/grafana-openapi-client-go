// Code generated by go-swagger; DO NOT EDIT.

package datasources

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

// DatasourceProxyGETByUIDcallsReader is a Reader for the DatasourceProxyGETByUIDcalls structure.
type DatasourceProxyGETByUIDcallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DatasourceProxyGETByUIDcallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDatasourceProxyGETByUIDcallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDatasourceProxyGETByUIDcallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDatasourceProxyGETByUIDcallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDatasourceProxyGETByUIDcallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDatasourceProxyGETByUIDcallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDatasourceProxyGETByUIDcallsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}] datasourceProxyGETByUIDcalls", response, response.Code())
	}
}

// NewDatasourceProxyGETByUIDcallsOK creates a DatasourceProxyGETByUIDcallsOK with default headers values
func NewDatasourceProxyGETByUIDcallsOK() *DatasourceProxyGETByUIDcallsOK {
	return &DatasourceProxyGETByUIDcallsOK{}
}

/*
DatasourceProxyGETByUIDcallsOK describes a response with status code 200, with default header values.

(empty)
*/
type DatasourceProxyGETByUIDcallsOK struct {
}

// IsSuccess returns true when this datasource proxy g e t by Ui dcalls Ok response has a 2xx status code
func (o *DatasourceProxyGETByUIDcallsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this datasource proxy g e t by Ui dcalls Ok response has a 3xx status code
func (o *DatasourceProxyGETByUIDcallsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e t by Ui dcalls Ok response has a 4xx status code
func (o *DatasourceProxyGETByUIDcallsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this datasource proxy g e t by Ui dcalls Ok response has a 5xx status code
func (o *DatasourceProxyGETByUIDcallsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e t by Ui dcalls Ok response a status code equal to that given
func (o *DatasourceProxyGETByUIDcallsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the datasource proxy g e t by Ui dcalls Ok response
func (o *DatasourceProxyGETByUIDcallsOK) Code() int {
	return 200
}

func (o *DatasourceProxyGETByUIDcallsOK) Error() string {
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsOk", 200)
}

func (o *DatasourceProxyGETByUIDcallsOK) String() string {
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsOk", 200)
}

func (o *DatasourceProxyGETByUIDcallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDatasourceProxyGETByUIDcallsBadRequest creates a DatasourceProxyGETByUIDcallsBadRequest with default headers values
func NewDatasourceProxyGETByUIDcallsBadRequest() *DatasourceProxyGETByUIDcallsBadRequest {
	return &DatasourceProxyGETByUIDcallsBadRequest{}
}

/*
DatasourceProxyGETByUIDcallsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DatasourceProxyGETByUIDcallsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e t by Ui dcalls bad request response has a 2xx status code
func (o *DatasourceProxyGETByUIDcallsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e t by Ui dcalls bad request response has a 3xx status code
func (o *DatasourceProxyGETByUIDcallsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e t by Ui dcalls bad request response has a 4xx status code
func (o *DatasourceProxyGETByUIDcallsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy g e t by Ui dcalls bad request response has a 5xx status code
func (o *DatasourceProxyGETByUIDcallsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e t by Ui dcalls bad request response a status code equal to that given
func (o *DatasourceProxyGETByUIDcallsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the datasource proxy g e t by Ui dcalls bad request response
func (o *DatasourceProxyGETByUIDcallsBadRequest) Code() int {
	return 400
}

func (o *DatasourceProxyGETByUIDcallsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsBadRequest %s", 400, payload)
}

func (o *DatasourceProxyGETByUIDcallsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsBadRequest %s", 400, payload)
}

func (o *DatasourceProxyGETByUIDcallsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETByUIDcallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyGETByUIDcallsUnauthorized creates a DatasourceProxyGETByUIDcallsUnauthorized with default headers values
func NewDatasourceProxyGETByUIDcallsUnauthorized() *DatasourceProxyGETByUIDcallsUnauthorized {
	return &DatasourceProxyGETByUIDcallsUnauthorized{}
}

/*
DatasourceProxyGETByUIDcallsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DatasourceProxyGETByUIDcallsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e t by Ui dcalls unauthorized response has a 2xx status code
func (o *DatasourceProxyGETByUIDcallsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e t by Ui dcalls unauthorized response has a 3xx status code
func (o *DatasourceProxyGETByUIDcallsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e t by Ui dcalls unauthorized response has a 4xx status code
func (o *DatasourceProxyGETByUIDcallsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy g e t by Ui dcalls unauthorized response has a 5xx status code
func (o *DatasourceProxyGETByUIDcallsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e t by Ui dcalls unauthorized response a status code equal to that given
func (o *DatasourceProxyGETByUIDcallsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the datasource proxy g e t by Ui dcalls unauthorized response
func (o *DatasourceProxyGETByUIDcallsUnauthorized) Code() int {
	return 401
}

func (o *DatasourceProxyGETByUIDcallsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsUnauthorized %s", 401, payload)
}

func (o *DatasourceProxyGETByUIDcallsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsUnauthorized %s", 401, payload)
}

func (o *DatasourceProxyGETByUIDcallsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETByUIDcallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyGETByUIDcallsForbidden creates a DatasourceProxyGETByUIDcallsForbidden with default headers values
func NewDatasourceProxyGETByUIDcallsForbidden() *DatasourceProxyGETByUIDcallsForbidden {
	return &DatasourceProxyGETByUIDcallsForbidden{}
}

/*
DatasourceProxyGETByUIDcallsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DatasourceProxyGETByUIDcallsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e t by Ui dcalls forbidden response has a 2xx status code
func (o *DatasourceProxyGETByUIDcallsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e t by Ui dcalls forbidden response has a 3xx status code
func (o *DatasourceProxyGETByUIDcallsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e t by Ui dcalls forbidden response has a 4xx status code
func (o *DatasourceProxyGETByUIDcallsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy g e t by Ui dcalls forbidden response has a 5xx status code
func (o *DatasourceProxyGETByUIDcallsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e t by Ui dcalls forbidden response a status code equal to that given
func (o *DatasourceProxyGETByUIDcallsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the datasource proxy g e t by Ui dcalls forbidden response
func (o *DatasourceProxyGETByUIDcallsForbidden) Code() int {
	return 403
}

func (o *DatasourceProxyGETByUIDcallsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsForbidden %s", 403, payload)
}

func (o *DatasourceProxyGETByUIDcallsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsForbidden %s", 403, payload)
}

func (o *DatasourceProxyGETByUIDcallsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETByUIDcallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyGETByUIDcallsNotFound creates a DatasourceProxyGETByUIDcallsNotFound with default headers values
func NewDatasourceProxyGETByUIDcallsNotFound() *DatasourceProxyGETByUIDcallsNotFound {
	return &DatasourceProxyGETByUIDcallsNotFound{}
}

/*
DatasourceProxyGETByUIDcallsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DatasourceProxyGETByUIDcallsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e t by Ui dcalls not found response has a 2xx status code
func (o *DatasourceProxyGETByUIDcallsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e t by Ui dcalls not found response has a 3xx status code
func (o *DatasourceProxyGETByUIDcallsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e t by Ui dcalls not found response has a 4xx status code
func (o *DatasourceProxyGETByUIDcallsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy g e t by Ui dcalls not found response has a 5xx status code
func (o *DatasourceProxyGETByUIDcallsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy g e t by Ui dcalls not found response a status code equal to that given
func (o *DatasourceProxyGETByUIDcallsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the datasource proxy g e t by Ui dcalls not found response
func (o *DatasourceProxyGETByUIDcallsNotFound) Code() int {
	return 404
}

func (o *DatasourceProxyGETByUIDcallsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsNotFound %s", 404, payload)
}

func (o *DatasourceProxyGETByUIDcallsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsNotFound %s", 404, payload)
}

func (o *DatasourceProxyGETByUIDcallsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETByUIDcallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyGETByUIDcallsInternalServerError creates a DatasourceProxyGETByUIDcallsInternalServerError with default headers values
func NewDatasourceProxyGETByUIDcallsInternalServerError() *DatasourceProxyGETByUIDcallsInternalServerError {
	return &DatasourceProxyGETByUIDcallsInternalServerError{}
}

/*
DatasourceProxyGETByUIDcallsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DatasourceProxyGETByUIDcallsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy g e t by Ui dcalls internal server error response has a 2xx status code
func (o *DatasourceProxyGETByUIDcallsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy g e t by Ui dcalls internal server error response has a 3xx status code
func (o *DatasourceProxyGETByUIDcallsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy g e t by Ui dcalls internal server error response has a 4xx status code
func (o *DatasourceProxyGETByUIDcallsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this datasource proxy g e t by Ui dcalls internal server error response has a 5xx status code
func (o *DatasourceProxyGETByUIDcallsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this datasource proxy g e t by Ui dcalls internal server error response a status code equal to that given
func (o *DatasourceProxyGETByUIDcallsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the datasource proxy g e t by Ui dcalls internal server error response
func (o *DatasourceProxyGETByUIDcallsInternalServerError) Code() int {
	return 500
}

func (o *DatasourceProxyGETByUIDcallsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsInternalServerError %s", 500, payload)
}

func (o *DatasourceProxyGETByUIDcallsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyGETByUiDcallsInternalServerError %s", 500, payload)
}

func (o *DatasourceProxyGETByUIDcallsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyGETByUIDcallsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
