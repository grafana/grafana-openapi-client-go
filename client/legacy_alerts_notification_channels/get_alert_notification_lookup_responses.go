// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts_notification_channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetAlertNotificationLookupReader is a Reader for the GetAlertNotificationLookup structure.
type GetAlertNotificationLookupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertNotificationLookupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertNotificationLookupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAlertNotificationLookupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertNotificationLookupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertNotificationLookupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /alert-notifications/lookup] getAlertNotificationLookup", response, response.Code())
	}
}

// NewGetAlertNotificationLookupOK creates a GetAlertNotificationLookupOK with default headers values
func NewGetAlertNotificationLookupOK() *GetAlertNotificationLookupOK {
	return &GetAlertNotificationLookupOK{}
}

/*
GetAlertNotificationLookupOK describes a response with status code 200, with default header values.

(empty)
*/
type GetAlertNotificationLookupOK struct {
	Payload []*models.AlertNotificationLookup
}

// IsSuccess returns true when this get alert notification lookup Ok response has a 2xx status code
func (o *GetAlertNotificationLookupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert notification lookup Ok response has a 3xx status code
func (o *GetAlertNotificationLookupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification lookup Ok response has a 4xx status code
func (o *GetAlertNotificationLookupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert notification lookup Ok response has a 5xx status code
func (o *GetAlertNotificationLookupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert notification lookup Ok response a status code equal to that given
func (o *GetAlertNotificationLookupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert notification lookup Ok response
func (o *GetAlertNotificationLookupOK) Code() int {
	return 200
}

func (o *GetAlertNotificationLookupOK) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/lookup][%d] getAlertNotificationLookupOk  %+v", 200, o.Payload)
}

func (o *GetAlertNotificationLookupOK) String() string {
	return fmt.Sprintf("[GET /alert-notifications/lookup][%d] getAlertNotificationLookupOk  %+v", 200, o.Payload)
}

func (o *GetAlertNotificationLookupOK) GetPayload() []*models.AlertNotificationLookup {
	return o.Payload
}

func (o *GetAlertNotificationLookupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationLookupUnauthorized creates a GetAlertNotificationLookupUnauthorized with default headers values
func NewGetAlertNotificationLookupUnauthorized() *GetAlertNotificationLookupUnauthorized {
	return &GetAlertNotificationLookupUnauthorized{}
}

/*
GetAlertNotificationLookupUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetAlertNotificationLookupUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alert notification lookup unauthorized response has a 2xx status code
func (o *GetAlertNotificationLookupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert notification lookup unauthorized response has a 3xx status code
func (o *GetAlertNotificationLookupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification lookup unauthorized response has a 4xx status code
func (o *GetAlertNotificationLookupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert notification lookup unauthorized response has a 5xx status code
func (o *GetAlertNotificationLookupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert notification lookup unauthorized response a status code equal to that given
func (o *GetAlertNotificationLookupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get alert notification lookup unauthorized response
func (o *GetAlertNotificationLookupUnauthorized) Code() int {
	return 401
}

func (o *GetAlertNotificationLookupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/lookup][%d] getAlertNotificationLookupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertNotificationLookupUnauthorized) String() string {
	return fmt.Sprintf("[GET /alert-notifications/lookup][%d] getAlertNotificationLookupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertNotificationLookupUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationLookupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationLookupForbidden creates a GetAlertNotificationLookupForbidden with default headers values
func NewGetAlertNotificationLookupForbidden() *GetAlertNotificationLookupForbidden {
	return &GetAlertNotificationLookupForbidden{}
}

/*
GetAlertNotificationLookupForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetAlertNotificationLookupForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alert notification lookup forbidden response has a 2xx status code
func (o *GetAlertNotificationLookupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert notification lookup forbidden response has a 3xx status code
func (o *GetAlertNotificationLookupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification lookup forbidden response has a 4xx status code
func (o *GetAlertNotificationLookupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert notification lookup forbidden response has a 5xx status code
func (o *GetAlertNotificationLookupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert notification lookup forbidden response a status code equal to that given
func (o *GetAlertNotificationLookupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get alert notification lookup forbidden response
func (o *GetAlertNotificationLookupForbidden) Code() int {
	return 403
}

func (o *GetAlertNotificationLookupForbidden) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/lookup][%d] getAlertNotificationLookupForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertNotificationLookupForbidden) String() string {
	return fmt.Sprintf("[GET /alert-notifications/lookup][%d] getAlertNotificationLookupForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertNotificationLookupForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationLookupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationLookupInternalServerError creates a GetAlertNotificationLookupInternalServerError with default headers values
func NewGetAlertNotificationLookupInternalServerError() *GetAlertNotificationLookupInternalServerError {
	return &GetAlertNotificationLookupInternalServerError{}
}

/*
GetAlertNotificationLookupInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetAlertNotificationLookupInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alert notification lookup internal server error response has a 2xx status code
func (o *GetAlertNotificationLookupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert notification lookup internal server error response has a 3xx status code
func (o *GetAlertNotificationLookupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification lookup internal server error response has a 4xx status code
func (o *GetAlertNotificationLookupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert notification lookup internal server error response has a 5xx status code
func (o *GetAlertNotificationLookupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alert notification lookup internal server error response a status code equal to that given
func (o *GetAlertNotificationLookupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get alert notification lookup internal server error response
func (o *GetAlertNotificationLookupInternalServerError) Code() int {
	return 500
}

func (o *GetAlertNotificationLookupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/lookup][%d] getAlertNotificationLookupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertNotificationLookupInternalServerError) String() string {
	return fmt.Sprintf("[GET /alert-notifications/lookup][%d] getAlertNotificationLookupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertNotificationLookupInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationLookupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
