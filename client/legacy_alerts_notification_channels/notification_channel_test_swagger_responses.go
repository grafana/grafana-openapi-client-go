// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts_notification_channels

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

// NotificationChannelTestReader is a Reader for the NotificationChannelTest structure.
type NotificationChannelTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationChannelTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationChannelTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationChannelTestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationChannelTestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationChannelTestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewNotificationChannelTestPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationChannelTestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /alert-notifications/test] notificationChannelTest", response, response.Code())
	}
}

// NewNotificationChannelTestOK creates a NotificationChannelTestOK with default headers values
func NewNotificationChannelTestOK() *NotificationChannelTestOK {
	return &NotificationChannelTestOK{}
}

/*
NotificationChannelTestOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type NotificationChannelTestOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this notification channel test Ok response has a 2xx status code
func (o *NotificationChannelTestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notification channel test Ok response has a 3xx status code
func (o *NotificationChannelTestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification channel test Ok response has a 4xx status code
func (o *NotificationChannelTestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification channel test Ok response has a 5xx status code
func (o *NotificationChannelTestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notification channel test Ok response a status code equal to that given
func (o *NotificationChannelTestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notification channel test Ok response
func (o *NotificationChannelTestOK) Code() int {
	return 200
}

func (o *NotificationChannelTestOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestOk %s", 200, payload)
}

func (o *NotificationChannelTestOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestOk %s", 200, payload)
}

func (o *NotificationChannelTestOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *NotificationChannelTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationChannelTestBadRequest creates a NotificationChannelTestBadRequest with default headers values
func NewNotificationChannelTestBadRequest() *NotificationChannelTestBadRequest {
	return &NotificationChannelTestBadRequest{}
}

/*
NotificationChannelTestBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type NotificationChannelTestBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this notification channel test bad request response has a 2xx status code
func (o *NotificationChannelTestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification channel test bad request response has a 3xx status code
func (o *NotificationChannelTestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification channel test bad request response has a 4xx status code
func (o *NotificationChannelTestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification channel test bad request response has a 5xx status code
func (o *NotificationChannelTestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this notification channel test bad request response a status code equal to that given
func (o *NotificationChannelTestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the notification channel test bad request response
func (o *NotificationChannelTestBadRequest) Code() int {
	return 400
}

func (o *NotificationChannelTestBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestBadRequest %s", 400, payload)
}

func (o *NotificationChannelTestBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestBadRequest %s", 400, payload)
}

func (o *NotificationChannelTestBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *NotificationChannelTestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationChannelTestUnauthorized creates a NotificationChannelTestUnauthorized with default headers values
func NewNotificationChannelTestUnauthorized() *NotificationChannelTestUnauthorized {
	return &NotificationChannelTestUnauthorized{}
}

/*
NotificationChannelTestUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type NotificationChannelTestUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this notification channel test unauthorized response has a 2xx status code
func (o *NotificationChannelTestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification channel test unauthorized response has a 3xx status code
func (o *NotificationChannelTestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification channel test unauthorized response has a 4xx status code
func (o *NotificationChannelTestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification channel test unauthorized response has a 5xx status code
func (o *NotificationChannelTestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notification channel test unauthorized response a status code equal to that given
func (o *NotificationChannelTestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the notification channel test unauthorized response
func (o *NotificationChannelTestUnauthorized) Code() int {
	return 401
}

func (o *NotificationChannelTestUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestUnauthorized %s", 401, payload)
}

func (o *NotificationChannelTestUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestUnauthorized %s", 401, payload)
}

func (o *NotificationChannelTestUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *NotificationChannelTestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationChannelTestForbidden creates a NotificationChannelTestForbidden with default headers values
func NewNotificationChannelTestForbidden() *NotificationChannelTestForbidden {
	return &NotificationChannelTestForbidden{}
}

/*
NotificationChannelTestForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type NotificationChannelTestForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this notification channel test forbidden response has a 2xx status code
func (o *NotificationChannelTestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification channel test forbidden response has a 3xx status code
func (o *NotificationChannelTestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification channel test forbidden response has a 4xx status code
func (o *NotificationChannelTestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification channel test forbidden response has a 5xx status code
func (o *NotificationChannelTestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notification channel test forbidden response a status code equal to that given
func (o *NotificationChannelTestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the notification channel test forbidden response
func (o *NotificationChannelTestForbidden) Code() int {
	return 403
}

func (o *NotificationChannelTestForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestForbidden %s", 403, payload)
}

func (o *NotificationChannelTestForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestForbidden %s", 403, payload)
}

func (o *NotificationChannelTestForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *NotificationChannelTestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationChannelTestPreconditionFailed creates a NotificationChannelTestPreconditionFailed with default headers values
func NewNotificationChannelTestPreconditionFailed() *NotificationChannelTestPreconditionFailed {
	return &NotificationChannelTestPreconditionFailed{}
}

/*
NotificationChannelTestPreconditionFailed describes a response with status code 412, with default header values.

(empty)
*/
type NotificationChannelTestPreconditionFailed struct {
}

// IsSuccess returns true when this notification channel test precondition failed response has a 2xx status code
func (o *NotificationChannelTestPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification channel test precondition failed response has a 3xx status code
func (o *NotificationChannelTestPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification channel test precondition failed response has a 4xx status code
func (o *NotificationChannelTestPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification channel test precondition failed response has a 5xx status code
func (o *NotificationChannelTestPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this notification channel test precondition failed response a status code equal to that given
func (o *NotificationChannelTestPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the notification channel test precondition failed response
func (o *NotificationChannelTestPreconditionFailed) Code() int {
	return 412
}

func (o *NotificationChannelTestPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestPreconditionFailed", 412)
}

func (o *NotificationChannelTestPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestPreconditionFailed", 412)
}

func (o *NotificationChannelTestPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationChannelTestInternalServerError creates a NotificationChannelTestInternalServerError with default headers values
func NewNotificationChannelTestInternalServerError() *NotificationChannelTestInternalServerError {
	return &NotificationChannelTestInternalServerError{}
}

/*
NotificationChannelTestInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type NotificationChannelTestInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this notification channel test internal server error response has a 2xx status code
func (o *NotificationChannelTestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification channel test internal server error response has a 3xx status code
func (o *NotificationChannelTestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification channel test internal server error response has a 4xx status code
func (o *NotificationChannelTestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification channel test internal server error response has a 5xx status code
func (o *NotificationChannelTestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this notification channel test internal server error response a status code equal to that given
func (o *NotificationChannelTestInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the notification channel test internal server error response
func (o *NotificationChannelTestInternalServerError) Code() int {
	return 500
}

func (o *NotificationChannelTestInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestInternalServerError %s", 500, payload)
}

func (o *NotificationChannelTestInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /alert-notifications/test][%d] notificationChannelTestInternalServerError %s", 500, payload)
}

func (o *NotificationChannelTestInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *NotificationChannelTestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
