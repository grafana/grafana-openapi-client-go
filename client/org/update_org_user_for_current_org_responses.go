// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// UpdateOrgUserForCurrentOrgReader is a Reader for the UpdateOrgUserForCurrentOrg structure.
type UpdateOrgUserForCurrentOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrgUserForCurrentOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrgUserForCurrentOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrgUserForCurrentOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrgUserForCurrentOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrgUserForCurrentOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateOrgUserForCurrentOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /org/users/{user_id}] updateOrgUserForCurrentOrg", response, response.Code())
	}
}

// NewUpdateOrgUserForCurrentOrgOK creates a UpdateOrgUserForCurrentOrgOK with default headers values
func NewUpdateOrgUserForCurrentOrgOK() *UpdateOrgUserForCurrentOrgOK {
	return &UpdateOrgUserForCurrentOrgOK{}
}

/*
UpdateOrgUserForCurrentOrgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateOrgUserForCurrentOrgOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this update org user for current org o k response has a 2xx status code
func (o *UpdateOrgUserForCurrentOrgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update org user for current org o k response has a 3xx status code
func (o *UpdateOrgUserForCurrentOrgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org user for current org o k response has a 4xx status code
func (o *UpdateOrgUserForCurrentOrgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update org user for current org o k response has a 5xx status code
func (o *UpdateOrgUserForCurrentOrgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update org user for current org o k response a status code equal to that given
func (o *UpdateOrgUserForCurrentOrgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update org user for current org o k response
func (o *UpdateOrgUserForCurrentOrgOK) Code() int {
	return 200
}

func (o *UpdateOrgUserForCurrentOrgOK) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgOK  %+v", 200, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgOK) String() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgOK  %+v", 200, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForCurrentOrgBadRequest creates a UpdateOrgUserForCurrentOrgBadRequest with default headers values
func NewUpdateOrgUserForCurrentOrgBadRequest() *UpdateOrgUserForCurrentOrgBadRequest {
	return &UpdateOrgUserForCurrentOrgBadRequest{}
}

/*
UpdateOrgUserForCurrentOrgBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateOrgUserForCurrentOrgBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update org user for current org bad request response has a 2xx status code
func (o *UpdateOrgUserForCurrentOrgBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org user for current org bad request response has a 3xx status code
func (o *UpdateOrgUserForCurrentOrgBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org user for current org bad request response has a 4xx status code
func (o *UpdateOrgUserForCurrentOrgBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update org user for current org bad request response has a 5xx status code
func (o *UpdateOrgUserForCurrentOrgBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update org user for current org bad request response a status code equal to that given
func (o *UpdateOrgUserForCurrentOrgBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update org user for current org bad request response
func (o *UpdateOrgUserForCurrentOrgBadRequest) Code() int {
	return 400
}

func (o *UpdateOrgUserForCurrentOrgBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgBadRequest) String() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForCurrentOrgUnauthorized creates a UpdateOrgUserForCurrentOrgUnauthorized with default headers values
func NewUpdateOrgUserForCurrentOrgUnauthorized() *UpdateOrgUserForCurrentOrgUnauthorized {
	return &UpdateOrgUserForCurrentOrgUnauthorized{}
}

/*
UpdateOrgUserForCurrentOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateOrgUserForCurrentOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update org user for current org unauthorized response has a 2xx status code
func (o *UpdateOrgUserForCurrentOrgUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org user for current org unauthorized response has a 3xx status code
func (o *UpdateOrgUserForCurrentOrgUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org user for current org unauthorized response has a 4xx status code
func (o *UpdateOrgUserForCurrentOrgUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update org user for current org unauthorized response has a 5xx status code
func (o *UpdateOrgUserForCurrentOrgUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update org user for current org unauthorized response a status code equal to that given
func (o *UpdateOrgUserForCurrentOrgUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update org user for current org unauthorized response
func (o *UpdateOrgUserForCurrentOrgUnauthorized) Code() int {
	return 401
}

func (o *UpdateOrgUserForCurrentOrgUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForCurrentOrgForbidden creates a UpdateOrgUserForCurrentOrgForbidden with default headers values
func NewUpdateOrgUserForCurrentOrgForbidden() *UpdateOrgUserForCurrentOrgForbidden {
	return &UpdateOrgUserForCurrentOrgForbidden{}
}

/*
UpdateOrgUserForCurrentOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateOrgUserForCurrentOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update org user for current org forbidden response has a 2xx status code
func (o *UpdateOrgUserForCurrentOrgForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org user for current org forbidden response has a 3xx status code
func (o *UpdateOrgUserForCurrentOrgForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org user for current org forbidden response has a 4xx status code
func (o *UpdateOrgUserForCurrentOrgForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update org user for current org forbidden response has a 5xx status code
func (o *UpdateOrgUserForCurrentOrgForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update org user for current org forbidden response a status code equal to that given
func (o *UpdateOrgUserForCurrentOrgForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update org user for current org forbidden response
func (o *UpdateOrgUserForCurrentOrgForbidden) Code() int {
	return 403
}

func (o *UpdateOrgUserForCurrentOrgForbidden) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgForbidden) String() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForCurrentOrgInternalServerError creates a UpdateOrgUserForCurrentOrgInternalServerError with default headers values
func NewUpdateOrgUserForCurrentOrgInternalServerError() *UpdateOrgUserForCurrentOrgInternalServerError {
	return &UpdateOrgUserForCurrentOrgInternalServerError{}
}

/*
UpdateOrgUserForCurrentOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateOrgUserForCurrentOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update org user for current org internal server error response has a 2xx status code
func (o *UpdateOrgUserForCurrentOrgInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update org user for current org internal server error response has a 3xx status code
func (o *UpdateOrgUserForCurrentOrgInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update org user for current org internal server error response has a 4xx status code
func (o *UpdateOrgUserForCurrentOrgInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update org user for current org internal server error response has a 5xx status code
func (o *UpdateOrgUserForCurrentOrgInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update org user for current org internal server error response a status code equal to that given
func (o *UpdateOrgUserForCurrentOrgInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update org user for current org internal server error response
func (o *UpdateOrgUserForCurrentOrgInternalServerError) Code() int {
	return 500
}

func (o *UpdateOrgUserForCurrentOrgInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /org/users/{user_id}][%d] updateOrgUserForCurrentOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOrgUserForCurrentOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForCurrentOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}