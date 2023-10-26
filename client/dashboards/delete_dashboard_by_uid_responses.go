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

// DeleteDashboardByUIDReader is a Reader for the DeleteDashboardByUID structure.
type DeleteDashboardByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDashboardByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDashboardByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDashboardByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDashboardByUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDashboardByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /dashboards/uid/{uid}] deleteDashboardByUID", response, response.Code())
	}
}

// NewDeleteDashboardByUIDOK creates a DeleteDashboardByUIDOK with default headers values
func NewDeleteDashboardByUIDOK() *DeleteDashboardByUIDOK {
	return &DeleteDashboardByUIDOK{}
}

/*
DeleteDashboardByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type DeleteDashboardByUIDOK struct {
	Payload *models.DeleteDashboardByUIDOKBody
}

// IsSuccess returns true when this delete dashboard by Uid Ok response has a 2xx status code
func (o *DeleteDashboardByUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete dashboard by Uid Ok response has a 3xx status code
func (o *DeleteDashboardByUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard by Uid Ok response has a 4xx status code
func (o *DeleteDashboardByUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete dashboard by Uid Ok response has a 5xx status code
func (o *DeleteDashboardByUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard by Uid Ok response a status code equal to that given
func (o *DeleteDashboardByUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete dashboard by Uid Ok response
func (o *DeleteDashboardByUIDOK) Code() int {
	return 200
}

func (o *DeleteDashboardByUIDOK) Error() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidOk  %+v", 200, o.Payload)
}

func (o *DeleteDashboardByUIDOK) String() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidOk  %+v", 200, o.Payload)
}

func (o *DeleteDashboardByUIDOK) GetPayload() *models.DeleteDashboardByUIDOKBody {
	return o.Payload
}

func (o *DeleteDashboardByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDashboardByUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardByUIDUnauthorized creates a DeleteDashboardByUIDUnauthorized with default headers values
func NewDeleteDashboardByUIDUnauthorized() *DeleteDashboardByUIDUnauthorized {
	return &DeleteDashboardByUIDUnauthorized{}
}

/*
DeleteDashboardByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteDashboardByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete dashboard by Uid unauthorized response has a 2xx status code
func (o *DeleteDashboardByUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard by Uid unauthorized response has a 3xx status code
func (o *DeleteDashboardByUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard by Uid unauthorized response has a 4xx status code
func (o *DeleteDashboardByUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete dashboard by Uid unauthorized response has a 5xx status code
func (o *DeleteDashboardByUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard by Uid unauthorized response a status code equal to that given
func (o *DeleteDashboardByUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete dashboard by Uid unauthorized response
func (o *DeleteDashboardByUIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteDashboardByUIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDashboardByUIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDashboardByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardByUIDForbidden creates a DeleteDashboardByUIDForbidden with default headers values
func NewDeleteDashboardByUIDForbidden() *DeleteDashboardByUIDForbidden {
	return &DeleteDashboardByUIDForbidden{}
}

/*
DeleteDashboardByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteDashboardByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete dashboard by Uid forbidden response has a 2xx status code
func (o *DeleteDashboardByUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard by Uid forbidden response has a 3xx status code
func (o *DeleteDashboardByUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard by Uid forbidden response has a 4xx status code
func (o *DeleteDashboardByUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete dashboard by Uid forbidden response has a 5xx status code
func (o *DeleteDashboardByUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard by Uid forbidden response a status code equal to that given
func (o *DeleteDashboardByUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete dashboard by Uid forbidden response
func (o *DeleteDashboardByUIDForbidden) Code() int {
	return 403
}

func (o *DeleteDashboardByUIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDashboardByUIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDashboardByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardByUIDNotFound creates a DeleteDashboardByUIDNotFound with default headers values
func NewDeleteDashboardByUIDNotFound() *DeleteDashboardByUIDNotFound {
	return &DeleteDashboardByUIDNotFound{}
}

/*
DeleteDashboardByUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteDashboardByUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete dashboard by Uid not found response has a 2xx status code
func (o *DeleteDashboardByUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard by Uid not found response has a 3xx status code
func (o *DeleteDashboardByUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard by Uid not found response has a 4xx status code
func (o *DeleteDashboardByUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete dashboard by Uid not found response has a 5xx status code
func (o *DeleteDashboardByUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard by Uid not found response a status code equal to that given
func (o *DeleteDashboardByUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete dashboard by Uid not found response
func (o *DeleteDashboardByUIDNotFound) Code() int {
	return 404
}

func (o *DeleteDashboardByUIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDashboardByUIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDashboardByUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardByUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardByUIDInternalServerError creates a DeleteDashboardByUIDInternalServerError with default headers values
func NewDeleteDashboardByUIDInternalServerError() *DeleteDashboardByUIDInternalServerError {
	return &DeleteDashboardByUIDInternalServerError{}
}

/*
DeleteDashboardByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteDashboardByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete dashboard by Uid internal server error response has a 2xx status code
func (o *DeleteDashboardByUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard by Uid internal server error response has a 3xx status code
func (o *DeleteDashboardByUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard by Uid internal server error response has a 4xx status code
func (o *DeleteDashboardByUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete dashboard by Uid internal server error response has a 5xx status code
func (o *DeleteDashboardByUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete dashboard by Uid internal server error response a status code equal to that given
func (o *DeleteDashboardByUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete dashboard by Uid internal server error response
func (o *DeleteDashboardByUIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteDashboardByUIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDashboardByUIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /dashboards/uid/{uid}][%d] deleteDashboardByUidInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDashboardByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
