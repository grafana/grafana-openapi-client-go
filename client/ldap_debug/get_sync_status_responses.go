// Code generated by go-swagger; DO NOT EDIT.

package ldap_debug

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

// GetSyncStatusReader is a Reader for the GetSyncStatus structure.
type GetSyncStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyncStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSyncStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSyncStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSyncStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSyncStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /admin/ldap-sync-status] getSyncStatus", response, response.Code())
	}
}

// NewGetSyncStatusOK creates a GetSyncStatusOK with default headers values
func NewGetSyncStatusOK() *GetSyncStatusOK {
	return &GetSyncStatusOK{}
}

/*
GetSyncStatusOK describes a response with status code 200, with default header values.

(empty)
*/
type GetSyncStatusOK struct {
	Payload *models.ActiveSyncStatusDTO
}

// IsSuccess returns true when this get sync status Ok response has a 2xx status code
func (o *GetSyncStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sync status Ok response has a 3xx status code
func (o *GetSyncStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sync status Ok response has a 4xx status code
func (o *GetSyncStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sync status Ok response has a 5xx status code
func (o *GetSyncStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sync status Ok response a status code equal to that given
func (o *GetSyncStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sync status Ok response
func (o *GetSyncStatusOK) Code() int {
	return 200
}

func (o *GetSyncStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusOk %s", 200, payload)
}

func (o *GetSyncStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusOk %s", 200, payload)
}

func (o *GetSyncStatusOK) GetPayload() *models.ActiveSyncStatusDTO {
	return o.Payload
}

func (o *GetSyncStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveSyncStatusDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyncStatusUnauthorized creates a GetSyncStatusUnauthorized with default headers values
func NewGetSyncStatusUnauthorized() *GetSyncStatusUnauthorized {
	return &GetSyncStatusUnauthorized{}
}

/*
GetSyncStatusUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetSyncStatusUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get sync status unauthorized response has a 2xx status code
func (o *GetSyncStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sync status unauthorized response has a 3xx status code
func (o *GetSyncStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sync status unauthorized response has a 4xx status code
func (o *GetSyncStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sync status unauthorized response has a 5xx status code
func (o *GetSyncStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get sync status unauthorized response a status code equal to that given
func (o *GetSyncStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get sync status unauthorized response
func (o *GetSyncStatusUnauthorized) Code() int {
	return 401
}

func (o *GetSyncStatusUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusUnauthorized %s", 401, payload)
}

func (o *GetSyncStatusUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusUnauthorized %s", 401, payload)
}

func (o *GetSyncStatusUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSyncStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyncStatusForbidden creates a GetSyncStatusForbidden with default headers values
func NewGetSyncStatusForbidden() *GetSyncStatusForbidden {
	return &GetSyncStatusForbidden{}
}

/*
GetSyncStatusForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetSyncStatusForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get sync status forbidden response has a 2xx status code
func (o *GetSyncStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sync status forbidden response has a 3xx status code
func (o *GetSyncStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sync status forbidden response has a 4xx status code
func (o *GetSyncStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sync status forbidden response has a 5xx status code
func (o *GetSyncStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get sync status forbidden response a status code equal to that given
func (o *GetSyncStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get sync status forbidden response
func (o *GetSyncStatusForbidden) Code() int {
	return 403
}

func (o *GetSyncStatusForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusForbidden %s", 403, payload)
}

func (o *GetSyncStatusForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusForbidden %s", 403, payload)
}

func (o *GetSyncStatusForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSyncStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyncStatusInternalServerError creates a GetSyncStatusInternalServerError with default headers values
func NewGetSyncStatusInternalServerError() *GetSyncStatusInternalServerError {
	return &GetSyncStatusInternalServerError{}
}

/*
GetSyncStatusInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetSyncStatusInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get sync status internal server error response has a 2xx status code
func (o *GetSyncStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sync status internal server error response has a 3xx status code
func (o *GetSyncStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sync status internal server error response has a 4xx status code
func (o *GetSyncStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sync status internal server error response has a 5xx status code
func (o *GetSyncStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sync status internal server error response a status code equal to that given
func (o *GetSyncStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get sync status internal server error response
func (o *GetSyncStatusInternalServerError) Code() int {
	return 500
}

func (o *GetSyncStatusInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusInternalServerError %s", 500, payload)
}

func (o *GetSyncStatusInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusInternalServerError %s", 500, payload)
}

func (o *GetSyncStatusInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSyncStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
