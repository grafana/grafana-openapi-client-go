// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// SaveReportSettingsReader is a Reader for the SaveReportSettings structure.
type SaveReportSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveReportSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSaveReportSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSaveReportSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSaveReportSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSaveReportSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSaveReportSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /reports/settings] saveReportSettings", response, response.Code())
	}
}

// NewSaveReportSettingsOK creates a SaveReportSettingsOK with default headers values
func NewSaveReportSettingsOK() *SaveReportSettingsOK {
	return &SaveReportSettingsOK{}
}

/*
SaveReportSettingsOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type SaveReportSettingsOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this save report settings Ok response has a 2xx status code
func (o *SaveReportSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this save report settings Ok response has a 3xx status code
func (o *SaveReportSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save report settings Ok response has a 4xx status code
func (o *SaveReportSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this save report settings Ok response has a 5xx status code
func (o *SaveReportSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this save report settings Ok response a status code equal to that given
func (o *SaveReportSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the save report settings Ok response
func (o *SaveReportSettingsOK) Code() int {
	return 200
}

func (o *SaveReportSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsOk %s", 200, payload)
}

func (o *SaveReportSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsOk %s", 200, payload)
}

func (o *SaveReportSettingsOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *SaveReportSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveReportSettingsBadRequest creates a SaveReportSettingsBadRequest with default headers values
func NewSaveReportSettingsBadRequest() *SaveReportSettingsBadRequest {
	return &SaveReportSettingsBadRequest{}
}

/*
SaveReportSettingsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type SaveReportSettingsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this save report settings bad request response has a 2xx status code
func (o *SaveReportSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save report settings bad request response has a 3xx status code
func (o *SaveReportSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save report settings bad request response has a 4xx status code
func (o *SaveReportSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this save report settings bad request response has a 5xx status code
func (o *SaveReportSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this save report settings bad request response a status code equal to that given
func (o *SaveReportSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the save report settings bad request response
func (o *SaveReportSettingsBadRequest) Code() int {
	return 400
}

func (o *SaveReportSettingsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsBadRequest %s", 400, payload)
}

func (o *SaveReportSettingsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsBadRequest %s", 400, payload)
}

func (o *SaveReportSettingsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SaveReportSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveReportSettingsUnauthorized creates a SaveReportSettingsUnauthorized with default headers values
func NewSaveReportSettingsUnauthorized() *SaveReportSettingsUnauthorized {
	return &SaveReportSettingsUnauthorized{}
}

/*
SaveReportSettingsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type SaveReportSettingsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this save report settings unauthorized response has a 2xx status code
func (o *SaveReportSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save report settings unauthorized response has a 3xx status code
func (o *SaveReportSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save report settings unauthorized response has a 4xx status code
func (o *SaveReportSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this save report settings unauthorized response has a 5xx status code
func (o *SaveReportSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this save report settings unauthorized response a status code equal to that given
func (o *SaveReportSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the save report settings unauthorized response
func (o *SaveReportSettingsUnauthorized) Code() int {
	return 401
}

func (o *SaveReportSettingsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsUnauthorized %s", 401, payload)
}

func (o *SaveReportSettingsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsUnauthorized %s", 401, payload)
}

func (o *SaveReportSettingsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SaveReportSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveReportSettingsForbidden creates a SaveReportSettingsForbidden with default headers values
func NewSaveReportSettingsForbidden() *SaveReportSettingsForbidden {
	return &SaveReportSettingsForbidden{}
}

/*
SaveReportSettingsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SaveReportSettingsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this save report settings forbidden response has a 2xx status code
func (o *SaveReportSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save report settings forbidden response has a 3xx status code
func (o *SaveReportSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save report settings forbidden response has a 4xx status code
func (o *SaveReportSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this save report settings forbidden response has a 5xx status code
func (o *SaveReportSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this save report settings forbidden response a status code equal to that given
func (o *SaveReportSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the save report settings forbidden response
func (o *SaveReportSettingsForbidden) Code() int {
	return 403
}

func (o *SaveReportSettingsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsForbidden %s", 403, payload)
}

func (o *SaveReportSettingsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsForbidden %s", 403, payload)
}

func (o *SaveReportSettingsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SaveReportSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveReportSettingsInternalServerError creates a SaveReportSettingsInternalServerError with default headers values
func NewSaveReportSettingsInternalServerError() *SaveReportSettingsInternalServerError {
	return &SaveReportSettingsInternalServerError{}
}

/*
SaveReportSettingsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SaveReportSettingsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this save report settings internal server error response has a 2xx status code
func (o *SaveReportSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save report settings internal server error response has a 3xx status code
func (o *SaveReportSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save report settings internal server error response has a 4xx status code
func (o *SaveReportSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this save report settings internal server error response has a 5xx status code
func (o *SaveReportSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this save report settings internal server error response a status code equal to that given
func (o *SaveReportSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the save report settings internal server error response
func (o *SaveReportSettingsInternalServerError) Code() int {
	return 500
}

func (o *SaveReportSettingsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsInternalServerError %s", 500, payload)
}

func (o *SaveReportSettingsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /reports/settings][%d] saveReportSettingsInternalServerError %s", 500, payload)
}

func (o *SaveReportSettingsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SaveReportSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
