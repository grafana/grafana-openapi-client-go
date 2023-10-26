// Code generated by go-swagger; DO NOT EDIT.

package dashboard_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetDashboardVersionsByIDReader is a Reader for the GetDashboardVersionsByID structure.
type GetDashboardVersionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardVersionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardVersionsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDashboardVersionsByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDashboardVersionsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDashboardVersionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDashboardVersionsByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /dashboards/id/{DashboardID}/versions] getDashboardVersionsByID", response, response.Code())
	}
}

// NewGetDashboardVersionsByIDOK creates a GetDashboardVersionsByIDOK with default headers values
func NewGetDashboardVersionsByIDOK() *GetDashboardVersionsByIDOK {
	return &GetDashboardVersionsByIDOK{}
}

/*
GetDashboardVersionsByIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetDashboardVersionsByIDOK struct {
	Payload []*models.DashboardVersionMeta
}

// IsSuccess returns true when this get dashboard versions by Id Ok response has a 2xx status code
func (o *GetDashboardVersionsByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dashboard versions by Id Ok response has a 3xx status code
func (o *GetDashboardVersionsByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard versions by Id Ok response has a 4xx status code
func (o *GetDashboardVersionsByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard versions by Id Ok response has a 5xx status code
func (o *GetDashboardVersionsByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard versions by Id Ok response a status code equal to that given
func (o *GetDashboardVersionsByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get dashboard versions by Id Ok response
func (o *GetDashboardVersionsByIDOK) Code() int {
	return 200
}

func (o *GetDashboardVersionsByIDOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdOk  %+v", 200, o.Payload)
}

func (o *GetDashboardVersionsByIDOK) String() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdOk  %+v", 200, o.Payload)
}

func (o *GetDashboardVersionsByIDOK) GetPayload() []*models.DashboardVersionMeta {
	return o.Payload
}

func (o *GetDashboardVersionsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionsByIDUnauthorized creates a GetDashboardVersionsByIDUnauthorized with default headers values
func NewGetDashboardVersionsByIDUnauthorized() *GetDashboardVersionsByIDUnauthorized {
	return &GetDashboardVersionsByIDUnauthorized{}
}

/*
GetDashboardVersionsByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetDashboardVersionsByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard versions by Id unauthorized response has a 2xx status code
func (o *GetDashboardVersionsByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard versions by Id unauthorized response has a 3xx status code
func (o *GetDashboardVersionsByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard versions by Id unauthorized response has a 4xx status code
func (o *GetDashboardVersionsByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard versions by Id unauthorized response has a 5xx status code
func (o *GetDashboardVersionsByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard versions by Id unauthorized response a status code equal to that given
func (o *GetDashboardVersionsByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get dashboard versions by Id unauthorized response
func (o *GetDashboardVersionsByIDUnauthorized) Code() int {
	return 401
}

func (o *GetDashboardVersionsByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDashboardVersionsByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDashboardVersionsByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionsByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionsByIDForbidden creates a GetDashboardVersionsByIDForbidden with default headers values
func NewGetDashboardVersionsByIDForbidden() *GetDashboardVersionsByIDForbidden {
	return &GetDashboardVersionsByIDForbidden{}
}

/*
GetDashboardVersionsByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetDashboardVersionsByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard versions by Id forbidden response has a 2xx status code
func (o *GetDashboardVersionsByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard versions by Id forbidden response has a 3xx status code
func (o *GetDashboardVersionsByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard versions by Id forbidden response has a 4xx status code
func (o *GetDashboardVersionsByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard versions by Id forbidden response has a 5xx status code
func (o *GetDashboardVersionsByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard versions by Id forbidden response a status code equal to that given
func (o *GetDashboardVersionsByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get dashboard versions by Id forbidden response
func (o *GetDashboardVersionsByIDForbidden) Code() int {
	return 403
}

func (o *GetDashboardVersionsByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetDashboardVersionsByIDForbidden) String() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetDashboardVersionsByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionsByIDNotFound creates a GetDashboardVersionsByIDNotFound with default headers values
func NewGetDashboardVersionsByIDNotFound() *GetDashboardVersionsByIDNotFound {
	return &GetDashboardVersionsByIDNotFound{}
}

/*
GetDashboardVersionsByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetDashboardVersionsByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard versions by Id not found response has a 2xx status code
func (o *GetDashboardVersionsByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard versions by Id not found response has a 3xx status code
func (o *GetDashboardVersionsByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard versions by Id not found response has a 4xx status code
func (o *GetDashboardVersionsByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard versions by Id not found response has a 5xx status code
func (o *GetDashboardVersionsByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard versions by Id not found response a status code equal to that given
func (o *GetDashboardVersionsByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get dashboard versions by Id not found response
func (o *GetDashboardVersionsByIDNotFound) Code() int {
	return 404
}

func (o *GetDashboardVersionsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDashboardVersionsByIDNotFound) String() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDashboardVersionsByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionsByIDInternalServerError creates a GetDashboardVersionsByIDInternalServerError with default headers values
func NewGetDashboardVersionsByIDInternalServerError() *GetDashboardVersionsByIDInternalServerError {
	return &GetDashboardVersionsByIDInternalServerError{}
}

/*
GetDashboardVersionsByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetDashboardVersionsByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard versions by Id internal server error response has a 2xx status code
func (o *GetDashboardVersionsByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard versions by Id internal server error response has a 3xx status code
func (o *GetDashboardVersionsByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard versions by Id internal server error response has a 4xx status code
func (o *GetDashboardVersionsByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard versions by Id internal server error response has a 5xx status code
func (o *GetDashboardVersionsByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dashboard versions by Id internal server error response a status code equal to that given
func (o *GetDashboardVersionsByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get dashboard versions by Id internal server error response
func (o *GetDashboardVersionsByIDInternalServerError) Code() int {
	return 500
}

func (o *GetDashboardVersionsByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDashboardVersionsByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /dashboards/id/{DashboardID}/versions][%d] getDashboardVersionsByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDashboardVersionsByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionsByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
