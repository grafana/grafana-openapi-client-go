// Code generated by go-swagger; DO NOT EDIT.

package folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetFolderByUIDReader is a Reader for the GetFolderByUID structure.
type GetFolderByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFolderByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFolderByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFolderByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFolderByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFolderByUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFolderByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /folders/{folder_uid}] getFolderByUID", response, response.Code())
	}
}

// NewGetFolderByUIDOK creates a GetFolderByUIDOK with default headers values
func NewGetFolderByUIDOK() *GetFolderByUIDOK {
	return &GetFolderByUIDOK{}
}

/*
GetFolderByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetFolderByUIDOK struct {
	Payload *models.Folder
}

// IsSuccess returns true when this get folder by Uid Ok response has a 2xx status code
func (o *GetFolderByUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get folder by Uid Ok response has a 3xx status code
func (o *GetFolderByUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Uid Ok response has a 4xx status code
func (o *GetFolderByUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folder by Uid Ok response has a 5xx status code
func (o *GetFolderByUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder by Uid Ok response a status code equal to that given
func (o *GetFolderByUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get folder by Uid Ok response
func (o *GetFolderByUIDOK) Code() int {
	return 200
}

func (o *GetFolderByUIDOK) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidOk  %+v", 200, o.Payload)
}

func (o *GetFolderByUIDOK) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidOk  %+v", 200, o.Payload)
}

func (o *GetFolderByUIDOK) GetPayload() *models.Folder {
	return o.Payload
}

func (o *GetFolderByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderByUIDUnauthorized creates a GetFolderByUIDUnauthorized with default headers values
func NewGetFolderByUIDUnauthorized() *GetFolderByUIDUnauthorized {
	return &GetFolderByUIDUnauthorized{}
}

/*
GetFolderByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetFolderByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder by Uid unauthorized response has a 2xx status code
func (o *GetFolderByUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder by Uid unauthorized response has a 3xx status code
func (o *GetFolderByUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Uid unauthorized response has a 4xx status code
func (o *GetFolderByUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder by Uid unauthorized response has a 5xx status code
func (o *GetFolderByUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder by Uid unauthorized response a status code equal to that given
func (o *GetFolderByUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get folder by Uid unauthorized response
func (o *GetFolderByUIDUnauthorized) Code() int {
	return 401
}

func (o *GetFolderByUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFolderByUIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFolderByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderByUIDForbidden creates a GetFolderByUIDForbidden with default headers values
func NewGetFolderByUIDForbidden() *GetFolderByUIDForbidden {
	return &GetFolderByUIDForbidden{}
}

/*
GetFolderByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetFolderByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder by Uid forbidden response has a 2xx status code
func (o *GetFolderByUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder by Uid forbidden response has a 3xx status code
func (o *GetFolderByUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Uid forbidden response has a 4xx status code
func (o *GetFolderByUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder by Uid forbidden response has a 5xx status code
func (o *GetFolderByUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder by Uid forbidden response a status code equal to that given
func (o *GetFolderByUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get folder by Uid forbidden response
func (o *GetFolderByUIDForbidden) Code() int {
	return 403
}

func (o *GetFolderByUIDForbidden) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidForbidden  %+v", 403, o.Payload)
}

func (o *GetFolderByUIDForbidden) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidForbidden  %+v", 403, o.Payload)
}

func (o *GetFolderByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderByUIDNotFound creates a GetFolderByUIDNotFound with default headers values
func NewGetFolderByUIDNotFound() *GetFolderByUIDNotFound {
	return &GetFolderByUIDNotFound{}
}

/*
GetFolderByUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetFolderByUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder by Uid not found response has a 2xx status code
func (o *GetFolderByUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder by Uid not found response has a 3xx status code
func (o *GetFolderByUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Uid not found response has a 4xx status code
func (o *GetFolderByUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder by Uid not found response has a 5xx status code
func (o *GetFolderByUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder by Uid not found response a status code equal to that given
func (o *GetFolderByUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get folder by Uid not found response
func (o *GetFolderByUIDNotFound) Code() int {
	return 404
}

func (o *GetFolderByUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidNotFound  %+v", 404, o.Payload)
}

func (o *GetFolderByUIDNotFound) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidNotFound  %+v", 404, o.Payload)
}

func (o *GetFolderByUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderByUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderByUIDInternalServerError creates a GetFolderByUIDInternalServerError with default headers values
func NewGetFolderByUIDInternalServerError() *GetFolderByUIDInternalServerError {
	return &GetFolderByUIDInternalServerError{}
}

/*
GetFolderByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetFolderByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder by Uid internal server error response has a 2xx status code
func (o *GetFolderByUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder by Uid internal server error response has a 3xx status code
func (o *GetFolderByUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Uid internal server error response has a 4xx status code
func (o *GetFolderByUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folder by Uid internal server error response has a 5xx status code
func (o *GetFolderByUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get folder by Uid internal server error response a status code equal to that given
func (o *GetFolderByUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get folder by Uid internal server error response
func (o *GetFolderByUIDInternalServerError) Code() int {
	return 500
}

func (o *GetFolderByUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFolderByUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /folders/{folder_uid}][%d] getFolderByUidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFolderByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
