// Code generated by go-swagger; DO NOT EDIT.

package migrations

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

// DeleteCloudMigrationTokenReader is a Reader for the DeleteCloudMigrationToken structure.
type DeleteCloudMigrationTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudMigrationTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCloudMigrationTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCloudMigrationTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCloudMigrationTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCloudMigrationTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCloudMigrationTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cloudmigration/token/{uid}] deleteCloudMigrationToken", response, response.Code())
	}
}

// NewDeleteCloudMigrationTokenNoContent creates a DeleteCloudMigrationTokenNoContent with default headers values
func NewDeleteCloudMigrationTokenNoContent() *DeleteCloudMigrationTokenNoContent {
	return &DeleteCloudMigrationTokenNoContent{}
}

/*
DeleteCloudMigrationTokenNoContent describes a response with status code 204, with default header values.

(empty)
*/
type DeleteCloudMigrationTokenNoContent struct {
}

// IsSuccess returns true when this delete cloud migration token no content response has a 2xx status code
func (o *DeleteCloudMigrationTokenNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cloud migration token no content response has a 3xx status code
func (o *DeleteCloudMigrationTokenNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cloud migration token no content response has a 4xx status code
func (o *DeleteCloudMigrationTokenNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cloud migration token no content response has a 5xx status code
func (o *DeleteCloudMigrationTokenNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cloud migration token no content response a status code equal to that given
func (o *DeleteCloudMigrationTokenNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete cloud migration token no content response
func (o *DeleteCloudMigrationTokenNoContent) Code() int {
	return 204
}

func (o *DeleteCloudMigrationTokenNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenNoContent", 204)
}

func (o *DeleteCloudMigrationTokenNoContent) String() string {
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenNoContent", 204)
}

func (o *DeleteCloudMigrationTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudMigrationTokenBadRequest creates a DeleteCloudMigrationTokenBadRequest with default headers values
func NewDeleteCloudMigrationTokenBadRequest() *DeleteCloudMigrationTokenBadRequest {
	return &DeleteCloudMigrationTokenBadRequest{}
}

/*
DeleteCloudMigrationTokenBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DeleteCloudMigrationTokenBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete cloud migration token bad request response has a 2xx status code
func (o *DeleteCloudMigrationTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cloud migration token bad request response has a 3xx status code
func (o *DeleteCloudMigrationTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cloud migration token bad request response has a 4xx status code
func (o *DeleteCloudMigrationTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cloud migration token bad request response has a 5xx status code
func (o *DeleteCloudMigrationTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cloud migration token bad request response a status code equal to that given
func (o *DeleteCloudMigrationTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete cloud migration token bad request response
func (o *DeleteCloudMigrationTokenBadRequest) Code() int {
	return 400
}

func (o *DeleteCloudMigrationTokenBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenBadRequest %s", 400, payload)
}

func (o *DeleteCloudMigrationTokenBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenBadRequest %s", 400, payload)
}

func (o *DeleteCloudMigrationTokenBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteCloudMigrationTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCloudMigrationTokenUnauthorized creates a DeleteCloudMigrationTokenUnauthorized with default headers values
func NewDeleteCloudMigrationTokenUnauthorized() *DeleteCloudMigrationTokenUnauthorized {
	return &DeleteCloudMigrationTokenUnauthorized{}
}

/*
DeleteCloudMigrationTokenUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteCloudMigrationTokenUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete cloud migration token unauthorized response has a 2xx status code
func (o *DeleteCloudMigrationTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cloud migration token unauthorized response has a 3xx status code
func (o *DeleteCloudMigrationTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cloud migration token unauthorized response has a 4xx status code
func (o *DeleteCloudMigrationTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cloud migration token unauthorized response has a 5xx status code
func (o *DeleteCloudMigrationTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cloud migration token unauthorized response a status code equal to that given
func (o *DeleteCloudMigrationTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete cloud migration token unauthorized response
func (o *DeleteCloudMigrationTokenUnauthorized) Code() int {
	return 401
}

func (o *DeleteCloudMigrationTokenUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenUnauthorized %s", 401, payload)
}

func (o *DeleteCloudMigrationTokenUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenUnauthorized %s", 401, payload)
}

func (o *DeleteCloudMigrationTokenUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteCloudMigrationTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCloudMigrationTokenForbidden creates a DeleteCloudMigrationTokenForbidden with default headers values
func NewDeleteCloudMigrationTokenForbidden() *DeleteCloudMigrationTokenForbidden {
	return &DeleteCloudMigrationTokenForbidden{}
}

/*
DeleteCloudMigrationTokenForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteCloudMigrationTokenForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete cloud migration token forbidden response has a 2xx status code
func (o *DeleteCloudMigrationTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cloud migration token forbidden response has a 3xx status code
func (o *DeleteCloudMigrationTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cloud migration token forbidden response has a 4xx status code
func (o *DeleteCloudMigrationTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cloud migration token forbidden response has a 5xx status code
func (o *DeleteCloudMigrationTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cloud migration token forbidden response a status code equal to that given
func (o *DeleteCloudMigrationTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete cloud migration token forbidden response
func (o *DeleteCloudMigrationTokenForbidden) Code() int {
	return 403
}

func (o *DeleteCloudMigrationTokenForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenForbidden %s", 403, payload)
}

func (o *DeleteCloudMigrationTokenForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenForbidden %s", 403, payload)
}

func (o *DeleteCloudMigrationTokenForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteCloudMigrationTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCloudMigrationTokenInternalServerError creates a DeleteCloudMigrationTokenInternalServerError with default headers values
func NewDeleteCloudMigrationTokenInternalServerError() *DeleteCloudMigrationTokenInternalServerError {
	return &DeleteCloudMigrationTokenInternalServerError{}
}

/*
DeleteCloudMigrationTokenInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteCloudMigrationTokenInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete cloud migration token internal server error response has a 2xx status code
func (o *DeleteCloudMigrationTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cloud migration token internal server error response has a 3xx status code
func (o *DeleteCloudMigrationTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cloud migration token internal server error response has a 4xx status code
func (o *DeleteCloudMigrationTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cloud migration token internal server error response has a 5xx status code
func (o *DeleteCloudMigrationTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete cloud migration token internal server error response a status code equal to that given
func (o *DeleteCloudMigrationTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete cloud migration token internal server error response
func (o *DeleteCloudMigrationTokenInternalServerError) Code() int {
	return 500
}

func (o *DeleteCloudMigrationTokenInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenInternalServerError %s", 500, payload)
}

func (o *DeleteCloudMigrationTokenInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloudmigration/token/{uid}][%d] deleteCloudMigrationTokenInternalServerError %s", 500, payload)
}

func (o *DeleteCloudMigrationTokenInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteCloudMigrationTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
