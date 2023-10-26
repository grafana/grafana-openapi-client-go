// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// ChangeUserPasswordReader is a Reader for the ChangeUserPassword structure.
type ChangeUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChangeUserPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewChangeUserPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewChangeUserPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewChangeUserPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /user/password] changeUserPassword", response, response.Code())
	}
}

// NewChangeUserPasswordOK creates a ChangeUserPasswordOK with default headers values
func NewChangeUserPasswordOK() *ChangeUserPasswordOK {
	return &ChangeUserPasswordOK{}
}

/*
ChangeUserPasswordOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type ChangeUserPasswordOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this change user password Ok response has a 2xx status code
func (o *ChangeUserPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change user password Ok response has a 3xx status code
func (o *ChangeUserPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change user password Ok response has a 4xx status code
func (o *ChangeUserPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this change user password Ok response has a 5xx status code
func (o *ChangeUserPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this change user password Ok response a status code equal to that given
func (o *ChangeUserPasswordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the change user password Ok response
func (o *ChangeUserPasswordOK) Code() int {
	return 200
}

func (o *ChangeUserPasswordOK) Error() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordOk  %+v", 200, o.Payload)
}

func (o *ChangeUserPasswordOK) String() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordOk  %+v", 200, o.Payload)
}

func (o *ChangeUserPasswordOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *ChangeUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeUserPasswordBadRequest creates a ChangeUserPasswordBadRequest with default headers values
func NewChangeUserPasswordBadRequest() *ChangeUserPasswordBadRequest {
	return &ChangeUserPasswordBadRequest{}
}

/*
ChangeUserPasswordBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type ChangeUserPasswordBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this change user password bad request response has a 2xx status code
func (o *ChangeUserPasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change user password bad request response has a 3xx status code
func (o *ChangeUserPasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change user password bad request response has a 4xx status code
func (o *ChangeUserPasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this change user password bad request response has a 5xx status code
func (o *ChangeUserPasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this change user password bad request response a status code equal to that given
func (o *ChangeUserPasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the change user password bad request response
func (o *ChangeUserPasswordBadRequest) Code() int {
	return 400
}

func (o *ChangeUserPasswordBadRequest) Error() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeUserPasswordBadRequest) String() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeUserPasswordBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ChangeUserPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeUserPasswordUnauthorized creates a ChangeUserPasswordUnauthorized with default headers values
func NewChangeUserPasswordUnauthorized() *ChangeUserPasswordUnauthorized {
	return &ChangeUserPasswordUnauthorized{}
}

/*
ChangeUserPasswordUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type ChangeUserPasswordUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this change user password unauthorized response has a 2xx status code
func (o *ChangeUserPasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change user password unauthorized response has a 3xx status code
func (o *ChangeUserPasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change user password unauthorized response has a 4xx status code
func (o *ChangeUserPasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this change user password unauthorized response has a 5xx status code
func (o *ChangeUserPasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this change user password unauthorized response a status code equal to that given
func (o *ChangeUserPasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the change user password unauthorized response
func (o *ChangeUserPasswordUnauthorized) Code() int {
	return 401
}

func (o *ChangeUserPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *ChangeUserPasswordUnauthorized) String() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *ChangeUserPasswordUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ChangeUserPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeUserPasswordForbidden creates a ChangeUserPasswordForbidden with default headers values
func NewChangeUserPasswordForbidden() *ChangeUserPasswordForbidden {
	return &ChangeUserPasswordForbidden{}
}

/*
ChangeUserPasswordForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ChangeUserPasswordForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this change user password forbidden response has a 2xx status code
func (o *ChangeUserPasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change user password forbidden response has a 3xx status code
func (o *ChangeUserPasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change user password forbidden response has a 4xx status code
func (o *ChangeUserPasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this change user password forbidden response has a 5xx status code
func (o *ChangeUserPasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this change user password forbidden response a status code equal to that given
func (o *ChangeUserPasswordForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the change user password forbidden response
func (o *ChangeUserPasswordForbidden) Code() int {
	return 403
}

func (o *ChangeUserPasswordForbidden) Error() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordForbidden  %+v", 403, o.Payload)
}

func (o *ChangeUserPasswordForbidden) String() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordForbidden  %+v", 403, o.Payload)
}

func (o *ChangeUserPasswordForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ChangeUserPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeUserPasswordInternalServerError creates a ChangeUserPasswordInternalServerError with default headers values
func NewChangeUserPasswordInternalServerError() *ChangeUserPasswordInternalServerError {
	return &ChangeUserPasswordInternalServerError{}
}

/*
ChangeUserPasswordInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ChangeUserPasswordInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this change user password internal server error response has a 2xx status code
func (o *ChangeUserPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change user password internal server error response has a 3xx status code
func (o *ChangeUserPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change user password internal server error response has a 4xx status code
func (o *ChangeUserPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this change user password internal server error response has a 5xx status code
func (o *ChangeUserPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this change user password internal server error response a status code equal to that given
func (o *ChangeUserPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the change user password internal server error response
func (o *ChangeUserPasswordInternalServerError) Code() int {
	return 500
}

func (o *ChangeUserPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *ChangeUserPasswordInternalServerError) String() string {
	return fmt.Sprintf("[PUT /user/password][%d] changeUserPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *ChangeUserPasswordInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ChangeUserPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
