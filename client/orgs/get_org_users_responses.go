// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

// GetOrgUsersReader is a Reader for the GetOrgUsers structure.
type GetOrgUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrgUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org_id}/users] getOrgUsers", response, response.Code())
	}
}

// NewGetOrgUsersOK creates a GetOrgUsersOK with default headers values
func NewGetOrgUsersOK() *GetOrgUsersOK {
	return &GetOrgUsersOK{}
}

/*
GetOrgUsersOK describes a response with status code 200, with default header values.

(empty)
*/
type GetOrgUsersOK struct {
	Payload []*models.OrgUserDTO
}

// IsSuccess returns true when this get org users Ok response has a 2xx status code
func (o *GetOrgUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org users Ok response has a 3xx status code
func (o *GetOrgUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org users Ok response has a 4xx status code
func (o *GetOrgUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org users Ok response has a 5xx status code
func (o *GetOrgUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org users Ok response a status code equal to that given
func (o *GetOrgUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org users Ok response
func (o *GetOrgUsersOK) Code() int {
	return 200
}

func (o *GetOrgUsersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/users][%d] getOrgUsersOk %s", 200, payload)
}

func (o *GetOrgUsersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/users][%d] getOrgUsersOk %s", 200, payload)
}

func (o *GetOrgUsersOK) GetPayload() []*models.OrgUserDTO {
	return o.Payload
}

func (o *GetOrgUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgUsersUnauthorized creates a GetOrgUsersUnauthorized with default headers values
func NewGetOrgUsersUnauthorized() *GetOrgUsersUnauthorized {
	return &GetOrgUsersUnauthorized{}
}

/*
GetOrgUsersUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetOrgUsersUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org users unauthorized response has a 2xx status code
func (o *GetOrgUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org users unauthorized response has a 3xx status code
func (o *GetOrgUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org users unauthorized response has a 4xx status code
func (o *GetOrgUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org users unauthorized response has a 5xx status code
func (o *GetOrgUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get org users unauthorized response a status code equal to that given
func (o *GetOrgUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get org users unauthorized response
func (o *GetOrgUsersUnauthorized) Code() int {
	return 401
}

func (o *GetOrgUsersUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/users][%d] getOrgUsersUnauthorized %s", 401, payload)
}

func (o *GetOrgUsersUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/users][%d] getOrgUsersUnauthorized %s", 401, payload)
}

func (o *GetOrgUsersUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgUsersForbidden creates a GetOrgUsersForbidden with default headers values
func NewGetOrgUsersForbidden() *GetOrgUsersForbidden {
	return &GetOrgUsersForbidden{}
}

/*
GetOrgUsersForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetOrgUsersForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org users forbidden response has a 2xx status code
func (o *GetOrgUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org users forbidden response has a 3xx status code
func (o *GetOrgUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org users forbidden response has a 4xx status code
func (o *GetOrgUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org users forbidden response has a 5xx status code
func (o *GetOrgUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org users forbidden response a status code equal to that given
func (o *GetOrgUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org users forbidden response
func (o *GetOrgUsersForbidden) Code() int {
	return 403
}

func (o *GetOrgUsersForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/users][%d] getOrgUsersForbidden %s", 403, payload)
}

func (o *GetOrgUsersForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/users][%d] getOrgUsersForbidden %s", 403, payload)
}

func (o *GetOrgUsersForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgUsersInternalServerError creates a GetOrgUsersInternalServerError with default headers values
func NewGetOrgUsersInternalServerError() *GetOrgUsersInternalServerError {
	return &GetOrgUsersInternalServerError{}
}

/*
GetOrgUsersInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetOrgUsersInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org users internal server error response has a 2xx status code
func (o *GetOrgUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org users internal server error response has a 3xx status code
func (o *GetOrgUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org users internal server error response has a 4xx status code
func (o *GetOrgUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org users internal server error response has a 5xx status code
func (o *GetOrgUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org users internal server error response a status code equal to that given
func (o *GetOrgUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org users internal server error response
func (o *GetOrgUsersInternalServerError) Code() int {
	return 500
}

func (o *GetOrgUsersInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/users][%d] getOrgUsersInternalServerError %s", 500, payload)
}

func (o *GetOrgUsersInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/users][%d] getOrgUsersInternalServerError %s", 500, payload)
}

func (o *GetOrgUsersInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
