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

// DeleteAlertNotificationChannelByUIDReader is a Reader for the DeleteAlertNotificationChannelByUID structure.
type DeleteAlertNotificationChannelByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertNotificationChannelByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAlertNotificationChannelByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAlertNotificationChannelByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAlertNotificationChannelByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAlertNotificationChannelByUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAlertNotificationChannelByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /alert-notifications/uid/{notification_channel_uid}] deleteAlertNotificationChannelByUID", response, response.Code())
	}
}

// NewDeleteAlertNotificationChannelByUIDOK creates a DeleteAlertNotificationChannelByUIDOK with default headers values
func NewDeleteAlertNotificationChannelByUIDOK() *DeleteAlertNotificationChannelByUIDOK {
	return &DeleteAlertNotificationChannelByUIDOK{}
}

/*
DeleteAlertNotificationChannelByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type DeleteAlertNotificationChannelByUIDOK struct {
	Payload *models.DeleteAlertNotificationChannelByUIDOKBody
}

// IsSuccess returns true when this delete alert notification channel by Uid Ok response has a 2xx status code
func (o *DeleteAlertNotificationChannelByUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete alert notification channel by Uid Ok response has a 3xx status code
func (o *DeleteAlertNotificationChannelByUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel by Uid Ok response has a 4xx status code
func (o *DeleteAlertNotificationChannelByUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alert notification channel by Uid Ok response has a 5xx status code
func (o *DeleteAlertNotificationChannelByUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert notification channel by Uid Ok response a status code equal to that given
func (o *DeleteAlertNotificationChannelByUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete alert notification channel by Uid Ok response
func (o *DeleteAlertNotificationChannelByUIDOK) Code() int {
	return 200
}

func (o *DeleteAlertNotificationChannelByUIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidOk %s", 200, payload)
}

func (o *DeleteAlertNotificationChannelByUIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidOk %s", 200, payload)
}

func (o *DeleteAlertNotificationChannelByUIDOK) GetPayload() *models.DeleteAlertNotificationChannelByUIDOKBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteAlertNotificationChannelByUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertNotificationChannelByUIDUnauthorized creates a DeleteAlertNotificationChannelByUIDUnauthorized with default headers values
func NewDeleteAlertNotificationChannelByUIDUnauthorized() *DeleteAlertNotificationChannelByUIDUnauthorized {
	return &DeleteAlertNotificationChannelByUIDUnauthorized{}
}

/*
DeleteAlertNotificationChannelByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteAlertNotificationChannelByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete alert notification channel by Uid unauthorized response has a 2xx status code
func (o *DeleteAlertNotificationChannelByUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alert notification channel by Uid unauthorized response has a 3xx status code
func (o *DeleteAlertNotificationChannelByUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel by Uid unauthorized response has a 4xx status code
func (o *DeleteAlertNotificationChannelByUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alert notification channel by Uid unauthorized response has a 5xx status code
func (o *DeleteAlertNotificationChannelByUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert notification channel by Uid unauthorized response a status code equal to that given
func (o *DeleteAlertNotificationChannelByUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete alert notification channel by Uid unauthorized response
func (o *DeleteAlertNotificationChannelByUIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteAlertNotificationChannelByUIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidUnauthorized %s", 401, payload)
}

func (o *DeleteAlertNotificationChannelByUIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidUnauthorized %s", 401, payload)
}

func (o *DeleteAlertNotificationChannelByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertNotificationChannelByUIDForbidden creates a DeleteAlertNotificationChannelByUIDForbidden with default headers values
func NewDeleteAlertNotificationChannelByUIDForbidden() *DeleteAlertNotificationChannelByUIDForbidden {
	return &DeleteAlertNotificationChannelByUIDForbidden{}
}

/*
DeleteAlertNotificationChannelByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteAlertNotificationChannelByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete alert notification channel by Uid forbidden response has a 2xx status code
func (o *DeleteAlertNotificationChannelByUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alert notification channel by Uid forbidden response has a 3xx status code
func (o *DeleteAlertNotificationChannelByUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel by Uid forbidden response has a 4xx status code
func (o *DeleteAlertNotificationChannelByUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alert notification channel by Uid forbidden response has a 5xx status code
func (o *DeleteAlertNotificationChannelByUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert notification channel by Uid forbidden response a status code equal to that given
func (o *DeleteAlertNotificationChannelByUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete alert notification channel by Uid forbidden response
func (o *DeleteAlertNotificationChannelByUIDForbidden) Code() int {
	return 403
}

func (o *DeleteAlertNotificationChannelByUIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidForbidden %s", 403, payload)
}

func (o *DeleteAlertNotificationChannelByUIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidForbidden %s", 403, payload)
}

func (o *DeleteAlertNotificationChannelByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertNotificationChannelByUIDNotFound creates a DeleteAlertNotificationChannelByUIDNotFound with default headers values
func NewDeleteAlertNotificationChannelByUIDNotFound() *DeleteAlertNotificationChannelByUIDNotFound {
	return &DeleteAlertNotificationChannelByUIDNotFound{}
}

/*
DeleteAlertNotificationChannelByUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteAlertNotificationChannelByUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete alert notification channel by Uid not found response has a 2xx status code
func (o *DeleteAlertNotificationChannelByUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alert notification channel by Uid not found response has a 3xx status code
func (o *DeleteAlertNotificationChannelByUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel by Uid not found response has a 4xx status code
func (o *DeleteAlertNotificationChannelByUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alert notification channel by Uid not found response has a 5xx status code
func (o *DeleteAlertNotificationChannelByUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert notification channel by Uid not found response a status code equal to that given
func (o *DeleteAlertNotificationChannelByUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete alert notification channel by Uid not found response
func (o *DeleteAlertNotificationChannelByUIDNotFound) Code() int {
	return 404
}

func (o *DeleteAlertNotificationChannelByUIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidNotFound %s", 404, payload)
}

func (o *DeleteAlertNotificationChannelByUIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidNotFound %s", 404, payload)
}

func (o *DeleteAlertNotificationChannelByUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelByUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertNotificationChannelByUIDInternalServerError creates a DeleteAlertNotificationChannelByUIDInternalServerError with default headers values
func NewDeleteAlertNotificationChannelByUIDInternalServerError() *DeleteAlertNotificationChannelByUIDInternalServerError {
	return &DeleteAlertNotificationChannelByUIDInternalServerError{}
}

/*
DeleteAlertNotificationChannelByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteAlertNotificationChannelByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete alert notification channel by Uid internal server error response has a 2xx status code
func (o *DeleteAlertNotificationChannelByUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alert notification channel by Uid internal server error response has a 3xx status code
func (o *DeleteAlertNotificationChannelByUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel by Uid internal server error response has a 4xx status code
func (o *DeleteAlertNotificationChannelByUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alert notification channel by Uid internal server error response has a 5xx status code
func (o *DeleteAlertNotificationChannelByUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete alert notification channel by Uid internal server error response a status code equal to that given
func (o *DeleteAlertNotificationChannelByUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete alert notification channel by Uid internal server error response
func (o *DeleteAlertNotificationChannelByUIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteAlertNotificationChannelByUIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidInternalServerError %s", 500, payload)
}

func (o *DeleteAlertNotificationChannelByUIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /alert-notifications/uid/{notification_channel_uid}][%d] deleteAlertNotificationChannelByUidInternalServerError %s", 500, payload)
}

func (o *DeleteAlertNotificationChannelByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
