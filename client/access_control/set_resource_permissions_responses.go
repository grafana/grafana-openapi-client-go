// Code generated by go-swagger; DO NOT EDIT.

package access_control

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

// SetResourcePermissionsReader is a Reader for the SetResourcePermissions structure.
type SetResourcePermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetResourcePermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetResourcePermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetResourcePermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetResourcePermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetResourcePermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetResourcePermissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /access-control/{resource}/{resourceID}] setResourcePermissions", response, response.Code())
	}
}

// NewSetResourcePermissionsOK creates a SetResourcePermissionsOK with default headers values
func NewSetResourcePermissionsOK() *SetResourcePermissionsOK {
	return &SetResourcePermissionsOK{}
}

/*
SetResourcePermissionsOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type SetResourcePermissionsOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this set resource permissions Ok response has a 2xx status code
func (o *SetResourcePermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set resource permissions Ok response has a 3xx status code
func (o *SetResourcePermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set resource permissions Ok response has a 4xx status code
func (o *SetResourcePermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set resource permissions Ok response has a 5xx status code
func (o *SetResourcePermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set resource permissions Ok response a status code equal to that given
func (o *SetResourcePermissionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set resource permissions Ok response
func (o *SetResourcePermissionsOK) Code() int {
	return 200
}

func (o *SetResourcePermissionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsOk %s", 200, payload)
}

func (o *SetResourcePermissionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsOk %s", 200, payload)
}

func (o *SetResourcePermissionsOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *SetResourcePermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetResourcePermissionsBadRequest creates a SetResourcePermissionsBadRequest with default headers values
func NewSetResourcePermissionsBadRequest() *SetResourcePermissionsBadRequest {
	return &SetResourcePermissionsBadRequest{}
}

/*
SetResourcePermissionsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type SetResourcePermissionsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set resource permissions bad request response has a 2xx status code
func (o *SetResourcePermissionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set resource permissions bad request response has a 3xx status code
func (o *SetResourcePermissionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set resource permissions bad request response has a 4xx status code
func (o *SetResourcePermissionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this set resource permissions bad request response has a 5xx status code
func (o *SetResourcePermissionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this set resource permissions bad request response a status code equal to that given
func (o *SetResourcePermissionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the set resource permissions bad request response
func (o *SetResourcePermissionsBadRequest) Code() int {
	return 400
}

func (o *SetResourcePermissionsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsBadRequest %s", 400, payload)
}

func (o *SetResourcePermissionsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsBadRequest %s", 400, payload)
}

func (o *SetResourcePermissionsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetResourcePermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetResourcePermissionsForbidden creates a SetResourcePermissionsForbidden with default headers values
func NewSetResourcePermissionsForbidden() *SetResourcePermissionsForbidden {
	return &SetResourcePermissionsForbidden{}
}

/*
SetResourcePermissionsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SetResourcePermissionsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set resource permissions forbidden response has a 2xx status code
func (o *SetResourcePermissionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set resource permissions forbidden response has a 3xx status code
func (o *SetResourcePermissionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set resource permissions forbidden response has a 4xx status code
func (o *SetResourcePermissionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set resource permissions forbidden response has a 5xx status code
func (o *SetResourcePermissionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set resource permissions forbidden response a status code equal to that given
func (o *SetResourcePermissionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the set resource permissions forbidden response
func (o *SetResourcePermissionsForbidden) Code() int {
	return 403
}

func (o *SetResourcePermissionsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsForbidden %s", 403, payload)
}

func (o *SetResourcePermissionsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsForbidden %s", 403, payload)
}

func (o *SetResourcePermissionsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetResourcePermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetResourcePermissionsNotFound creates a SetResourcePermissionsNotFound with default headers values
func NewSetResourcePermissionsNotFound() *SetResourcePermissionsNotFound {
	return &SetResourcePermissionsNotFound{}
}

/*
SetResourcePermissionsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type SetResourcePermissionsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set resource permissions not found response has a 2xx status code
func (o *SetResourcePermissionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set resource permissions not found response has a 3xx status code
func (o *SetResourcePermissionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set resource permissions not found response has a 4xx status code
func (o *SetResourcePermissionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set resource permissions not found response has a 5xx status code
func (o *SetResourcePermissionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set resource permissions not found response a status code equal to that given
func (o *SetResourcePermissionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set resource permissions not found response
func (o *SetResourcePermissionsNotFound) Code() int {
	return 404
}

func (o *SetResourcePermissionsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsNotFound %s", 404, payload)
}

func (o *SetResourcePermissionsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsNotFound %s", 404, payload)
}

func (o *SetResourcePermissionsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetResourcePermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetResourcePermissionsInternalServerError creates a SetResourcePermissionsInternalServerError with default headers values
func NewSetResourcePermissionsInternalServerError() *SetResourcePermissionsInternalServerError {
	return &SetResourcePermissionsInternalServerError{}
}

/*
SetResourcePermissionsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SetResourcePermissionsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set resource permissions internal server error response has a 2xx status code
func (o *SetResourcePermissionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set resource permissions internal server error response has a 3xx status code
func (o *SetResourcePermissionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set resource permissions internal server error response has a 4xx status code
func (o *SetResourcePermissionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set resource permissions internal server error response has a 5xx status code
func (o *SetResourcePermissionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set resource permissions internal server error response a status code equal to that given
func (o *SetResourcePermissionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set resource permissions internal server error response
func (o *SetResourcePermissionsInternalServerError) Code() int {
	return 500
}

func (o *SetResourcePermissionsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsInternalServerError %s", 500, payload)
}

func (o *SetResourcePermissionsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/{resource}/{resourceID}][%d] setResourcePermissionsInternalServerError %s", 500, payload)
}

func (o *SetResourcePermissionsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetResourcePermissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
