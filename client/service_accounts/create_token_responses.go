// Code generated by go-swagger; DO NOT EDIT.

package service_accounts

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

// CreateTokenReader is a Reader for the CreateToken structure.
type CreateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTokenConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /serviceaccounts/{serviceAccountId}/tokens] createToken", response, response.Code())
	}
}

// NewCreateTokenOK creates a CreateTokenOK with default headers values
func NewCreateTokenOK() *CreateTokenOK {
	return &CreateTokenOK{}
}

/*
CreateTokenOK describes a response with status code 200, with default header values.

(empty)
*/
type CreateTokenOK struct {
	Payload *models.NewAPIKeyResult
}

// IsSuccess returns true when this create token Ok response has a 2xx status code
func (o *CreateTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create token Ok response has a 3xx status code
func (o *CreateTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token Ok response has a 4xx status code
func (o *CreateTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create token Ok response has a 5xx status code
func (o *CreateTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create token Ok response a status code equal to that given
func (o *CreateTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create token Ok response
func (o *CreateTokenOK) Code() int {
	return 200
}

func (o *CreateTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenOk %s", 200, payload)
}

func (o *CreateTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenOk %s", 200, payload)
}

func (o *CreateTokenOK) GetPayload() *models.NewAPIKeyResult {
	return o.Payload
}

func (o *CreateTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewAPIKeyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenBadRequest creates a CreateTokenBadRequest with default headers values
func NewCreateTokenBadRequest() *CreateTokenBadRequest {
	return &CreateTokenBadRequest{}
}

/*
CreateTokenBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CreateTokenBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create token bad request response has a 2xx status code
func (o *CreateTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token bad request response has a 3xx status code
func (o *CreateTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token bad request response has a 4xx status code
func (o *CreateTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create token bad request response has a 5xx status code
func (o *CreateTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create token bad request response a status code equal to that given
func (o *CreateTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create token bad request response
func (o *CreateTokenBadRequest) Code() int {
	return 400
}

func (o *CreateTokenBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenBadRequest %s", 400, payload)
}

func (o *CreateTokenBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenBadRequest %s", 400, payload)
}

func (o *CreateTokenBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenUnauthorized creates a CreateTokenUnauthorized with default headers values
func NewCreateTokenUnauthorized() *CreateTokenUnauthorized {
	return &CreateTokenUnauthorized{}
}

/*
CreateTokenUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreateTokenUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create token unauthorized response has a 2xx status code
func (o *CreateTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token unauthorized response has a 3xx status code
func (o *CreateTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token unauthorized response has a 4xx status code
func (o *CreateTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create token unauthorized response has a 5xx status code
func (o *CreateTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create token unauthorized response a status code equal to that given
func (o *CreateTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create token unauthorized response
func (o *CreateTokenUnauthorized) Code() int {
	return 401
}

func (o *CreateTokenUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenUnauthorized %s", 401, payload)
}

func (o *CreateTokenUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenUnauthorized %s", 401, payload)
}

func (o *CreateTokenUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenForbidden creates a CreateTokenForbidden with default headers values
func NewCreateTokenForbidden() *CreateTokenForbidden {
	return &CreateTokenForbidden{}
}

/*
CreateTokenForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateTokenForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create token forbidden response has a 2xx status code
func (o *CreateTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token forbidden response has a 3xx status code
func (o *CreateTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token forbidden response has a 4xx status code
func (o *CreateTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create token forbidden response has a 5xx status code
func (o *CreateTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create token forbidden response a status code equal to that given
func (o *CreateTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create token forbidden response
func (o *CreateTokenForbidden) Code() int {
	return 403
}

func (o *CreateTokenForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenForbidden %s", 403, payload)
}

func (o *CreateTokenForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenForbidden %s", 403, payload)
}

func (o *CreateTokenForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenNotFound creates a CreateTokenNotFound with default headers values
func NewCreateTokenNotFound() *CreateTokenNotFound {
	return &CreateTokenNotFound{}
}

/*
CreateTokenNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type CreateTokenNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create token not found response has a 2xx status code
func (o *CreateTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token not found response has a 3xx status code
func (o *CreateTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token not found response has a 4xx status code
func (o *CreateTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create token not found response has a 5xx status code
func (o *CreateTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create token not found response a status code equal to that given
func (o *CreateTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create token not found response
func (o *CreateTokenNotFound) Code() int {
	return 404
}

func (o *CreateTokenNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenNotFound %s", 404, payload)
}

func (o *CreateTokenNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenNotFound %s", 404, payload)
}

func (o *CreateTokenNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenConflict creates a CreateTokenConflict with default headers values
func NewCreateTokenConflict() *CreateTokenConflict {
	return &CreateTokenConflict{}
}

/*
CreateTokenConflict describes a response with status code 409, with default header values.

ConflictError
*/
type CreateTokenConflict struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create token conflict response has a 2xx status code
func (o *CreateTokenConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token conflict response has a 3xx status code
func (o *CreateTokenConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token conflict response has a 4xx status code
func (o *CreateTokenConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create token conflict response has a 5xx status code
func (o *CreateTokenConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create token conflict response a status code equal to that given
func (o *CreateTokenConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create token conflict response
func (o *CreateTokenConflict) Code() int {
	return 409
}

func (o *CreateTokenConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenConflict %s", 409, payload)
}

func (o *CreateTokenConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenConflict %s", 409, payload)
}

func (o *CreateTokenConflict) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenInternalServerError creates a CreateTokenInternalServerError with default headers values
func NewCreateTokenInternalServerError() *CreateTokenInternalServerError {
	return &CreateTokenInternalServerError{}
}

/*
CreateTokenInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateTokenInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create token internal server error response has a 2xx status code
func (o *CreateTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token internal server error response has a 3xx status code
func (o *CreateTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token internal server error response has a 4xx status code
func (o *CreateTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create token internal server error response has a 5xx status code
func (o *CreateTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create token internal server error response a status code equal to that given
func (o *CreateTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create token internal server error response
func (o *CreateTokenInternalServerError) Code() int {
	return 500
}

func (o *CreateTokenInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenInternalServerError %s", 500, payload)
}

func (o *CreateTokenInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts/{serviceAccountId}/tokens][%d] createTokenInternalServerError %s", 500, payload)
}

func (o *CreateTokenInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
