// Code generated by go-swagger; DO NOT EDIT.

package admin_users

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

// GetUserQuotaReader is a Reader for the GetUserQuota structure.
type GetUserQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /admin/users/{user_id}/quotas] getUserQuota", response, response.Code())
	}
}

// NewGetUserQuotaOK creates a GetUserQuotaOK with default headers values
func NewGetUserQuotaOK() *GetUserQuotaOK {
	return &GetUserQuotaOK{}
}

/*
GetUserQuotaOK describes a response with status code 200, with default header values.

(empty)
*/
type GetUserQuotaOK struct {
	Payload []*models.QuotaDTO
}

// IsSuccess returns true when this get user quota Ok response has a 2xx status code
func (o *GetUserQuotaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user quota Ok response has a 3xx status code
func (o *GetUserQuotaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user quota Ok response has a 4xx status code
func (o *GetUserQuotaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user quota Ok response has a 5xx status code
func (o *GetUserQuotaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user quota Ok response a status code equal to that given
func (o *GetUserQuotaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user quota Ok response
func (o *GetUserQuotaOK) Code() int {
	return 200
}

func (o *GetUserQuotaOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaOk %s", 200, payload)
}

func (o *GetUserQuotaOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaOk %s", 200, payload)
}

func (o *GetUserQuotaOK) GetPayload() []*models.QuotaDTO {
	return o.Payload
}

func (o *GetUserQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQuotaUnauthorized creates a GetUserQuotaUnauthorized with default headers values
func NewGetUserQuotaUnauthorized() *GetUserQuotaUnauthorized {
	return &GetUserQuotaUnauthorized{}
}

/*
GetUserQuotaUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetUserQuotaUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user quota unauthorized response has a 2xx status code
func (o *GetUserQuotaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user quota unauthorized response has a 3xx status code
func (o *GetUserQuotaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user quota unauthorized response has a 4xx status code
func (o *GetUserQuotaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user quota unauthorized response has a 5xx status code
func (o *GetUserQuotaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user quota unauthorized response a status code equal to that given
func (o *GetUserQuotaUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get user quota unauthorized response
func (o *GetUserQuotaUnauthorized) Code() int {
	return 401
}

func (o *GetUserQuotaUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaUnauthorized %s", 401, payload)
}

func (o *GetUserQuotaUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaUnauthorized %s", 401, payload)
}

func (o *GetUserQuotaUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQuotaForbidden creates a GetUserQuotaForbidden with default headers values
func NewGetUserQuotaForbidden() *GetUserQuotaForbidden {
	return &GetUserQuotaForbidden{}
}

/*
GetUserQuotaForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetUserQuotaForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user quota forbidden response has a 2xx status code
func (o *GetUserQuotaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user quota forbidden response has a 3xx status code
func (o *GetUserQuotaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user quota forbidden response has a 4xx status code
func (o *GetUserQuotaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user quota forbidden response has a 5xx status code
func (o *GetUserQuotaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user quota forbidden response a status code equal to that given
func (o *GetUserQuotaForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user quota forbidden response
func (o *GetUserQuotaForbidden) Code() int {
	return 403
}

func (o *GetUserQuotaForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaForbidden %s", 403, payload)
}

func (o *GetUserQuotaForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaForbidden %s", 403, payload)
}

func (o *GetUserQuotaForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQuotaNotFound creates a GetUserQuotaNotFound with default headers values
func NewGetUserQuotaNotFound() *GetUserQuotaNotFound {
	return &GetUserQuotaNotFound{}
}

/*
GetUserQuotaNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetUserQuotaNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user quota not found response has a 2xx status code
func (o *GetUserQuotaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user quota not found response has a 3xx status code
func (o *GetUserQuotaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user quota not found response has a 4xx status code
func (o *GetUserQuotaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user quota not found response has a 5xx status code
func (o *GetUserQuotaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user quota not found response a status code equal to that given
func (o *GetUserQuotaNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user quota not found response
func (o *GetUserQuotaNotFound) Code() int {
	return 404
}

func (o *GetUserQuotaNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaNotFound %s", 404, payload)
}

func (o *GetUserQuotaNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaNotFound %s", 404, payload)
}

func (o *GetUserQuotaNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQuotaInternalServerError creates a GetUserQuotaInternalServerError with default headers values
func NewGetUserQuotaInternalServerError() *GetUserQuotaInternalServerError {
	return &GetUserQuotaInternalServerError{}
}

/*
GetUserQuotaInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetUserQuotaInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user quota internal server error response has a 2xx status code
func (o *GetUserQuotaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user quota internal server error response has a 3xx status code
func (o *GetUserQuotaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user quota internal server error response has a 4xx status code
func (o *GetUserQuotaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user quota internal server error response has a 5xx status code
func (o *GetUserQuotaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user quota internal server error response a status code equal to that given
func (o *GetUserQuotaInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user quota internal server error response
func (o *GetUserQuotaInternalServerError) Code() int {
	return 500
}

func (o *GetUserQuotaInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaInternalServerError %s", 500, payload)
}

func (o *GetUserQuotaInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users/{user_id}/quotas][%d] getUserQuotaInternalServerError %s", 500, payload)
}

func (o *GetUserQuotaInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
