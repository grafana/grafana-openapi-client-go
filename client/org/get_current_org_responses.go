// Code generated by go-swagger; DO NOT EDIT.

package org

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

// GetCurrentOrgReader is a Reader for the GetCurrentOrg structure.
type GetCurrentOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCurrentOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /org] getCurrentOrg", response, response.Code())
	}
}

// NewGetCurrentOrgOK creates a GetCurrentOrgOK with default headers values
func NewGetCurrentOrgOK() *GetCurrentOrgOK {
	return &GetCurrentOrgOK{}
}

/*
GetCurrentOrgOK describes a response with status code 200, with default header values.

(empty)
*/
type GetCurrentOrgOK struct {
	Payload *models.OrgDetailsDTO
}

// IsSuccess returns true when this get current org Ok response has a 2xx status code
func (o *GetCurrentOrgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get current org Ok response has a 3xx status code
func (o *GetCurrentOrgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current org Ok response has a 4xx status code
func (o *GetCurrentOrgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current org Ok response has a 5xx status code
func (o *GetCurrentOrgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get current org Ok response a status code equal to that given
func (o *GetCurrentOrgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get current org Ok response
func (o *GetCurrentOrgOK) Code() int {
	return 200
}

func (o *GetCurrentOrgOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgOk %s", 200, payload)
}

func (o *GetCurrentOrgOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgOk %s", 200, payload)
}

func (o *GetCurrentOrgOK) GetPayload() *models.OrgDetailsDTO {
	return o.Payload
}

func (o *GetCurrentOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrgDetailsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgUnauthorized creates a GetCurrentOrgUnauthorized with default headers values
func NewGetCurrentOrgUnauthorized() *GetCurrentOrgUnauthorized {
	return &GetCurrentOrgUnauthorized{}
}

/*
GetCurrentOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetCurrentOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get current org unauthorized response has a 2xx status code
func (o *GetCurrentOrgUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current org unauthorized response has a 3xx status code
func (o *GetCurrentOrgUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current org unauthorized response has a 4xx status code
func (o *GetCurrentOrgUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current org unauthorized response has a 5xx status code
func (o *GetCurrentOrgUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get current org unauthorized response a status code equal to that given
func (o *GetCurrentOrgUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get current org unauthorized response
func (o *GetCurrentOrgUnauthorized) Code() int {
	return 401
}

func (o *GetCurrentOrgUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgUnauthorized %s", 401, payload)
}

func (o *GetCurrentOrgUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgUnauthorized %s", 401, payload)
}

func (o *GetCurrentOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgForbidden creates a GetCurrentOrgForbidden with default headers values
func NewGetCurrentOrgForbidden() *GetCurrentOrgForbidden {
	return &GetCurrentOrgForbidden{}
}

/*
GetCurrentOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetCurrentOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get current org forbidden response has a 2xx status code
func (o *GetCurrentOrgForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current org forbidden response has a 3xx status code
func (o *GetCurrentOrgForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current org forbidden response has a 4xx status code
func (o *GetCurrentOrgForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current org forbidden response has a 5xx status code
func (o *GetCurrentOrgForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get current org forbidden response a status code equal to that given
func (o *GetCurrentOrgForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get current org forbidden response
func (o *GetCurrentOrgForbidden) Code() int {
	return 403
}

func (o *GetCurrentOrgForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgForbidden %s", 403, payload)
}

func (o *GetCurrentOrgForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgForbidden %s", 403, payload)
}

func (o *GetCurrentOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgInternalServerError creates a GetCurrentOrgInternalServerError with default headers values
func NewGetCurrentOrgInternalServerError() *GetCurrentOrgInternalServerError {
	return &GetCurrentOrgInternalServerError{}
}

/*
GetCurrentOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCurrentOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get current org internal server error response has a 2xx status code
func (o *GetCurrentOrgInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current org internal server error response has a 3xx status code
func (o *GetCurrentOrgInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current org internal server error response has a 4xx status code
func (o *GetCurrentOrgInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current org internal server error response has a 5xx status code
func (o *GetCurrentOrgInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get current org internal server error response a status code equal to that given
func (o *GetCurrentOrgInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get current org internal server error response
func (o *GetCurrentOrgInternalServerError) Code() int {
	return 500
}

func (o *GetCurrentOrgInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgInternalServerError %s", 500, payload)
}

func (o *GetCurrentOrgInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgInternalServerError %s", 500, payload)
}

func (o *GetCurrentOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
