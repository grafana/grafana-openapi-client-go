// Code generated by go-swagger; DO NOT EDIT.

package library_elements

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

// GetLibraryElementConnectionsReader is a Reader for the GetLibraryElementConnections structure.
type GetLibraryElementConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLibraryElementConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLibraryElementConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLibraryElementConnectionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLibraryElementConnectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLibraryElementConnectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLibraryElementConnectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /library-elements/{library_element_uid}/connections/] getLibraryElementConnections", response, response.Code())
	}
}

// NewGetLibraryElementConnectionsOK creates a GetLibraryElementConnectionsOK with default headers values
func NewGetLibraryElementConnectionsOK() *GetLibraryElementConnectionsOK {
	return &GetLibraryElementConnectionsOK{}
}

/*
GetLibraryElementConnectionsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetLibraryElementConnectionsOK struct {
	Payload *models.LibraryElementConnectionsResponse
}

// IsSuccess returns true when this get library element connections Ok response has a 2xx status code
func (o *GetLibraryElementConnectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get library element connections Ok response has a 3xx status code
func (o *GetLibraryElementConnectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element connections Ok response has a 4xx status code
func (o *GetLibraryElementConnectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get library element connections Ok response has a 5xx status code
func (o *GetLibraryElementConnectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get library element connections Ok response a status code equal to that given
func (o *GetLibraryElementConnectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get library element connections Ok response
func (o *GetLibraryElementConnectionsOK) Code() int {
	return 200
}

func (o *GetLibraryElementConnectionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsOk %s", 200, payload)
}

func (o *GetLibraryElementConnectionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsOk %s", 200, payload)
}

func (o *GetLibraryElementConnectionsOK) GetPayload() *models.LibraryElementConnectionsResponse {
	return o.Payload
}

func (o *GetLibraryElementConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LibraryElementConnectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementConnectionsUnauthorized creates a GetLibraryElementConnectionsUnauthorized with default headers values
func NewGetLibraryElementConnectionsUnauthorized() *GetLibraryElementConnectionsUnauthorized {
	return &GetLibraryElementConnectionsUnauthorized{}
}

/*
GetLibraryElementConnectionsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetLibraryElementConnectionsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get library element connections unauthorized response has a 2xx status code
func (o *GetLibraryElementConnectionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get library element connections unauthorized response has a 3xx status code
func (o *GetLibraryElementConnectionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element connections unauthorized response has a 4xx status code
func (o *GetLibraryElementConnectionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get library element connections unauthorized response has a 5xx status code
func (o *GetLibraryElementConnectionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get library element connections unauthorized response a status code equal to that given
func (o *GetLibraryElementConnectionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get library element connections unauthorized response
func (o *GetLibraryElementConnectionsUnauthorized) Code() int {
	return 401
}

func (o *GetLibraryElementConnectionsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsUnauthorized %s", 401, payload)
}

func (o *GetLibraryElementConnectionsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsUnauthorized %s", 401, payload)
}

func (o *GetLibraryElementConnectionsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementConnectionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementConnectionsForbidden creates a GetLibraryElementConnectionsForbidden with default headers values
func NewGetLibraryElementConnectionsForbidden() *GetLibraryElementConnectionsForbidden {
	return &GetLibraryElementConnectionsForbidden{}
}

/*
GetLibraryElementConnectionsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetLibraryElementConnectionsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get library element connections forbidden response has a 2xx status code
func (o *GetLibraryElementConnectionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get library element connections forbidden response has a 3xx status code
func (o *GetLibraryElementConnectionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element connections forbidden response has a 4xx status code
func (o *GetLibraryElementConnectionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get library element connections forbidden response has a 5xx status code
func (o *GetLibraryElementConnectionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get library element connections forbidden response a status code equal to that given
func (o *GetLibraryElementConnectionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get library element connections forbidden response
func (o *GetLibraryElementConnectionsForbidden) Code() int {
	return 403
}

func (o *GetLibraryElementConnectionsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsForbidden %s", 403, payload)
}

func (o *GetLibraryElementConnectionsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsForbidden %s", 403, payload)
}

func (o *GetLibraryElementConnectionsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementConnectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementConnectionsNotFound creates a GetLibraryElementConnectionsNotFound with default headers values
func NewGetLibraryElementConnectionsNotFound() *GetLibraryElementConnectionsNotFound {
	return &GetLibraryElementConnectionsNotFound{}
}

/*
GetLibraryElementConnectionsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetLibraryElementConnectionsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get library element connections not found response has a 2xx status code
func (o *GetLibraryElementConnectionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get library element connections not found response has a 3xx status code
func (o *GetLibraryElementConnectionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element connections not found response has a 4xx status code
func (o *GetLibraryElementConnectionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get library element connections not found response has a 5xx status code
func (o *GetLibraryElementConnectionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get library element connections not found response a status code equal to that given
func (o *GetLibraryElementConnectionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get library element connections not found response
func (o *GetLibraryElementConnectionsNotFound) Code() int {
	return 404
}

func (o *GetLibraryElementConnectionsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsNotFound %s", 404, payload)
}

func (o *GetLibraryElementConnectionsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsNotFound %s", 404, payload)
}

func (o *GetLibraryElementConnectionsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementConnectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementConnectionsInternalServerError creates a GetLibraryElementConnectionsInternalServerError with default headers values
func NewGetLibraryElementConnectionsInternalServerError() *GetLibraryElementConnectionsInternalServerError {
	return &GetLibraryElementConnectionsInternalServerError{}
}

/*
GetLibraryElementConnectionsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetLibraryElementConnectionsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get library element connections internal server error response has a 2xx status code
func (o *GetLibraryElementConnectionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get library element connections internal server error response has a 3xx status code
func (o *GetLibraryElementConnectionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element connections internal server error response has a 4xx status code
func (o *GetLibraryElementConnectionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get library element connections internal server error response has a 5xx status code
func (o *GetLibraryElementConnectionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get library element connections internal server error response a status code equal to that given
func (o *GetLibraryElementConnectionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get library element connections internal server error response
func (o *GetLibraryElementConnectionsInternalServerError) Code() int {
	return 500
}

func (o *GetLibraryElementConnectionsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsInternalServerError %s", 500, payload)
}

func (o *GetLibraryElementConnectionsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /library-elements/{library_element_uid}/connections/][%d] getLibraryElementConnectionsInternalServerError %s", 500, payload)
}

func (o *GetLibraryElementConnectionsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementConnectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
