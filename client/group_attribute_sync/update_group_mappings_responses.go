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

// UpdateGroupMappingsReader is a Reader for the UpdateGroupMappings structure.
type UpdateGroupMappingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGroupMappingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateGroupMappingsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGroupMappingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateGroupMappingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGroupMappingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateGroupMappingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /groupsync/groups/{group_id}] updateGroupMappings", response, response.Code())
	}
}

// NewUpdateGroupMappingsCreated creates a UpdateGroupMappingsCreated with default headers values
func NewUpdateGroupMappingsCreated() *UpdateGroupMappingsCreated {
	return &UpdateGroupMappingsCreated{}
}

/*
UpdateGroupMappingsCreated describes a response with status code 201, with default header values.

(empty)
*/
type UpdateGroupMappingsCreated struct {
	Payload *models.MessageResponse
}

// IsSuccess returns true when this update group mappings created response has a 2xx status code
func (o *UpdateGroupMappingsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update group mappings created response has a 3xx status code
func (o *UpdateGroupMappingsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group mappings created response has a 4xx status code
func (o *UpdateGroupMappingsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update group mappings created response has a 5xx status code
func (o *UpdateGroupMappingsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update group mappings created response a status code equal to that given
func (o *UpdateGroupMappingsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update group mappings created response
func (o *UpdateGroupMappingsCreated) Code() int {
	return 201
}

func (o *UpdateGroupMappingsCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsCreated %s", 201, payload)
}

func (o *UpdateGroupMappingsCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsCreated %s", 201, payload)
}

func (o *UpdateGroupMappingsCreated) GetPayload() *models.MessageResponse {
	return o.Payload
}

func (o *UpdateGroupMappingsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupMappingsBadRequest creates a UpdateGroupMappingsBadRequest with default headers values
func NewUpdateGroupMappingsBadRequest() *UpdateGroupMappingsBadRequest {
	return &UpdateGroupMappingsBadRequest{}
}

/*
UpdateGroupMappingsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateGroupMappingsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update group mappings bad request response has a 2xx status code
func (o *UpdateGroupMappingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update group mappings bad request response has a 3xx status code
func (o *UpdateGroupMappingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group mappings bad request response has a 4xx status code
func (o *UpdateGroupMappingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update group mappings bad request response has a 5xx status code
func (o *UpdateGroupMappingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update group mappings bad request response a status code equal to that given
func (o *UpdateGroupMappingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update group mappings bad request response
func (o *UpdateGroupMappingsBadRequest) Code() int {
	return 400
}

func (o *UpdateGroupMappingsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsBadRequest %s", 400, payload)
}

func (o *UpdateGroupMappingsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsBadRequest %s", 400, payload)
}

func (o *UpdateGroupMappingsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateGroupMappingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupMappingsUnauthorized creates a UpdateGroupMappingsUnauthorized with default headers values
func NewUpdateGroupMappingsUnauthorized() *UpdateGroupMappingsUnauthorized {
	return &UpdateGroupMappingsUnauthorized{}
}

/*
UpdateGroupMappingsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateGroupMappingsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update group mappings unauthorized response has a 2xx status code
func (o *UpdateGroupMappingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update group mappings unauthorized response has a 3xx status code
func (o *UpdateGroupMappingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group mappings unauthorized response has a 4xx status code
func (o *UpdateGroupMappingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update group mappings unauthorized response has a 5xx status code
func (o *UpdateGroupMappingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update group mappings unauthorized response a status code equal to that given
func (o *UpdateGroupMappingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update group mappings unauthorized response
func (o *UpdateGroupMappingsUnauthorized) Code() int {
	return 401
}

func (o *UpdateGroupMappingsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsUnauthorized %s", 401, payload)
}

func (o *UpdateGroupMappingsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsUnauthorized %s", 401, payload)
}

func (o *UpdateGroupMappingsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateGroupMappingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupMappingsForbidden creates a UpdateGroupMappingsForbidden with default headers values
func NewUpdateGroupMappingsForbidden() *UpdateGroupMappingsForbidden {
	return &UpdateGroupMappingsForbidden{}
}

/*
UpdateGroupMappingsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateGroupMappingsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update group mappings forbidden response has a 2xx status code
func (o *UpdateGroupMappingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update group mappings forbidden response has a 3xx status code
func (o *UpdateGroupMappingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group mappings forbidden response has a 4xx status code
func (o *UpdateGroupMappingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update group mappings forbidden response has a 5xx status code
func (o *UpdateGroupMappingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update group mappings forbidden response a status code equal to that given
func (o *UpdateGroupMappingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update group mappings forbidden response
func (o *UpdateGroupMappingsForbidden) Code() int {
	return 403
}

func (o *UpdateGroupMappingsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsForbidden %s", 403, payload)
}

func (o *UpdateGroupMappingsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsForbidden %s", 403, payload)
}

func (o *UpdateGroupMappingsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateGroupMappingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupMappingsInternalServerError creates a UpdateGroupMappingsInternalServerError with default headers values
func NewUpdateGroupMappingsInternalServerError() *UpdateGroupMappingsInternalServerError {
	return &UpdateGroupMappingsInternalServerError{}
}

/*
UpdateGroupMappingsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateGroupMappingsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update group mappings internal server error response has a 2xx status code
func (o *UpdateGroupMappingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update group mappings internal server error response has a 3xx status code
func (o *UpdateGroupMappingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group mappings internal server error response has a 4xx status code
func (o *UpdateGroupMappingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update group mappings internal server error response has a 5xx status code
func (o *UpdateGroupMappingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update group mappings internal server error response a status code equal to that given
func (o *UpdateGroupMappingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update group mappings internal server error response
func (o *UpdateGroupMappingsInternalServerError) Code() int {
	return 500
}

func (o *UpdateGroupMappingsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsInternalServerError %s", 500, payload)
}

func (o *UpdateGroupMappingsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groupsync/groups/{group_id}][%d] updateGroupMappingsInternalServerError %s", 500, payload)
}

func (o *UpdateGroupMappingsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateGroupMappingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
