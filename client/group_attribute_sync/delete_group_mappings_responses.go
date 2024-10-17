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

// DeleteGroupMappingsReader is a Reader for the DeleteGroupMappings structure.
type DeleteGroupMappingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupMappingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGroupMappingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteGroupMappingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteGroupMappingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGroupMappingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteGroupMappingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /groupsync/groups/{group_id}] deleteGroupMappings", response, response.Code())
	}
}

// NewDeleteGroupMappingsNoContent creates a DeleteGroupMappingsNoContent with default headers values
func NewDeleteGroupMappingsNoContent() *DeleteGroupMappingsNoContent {
	return &DeleteGroupMappingsNoContent{}
}

/*
DeleteGroupMappingsNoContent describes a response with status code 204, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteGroupMappingsNoContent struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete group mappings no content response has a 2xx status code
func (o *DeleteGroupMappingsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete group mappings no content response has a 3xx status code
func (o *DeleteGroupMappingsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group mappings no content response has a 4xx status code
func (o *DeleteGroupMappingsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group mappings no content response has a 5xx status code
func (o *DeleteGroupMappingsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group mappings no content response a status code equal to that given
func (o *DeleteGroupMappingsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete group mappings no content response
func (o *DeleteGroupMappingsNoContent) Code() int {
	return 204
}

func (o *DeleteGroupMappingsNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsNoContent %s", 204, payload)
}

func (o *DeleteGroupMappingsNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsNoContent %s", 204, payload)
}

func (o *DeleteGroupMappingsNoContent) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteGroupMappingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMappingsBadRequest creates a DeleteGroupMappingsBadRequest with default headers values
func NewDeleteGroupMappingsBadRequest() *DeleteGroupMappingsBadRequest {
	return &DeleteGroupMappingsBadRequest{}
}

/*
DeleteGroupMappingsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DeleteGroupMappingsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete group mappings bad request response has a 2xx status code
func (o *DeleteGroupMappingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group mappings bad request response has a 3xx status code
func (o *DeleteGroupMappingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group mappings bad request response has a 4xx status code
func (o *DeleteGroupMappingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group mappings bad request response has a 5xx status code
func (o *DeleteGroupMappingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group mappings bad request response a status code equal to that given
func (o *DeleteGroupMappingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete group mappings bad request response
func (o *DeleteGroupMappingsBadRequest) Code() int {
	return 400
}

func (o *DeleteGroupMappingsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsBadRequest %s", 400, payload)
}

func (o *DeleteGroupMappingsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsBadRequest %s", 400, payload)
}

func (o *DeleteGroupMappingsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteGroupMappingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMappingsUnauthorized creates a DeleteGroupMappingsUnauthorized with default headers values
func NewDeleteGroupMappingsUnauthorized() *DeleteGroupMappingsUnauthorized {
	return &DeleteGroupMappingsUnauthorized{}
}

/*
DeleteGroupMappingsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteGroupMappingsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete group mappings unauthorized response has a 2xx status code
func (o *DeleteGroupMappingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group mappings unauthorized response has a 3xx status code
func (o *DeleteGroupMappingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group mappings unauthorized response has a 4xx status code
func (o *DeleteGroupMappingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group mappings unauthorized response has a 5xx status code
func (o *DeleteGroupMappingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group mappings unauthorized response a status code equal to that given
func (o *DeleteGroupMappingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete group mappings unauthorized response
func (o *DeleteGroupMappingsUnauthorized) Code() int {
	return 401
}

func (o *DeleteGroupMappingsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsUnauthorized %s", 401, payload)
}

func (o *DeleteGroupMappingsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsUnauthorized %s", 401, payload)
}

func (o *DeleteGroupMappingsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteGroupMappingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMappingsForbidden creates a DeleteGroupMappingsForbidden with default headers values
func NewDeleteGroupMappingsForbidden() *DeleteGroupMappingsForbidden {
	return &DeleteGroupMappingsForbidden{}
}

/*
DeleteGroupMappingsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteGroupMappingsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete group mappings forbidden response has a 2xx status code
func (o *DeleteGroupMappingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group mappings forbidden response has a 3xx status code
func (o *DeleteGroupMappingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group mappings forbidden response has a 4xx status code
func (o *DeleteGroupMappingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group mappings forbidden response has a 5xx status code
func (o *DeleteGroupMappingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group mappings forbidden response a status code equal to that given
func (o *DeleteGroupMappingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete group mappings forbidden response
func (o *DeleteGroupMappingsForbidden) Code() int {
	return 403
}

func (o *DeleteGroupMappingsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsForbidden %s", 403, payload)
}

func (o *DeleteGroupMappingsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsForbidden %s", 403, payload)
}

func (o *DeleteGroupMappingsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteGroupMappingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMappingsInternalServerError creates a DeleteGroupMappingsInternalServerError with default headers values
func NewDeleteGroupMappingsInternalServerError() *DeleteGroupMappingsInternalServerError {
	return &DeleteGroupMappingsInternalServerError{}
}

/*
DeleteGroupMappingsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteGroupMappingsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete group mappings internal server error response has a 2xx status code
func (o *DeleteGroupMappingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group mappings internal server error response has a 3xx status code
func (o *DeleteGroupMappingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group mappings internal server error response has a 4xx status code
func (o *DeleteGroupMappingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group mappings internal server error response has a 5xx status code
func (o *DeleteGroupMappingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete group mappings internal server error response a status code equal to that given
func (o *DeleteGroupMappingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete group mappings internal server error response
func (o *DeleteGroupMappingsInternalServerError) Code() int {
	return 500
}

func (o *DeleteGroupMappingsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsInternalServerError %s", 500, payload)
}

func (o *DeleteGroupMappingsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groupsync/groups/{group_id}][%d] deleteGroupMappingsInternalServerError %s", 500, payload)
}

func (o *DeleteGroupMappingsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteGroupMappingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
