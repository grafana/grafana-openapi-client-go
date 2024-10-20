// Code generated by go-swagger; DO NOT EDIT.

package group_attribute_sync

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

// GetGroupMappingsReader is a Reader for the GetGroupMappings structure.
type GetGroupMappingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupMappingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupMappingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupMappingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupMappingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGroupMappingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupMappingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /groupsync/mappings] getGroupMappings", response, response.Code())
	}
}

// NewGetGroupMappingsOK creates a GetGroupMappingsOK with default headers values
func NewGetGroupMappingsOK() *GetGroupMappingsOK {
	return &GetGroupMappingsOK{}
}

/*
GetGroupMappingsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetGroupMappingsOK struct {
	Payload *models.GetGroupMappings
}

// IsSuccess returns true when this get group mappings Ok response has a 2xx status code
func (o *GetGroupMappingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get group mappings Ok response has a 3xx status code
func (o *GetGroupMappingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group mappings Ok response has a 4xx status code
func (o *GetGroupMappingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get group mappings Ok response has a 5xx status code
func (o *GetGroupMappingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get group mappings Ok response a status code equal to that given
func (o *GetGroupMappingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get group mappings Ok response
func (o *GetGroupMappingsOK) Code() int {
	return 200
}

func (o *GetGroupMappingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsOk %s", 200, payload)
}

func (o *GetGroupMappingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsOk %s", 200, payload)
}

func (o *GetGroupMappingsOK) GetPayload() *models.GetGroupMappings {
	return o.Payload
}

func (o *GetGroupMappingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetGroupMappings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupMappingsBadRequest creates a GetGroupMappingsBadRequest with default headers values
func NewGetGroupMappingsBadRequest() *GetGroupMappingsBadRequest {
	return &GetGroupMappingsBadRequest{}
}

/*
GetGroupMappingsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type GetGroupMappingsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get group mappings bad request response has a 2xx status code
func (o *GetGroupMappingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group mappings bad request response has a 3xx status code
func (o *GetGroupMappingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group mappings bad request response has a 4xx status code
func (o *GetGroupMappingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group mappings bad request response has a 5xx status code
func (o *GetGroupMappingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get group mappings bad request response a status code equal to that given
func (o *GetGroupMappingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get group mappings bad request response
func (o *GetGroupMappingsBadRequest) Code() int {
	return 400
}

func (o *GetGroupMappingsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsBadRequest %s", 400, payload)
}

func (o *GetGroupMappingsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsBadRequest %s", 400, payload)
}

func (o *GetGroupMappingsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetGroupMappingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupMappingsUnauthorized creates a GetGroupMappingsUnauthorized with default headers values
func NewGetGroupMappingsUnauthorized() *GetGroupMappingsUnauthorized {
	return &GetGroupMappingsUnauthorized{}
}

/*
GetGroupMappingsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetGroupMappingsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get group mappings unauthorized response has a 2xx status code
func (o *GetGroupMappingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group mappings unauthorized response has a 3xx status code
func (o *GetGroupMappingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group mappings unauthorized response has a 4xx status code
func (o *GetGroupMappingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group mappings unauthorized response has a 5xx status code
func (o *GetGroupMappingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get group mappings unauthorized response a status code equal to that given
func (o *GetGroupMappingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get group mappings unauthorized response
func (o *GetGroupMappingsUnauthorized) Code() int {
	return 401
}

func (o *GetGroupMappingsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsUnauthorized %s", 401, payload)
}

func (o *GetGroupMappingsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsUnauthorized %s", 401, payload)
}

func (o *GetGroupMappingsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetGroupMappingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupMappingsForbidden creates a GetGroupMappingsForbidden with default headers values
func NewGetGroupMappingsForbidden() *GetGroupMappingsForbidden {
	return &GetGroupMappingsForbidden{}
}

/*
GetGroupMappingsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetGroupMappingsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get group mappings forbidden response has a 2xx status code
func (o *GetGroupMappingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group mappings forbidden response has a 3xx status code
func (o *GetGroupMappingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group mappings forbidden response has a 4xx status code
func (o *GetGroupMappingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group mappings forbidden response has a 5xx status code
func (o *GetGroupMappingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get group mappings forbidden response a status code equal to that given
func (o *GetGroupMappingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get group mappings forbidden response
func (o *GetGroupMappingsForbidden) Code() int {
	return 403
}

func (o *GetGroupMappingsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsForbidden %s", 403, payload)
}

func (o *GetGroupMappingsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsForbidden %s", 403, payload)
}

func (o *GetGroupMappingsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetGroupMappingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupMappingsInternalServerError creates a GetGroupMappingsInternalServerError with default headers values
func NewGetGroupMappingsInternalServerError() *GetGroupMappingsInternalServerError {
	return &GetGroupMappingsInternalServerError{}
}

/*
GetGroupMappingsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetGroupMappingsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get group mappings internal server error response has a 2xx status code
func (o *GetGroupMappingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group mappings internal server error response has a 3xx status code
func (o *GetGroupMappingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group mappings internal server error response has a 4xx status code
func (o *GetGroupMappingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get group mappings internal server error response has a 5xx status code
func (o *GetGroupMappingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get group mappings internal server error response a status code equal to that given
func (o *GetGroupMappingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get group mappings internal server error response
func (o *GetGroupMappingsInternalServerError) Code() int {
	return 500
}

func (o *GetGroupMappingsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsInternalServerError %s", 500, payload)
}

func (o *GetGroupMappingsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /groupsync/mappings][%d] getGroupMappingsInternalServerError %s", 500, payload)
}

func (o *GetGroupMappingsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetGroupMappingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
