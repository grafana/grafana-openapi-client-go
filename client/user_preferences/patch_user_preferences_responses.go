// Code generated by go-swagger; DO NOT EDIT.

package user_preferences

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

// PatchUserPreferencesReader is a Reader for the PatchUserPreferences structure.
type PatchUserPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUserPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUserPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchUserPreferencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchUserPreferencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUserPreferencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /user/preferences] patchUserPreferences", response, response.Code())
	}
}

// NewPatchUserPreferencesOK creates a PatchUserPreferencesOK with default headers values
func NewPatchUserPreferencesOK() *PatchUserPreferencesOK {
	return &PatchUserPreferencesOK{}
}

/*
PatchUserPreferencesOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type PatchUserPreferencesOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this patch user preferences Ok response has a 2xx status code
func (o *PatchUserPreferencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch user preferences Ok response has a 3xx status code
func (o *PatchUserPreferencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user preferences Ok response has a 4xx status code
func (o *PatchUserPreferencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch user preferences Ok response has a 5xx status code
func (o *PatchUserPreferencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user preferences Ok response a status code equal to that given
func (o *PatchUserPreferencesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch user preferences Ok response
func (o *PatchUserPreferencesOK) Code() int {
	return 200
}

func (o *PatchUserPreferencesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /user/preferences][%d] patchUserPreferencesOk %s", 200, payload)
}

func (o *PatchUserPreferencesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /user/preferences][%d] patchUserPreferencesOk %s", 200, payload)
}

func (o *PatchUserPreferencesOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *PatchUserPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserPreferencesBadRequest creates a PatchUserPreferencesBadRequest with default headers values
func NewPatchUserPreferencesBadRequest() *PatchUserPreferencesBadRequest {
	return &PatchUserPreferencesBadRequest{}
}

/*
PatchUserPreferencesBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type PatchUserPreferencesBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this patch user preferences bad request response has a 2xx status code
func (o *PatchUserPreferencesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user preferences bad request response has a 3xx status code
func (o *PatchUserPreferencesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user preferences bad request response has a 4xx status code
func (o *PatchUserPreferencesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user preferences bad request response has a 5xx status code
func (o *PatchUserPreferencesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user preferences bad request response a status code equal to that given
func (o *PatchUserPreferencesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch user preferences bad request response
func (o *PatchUserPreferencesBadRequest) Code() int {
	return 400
}

func (o *PatchUserPreferencesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /user/preferences][%d] patchUserPreferencesBadRequest %s", 400, payload)
}

func (o *PatchUserPreferencesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /user/preferences][%d] patchUserPreferencesBadRequest %s", 400, payload)
}

func (o *PatchUserPreferencesBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchUserPreferencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserPreferencesUnauthorized creates a PatchUserPreferencesUnauthorized with default headers values
func NewPatchUserPreferencesUnauthorized() *PatchUserPreferencesUnauthorized {
	return &PatchUserPreferencesUnauthorized{}
}

/*
PatchUserPreferencesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type PatchUserPreferencesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this patch user preferences unauthorized response has a 2xx status code
func (o *PatchUserPreferencesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user preferences unauthorized response has a 3xx status code
func (o *PatchUserPreferencesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user preferences unauthorized response has a 4xx status code
func (o *PatchUserPreferencesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user preferences unauthorized response has a 5xx status code
func (o *PatchUserPreferencesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user preferences unauthorized response a status code equal to that given
func (o *PatchUserPreferencesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch user preferences unauthorized response
func (o *PatchUserPreferencesUnauthorized) Code() int {
	return 401
}

func (o *PatchUserPreferencesUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /user/preferences][%d] patchUserPreferencesUnauthorized %s", 401, payload)
}

func (o *PatchUserPreferencesUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /user/preferences][%d] patchUserPreferencesUnauthorized %s", 401, payload)
}

func (o *PatchUserPreferencesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchUserPreferencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserPreferencesInternalServerError creates a PatchUserPreferencesInternalServerError with default headers values
func NewPatchUserPreferencesInternalServerError() *PatchUserPreferencesInternalServerError {
	return &PatchUserPreferencesInternalServerError{}
}

/*
PatchUserPreferencesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PatchUserPreferencesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this patch user preferences internal server error response has a 2xx status code
func (o *PatchUserPreferencesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user preferences internal server error response has a 3xx status code
func (o *PatchUserPreferencesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user preferences internal server error response has a 4xx status code
func (o *PatchUserPreferencesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch user preferences internal server error response has a 5xx status code
func (o *PatchUserPreferencesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch user preferences internal server error response a status code equal to that given
func (o *PatchUserPreferencesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch user preferences internal server error response
func (o *PatchUserPreferencesInternalServerError) Code() int {
	return 500
}

func (o *PatchUserPreferencesInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /user/preferences][%d] patchUserPreferencesInternalServerError %s", 500, payload)
}

func (o *PatchUserPreferencesInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /user/preferences][%d] patchUserPreferencesInternalServerError %s", 500, payload)
}

func (o *PatchUserPreferencesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchUserPreferencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
