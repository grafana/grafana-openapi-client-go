// Code generated by go-swagger; DO NOT EDIT.

package org_invites

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

// GetPendingOrgInvitesReader is a Reader for the GetPendingOrgInvites structure.
type GetPendingOrgInvitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPendingOrgInvitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPendingOrgInvitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPendingOrgInvitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPendingOrgInvitesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPendingOrgInvitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /org/invites] getPendingOrgInvites", response, response.Code())
	}
}

// NewGetPendingOrgInvitesOK creates a GetPendingOrgInvitesOK with default headers values
func NewGetPendingOrgInvitesOK() *GetPendingOrgInvitesOK {
	return &GetPendingOrgInvitesOK{}
}

/*
GetPendingOrgInvitesOK describes a response with status code 200, with default header values.

(empty)
*/
type GetPendingOrgInvitesOK struct {
	Payload []*models.TempUserDTO
}

// IsSuccess returns true when this get pending org invites Ok response has a 2xx status code
func (o *GetPendingOrgInvitesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pending org invites Ok response has a 3xx status code
func (o *GetPendingOrgInvitesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pending org invites Ok response has a 4xx status code
func (o *GetPendingOrgInvitesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pending org invites Ok response has a 5xx status code
func (o *GetPendingOrgInvitesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pending org invites Ok response a status code equal to that given
func (o *GetPendingOrgInvitesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get pending org invites Ok response
func (o *GetPendingOrgInvitesOK) Code() int {
	return 200
}

func (o *GetPendingOrgInvitesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org/invites][%d] getPendingOrgInvitesOk %s", 200, payload)
}

func (o *GetPendingOrgInvitesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org/invites][%d] getPendingOrgInvitesOk %s", 200, payload)
}

func (o *GetPendingOrgInvitesOK) GetPayload() []*models.TempUserDTO {
	return o.Payload
}

func (o *GetPendingOrgInvitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPendingOrgInvitesUnauthorized creates a GetPendingOrgInvitesUnauthorized with default headers values
func NewGetPendingOrgInvitesUnauthorized() *GetPendingOrgInvitesUnauthorized {
	return &GetPendingOrgInvitesUnauthorized{}
}

/*
GetPendingOrgInvitesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetPendingOrgInvitesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get pending org invites unauthorized response has a 2xx status code
func (o *GetPendingOrgInvitesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pending org invites unauthorized response has a 3xx status code
func (o *GetPendingOrgInvitesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pending org invites unauthorized response has a 4xx status code
func (o *GetPendingOrgInvitesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pending org invites unauthorized response has a 5xx status code
func (o *GetPendingOrgInvitesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get pending org invites unauthorized response a status code equal to that given
func (o *GetPendingOrgInvitesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get pending org invites unauthorized response
func (o *GetPendingOrgInvitesUnauthorized) Code() int {
	return 401
}

func (o *GetPendingOrgInvitesUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org/invites][%d] getPendingOrgInvitesUnauthorized %s", 401, payload)
}

func (o *GetPendingOrgInvitesUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org/invites][%d] getPendingOrgInvitesUnauthorized %s", 401, payload)
}

func (o *GetPendingOrgInvitesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetPendingOrgInvitesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPendingOrgInvitesForbidden creates a GetPendingOrgInvitesForbidden with default headers values
func NewGetPendingOrgInvitesForbidden() *GetPendingOrgInvitesForbidden {
	return &GetPendingOrgInvitesForbidden{}
}

/*
GetPendingOrgInvitesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetPendingOrgInvitesForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get pending org invites forbidden response has a 2xx status code
func (o *GetPendingOrgInvitesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pending org invites forbidden response has a 3xx status code
func (o *GetPendingOrgInvitesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pending org invites forbidden response has a 4xx status code
func (o *GetPendingOrgInvitesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pending org invites forbidden response has a 5xx status code
func (o *GetPendingOrgInvitesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get pending org invites forbidden response a status code equal to that given
func (o *GetPendingOrgInvitesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get pending org invites forbidden response
func (o *GetPendingOrgInvitesForbidden) Code() int {
	return 403
}

func (o *GetPendingOrgInvitesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org/invites][%d] getPendingOrgInvitesForbidden %s", 403, payload)
}

func (o *GetPendingOrgInvitesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org/invites][%d] getPendingOrgInvitesForbidden %s", 403, payload)
}

func (o *GetPendingOrgInvitesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetPendingOrgInvitesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPendingOrgInvitesInternalServerError creates a GetPendingOrgInvitesInternalServerError with default headers values
func NewGetPendingOrgInvitesInternalServerError() *GetPendingOrgInvitesInternalServerError {
	return &GetPendingOrgInvitesInternalServerError{}
}

/*
GetPendingOrgInvitesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetPendingOrgInvitesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get pending org invites internal server error response has a 2xx status code
func (o *GetPendingOrgInvitesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pending org invites internal server error response has a 3xx status code
func (o *GetPendingOrgInvitesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pending org invites internal server error response has a 4xx status code
func (o *GetPendingOrgInvitesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pending org invites internal server error response has a 5xx status code
func (o *GetPendingOrgInvitesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get pending org invites internal server error response a status code equal to that given
func (o *GetPendingOrgInvitesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get pending org invites internal server error response
func (o *GetPendingOrgInvitesInternalServerError) Code() int {
	return 500
}

func (o *GetPendingOrgInvitesInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org/invites][%d] getPendingOrgInvitesInternalServerError %s", 500, payload)
}

func (o *GetPendingOrgInvitesInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org/invites][%d] getPendingOrgInvitesInternalServerError %s", 500, payload)
}

func (o *GetPendingOrgInvitesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetPendingOrgInvitesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
