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

// GetCloudMigrationRunReader is a Reader for the GetCloudMigrationRun structure.
type GetCloudMigrationRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudMigrationRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudMigrationRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCloudMigrationRunUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCloudMigrationRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCloudMigrationRunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloudmigration/migration/{id}/run/{runID}] getCloudMigrationRun", response, response.Code())
	}
}

// NewGetCloudMigrationRunOK creates a GetCloudMigrationRunOK with default headers values
func NewGetCloudMigrationRunOK() *GetCloudMigrationRunOK {
	return &GetCloudMigrationRunOK{}
}

/*
GetCloudMigrationRunOK describes a response with status code 200, with default header values.

(empty)
*/
type GetCloudMigrationRunOK struct {
	Payload *models.MigrateDataResponseDTO
}

// IsSuccess returns true when this get cloud migration run Ok response has a 2xx status code
func (o *GetCloudMigrationRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cloud migration run Ok response has a 3xx status code
func (o *GetCloudMigrationRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud migration run Ok response has a 4xx status code
func (o *GetCloudMigrationRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud migration run Ok response has a 5xx status code
func (o *GetCloudMigrationRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud migration run Ok response a status code equal to that given
func (o *GetCloudMigrationRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cloud migration run Ok response
func (o *GetCloudMigrationRunOK) Code() int {
	return 200
}

func (o *GetCloudMigrationRunOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{id}/run/{runID}][%d] getCloudMigrationRunOk %s", 200, payload)
}

func (o *GetCloudMigrationRunOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{id}/run/{runID}][%d] getCloudMigrationRunOk %s", 200, payload)
}

func (o *GetCloudMigrationRunOK) GetPayload() *models.MigrateDataResponseDTO {
	return o.Payload
}

func (o *GetCloudMigrationRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MigrateDataResponseDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudMigrationRunUnauthorized creates a GetCloudMigrationRunUnauthorized with default headers values
func NewGetCloudMigrationRunUnauthorized() *GetCloudMigrationRunUnauthorized {
	return &GetCloudMigrationRunUnauthorized{}
}

/*
GetCloudMigrationRunUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetCloudMigrationRunUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get cloud migration run unauthorized response has a 2xx status code
func (o *GetCloudMigrationRunUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud migration run unauthorized response has a 3xx status code
func (o *GetCloudMigrationRunUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud migration run unauthorized response has a 4xx status code
func (o *GetCloudMigrationRunUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud migration run unauthorized response has a 5xx status code
func (o *GetCloudMigrationRunUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud migration run unauthorized response a status code equal to that given
func (o *GetCloudMigrationRunUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cloud migration run unauthorized response
func (o *GetCloudMigrationRunUnauthorized) Code() int {
	return 401
}

func (o *GetCloudMigrationRunUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{id}/run/{runID}][%d] getCloudMigrationRunUnauthorized %s", 401, payload)
}

func (o *GetCloudMigrationRunUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{id}/run/{runID}][%d] getCloudMigrationRunUnauthorized %s", 401, payload)
}

func (o *GetCloudMigrationRunUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCloudMigrationRunUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudMigrationRunForbidden creates a GetCloudMigrationRunForbidden with default headers values
func NewGetCloudMigrationRunForbidden() *GetCloudMigrationRunForbidden {
	return &GetCloudMigrationRunForbidden{}
}

/*
GetCloudMigrationRunForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetCloudMigrationRunForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get cloud migration run forbidden response has a 2xx status code
func (o *GetCloudMigrationRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud migration run forbidden response has a 3xx status code
func (o *GetCloudMigrationRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud migration run forbidden response has a 4xx status code
func (o *GetCloudMigrationRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud migration run forbidden response has a 5xx status code
func (o *GetCloudMigrationRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud migration run forbidden response a status code equal to that given
func (o *GetCloudMigrationRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cloud migration run forbidden response
func (o *GetCloudMigrationRunForbidden) Code() int {
	return 403
}

func (o *GetCloudMigrationRunForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{id}/run/{runID}][%d] getCloudMigrationRunForbidden %s", 403, payload)
}

func (o *GetCloudMigrationRunForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{id}/run/{runID}][%d] getCloudMigrationRunForbidden %s", 403, payload)
}

func (o *GetCloudMigrationRunForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCloudMigrationRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudMigrationRunInternalServerError creates a GetCloudMigrationRunInternalServerError with default headers values
func NewGetCloudMigrationRunInternalServerError() *GetCloudMigrationRunInternalServerError {
	return &GetCloudMigrationRunInternalServerError{}
}

/*
GetCloudMigrationRunInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCloudMigrationRunInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get cloud migration run internal server error response has a 2xx status code
func (o *GetCloudMigrationRunInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud migration run internal server error response has a 3xx status code
func (o *GetCloudMigrationRunInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud migration run internal server error response has a 4xx status code
func (o *GetCloudMigrationRunInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud migration run internal server error response has a 5xx status code
func (o *GetCloudMigrationRunInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cloud migration run internal server error response a status code equal to that given
func (o *GetCloudMigrationRunInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cloud migration run internal server error response
func (o *GetCloudMigrationRunInternalServerError) Code() int {
	return 500
}

func (o *GetCloudMigrationRunInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{id}/run/{runID}][%d] getCloudMigrationRunInternalServerError %s", 500, payload)
}

func (o *GetCloudMigrationRunInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cloudmigration/migration/{id}/run/{runID}][%d] getCloudMigrationRunInternalServerError %s", 500, payload)
}

func (o *GetCloudMigrationRunInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCloudMigrationRunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}