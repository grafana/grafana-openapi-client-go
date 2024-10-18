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

// SetTeamMembershipsReader is a Reader for the SetTeamMemberships structure.
type SetTeamMembershipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetTeamMembershipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetTeamMembershipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetTeamMembershipsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetTeamMembershipsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetTeamMembershipsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetTeamMembershipsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /teams/{team_id}/members] setTeamMemberships", response, response.Code())
	}
}

// NewSetTeamMembershipsOK creates a SetTeamMembershipsOK with default headers values
func NewSetTeamMembershipsOK() *SetTeamMembershipsOK {
	return &SetTeamMembershipsOK{}
}

/*
SetTeamMembershipsOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type SetTeamMembershipsOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this set team memberships Ok response has a 2xx status code
func (o *SetTeamMembershipsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set team memberships Ok response has a 3xx status code
func (o *SetTeamMembershipsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set team memberships Ok response has a 4xx status code
func (o *SetTeamMembershipsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set team memberships Ok response has a 5xx status code
func (o *SetTeamMembershipsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set team memberships Ok response a status code equal to that given
func (o *SetTeamMembershipsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set team memberships Ok response
func (o *SetTeamMembershipsOK) Code() int {
	return 200
}

func (o *SetTeamMembershipsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsOk %s", 200, payload)
}

func (o *SetTeamMembershipsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsOk %s", 200, payload)
}

func (o *SetTeamMembershipsOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *SetTeamMembershipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTeamMembershipsUnauthorized creates a SetTeamMembershipsUnauthorized with default headers values
func NewSetTeamMembershipsUnauthorized() *SetTeamMembershipsUnauthorized {
	return &SetTeamMembershipsUnauthorized{}
}

/*
SetTeamMembershipsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type SetTeamMembershipsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set team memberships unauthorized response has a 2xx status code
func (o *SetTeamMembershipsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set team memberships unauthorized response has a 3xx status code
func (o *SetTeamMembershipsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set team memberships unauthorized response has a 4xx status code
func (o *SetTeamMembershipsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this set team memberships unauthorized response has a 5xx status code
func (o *SetTeamMembershipsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this set team memberships unauthorized response a status code equal to that given
func (o *SetTeamMembershipsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the set team memberships unauthorized response
func (o *SetTeamMembershipsUnauthorized) Code() int {
	return 401
}

func (o *SetTeamMembershipsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsUnauthorized %s", 401, payload)
}

func (o *SetTeamMembershipsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsUnauthorized %s", 401, payload)
}

func (o *SetTeamMembershipsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetTeamMembershipsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTeamMembershipsForbidden creates a SetTeamMembershipsForbidden with default headers values
func NewSetTeamMembershipsForbidden() *SetTeamMembershipsForbidden {
	return &SetTeamMembershipsForbidden{}
}

/*
SetTeamMembershipsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SetTeamMembershipsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set team memberships forbidden response has a 2xx status code
func (o *SetTeamMembershipsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set team memberships forbidden response has a 3xx status code
func (o *SetTeamMembershipsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set team memberships forbidden response has a 4xx status code
func (o *SetTeamMembershipsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set team memberships forbidden response has a 5xx status code
func (o *SetTeamMembershipsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set team memberships forbidden response a status code equal to that given
func (o *SetTeamMembershipsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the set team memberships forbidden response
func (o *SetTeamMembershipsForbidden) Code() int {
	return 403
}

func (o *SetTeamMembershipsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsForbidden %s", 403, payload)
}

func (o *SetTeamMembershipsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsForbidden %s", 403, payload)
}

func (o *SetTeamMembershipsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetTeamMembershipsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTeamMembershipsNotFound creates a SetTeamMembershipsNotFound with default headers values
func NewSetTeamMembershipsNotFound() *SetTeamMembershipsNotFound {
	return &SetTeamMembershipsNotFound{}
}

/*
SetTeamMembershipsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type SetTeamMembershipsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set team memberships not found response has a 2xx status code
func (o *SetTeamMembershipsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set team memberships not found response has a 3xx status code
func (o *SetTeamMembershipsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set team memberships not found response has a 4xx status code
func (o *SetTeamMembershipsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set team memberships not found response has a 5xx status code
func (o *SetTeamMembershipsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set team memberships not found response a status code equal to that given
func (o *SetTeamMembershipsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set team memberships not found response
func (o *SetTeamMembershipsNotFound) Code() int {
	return 404
}

func (o *SetTeamMembershipsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsNotFound %s", 404, payload)
}

func (o *SetTeamMembershipsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsNotFound %s", 404, payload)
}

func (o *SetTeamMembershipsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetTeamMembershipsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetTeamMembershipsInternalServerError creates a SetTeamMembershipsInternalServerError with default headers values
func NewSetTeamMembershipsInternalServerError() *SetTeamMembershipsInternalServerError {
	return &SetTeamMembershipsInternalServerError{}
}

/*
SetTeamMembershipsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SetTeamMembershipsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set team memberships internal server error response has a 2xx status code
func (o *SetTeamMembershipsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set team memberships internal server error response has a 3xx status code
func (o *SetTeamMembershipsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set team memberships internal server error response has a 4xx status code
func (o *SetTeamMembershipsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set team memberships internal server error response has a 5xx status code
func (o *SetTeamMembershipsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set team memberships internal server error response a status code equal to that given
func (o *SetTeamMembershipsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set team memberships internal server error response
func (o *SetTeamMembershipsInternalServerError) Code() int {
	return 500
}

func (o *SetTeamMembershipsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsInternalServerError %s", 500, payload)
}

func (o *SetTeamMembershipsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /teams/{team_id}/members][%d] setTeamMembershipsInternalServerError %s", 500, payload)
}

func (o *SetTeamMembershipsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetTeamMembershipsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}