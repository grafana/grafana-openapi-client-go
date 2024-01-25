// Code generated by go-swagger; DO NOT EDIT.

package sso_settings

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

// RemoveProviderSettingsReader is a Reader for the RemoveProviderSettings structure.
type RemoveProviderSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveProviderSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveProviderSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveProviderSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveProviderSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveProviderSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveProviderSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveProviderSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/sso-settings/{key}] removeProviderSettings", response, response.Code())
	}
}

// NewRemoveProviderSettingsNoContent creates a RemoveProviderSettingsNoContent with default headers values
func NewRemoveProviderSettingsNoContent() *RemoveProviderSettingsNoContent {
	return &RemoveProviderSettingsNoContent{}
}

/*
RemoveProviderSettingsNoContent describes a response with status code 204, with default header values.

An OKResponse is returned if the request was successful.
*/
type RemoveProviderSettingsNoContent struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this remove provider settings no content response has a 2xx status code
func (o *RemoveProviderSettingsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove provider settings no content response has a 3xx status code
func (o *RemoveProviderSettingsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove provider settings no content response has a 4xx status code
func (o *RemoveProviderSettingsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove provider settings no content response has a 5xx status code
func (o *RemoveProviderSettingsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this remove provider settings no content response a status code equal to that given
func (o *RemoveProviderSettingsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the remove provider settings no content response
func (o *RemoveProviderSettingsNoContent) Code() int {
	return 204
}

func (o *RemoveProviderSettingsNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsNoContent %s", 204, payload)
}

func (o *RemoveProviderSettingsNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsNoContent %s", 204, payload)
}

func (o *RemoveProviderSettingsNoContent) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *RemoveProviderSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveProviderSettingsBadRequest creates a RemoveProviderSettingsBadRequest with default headers values
func NewRemoveProviderSettingsBadRequest() *RemoveProviderSettingsBadRequest {
	return &RemoveProviderSettingsBadRequest{}
}

/*
RemoveProviderSettingsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RemoveProviderSettingsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove provider settings bad request response has a 2xx status code
func (o *RemoveProviderSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove provider settings bad request response has a 3xx status code
func (o *RemoveProviderSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove provider settings bad request response has a 4xx status code
func (o *RemoveProviderSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove provider settings bad request response has a 5xx status code
func (o *RemoveProviderSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove provider settings bad request response a status code equal to that given
func (o *RemoveProviderSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove provider settings bad request response
func (o *RemoveProviderSettingsBadRequest) Code() int {
	return 400
}

func (o *RemoveProviderSettingsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsBadRequest %s", 400, payload)
}

func (o *RemoveProviderSettingsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsBadRequest %s", 400, payload)
}

func (o *RemoveProviderSettingsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveProviderSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveProviderSettingsUnauthorized creates a RemoveProviderSettingsUnauthorized with default headers values
func NewRemoveProviderSettingsUnauthorized() *RemoveProviderSettingsUnauthorized {
	return &RemoveProviderSettingsUnauthorized{}
}

/*
RemoveProviderSettingsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RemoveProviderSettingsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove provider settings unauthorized response has a 2xx status code
func (o *RemoveProviderSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove provider settings unauthorized response has a 3xx status code
func (o *RemoveProviderSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove provider settings unauthorized response has a 4xx status code
func (o *RemoveProviderSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove provider settings unauthorized response has a 5xx status code
func (o *RemoveProviderSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove provider settings unauthorized response a status code equal to that given
func (o *RemoveProviderSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove provider settings unauthorized response
func (o *RemoveProviderSettingsUnauthorized) Code() int {
	return 401
}

func (o *RemoveProviderSettingsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsUnauthorized %s", 401, payload)
}

func (o *RemoveProviderSettingsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsUnauthorized %s", 401, payload)
}

func (o *RemoveProviderSettingsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveProviderSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveProviderSettingsForbidden creates a RemoveProviderSettingsForbidden with default headers values
func NewRemoveProviderSettingsForbidden() *RemoveProviderSettingsForbidden {
	return &RemoveProviderSettingsForbidden{}
}

/*
RemoveProviderSettingsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RemoveProviderSettingsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove provider settings forbidden response has a 2xx status code
func (o *RemoveProviderSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove provider settings forbidden response has a 3xx status code
func (o *RemoveProviderSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove provider settings forbidden response has a 4xx status code
func (o *RemoveProviderSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove provider settings forbidden response has a 5xx status code
func (o *RemoveProviderSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove provider settings forbidden response a status code equal to that given
func (o *RemoveProviderSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove provider settings forbidden response
func (o *RemoveProviderSettingsForbidden) Code() int {
	return 403
}

func (o *RemoveProviderSettingsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsForbidden %s", 403, payload)
}

func (o *RemoveProviderSettingsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsForbidden %s", 403, payload)
}

func (o *RemoveProviderSettingsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveProviderSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveProviderSettingsNotFound creates a RemoveProviderSettingsNotFound with default headers values
func NewRemoveProviderSettingsNotFound() *RemoveProviderSettingsNotFound {
	return &RemoveProviderSettingsNotFound{}
}

/*
RemoveProviderSettingsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RemoveProviderSettingsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove provider settings not found response has a 2xx status code
func (o *RemoveProviderSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove provider settings not found response has a 3xx status code
func (o *RemoveProviderSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove provider settings not found response has a 4xx status code
func (o *RemoveProviderSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove provider settings not found response has a 5xx status code
func (o *RemoveProviderSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove provider settings not found response a status code equal to that given
func (o *RemoveProviderSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove provider settings not found response
func (o *RemoveProviderSettingsNotFound) Code() int {
	return 404
}

func (o *RemoveProviderSettingsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsNotFound %s", 404, payload)
}

func (o *RemoveProviderSettingsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsNotFound %s", 404, payload)
}

func (o *RemoveProviderSettingsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveProviderSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveProviderSettingsInternalServerError creates a RemoveProviderSettingsInternalServerError with default headers values
func NewRemoveProviderSettingsInternalServerError() *RemoveProviderSettingsInternalServerError {
	return &RemoveProviderSettingsInternalServerError{}
}

/*
RemoveProviderSettingsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RemoveProviderSettingsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove provider settings internal server error response has a 2xx status code
func (o *RemoveProviderSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove provider settings internal server error response has a 3xx status code
func (o *RemoveProviderSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove provider settings internal server error response has a 4xx status code
func (o *RemoveProviderSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove provider settings internal server error response has a 5xx status code
func (o *RemoveProviderSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove provider settings internal server error response a status code equal to that given
func (o *RemoveProviderSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove provider settings internal server error response
func (o *RemoveProviderSettingsInternalServerError) Code() int {
	return 500
}

func (o *RemoveProviderSettingsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsInternalServerError %s", 500, payload)
}

func (o *RemoveProviderSettingsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/sso-settings/{key}][%d] removeProviderSettingsInternalServerError %s", 500, payload)
}

func (o *RemoveProviderSettingsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveProviderSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
