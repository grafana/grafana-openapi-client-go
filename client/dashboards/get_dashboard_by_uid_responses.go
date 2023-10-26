// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetDashboardByUIDReader is a Reader for the GetDashboardByUID structure.
type GetDashboardByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDashboardByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDashboardByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDashboardByUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDashboardByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /dashboards/uid/{uid}] getDashboardByUID", response, response.Code())
	}
}

// NewGetDashboardByUIDOK creates a GetDashboardByUIDOK with default headers values
func NewGetDashboardByUIDOK() *GetDashboardByUIDOK {
	return &GetDashboardByUIDOK{}
}

/*
GetDashboardByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetDashboardByUIDOK struct {
	Payload *models.DashboardFullWithMeta
}

// IsSuccess returns true when this get dashboard by Uid Ok response has a 2xx status code
func (o *GetDashboardByUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dashboard by Uid Ok response has a 3xx status code
func (o *GetDashboardByUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard by Uid Ok response has a 4xx status code
func (o *GetDashboardByUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard by Uid Ok response has a 5xx status code
func (o *GetDashboardByUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard by Uid Ok response a status code equal to that given
func (o *GetDashboardByUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get dashboard by Uid Ok response
func (o *GetDashboardByUIDOK) Code() int {
	return 200
}

func (o *GetDashboardByUIDOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidOk  %+v", 200, o.Payload)
}

func (o *GetDashboardByUIDOK) String() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidOk  %+v", 200, o.Payload)
}

func (o *GetDashboardByUIDOK) GetPayload() *models.DashboardFullWithMeta {
	return o.Payload
}

func (o *GetDashboardByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardFullWithMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardByUIDUnauthorized creates a GetDashboardByUIDUnauthorized with default headers values
func NewGetDashboardByUIDUnauthorized() *GetDashboardByUIDUnauthorized {
	return &GetDashboardByUIDUnauthorized{}
}

/*
GetDashboardByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetDashboardByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard by Uid unauthorized response has a 2xx status code
func (o *GetDashboardByUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard by Uid unauthorized response has a 3xx status code
func (o *GetDashboardByUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard by Uid unauthorized response has a 4xx status code
func (o *GetDashboardByUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard by Uid unauthorized response has a 5xx status code
func (o *GetDashboardByUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard by Uid unauthorized response a status code equal to that given
func (o *GetDashboardByUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get dashboard by Uid unauthorized response
func (o *GetDashboardByUIDUnauthorized) Code() int {
	return 401
}

func (o *GetDashboardByUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDashboardByUIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDashboardByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardByUIDForbidden creates a GetDashboardByUIDForbidden with default headers values
func NewGetDashboardByUIDForbidden() *GetDashboardByUIDForbidden {
	return &GetDashboardByUIDForbidden{}
}

/*
GetDashboardByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetDashboardByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard by Uid forbidden response has a 2xx status code
func (o *GetDashboardByUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard by Uid forbidden response has a 3xx status code
func (o *GetDashboardByUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard by Uid forbidden response has a 4xx status code
func (o *GetDashboardByUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard by Uid forbidden response has a 5xx status code
func (o *GetDashboardByUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard by Uid forbidden response a status code equal to that given
func (o *GetDashboardByUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get dashboard by Uid forbidden response
func (o *GetDashboardByUIDForbidden) Code() int {
	return 403
}

func (o *GetDashboardByUIDForbidden) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidForbidden  %+v", 403, o.Payload)
}

func (o *GetDashboardByUIDForbidden) String() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidForbidden  %+v", 403, o.Payload)
}

func (o *GetDashboardByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardByUIDNotFound creates a GetDashboardByUIDNotFound with default headers values
func NewGetDashboardByUIDNotFound() *GetDashboardByUIDNotFound {
	return &GetDashboardByUIDNotFound{}
}

/*
GetDashboardByUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetDashboardByUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard by Uid not found response has a 2xx status code
func (o *GetDashboardByUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard by Uid not found response has a 3xx status code
func (o *GetDashboardByUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard by Uid not found response has a 4xx status code
func (o *GetDashboardByUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard by Uid not found response has a 5xx status code
func (o *GetDashboardByUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard by Uid not found response a status code equal to that given
func (o *GetDashboardByUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get dashboard by Uid not found response
func (o *GetDashboardByUIDNotFound) Code() int {
	return 404
}

func (o *GetDashboardByUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidNotFound  %+v", 404, o.Payload)
}

func (o *GetDashboardByUIDNotFound) String() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidNotFound  %+v", 404, o.Payload)
}

func (o *GetDashboardByUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardByUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardByUIDInternalServerError creates a GetDashboardByUIDInternalServerError with default headers values
func NewGetDashboardByUIDInternalServerError() *GetDashboardByUIDInternalServerError {
	return &GetDashboardByUIDInternalServerError{}
}

/*
GetDashboardByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetDashboardByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard by Uid internal server error response has a 2xx status code
func (o *GetDashboardByUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard by Uid internal server error response has a 3xx status code
func (o *GetDashboardByUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard by Uid internal server error response has a 4xx status code
func (o *GetDashboardByUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard by Uid internal server error response has a 5xx status code
func (o *GetDashboardByUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dashboard by Uid internal server error response a status code equal to that given
func (o *GetDashboardByUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get dashboard by Uid internal server error response
func (o *GetDashboardByUIDInternalServerError) Code() int {
	return 500
}

func (o *GetDashboardByUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDashboardByUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}][%d] getDashboardByUidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDashboardByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
