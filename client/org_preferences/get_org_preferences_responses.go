// Code generated by go-swagger; DO NOT EDIT.

package org_preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetOrgPreferencesReader is a Reader for the GetOrgPreferences structure.
type GetOrgPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrgPreferencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgPreferencesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgPreferencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /org/preferences] getOrgPreferences", response, response.Code())
	}
}

// NewGetOrgPreferencesOK creates a GetOrgPreferencesOK with default headers values
func NewGetOrgPreferencesOK() *GetOrgPreferencesOK {
	return &GetOrgPreferencesOK{}
}

/*
GetOrgPreferencesOK describes a response with status code 200, with default header values.

(empty)
*/
type GetOrgPreferencesOK struct {
	Payload *models.Preferences
}

// IsSuccess returns true when this get org preferences Ok response has a 2xx status code
func (o *GetOrgPreferencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org preferences Ok response has a 3xx status code
func (o *GetOrgPreferencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org preferences Ok response has a 4xx status code
func (o *GetOrgPreferencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org preferences Ok response has a 5xx status code
func (o *GetOrgPreferencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org preferences Ok response a status code equal to that given
func (o *GetOrgPreferencesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org preferences Ok response
func (o *GetOrgPreferencesOK) Code() int {
	return 200
}

func (o *GetOrgPreferencesOK) Error() string {
	return fmt.Sprintf("[GET /org/preferences][%d] getOrgPreferencesOk  %+v", 200, o.Payload)
}

func (o *GetOrgPreferencesOK) String() string {
	return fmt.Sprintf("[GET /org/preferences][%d] getOrgPreferencesOk  %+v", 200, o.Payload)
}

func (o *GetOrgPreferencesOK) GetPayload() *models.Preferences {
	return o.Payload
}

func (o *GetOrgPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Preferences)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgPreferencesUnauthorized creates a GetOrgPreferencesUnauthorized with default headers values
func NewGetOrgPreferencesUnauthorized() *GetOrgPreferencesUnauthorized {
	return &GetOrgPreferencesUnauthorized{}
}

/*
GetOrgPreferencesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetOrgPreferencesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org preferences unauthorized response has a 2xx status code
func (o *GetOrgPreferencesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org preferences unauthorized response has a 3xx status code
func (o *GetOrgPreferencesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org preferences unauthorized response has a 4xx status code
func (o *GetOrgPreferencesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org preferences unauthorized response has a 5xx status code
func (o *GetOrgPreferencesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get org preferences unauthorized response a status code equal to that given
func (o *GetOrgPreferencesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get org preferences unauthorized response
func (o *GetOrgPreferencesUnauthorized) Code() int {
	return 401
}

func (o *GetOrgPreferencesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /org/preferences][%d] getOrgPreferencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgPreferencesUnauthorized) String() string {
	return fmt.Sprintf("[GET /org/preferences][%d] getOrgPreferencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgPreferencesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgPreferencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgPreferencesForbidden creates a GetOrgPreferencesForbidden with default headers values
func NewGetOrgPreferencesForbidden() *GetOrgPreferencesForbidden {
	return &GetOrgPreferencesForbidden{}
}

/*
GetOrgPreferencesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetOrgPreferencesForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org preferences forbidden response has a 2xx status code
func (o *GetOrgPreferencesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org preferences forbidden response has a 3xx status code
func (o *GetOrgPreferencesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org preferences forbidden response has a 4xx status code
func (o *GetOrgPreferencesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org preferences forbidden response has a 5xx status code
func (o *GetOrgPreferencesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org preferences forbidden response a status code equal to that given
func (o *GetOrgPreferencesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org preferences forbidden response
func (o *GetOrgPreferencesForbidden) Code() int {
	return 403
}

func (o *GetOrgPreferencesForbidden) Error() string {
	return fmt.Sprintf("[GET /org/preferences][%d] getOrgPreferencesForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgPreferencesForbidden) String() string {
	return fmt.Sprintf("[GET /org/preferences][%d] getOrgPreferencesForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgPreferencesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgPreferencesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgPreferencesInternalServerError creates a GetOrgPreferencesInternalServerError with default headers values
func NewGetOrgPreferencesInternalServerError() *GetOrgPreferencesInternalServerError {
	return &GetOrgPreferencesInternalServerError{}
}

/*
GetOrgPreferencesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetOrgPreferencesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org preferences internal server error response has a 2xx status code
func (o *GetOrgPreferencesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org preferences internal server error response has a 3xx status code
func (o *GetOrgPreferencesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org preferences internal server error response has a 4xx status code
func (o *GetOrgPreferencesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org preferences internal server error response has a 5xx status code
func (o *GetOrgPreferencesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org preferences internal server error response a status code equal to that given
func (o *GetOrgPreferencesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org preferences internal server error response
func (o *GetOrgPreferencesInternalServerError) Code() int {
	return 500
}

func (o *GetOrgPreferencesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /org/preferences][%d] getOrgPreferencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgPreferencesInternalServerError) String() string {
	return fmt.Sprintf("[GET /org/preferences][%d] getOrgPreferencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgPreferencesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgPreferencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
