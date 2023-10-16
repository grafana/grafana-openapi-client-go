// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// AdminCreateUserReader is a Reader for the AdminCreateUser structure.
type AdminCreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminCreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAdminCreateUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/users] adminCreateUser", response, response.Code())
	}
}

// NewAdminCreateUserOK creates a AdminCreateUserOK with default headers values
func NewAdminCreateUserOK() *AdminCreateUserOK {
	return &AdminCreateUserOK{}
}

/*
AdminCreateUserOK describes a response with status code 200, with default header values.

(empty)
*/
type AdminCreateUserOK struct {
	Payload *models.AdminCreateUserResponse
}

// IsSuccess returns true when this admin create user o k response has a 2xx status code
func (o *AdminCreateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin create user o k response has a 3xx status code
func (o *AdminCreateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user o k response has a 4xx status code
func (o *AdminCreateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin create user o k response has a 5xx status code
func (o *AdminCreateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user o k response a status code equal to that given
func (o *AdminCreateUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin create user o k response
func (o *AdminCreateUserOK) Code() int {
	return 200
}

func (o *AdminCreateUserOK) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserOK  %+v", 200, o.Payload)
}

func (o *AdminCreateUserOK) String() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserOK  %+v", 200, o.Payload)
}

func (o *AdminCreateUserOK) GetPayload() *models.AdminCreateUserResponse {
	return o.Payload
}

func (o *AdminCreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminCreateUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserBadRequest creates a AdminCreateUserBadRequest with default headers values
func NewAdminCreateUserBadRequest() *AdminCreateUserBadRequest {
	return &AdminCreateUserBadRequest{}
}

/*
AdminCreateUserBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AdminCreateUserBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin create user bad request response has a 2xx status code
func (o *AdminCreateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user bad request response has a 3xx status code
func (o *AdminCreateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user bad request response has a 4xx status code
func (o *AdminCreateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create user bad request response has a 5xx status code
func (o *AdminCreateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user bad request response a status code equal to that given
func (o *AdminCreateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin create user bad request response
func (o *AdminCreateUserBadRequest) Code() int {
	return 400
}

func (o *AdminCreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserBadRequest  %+v", 400, o.Payload)
}

func (o *AdminCreateUserBadRequest) String() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserBadRequest  %+v", 400, o.Payload)
}

func (o *AdminCreateUserBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserUnauthorized creates a AdminCreateUserUnauthorized with default headers values
func NewAdminCreateUserUnauthorized() *AdminCreateUserUnauthorized {
	return &AdminCreateUserUnauthorized{}
}

/*
AdminCreateUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminCreateUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin create user unauthorized response has a 2xx status code
func (o *AdminCreateUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user unauthorized response has a 3xx status code
func (o *AdminCreateUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user unauthorized response has a 4xx status code
func (o *AdminCreateUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create user unauthorized response has a 5xx status code
func (o *AdminCreateUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user unauthorized response a status code equal to that given
func (o *AdminCreateUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin create user unauthorized response
func (o *AdminCreateUserUnauthorized) Code() int {
	return 401
}

func (o *AdminCreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminCreateUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminCreateUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserForbidden creates a AdminCreateUserForbidden with default headers values
func NewAdminCreateUserForbidden() *AdminCreateUserForbidden {
	return &AdminCreateUserForbidden{}
}

/*
AdminCreateUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminCreateUserForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin create user forbidden response has a 2xx status code
func (o *AdminCreateUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user forbidden response has a 3xx status code
func (o *AdminCreateUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user forbidden response has a 4xx status code
func (o *AdminCreateUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create user forbidden response has a 5xx status code
func (o *AdminCreateUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user forbidden response a status code equal to that given
func (o *AdminCreateUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin create user forbidden response
func (o *AdminCreateUserForbidden) Code() int {
	return 403
}

func (o *AdminCreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminCreateUserForbidden) String() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminCreateUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserPreconditionFailed creates a AdminCreateUserPreconditionFailed with default headers values
func NewAdminCreateUserPreconditionFailed() *AdminCreateUserPreconditionFailed {
	return &AdminCreateUserPreconditionFailed{}
}

/*
AdminCreateUserPreconditionFailed describes a response with status code 412, with default header values.

PreconditionFailedError
*/
type AdminCreateUserPreconditionFailed struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin create user precondition failed response has a 2xx status code
func (o *AdminCreateUserPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user precondition failed response has a 3xx status code
func (o *AdminCreateUserPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user precondition failed response has a 4xx status code
func (o *AdminCreateUserPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create user precondition failed response has a 5xx status code
func (o *AdminCreateUserPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user precondition failed response a status code equal to that given
func (o *AdminCreateUserPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the admin create user precondition failed response
func (o *AdminCreateUserPreconditionFailed) Code() int {
	return 412
}

func (o *AdminCreateUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AdminCreateUserPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AdminCreateUserPreconditionFailed) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserInternalServerError creates a AdminCreateUserInternalServerError with default headers values
func NewAdminCreateUserInternalServerError() *AdminCreateUserInternalServerError {
	return &AdminCreateUserInternalServerError{}
}

/*
AdminCreateUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminCreateUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin create user internal server error response has a 2xx status code
func (o *AdminCreateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user internal server error response has a 3xx status code
func (o *AdminCreateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user internal server error response has a 4xx status code
func (o *AdminCreateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin create user internal server error response has a 5xx status code
func (o *AdminCreateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin create user internal server error response a status code equal to that given
func (o *AdminCreateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin create user internal server error response
func (o *AdminCreateUserInternalServerError) Code() int {
	return 500
}

func (o *AdminCreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminCreateUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminCreateUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}