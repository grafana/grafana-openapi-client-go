// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// PauseAllAlertsReader is a Reader for the PauseAllAlerts structure.
type PauseAllAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseAllAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPauseAllAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPauseAllAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPauseAllAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPauseAllAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/pause-all-alerts] pauseAllAlerts", response, response.Code())
	}
}

// NewPauseAllAlertsOK creates a PauseAllAlertsOK with default headers values
func NewPauseAllAlertsOK() *PauseAllAlertsOK {
	return &PauseAllAlertsOK{}
}

/*
PauseAllAlertsOK describes a response with status code 200, with default header values.

(empty)
*/
type PauseAllAlertsOK struct {
	Payload *models.PauseAllAlertsOKBody
}

// IsSuccess returns true when this pause all alerts Ok response has a 2xx status code
func (o *PauseAllAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pause all alerts Ok response has a 3xx status code
func (o *PauseAllAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause all alerts Ok response has a 4xx status code
func (o *PauseAllAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pause all alerts Ok response has a 5xx status code
func (o *PauseAllAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pause all alerts Ok response a status code equal to that given
func (o *PauseAllAlertsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pause all alerts Ok response
func (o *PauseAllAlertsOK) Code() int {
	return 200
}

func (o *PauseAllAlertsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/pause-all-alerts][%d] pauseAllAlertsOk %s", 200, payload)
}

func (o *PauseAllAlertsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/pause-all-alerts][%d] pauseAllAlertsOk %s", 200, payload)
}

func (o *PauseAllAlertsOK) GetPayload() *models.PauseAllAlertsOKBody {
	return o.Payload
}

func (o *PauseAllAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PauseAllAlertsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseAllAlertsUnauthorized creates a PauseAllAlertsUnauthorized with default headers values
func NewPauseAllAlertsUnauthorized() *PauseAllAlertsUnauthorized {
	return &PauseAllAlertsUnauthorized{}
}

/*
PauseAllAlertsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type PauseAllAlertsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this pause all alerts unauthorized response has a 2xx status code
func (o *PauseAllAlertsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause all alerts unauthorized response has a 3xx status code
func (o *PauseAllAlertsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause all alerts unauthorized response has a 4xx status code
func (o *PauseAllAlertsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pause all alerts unauthorized response has a 5xx status code
func (o *PauseAllAlertsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pause all alerts unauthorized response a status code equal to that given
func (o *PauseAllAlertsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pause all alerts unauthorized response
func (o *PauseAllAlertsUnauthorized) Code() int {
	return 401
}

func (o *PauseAllAlertsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/pause-all-alerts][%d] pauseAllAlertsUnauthorized %s", 401, payload)
}

func (o *PauseAllAlertsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/pause-all-alerts][%d] pauseAllAlertsUnauthorized %s", 401, payload)
}

func (o *PauseAllAlertsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PauseAllAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseAllAlertsForbidden creates a PauseAllAlertsForbidden with default headers values
func NewPauseAllAlertsForbidden() *PauseAllAlertsForbidden {
	return &PauseAllAlertsForbidden{}
}

/*
PauseAllAlertsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type PauseAllAlertsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this pause all alerts forbidden response has a 2xx status code
func (o *PauseAllAlertsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause all alerts forbidden response has a 3xx status code
func (o *PauseAllAlertsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause all alerts forbidden response has a 4xx status code
func (o *PauseAllAlertsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pause all alerts forbidden response has a 5xx status code
func (o *PauseAllAlertsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pause all alerts forbidden response a status code equal to that given
func (o *PauseAllAlertsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pause all alerts forbidden response
func (o *PauseAllAlertsForbidden) Code() int {
	return 403
}

func (o *PauseAllAlertsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/pause-all-alerts][%d] pauseAllAlertsForbidden %s", 403, payload)
}

func (o *PauseAllAlertsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/pause-all-alerts][%d] pauseAllAlertsForbidden %s", 403, payload)
}

func (o *PauseAllAlertsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PauseAllAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseAllAlertsInternalServerError creates a PauseAllAlertsInternalServerError with default headers values
func NewPauseAllAlertsInternalServerError() *PauseAllAlertsInternalServerError {
	return &PauseAllAlertsInternalServerError{}
}

/*
PauseAllAlertsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PauseAllAlertsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this pause all alerts internal server error response has a 2xx status code
func (o *PauseAllAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause all alerts internal server error response has a 3xx status code
func (o *PauseAllAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause all alerts internal server error response has a 4xx status code
func (o *PauseAllAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pause all alerts internal server error response has a 5xx status code
func (o *PauseAllAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pause all alerts internal server error response a status code equal to that given
func (o *PauseAllAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pause all alerts internal server error response
func (o *PauseAllAlertsInternalServerError) Code() int {
	return 500
}

func (o *PauseAllAlertsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/pause-all-alerts][%d] pauseAllAlertsInternalServerError %s", 500, payload)
}

func (o *PauseAllAlertsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/pause-all-alerts][%d] pauseAllAlertsInternalServerError %s", 500, payload)
}

func (o *PauseAllAlertsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PauseAllAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
