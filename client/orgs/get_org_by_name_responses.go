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

// GetOrgByNameReader is a Reader for the GetOrgByName structure.
type GetOrgByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrgByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/name/{org_name}] getOrgByName", response, response.Code())
	}
}

// NewGetOrgByNameOK creates a GetOrgByNameOK with default headers values
func NewGetOrgByNameOK() *GetOrgByNameOK {
	return &GetOrgByNameOK{}
}

/*
GetOrgByNameOK describes a response with status code 200, with default header values.

(empty)
*/
type GetOrgByNameOK struct {
	Payload *models.OrgDetailsDTO
}

// IsSuccess returns true when this get org by name Ok response has a 2xx status code
func (o *GetOrgByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org by name Ok response has a 3xx status code
func (o *GetOrgByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org by name Ok response has a 4xx status code
func (o *GetOrgByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org by name Ok response has a 5xx status code
func (o *GetOrgByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org by name Ok response a status code equal to that given
func (o *GetOrgByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org by name Ok response
func (o *GetOrgByNameOK) Code() int {
	return 200
}

func (o *GetOrgByNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/name/{org_name}][%d] getOrgByNameOk %s", 200, payload)
}

func (o *GetOrgByNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/name/{org_name}][%d] getOrgByNameOk %s", 200, payload)
}

func (o *GetOrgByNameOK) GetPayload() *models.OrgDetailsDTO {
	return o.Payload
}

func (o *GetOrgByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrgDetailsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgByNameUnauthorized creates a GetOrgByNameUnauthorized with default headers values
func NewGetOrgByNameUnauthorized() *GetOrgByNameUnauthorized {
	return &GetOrgByNameUnauthorized{}
}

/*
GetOrgByNameUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetOrgByNameUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org by name unauthorized response has a 2xx status code
func (o *GetOrgByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org by name unauthorized response has a 3xx status code
func (o *GetOrgByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org by name unauthorized response has a 4xx status code
func (o *GetOrgByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org by name unauthorized response has a 5xx status code
func (o *GetOrgByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get org by name unauthorized response a status code equal to that given
func (o *GetOrgByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get org by name unauthorized response
func (o *GetOrgByNameUnauthorized) Code() int {
	return 401
}

func (o *GetOrgByNameUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/name/{org_name}][%d] getOrgByNameUnauthorized %s", 401, payload)
}

func (o *GetOrgByNameUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/name/{org_name}][%d] getOrgByNameUnauthorized %s", 401, payload)
}

func (o *GetOrgByNameUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgByNameForbidden creates a GetOrgByNameForbidden with default headers values
func NewGetOrgByNameForbidden() *GetOrgByNameForbidden {
	return &GetOrgByNameForbidden{}
}

/*
GetOrgByNameForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetOrgByNameForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org by name forbidden response has a 2xx status code
func (o *GetOrgByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org by name forbidden response has a 3xx status code
func (o *GetOrgByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org by name forbidden response has a 4xx status code
func (o *GetOrgByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org by name forbidden response has a 5xx status code
func (o *GetOrgByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org by name forbidden response a status code equal to that given
func (o *GetOrgByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org by name forbidden response
func (o *GetOrgByNameForbidden) Code() int {
	return 403
}

func (o *GetOrgByNameForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/name/{org_name}][%d] getOrgByNameForbidden %s", 403, payload)
}

func (o *GetOrgByNameForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/name/{org_name}][%d] getOrgByNameForbidden %s", 403, payload)
}

func (o *GetOrgByNameForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgByNameInternalServerError creates a GetOrgByNameInternalServerError with default headers values
func NewGetOrgByNameInternalServerError() *GetOrgByNameInternalServerError {
	return &GetOrgByNameInternalServerError{}
}

/*
GetOrgByNameInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetOrgByNameInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org by name internal server error response has a 2xx status code
func (o *GetOrgByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org by name internal server error response has a 3xx status code
func (o *GetOrgByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org by name internal server error response has a 4xx status code
func (o *GetOrgByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org by name internal server error response has a 5xx status code
func (o *GetOrgByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org by name internal server error response a status code equal to that given
func (o *GetOrgByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org by name internal server error response
func (o *GetOrgByNameInternalServerError) Code() int {
	return 500
}

func (o *GetOrgByNameInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/name/{org_name}][%d] getOrgByNameInternalServerError %s", 500, payload)
}

func (o *GetOrgByNameInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/name/{org_name}][%d] getOrgByNameInternalServerError %s", 500, payload)
}

func (o *GetOrgByNameInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
