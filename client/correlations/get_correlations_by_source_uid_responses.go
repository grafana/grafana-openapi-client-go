// Code generated by go-swagger; DO NOT EDIT.

package correlations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetCorrelationsBySourceUIDReader is a Reader for the GetCorrelationsBySourceUID structure.
type GetCorrelationsBySourceUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorrelationsBySourceUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorrelationsBySourceUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCorrelationsBySourceUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCorrelationsBySourceUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorrelationsBySourceUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /datasources/uid/{sourceUID}/correlations] getCorrelationsBySourceUID", response, response.Code())
	}
}

// NewGetCorrelationsBySourceUIDOK creates a GetCorrelationsBySourceUIDOK with default headers values
func NewGetCorrelationsBySourceUIDOK() *GetCorrelationsBySourceUIDOK {
	return &GetCorrelationsBySourceUIDOK{}
}

/*
GetCorrelationsBySourceUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetCorrelationsBySourceUIDOK struct {
	Payload []*models.Correlation
}

// IsSuccess returns true when this get correlations by source Uid o k response has a 2xx status code
func (o *GetCorrelationsBySourceUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get correlations by source Uid o k response has a 3xx status code
func (o *GetCorrelationsBySourceUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get correlations by source Uid o k response has a 4xx status code
func (o *GetCorrelationsBySourceUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get correlations by source Uid o k response has a 5xx status code
func (o *GetCorrelationsBySourceUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get correlations by source Uid o k response a status code equal to that given
func (o *GetCorrelationsBySourceUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get correlations by source Uid o k response
func (o *GetCorrelationsBySourceUIDOK) Code() int {
	return 200
}

func (o *GetCorrelationsBySourceUIDOK) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations][%d] getCorrelationsBySourceUidOK  %+v", 200, o.Payload)
}

func (o *GetCorrelationsBySourceUIDOK) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations][%d] getCorrelationsBySourceUidOK  %+v", 200, o.Payload)
}

func (o *GetCorrelationsBySourceUIDOK) GetPayload() []*models.Correlation {
	return o.Payload
}

func (o *GetCorrelationsBySourceUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorrelationsBySourceUIDUnauthorized creates a GetCorrelationsBySourceUIDUnauthorized with default headers values
func NewGetCorrelationsBySourceUIDUnauthorized() *GetCorrelationsBySourceUIDUnauthorized {
	return &GetCorrelationsBySourceUIDUnauthorized{}
}

/*
GetCorrelationsBySourceUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetCorrelationsBySourceUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get correlations by source Uid unauthorized response has a 2xx status code
func (o *GetCorrelationsBySourceUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get correlations by source Uid unauthorized response has a 3xx status code
func (o *GetCorrelationsBySourceUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get correlations by source Uid unauthorized response has a 4xx status code
func (o *GetCorrelationsBySourceUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get correlations by source Uid unauthorized response has a 5xx status code
func (o *GetCorrelationsBySourceUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get correlations by source Uid unauthorized response a status code equal to that given
func (o *GetCorrelationsBySourceUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get correlations by source Uid unauthorized response
func (o *GetCorrelationsBySourceUIDUnauthorized) Code() int {
	return 401
}

func (o *GetCorrelationsBySourceUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations][%d] getCorrelationsBySourceUidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorrelationsBySourceUIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations][%d] getCorrelationsBySourceUidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorrelationsBySourceUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCorrelationsBySourceUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorrelationsBySourceUIDNotFound creates a GetCorrelationsBySourceUIDNotFound with default headers values
func NewGetCorrelationsBySourceUIDNotFound() *GetCorrelationsBySourceUIDNotFound {
	return &GetCorrelationsBySourceUIDNotFound{}
}

/*
GetCorrelationsBySourceUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetCorrelationsBySourceUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get correlations by source Uid not found response has a 2xx status code
func (o *GetCorrelationsBySourceUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get correlations by source Uid not found response has a 3xx status code
func (o *GetCorrelationsBySourceUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get correlations by source Uid not found response has a 4xx status code
func (o *GetCorrelationsBySourceUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get correlations by source Uid not found response has a 5xx status code
func (o *GetCorrelationsBySourceUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get correlations by source Uid not found response a status code equal to that given
func (o *GetCorrelationsBySourceUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get correlations by source Uid not found response
func (o *GetCorrelationsBySourceUIDNotFound) Code() int {
	return 404
}

func (o *GetCorrelationsBySourceUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations][%d] getCorrelationsBySourceUidNotFound  %+v", 404, o.Payload)
}

func (o *GetCorrelationsBySourceUIDNotFound) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations][%d] getCorrelationsBySourceUidNotFound  %+v", 404, o.Payload)
}

func (o *GetCorrelationsBySourceUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCorrelationsBySourceUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorrelationsBySourceUIDInternalServerError creates a GetCorrelationsBySourceUIDInternalServerError with default headers values
func NewGetCorrelationsBySourceUIDInternalServerError() *GetCorrelationsBySourceUIDInternalServerError {
	return &GetCorrelationsBySourceUIDInternalServerError{}
}

/*
GetCorrelationsBySourceUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCorrelationsBySourceUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get correlations by source Uid internal server error response has a 2xx status code
func (o *GetCorrelationsBySourceUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get correlations by source Uid internal server error response has a 3xx status code
func (o *GetCorrelationsBySourceUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get correlations by source Uid internal server error response has a 4xx status code
func (o *GetCorrelationsBySourceUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get correlations by source Uid internal server error response has a 5xx status code
func (o *GetCorrelationsBySourceUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get correlations by source Uid internal server error response a status code equal to that given
func (o *GetCorrelationsBySourceUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get correlations by source Uid internal server error response
func (o *GetCorrelationsBySourceUIDInternalServerError) Code() int {
	return 500
}

func (o *GetCorrelationsBySourceUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations][%d] getCorrelationsBySourceUidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorrelationsBySourceUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations][%d] getCorrelationsBySourceUidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorrelationsBySourceUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCorrelationsBySourceUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}