// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

// DeleteAPIkeyReader is a Reader for the DeleteAPIkey structure.
type DeleteAPIkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAPIkeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIkeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIkeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAPIkeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /auth/keys/{id}] deleteAPIkey", response, response.Code())
	}
}

// NewDeleteAPIkeyOK creates a DeleteAPIkeyOK with default headers values
func NewDeleteAPIkeyOK() *DeleteAPIkeyOK {
	return &DeleteAPIkeyOK{}
}

/*
DeleteAPIkeyOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteAPIkeyOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete a p ikey Ok response has a 2xx status code
func (o *DeleteAPIkeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete a p ikey Ok response has a 3xx status code
func (o *DeleteAPIkeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete a p ikey Ok response has a 4xx status code
func (o *DeleteAPIkeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete a p ikey Ok response has a 5xx status code
func (o *DeleteAPIkeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete a p ikey Ok response a status code equal to that given
func (o *DeleteAPIkeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete a p ikey Ok response
func (o *DeleteAPIkeyOK) Code() int {
	return 200
}

func (o *DeleteAPIkeyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyOk %s", 200, payload)
}

func (o *DeleteAPIkeyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyOk %s", 200, payload)
}

func (o *DeleteAPIkeyOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteAPIkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIkeyUnauthorized creates a DeleteAPIkeyUnauthorized with default headers values
func NewDeleteAPIkeyUnauthorized() *DeleteAPIkeyUnauthorized {
	return &DeleteAPIkeyUnauthorized{}
}

/*
DeleteAPIkeyUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteAPIkeyUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete a p ikey unauthorized response has a 2xx status code
func (o *DeleteAPIkeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete a p ikey unauthorized response has a 3xx status code
func (o *DeleteAPIkeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete a p ikey unauthorized response has a 4xx status code
func (o *DeleteAPIkeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete a p ikey unauthorized response has a 5xx status code
func (o *DeleteAPIkeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete a p ikey unauthorized response a status code equal to that given
func (o *DeleteAPIkeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete a p ikey unauthorized response
func (o *DeleteAPIkeyUnauthorized) Code() int {
	return 401
}

func (o *DeleteAPIkeyUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyUnauthorized %s", 401, payload)
}

func (o *DeleteAPIkeyUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyUnauthorized %s", 401, payload)
}

func (o *DeleteAPIkeyUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAPIkeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIkeyForbidden creates a DeleteAPIkeyForbidden with default headers values
func NewDeleteAPIkeyForbidden() *DeleteAPIkeyForbidden {
	return &DeleteAPIkeyForbidden{}
}

/*
DeleteAPIkeyForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteAPIkeyForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete a p ikey forbidden response has a 2xx status code
func (o *DeleteAPIkeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete a p ikey forbidden response has a 3xx status code
func (o *DeleteAPIkeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete a p ikey forbidden response has a 4xx status code
func (o *DeleteAPIkeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete a p ikey forbidden response has a 5xx status code
func (o *DeleteAPIkeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete a p ikey forbidden response a status code equal to that given
func (o *DeleteAPIkeyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete a p ikey forbidden response
func (o *DeleteAPIkeyForbidden) Code() int {
	return 403
}

func (o *DeleteAPIkeyForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyForbidden %s", 403, payload)
}

func (o *DeleteAPIkeyForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyForbidden %s", 403, payload)
}

func (o *DeleteAPIkeyForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAPIkeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIkeyNotFound creates a DeleteAPIkeyNotFound with default headers values
func NewDeleteAPIkeyNotFound() *DeleteAPIkeyNotFound {
	return &DeleteAPIkeyNotFound{}
}

/*
DeleteAPIkeyNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteAPIkeyNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete a p ikey not found response has a 2xx status code
func (o *DeleteAPIkeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete a p ikey not found response has a 3xx status code
func (o *DeleteAPIkeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete a p ikey not found response has a 4xx status code
func (o *DeleteAPIkeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete a p ikey not found response has a 5xx status code
func (o *DeleteAPIkeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete a p ikey not found response a status code equal to that given
func (o *DeleteAPIkeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete a p ikey not found response
func (o *DeleteAPIkeyNotFound) Code() int {
	return 404
}

func (o *DeleteAPIkeyNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyNotFound %s", 404, payload)
}

func (o *DeleteAPIkeyNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyNotFound %s", 404, payload)
}

func (o *DeleteAPIkeyNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAPIkeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIkeyInternalServerError creates a DeleteAPIkeyInternalServerError with default headers values
func NewDeleteAPIkeyInternalServerError() *DeleteAPIkeyInternalServerError {
	return &DeleteAPIkeyInternalServerError{}
}

/*
DeleteAPIkeyInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteAPIkeyInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete a p ikey internal server error response has a 2xx status code
func (o *DeleteAPIkeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete a p ikey internal server error response has a 3xx status code
func (o *DeleteAPIkeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete a p ikey internal server error response has a 4xx status code
func (o *DeleteAPIkeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete a p ikey internal server error response has a 5xx status code
func (o *DeleteAPIkeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete a p ikey internal server error response a status code equal to that given
func (o *DeleteAPIkeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete a p ikey internal server error response
func (o *DeleteAPIkeyInternalServerError) Code() int {
	return 500
}

func (o *DeleteAPIkeyInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyInternalServerError %s", 500, payload)
}

func (o *DeleteAPIkeyInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /auth/keys/{id}][%d] deleteAPIkeyInternalServerError %s", 500, payload)
}

func (o *DeleteAPIkeyInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAPIkeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
