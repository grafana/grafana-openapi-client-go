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

// RunCloudMigrationReader is a Reader for the RunCloudMigration structure.
type RunCloudMigrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunCloudMigrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunCloudMigrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunCloudMigrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRunCloudMigrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRunCloudMigrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRunCloudMigrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloudmigration/migration/{uid}/run] runCloudMigration", response, response.Code())
	}
}

// NewRunCloudMigrationOK creates a RunCloudMigrationOK with default headers values
func NewRunCloudMigrationOK() *RunCloudMigrationOK {
	return &RunCloudMigrationOK{}
}

/*
RunCloudMigrationOK describes a response with status code 200, with default header values.

(empty)
*/
type RunCloudMigrationOK struct {
	Payload *models.MigrateDataResponseDTO
}

// IsSuccess returns true when this run cloud migration Ok response has a 2xx status code
func (o *RunCloudMigrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this run cloud migration Ok response has a 3xx status code
func (o *RunCloudMigrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run cloud migration Ok response has a 4xx status code
func (o *RunCloudMigrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this run cloud migration Ok response has a 5xx status code
func (o *RunCloudMigrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this run cloud migration Ok response a status code equal to that given
func (o *RunCloudMigrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the run cloud migration Ok response
func (o *RunCloudMigrationOK) Code() int {
	return 200
}

func (o *RunCloudMigrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationOk %s", 200, payload)
}

func (o *RunCloudMigrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationOk %s", 200, payload)
}

func (o *RunCloudMigrationOK) GetPayload() *models.MigrateDataResponseDTO {
	return o.Payload
}

func (o *RunCloudMigrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MigrateDataResponseDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunCloudMigrationBadRequest creates a RunCloudMigrationBadRequest with default headers values
func NewRunCloudMigrationBadRequest() *RunCloudMigrationBadRequest {
	return &RunCloudMigrationBadRequest{}
}

/*
RunCloudMigrationBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RunCloudMigrationBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this run cloud migration bad request response has a 2xx status code
func (o *RunCloudMigrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this run cloud migration bad request response has a 3xx status code
func (o *RunCloudMigrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run cloud migration bad request response has a 4xx status code
func (o *RunCloudMigrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this run cloud migration bad request response has a 5xx status code
func (o *RunCloudMigrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this run cloud migration bad request response a status code equal to that given
func (o *RunCloudMigrationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the run cloud migration bad request response
func (o *RunCloudMigrationBadRequest) Code() int {
	return 400
}

func (o *RunCloudMigrationBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationBadRequest %s", 400, payload)
}

func (o *RunCloudMigrationBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationBadRequest %s", 400, payload)
}

func (o *RunCloudMigrationBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RunCloudMigrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunCloudMigrationUnauthorized creates a RunCloudMigrationUnauthorized with default headers values
func NewRunCloudMigrationUnauthorized() *RunCloudMigrationUnauthorized {
	return &RunCloudMigrationUnauthorized{}
}

/*
RunCloudMigrationUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RunCloudMigrationUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this run cloud migration unauthorized response has a 2xx status code
func (o *RunCloudMigrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this run cloud migration unauthorized response has a 3xx status code
func (o *RunCloudMigrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run cloud migration unauthorized response has a 4xx status code
func (o *RunCloudMigrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this run cloud migration unauthorized response has a 5xx status code
func (o *RunCloudMigrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this run cloud migration unauthorized response a status code equal to that given
func (o *RunCloudMigrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the run cloud migration unauthorized response
func (o *RunCloudMigrationUnauthorized) Code() int {
	return 401
}

func (o *RunCloudMigrationUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationUnauthorized %s", 401, payload)
}

func (o *RunCloudMigrationUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationUnauthorized %s", 401, payload)
}

func (o *RunCloudMigrationUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RunCloudMigrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunCloudMigrationForbidden creates a RunCloudMigrationForbidden with default headers values
func NewRunCloudMigrationForbidden() *RunCloudMigrationForbidden {
	return &RunCloudMigrationForbidden{}
}

/*
RunCloudMigrationForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RunCloudMigrationForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this run cloud migration forbidden response has a 2xx status code
func (o *RunCloudMigrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this run cloud migration forbidden response has a 3xx status code
func (o *RunCloudMigrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run cloud migration forbidden response has a 4xx status code
func (o *RunCloudMigrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this run cloud migration forbidden response has a 5xx status code
func (o *RunCloudMigrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this run cloud migration forbidden response a status code equal to that given
func (o *RunCloudMigrationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the run cloud migration forbidden response
func (o *RunCloudMigrationForbidden) Code() int {
	return 403
}

func (o *RunCloudMigrationForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationForbidden %s", 403, payload)
}

func (o *RunCloudMigrationForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationForbidden %s", 403, payload)
}

func (o *RunCloudMigrationForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RunCloudMigrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunCloudMigrationInternalServerError creates a RunCloudMigrationInternalServerError with default headers values
func NewRunCloudMigrationInternalServerError() *RunCloudMigrationInternalServerError {
	return &RunCloudMigrationInternalServerError{}
}

/*
RunCloudMigrationInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RunCloudMigrationInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this run cloud migration internal server error response has a 2xx status code
func (o *RunCloudMigrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this run cloud migration internal server error response has a 3xx status code
func (o *RunCloudMigrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run cloud migration internal server error response has a 4xx status code
func (o *RunCloudMigrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this run cloud migration internal server error response has a 5xx status code
func (o *RunCloudMigrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this run cloud migration internal server error response a status code equal to that given
func (o *RunCloudMigrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the run cloud migration internal server error response
func (o *RunCloudMigrationInternalServerError) Code() int {
	return 500
}

func (o *RunCloudMigrationInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationInternalServerError %s", 500, payload)
}

func (o *RunCloudMigrationInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloudmigration/migration/{uid}/run][%d] runCloudMigrationInternalServerError %s", 500, payload)
}

func (o *RunCloudMigrationInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RunCloudMigrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
