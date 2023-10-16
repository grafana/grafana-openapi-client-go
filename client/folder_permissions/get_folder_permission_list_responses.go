// Code generated by go-swagger; DO NOT EDIT.

package folder_permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetFolderPermissionListReader is a Reader for the GetFolderPermissionList structure.
type GetFolderPermissionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFolderPermissionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFolderPermissionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFolderPermissionListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFolderPermissionListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFolderPermissionListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFolderPermissionListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /folders/{folder_uid}/permissions] getFolderPermissionList", response, response.Code())
	}
}

// NewGetFolderPermissionListOK creates a GetFolderPermissionListOK with default headers values
func NewGetFolderPermissionListOK() *GetFolderPermissionListOK {
	return &GetFolderPermissionListOK{}
}

/*
GetFolderPermissionListOK describes a response with status code 200, with default header values.

(empty)
*/
type GetFolderPermissionListOK struct {
	Payload []*models.DashboardACLInfoDTO
}

// IsSuccess returns true when this get folder permission list o k response has a 2xx status code
func (o *GetFolderPermissionListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get folder permission list o k response has a 3xx status code
func (o *GetFolderPermissionListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder permission list o k response has a 4xx status code
func (o *GetFolderPermissionListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folder permission list o k response has a 5xx status code
func (o *GetFolderPermissionListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder permission list o k response a status code equal to that given
func (o *GetFolderPermissionListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get folder permission list o k response
func (o *GetFolderPermissionListOK) Code() int {
	return 200
}

func (o *GetFolderPermissionListOK) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListOK  %+v", 200, o.Payload)
}

func (o *GetFolderPermissionListOK) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListOK  %+v", 200, o.Payload)
}

func (o *GetFolderPermissionListOK) GetPayload() []*models.DashboardACLInfoDTO {
	return o.Payload
}

func (o *GetFolderPermissionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderPermissionListUnauthorized creates a GetFolderPermissionListUnauthorized with default headers values
func NewGetFolderPermissionListUnauthorized() *GetFolderPermissionListUnauthorized {
	return &GetFolderPermissionListUnauthorized{}
}

/*
GetFolderPermissionListUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetFolderPermissionListUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder permission list unauthorized response has a 2xx status code
func (o *GetFolderPermissionListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder permission list unauthorized response has a 3xx status code
func (o *GetFolderPermissionListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder permission list unauthorized response has a 4xx status code
func (o *GetFolderPermissionListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder permission list unauthorized response has a 5xx status code
func (o *GetFolderPermissionListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder permission list unauthorized response a status code equal to that given
func (o *GetFolderPermissionListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get folder permission list unauthorized response
func (o *GetFolderPermissionListUnauthorized) Code() int {
	return 401
}

func (o *GetFolderPermissionListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFolderPermissionListUnauthorized) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFolderPermissionListUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderPermissionListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderPermissionListForbidden creates a GetFolderPermissionListForbidden with default headers values
func NewGetFolderPermissionListForbidden() *GetFolderPermissionListForbidden {
	return &GetFolderPermissionListForbidden{}
}

/*
GetFolderPermissionListForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetFolderPermissionListForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder permission list forbidden response has a 2xx status code
func (o *GetFolderPermissionListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder permission list forbidden response has a 3xx status code
func (o *GetFolderPermissionListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder permission list forbidden response has a 4xx status code
func (o *GetFolderPermissionListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder permission list forbidden response has a 5xx status code
func (o *GetFolderPermissionListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder permission list forbidden response a status code equal to that given
func (o *GetFolderPermissionListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get folder permission list forbidden response
func (o *GetFolderPermissionListForbidden) Code() int {
	return 403
}

func (o *GetFolderPermissionListForbidden) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListForbidden  %+v", 403, o.Payload)
}

func (o *GetFolderPermissionListForbidden) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListForbidden  %+v", 403, o.Payload)
}

func (o *GetFolderPermissionListForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderPermissionListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderPermissionListNotFound creates a GetFolderPermissionListNotFound with default headers values
func NewGetFolderPermissionListNotFound() *GetFolderPermissionListNotFound {
	return &GetFolderPermissionListNotFound{}
}

/*
GetFolderPermissionListNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetFolderPermissionListNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder permission list not found response has a 2xx status code
func (o *GetFolderPermissionListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder permission list not found response has a 3xx status code
func (o *GetFolderPermissionListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder permission list not found response has a 4xx status code
func (o *GetFolderPermissionListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder permission list not found response has a 5xx status code
func (o *GetFolderPermissionListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder permission list not found response a status code equal to that given
func (o *GetFolderPermissionListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get folder permission list not found response
func (o *GetFolderPermissionListNotFound) Code() int {
	return 404
}

func (o *GetFolderPermissionListNotFound) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListNotFound  %+v", 404, o.Payload)
}

func (o *GetFolderPermissionListNotFound) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListNotFound  %+v", 404, o.Payload)
}

func (o *GetFolderPermissionListNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderPermissionListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderPermissionListInternalServerError creates a GetFolderPermissionListInternalServerError with default headers values
func NewGetFolderPermissionListInternalServerError() *GetFolderPermissionListInternalServerError {
	return &GetFolderPermissionListInternalServerError{}
}

/*
GetFolderPermissionListInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetFolderPermissionListInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder permission list internal server error response has a 2xx status code
func (o *GetFolderPermissionListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder permission list internal server error response has a 3xx status code
func (o *GetFolderPermissionListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder permission list internal server error response has a 4xx status code
func (o *GetFolderPermissionListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folder permission list internal server error response has a 5xx status code
func (o *GetFolderPermissionListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get folder permission list internal server error response a status code equal to that given
func (o *GetFolderPermissionListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get folder permission list internal server error response
func (o *GetFolderPermissionListInternalServerError) Code() int {
	return 500
}

func (o *GetFolderPermissionListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFolderPermissionListInternalServerError) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}/permissions][%d] getFolderPermissionListInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFolderPermissionListInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderPermissionListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}