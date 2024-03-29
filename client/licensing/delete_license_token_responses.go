// Code generated by go-swagger; DO NOT EDIT.

package licensing

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

// DeleteLicenseTokenReader is a Reader for the DeleteLicenseToken structure.
type DeleteLicenseTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLicenseTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteLicenseTokenAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLicenseTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLicenseTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLicenseTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteLicenseTokenUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLicenseTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /licensing/token] deleteLicenseToken", response, response.Code())
	}
}

// NewDeleteLicenseTokenAccepted creates a DeleteLicenseTokenAccepted with default headers values
func NewDeleteLicenseTokenAccepted() *DeleteLicenseTokenAccepted {
	return &DeleteLicenseTokenAccepted{}
}

/*
DeleteLicenseTokenAccepted describes a response with status code 202, with default header values.

AcceptedResponse
*/
type DeleteLicenseTokenAccepted struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete license token accepted response has a 2xx status code
func (o *DeleteLicenseTokenAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete license token accepted response has a 3xx status code
func (o *DeleteLicenseTokenAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license token accepted response has a 4xx status code
func (o *DeleteLicenseTokenAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete license token accepted response has a 5xx status code
func (o *DeleteLicenseTokenAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete license token accepted response a status code equal to that given
func (o *DeleteLicenseTokenAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete license token accepted response
func (o *DeleteLicenseTokenAccepted) Code() int {
	return 202
}

func (o *DeleteLicenseTokenAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenAccepted %s", 202, payload)
}

func (o *DeleteLicenseTokenAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenAccepted %s", 202, payload)
}

func (o *DeleteLicenseTokenAccepted) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteLicenseTokenAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLicenseTokenBadRequest creates a DeleteLicenseTokenBadRequest with default headers values
func NewDeleteLicenseTokenBadRequest() *DeleteLicenseTokenBadRequest {
	return &DeleteLicenseTokenBadRequest{}
}

/*
DeleteLicenseTokenBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DeleteLicenseTokenBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete license token bad request response has a 2xx status code
func (o *DeleteLicenseTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete license token bad request response has a 3xx status code
func (o *DeleteLicenseTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license token bad request response has a 4xx status code
func (o *DeleteLicenseTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete license token bad request response has a 5xx status code
func (o *DeleteLicenseTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete license token bad request response a status code equal to that given
func (o *DeleteLicenseTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete license token bad request response
func (o *DeleteLicenseTokenBadRequest) Code() int {
	return 400
}

func (o *DeleteLicenseTokenBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenBadRequest %s", 400, payload)
}

func (o *DeleteLicenseTokenBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenBadRequest %s", 400, payload)
}

func (o *DeleteLicenseTokenBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteLicenseTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLicenseTokenUnauthorized creates a DeleteLicenseTokenUnauthorized with default headers values
func NewDeleteLicenseTokenUnauthorized() *DeleteLicenseTokenUnauthorized {
	return &DeleteLicenseTokenUnauthorized{}
}

/*
DeleteLicenseTokenUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteLicenseTokenUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete license token unauthorized response has a 2xx status code
func (o *DeleteLicenseTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete license token unauthorized response has a 3xx status code
func (o *DeleteLicenseTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license token unauthorized response has a 4xx status code
func (o *DeleteLicenseTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete license token unauthorized response has a 5xx status code
func (o *DeleteLicenseTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete license token unauthorized response a status code equal to that given
func (o *DeleteLicenseTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete license token unauthorized response
func (o *DeleteLicenseTokenUnauthorized) Code() int {
	return 401
}

func (o *DeleteLicenseTokenUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenUnauthorized %s", 401, payload)
}

func (o *DeleteLicenseTokenUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenUnauthorized %s", 401, payload)
}

func (o *DeleteLicenseTokenUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteLicenseTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLicenseTokenForbidden creates a DeleteLicenseTokenForbidden with default headers values
func NewDeleteLicenseTokenForbidden() *DeleteLicenseTokenForbidden {
	return &DeleteLicenseTokenForbidden{}
}

/*
DeleteLicenseTokenForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteLicenseTokenForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete license token forbidden response has a 2xx status code
func (o *DeleteLicenseTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete license token forbidden response has a 3xx status code
func (o *DeleteLicenseTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license token forbidden response has a 4xx status code
func (o *DeleteLicenseTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete license token forbidden response has a 5xx status code
func (o *DeleteLicenseTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete license token forbidden response a status code equal to that given
func (o *DeleteLicenseTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete license token forbidden response
func (o *DeleteLicenseTokenForbidden) Code() int {
	return 403
}

func (o *DeleteLicenseTokenForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenForbidden %s", 403, payload)
}

func (o *DeleteLicenseTokenForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenForbidden %s", 403, payload)
}

func (o *DeleteLicenseTokenForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteLicenseTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLicenseTokenUnprocessableEntity creates a DeleteLicenseTokenUnprocessableEntity with default headers values
func NewDeleteLicenseTokenUnprocessableEntity() *DeleteLicenseTokenUnprocessableEntity {
	return &DeleteLicenseTokenUnprocessableEntity{}
}

/*
DeleteLicenseTokenUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntityError
*/
type DeleteLicenseTokenUnprocessableEntity struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete license token unprocessable entity response has a 2xx status code
func (o *DeleteLicenseTokenUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete license token unprocessable entity response has a 3xx status code
func (o *DeleteLicenseTokenUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license token unprocessable entity response has a 4xx status code
func (o *DeleteLicenseTokenUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete license token unprocessable entity response has a 5xx status code
func (o *DeleteLicenseTokenUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete license token unprocessable entity response a status code equal to that given
func (o *DeleteLicenseTokenUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete license token unprocessable entity response
func (o *DeleteLicenseTokenUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteLicenseTokenUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenUnprocessableEntity %s", 422, payload)
}

func (o *DeleteLicenseTokenUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenUnprocessableEntity %s", 422, payload)
}

func (o *DeleteLicenseTokenUnprocessableEntity) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteLicenseTokenUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLicenseTokenInternalServerError creates a DeleteLicenseTokenInternalServerError with default headers values
func NewDeleteLicenseTokenInternalServerError() *DeleteLicenseTokenInternalServerError {
	return &DeleteLicenseTokenInternalServerError{}
}

/*
DeleteLicenseTokenInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteLicenseTokenInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete license token internal server error response has a 2xx status code
func (o *DeleteLicenseTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete license token internal server error response has a 3xx status code
func (o *DeleteLicenseTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license token internal server error response has a 4xx status code
func (o *DeleteLicenseTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete license token internal server error response has a 5xx status code
func (o *DeleteLicenseTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete license token internal server error response a status code equal to that given
func (o *DeleteLicenseTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete license token internal server error response
func (o *DeleteLicenseTokenInternalServerError) Code() int {
	return 500
}

func (o *DeleteLicenseTokenInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenInternalServerError %s", 500, payload)
}

func (o *DeleteLicenseTokenInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /licensing/token][%d] deleteLicenseTokenInternalServerError %s", 500, payload)
}

func (o *DeleteLicenseTokenInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteLicenseTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
