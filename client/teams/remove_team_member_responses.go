// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// RemoveTeamMemberReader is a Reader for the RemoveTeamMember structure.
type RemoveTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveTeamMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveTeamMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /teams/{team_id}/members/{user_id}] removeTeamMember", response, response.Code())
	}
}

// NewRemoveTeamMemberOK creates a RemoveTeamMemberOK with default headers values
func NewRemoveTeamMemberOK() *RemoveTeamMemberOK {
	return &RemoveTeamMemberOK{}
}

/*
RemoveTeamMemberOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type RemoveTeamMemberOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this remove team member Ok response has a 2xx status code
func (o *RemoveTeamMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove team member Ok response has a 3xx status code
func (o *RemoveTeamMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team member Ok response has a 4xx status code
func (o *RemoveTeamMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove team member Ok response has a 5xx status code
func (o *RemoveTeamMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team member Ok response a status code equal to that given
func (o *RemoveTeamMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove team member Ok response
func (o *RemoveTeamMemberOK) Code() int {
	return 200
}

func (o *RemoveTeamMemberOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberOk %s", 200, payload)
}

func (o *RemoveTeamMemberOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberOk %s", 200, payload)
}

func (o *RemoveTeamMemberOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamMemberUnauthorized creates a RemoveTeamMemberUnauthorized with default headers values
func NewRemoveTeamMemberUnauthorized() *RemoveTeamMemberUnauthorized {
	return &RemoveTeamMemberUnauthorized{}
}

/*
RemoveTeamMemberUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RemoveTeamMemberUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team member unauthorized response has a 2xx status code
func (o *RemoveTeamMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team member unauthorized response has a 3xx status code
func (o *RemoveTeamMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team member unauthorized response has a 4xx status code
func (o *RemoveTeamMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove team member unauthorized response has a 5xx status code
func (o *RemoveTeamMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team member unauthorized response a status code equal to that given
func (o *RemoveTeamMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove team member unauthorized response
func (o *RemoveTeamMemberUnauthorized) Code() int {
	return 401
}

func (o *RemoveTeamMemberUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberUnauthorized %s", 401, payload)
}

func (o *RemoveTeamMemberUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberUnauthorized %s", 401, payload)
}

func (o *RemoveTeamMemberUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamMemberForbidden creates a RemoveTeamMemberForbidden with default headers values
func NewRemoveTeamMemberForbidden() *RemoveTeamMemberForbidden {
	return &RemoveTeamMemberForbidden{}
}

/*
RemoveTeamMemberForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RemoveTeamMemberForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team member forbidden response has a 2xx status code
func (o *RemoveTeamMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team member forbidden response has a 3xx status code
func (o *RemoveTeamMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team member forbidden response has a 4xx status code
func (o *RemoveTeamMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove team member forbidden response has a 5xx status code
func (o *RemoveTeamMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team member forbidden response a status code equal to that given
func (o *RemoveTeamMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove team member forbidden response
func (o *RemoveTeamMemberForbidden) Code() int {
	return 403
}

func (o *RemoveTeamMemberForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberForbidden %s", 403, payload)
}

func (o *RemoveTeamMemberForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberForbidden %s", 403, payload)
}

func (o *RemoveTeamMemberForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamMemberNotFound creates a RemoveTeamMemberNotFound with default headers values
func NewRemoveTeamMemberNotFound() *RemoveTeamMemberNotFound {
	return &RemoveTeamMemberNotFound{}
}

/*
RemoveTeamMemberNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RemoveTeamMemberNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team member not found response has a 2xx status code
func (o *RemoveTeamMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team member not found response has a 3xx status code
func (o *RemoveTeamMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team member not found response has a 4xx status code
func (o *RemoveTeamMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove team member not found response has a 5xx status code
func (o *RemoveTeamMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team member not found response a status code equal to that given
func (o *RemoveTeamMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove team member not found response
func (o *RemoveTeamMemberNotFound) Code() int {
	return 404
}

func (o *RemoveTeamMemberNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberNotFound %s", 404, payload)
}

func (o *RemoveTeamMemberNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberNotFound %s", 404, payload)
}

func (o *RemoveTeamMemberNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamMemberInternalServerError creates a RemoveTeamMemberInternalServerError with default headers values
func NewRemoveTeamMemberInternalServerError() *RemoveTeamMemberInternalServerError {
	return &RemoveTeamMemberInternalServerError{}
}

/*
RemoveTeamMemberInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RemoveTeamMemberInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team member internal server error response has a 2xx status code
func (o *RemoveTeamMemberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team member internal server error response has a 3xx status code
func (o *RemoveTeamMemberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team member internal server error response has a 4xx status code
func (o *RemoveTeamMemberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove team member internal server error response has a 5xx status code
func (o *RemoveTeamMemberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove team member internal server error response a status code equal to that given
func (o *RemoveTeamMemberInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove team member internal server error response
func (o *RemoveTeamMemberInternalServerError) Code() int {
	return 500
}

func (o *RemoveTeamMemberInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberInternalServerError %s", 500, payload)
}

func (o *RemoveTeamMemberInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{team_id}/members/{user_id}][%d] removeTeamMemberInternalServerError %s", 500, payload)
}

func (o *RemoveTeamMemberInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
