// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

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

// GetSignedInUserTeamListReader is a Reader for the GetSignedInUserTeamList structure.
type GetSignedInUserTeamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSignedInUserTeamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSignedInUserTeamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSignedInUserTeamListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSignedInUserTeamListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSignedInUserTeamListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user/teams] getSignedInUserTeamList", response, response.Code())
	}
}

// NewGetSignedInUserTeamListOK creates a GetSignedInUserTeamListOK with default headers values
func NewGetSignedInUserTeamListOK() *GetSignedInUserTeamListOK {
	return &GetSignedInUserTeamListOK{}
}

/*
GetSignedInUserTeamListOK describes a response with status code 200, with default header values.

(empty)
*/
type GetSignedInUserTeamListOK struct {
	Payload []*models.TeamDTO
}

// IsSuccess returns true when this get signed in user team list Ok response has a 2xx status code
func (o *GetSignedInUserTeamListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get signed in user team list Ok response has a 3xx status code
func (o *GetSignedInUserTeamListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user team list Ok response has a 4xx status code
func (o *GetSignedInUserTeamListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get signed in user team list Ok response has a 5xx status code
func (o *GetSignedInUserTeamListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get signed in user team list Ok response a status code equal to that given
func (o *GetSignedInUserTeamListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get signed in user team list Ok response
func (o *GetSignedInUserTeamListOK) Code() int {
	return 200
}

func (o *GetSignedInUserTeamListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListOk %s", 200, payload)
}

func (o *GetSignedInUserTeamListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListOk %s", 200, payload)
}

func (o *GetSignedInUserTeamListOK) GetPayload() []*models.TeamDTO {
	return o.Payload
}

func (o *GetSignedInUserTeamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserTeamListUnauthorized creates a GetSignedInUserTeamListUnauthorized with default headers values
func NewGetSignedInUserTeamListUnauthorized() *GetSignedInUserTeamListUnauthorized {
	return &GetSignedInUserTeamListUnauthorized{}
}

/*
GetSignedInUserTeamListUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetSignedInUserTeamListUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get signed in user team list unauthorized response has a 2xx status code
func (o *GetSignedInUserTeamListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get signed in user team list unauthorized response has a 3xx status code
func (o *GetSignedInUserTeamListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user team list unauthorized response has a 4xx status code
func (o *GetSignedInUserTeamListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get signed in user team list unauthorized response has a 5xx status code
func (o *GetSignedInUserTeamListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get signed in user team list unauthorized response a status code equal to that given
func (o *GetSignedInUserTeamListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get signed in user team list unauthorized response
func (o *GetSignedInUserTeamListUnauthorized) Code() int {
	return 401
}

func (o *GetSignedInUserTeamListUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListUnauthorized %s", 401, payload)
}

func (o *GetSignedInUserTeamListUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListUnauthorized %s", 401, payload)
}

func (o *GetSignedInUserTeamListUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserTeamListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserTeamListForbidden creates a GetSignedInUserTeamListForbidden with default headers values
func NewGetSignedInUserTeamListForbidden() *GetSignedInUserTeamListForbidden {
	return &GetSignedInUserTeamListForbidden{}
}

/*
GetSignedInUserTeamListForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetSignedInUserTeamListForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get signed in user team list forbidden response has a 2xx status code
func (o *GetSignedInUserTeamListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get signed in user team list forbidden response has a 3xx status code
func (o *GetSignedInUserTeamListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user team list forbidden response has a 4xx status code
func (o *GetSignedInUserTeamListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get signed in user team list forbidden response has a 5xx status code
func (o *GetSignedInUserTeamListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get signed in user team list forbidden response a status code equal to that given
func (o *GetSignedInUserTeamListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get signed in user team list forbidden response
func (o *GetSignedInUserTeamListForbidden) Code() int {
	return 403
}

func (o *GetSignedInUserTeamListForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListForbidden %s", 403, payload)
}

func (o *GetSignedInUserTeamListForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListForbidden %s", 403, payload)
}

func (o *GetSignedInUserTeamListForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserTeamListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserTeamListInternalServerError creates a GetSignedInUserTeamListInternalServerError with default headers values
func NewGetSignedInUserTeamListInternalServerError() *GetSignedInUserTeamListInternalServerError {
	return &GetSignedInUserTeamListInternalServerError{}
}

/*
GetSignedInUserTeamListInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetSignedInUserTeamListInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get signed in user team list internal server error response has a 2xx status code
func (o *GetSignedInUserTeamListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get signed in user team list internal server error response has a 3xx status code
func (o *GetSignedInUserTeamListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signed in user team list internal server error response has a 4xx status code
func (o *GetSignedInUserTeamListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get signed in user team list internal server error response has a 5xx status code
func (o *GetSignedInUserTeamListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get signed in user team list internal server error response a status code equal to that given
func (o *GetSignedInUserTeamListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get signed in user team list internal server error response
func (o *GetSignedInUserTeamListInternalServerError) Code() int {
	return 500
}

func (o *GetSignedInUserTeamListInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListInternalServerError %s", 500, payload)
}

func (o *GetSignedInUserTeamListInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListInternalServerError %s", 500, payload)
}

func (o *GetSignedInUserTeamListInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserTeamListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
