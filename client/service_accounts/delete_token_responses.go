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

// DeleteTokenReader is a Reader for the DeleteToken structure.
type DeleteTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}] deleteToken", response, response.Code())
	}
}

// NewDeleteTokenOK creates a DeleteTokenOK with default headers values
func NewDeleteTokenOK() *DeleteTokenOK {
	return &DeleteTokenOK{}
}

/*
DeleteTokenOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteTokenOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete token Ok response has a 2xx status code
func (o *DeleteTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete token Ok response has a 3xx status code
func (o *DeleteTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token Ok response has a 4xx status code
func (o *DeleteTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete token Ok response has a 5xx status code
func (o *DeleteTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token Ok response a status code equal to that given
func (o *DeleteTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete token Ok response
func (o *DeleteTokenOK) Code() int {
	return 200
}

func (o *DeleteTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenOk %s", 200, payload)
}

func (o *DeleteTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenOk %s", 200, payload)
}

func (o *DeleteTokenOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenBadRequest creates a DeleteTokenBadRequest with default headers values
func NewDeleteTokenBadRequest() *DeleteTokenBadRequest {
	return &DeleteTokenBadRequest{}
}

/*
DeleteTokenBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DeleteTokenBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete token bad request response has a 2xx status code
func (o *DeleteTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete token bad request response has a 3xx status code
func (o *DeleteTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token bad request response has a 4xx status code
func (o *DeleteTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete token bad request response has a 5xx status code
func (o *DeleteTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token bad request response a status code equal to that given
func (o *DeleteTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete token bad request response
func (o *DeleteTokenBadRequest) Code() int {
	return 400
}

func (o *DeleteTokenBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenBadRequest %s", 400, payload)
}

func (o *DeleteTokenBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenBadRequest %s", 400, payload)
}

func (o *DeleteTokenBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenUnauthorized creates a DeleteTokenUnauthorized with default headers values
func NewDeleteTokenUnauthorized() *DeleteTokenUnauthorized {
	return &DeleteTokenUnauthorized{}
}

/*
DeleteTokenUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteTokenUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete token unauthorized response has a 2xx status code
func (o *DeleteTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete token unauthorized response has a 3xx status code
func (o *DeleteTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token unauthorized response has a 4xx status code
func (o *DeleteTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete token unauthorized response has a 5xx status code
func (o *DeleteTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token unauthorized response a status code equal to that given
func (o *DeleteTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete token unauthorized response
func (o *DeleteTokenUnauthorized) Code() int {
	return 401
}

func (o *DeleteTokenUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenUnauthorized %s", 401, payload)
}

func (o *DeleteTokenUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenUnauthorized %s", 401, payload)
}

func (o *DeleteTokenUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenForbidden creates a DeleteTokenForbidden with default headers values
func NewDeleteTokenForbidden() *DeleteTokenForbidden {
	return &DeleteTokenForbidden{}
}

/*
DeleteTokenForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteTokenForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete token forbidden response has a 2xx status code
func (o *DeleteTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete token forbidden response has a 3xx status code
func (o *DeleteTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token forbidden response has a 4xx status code
func (o *DeleteTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete token forbidden response has a 5xx status code
func (o *DeleteTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token forbidden response a status code equal to that given
func (o *DeleteTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete token forbidden response
func (o *DeleteTokenForbidden) Code() int {
	return 403
}

func (o *DeleteTokenForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenForbidden %s", 403, payload)
}

func (o *DeleteTokenForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenForbidden %s", 403, payload)
}

func (o *DeleteTokenForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenNotFound creates a DeleteTokenNotFound with default headers values
func NewDeleteTokenNotFound() *DeleteTokenNotFound {
	return &DeleteTokenNotFound{}
}

/*
DeleteTokenNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteTokenNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete token not found response has a 2xx status code
func (o *DeleteTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete token not found response has a 3xx status code
func (o *DeleteTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token not found response has a 4xx status code
func (o *DeleteTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete token not found response has a 5xx status code
func (o *DeleteTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token not found response a status code equal to that given
func (o *DeleteTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete token not found response
func (o *DeleteTokenNotFound) Code() int {
	return 404
}

func (o *DeleteTokenNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenNotFound %s", 404, payload)
}

func (o *DeleteTokenNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenNotFound %s", 404, payload)
}

func (o *DeleteTokenNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenInternalServerError creates a DeleteTokenInternalServerError with default headers values
func NewDeleteTokenInternalServerError() *DeleteTokenInternalServerError {
	return &DeleteTokenInternalServerError{}
}

/*
DeleteTokenInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteTokenInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete token internal server error response has a 2xx status code
func (o *DeleteTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete token internal server error response has a 3xx status code
func (o *DeleteTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token internal server error response has a 4xx status code
func (o *DeleteTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete token internal server error response has a 5xx status code
func (o *DeleteTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete token internal server error response a status code equal to that given
func (o *DeleteTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete token internal server error response
func (o *DeleteTokenInternalServerError) Code() int {
	return 500
}

func (o *DeleteTokenInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenInternalServerError %s", 500, payload)
}

func (o *DeleteTokenInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}/tokens/{tokenId}][%d] deleteTokenInternalServerError %s", 500, payload)
}

func (o *DeleteTokenInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
