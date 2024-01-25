// Code generated by go-swagger; DO NOT EDIT.

package folders

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

// GetFoldersReader is a Reader for the GetFolders structure.
type GetFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFoldersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFoldersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFoldersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFoldersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /folders] getFolders", response, response.Code())
	}
}

// NewGetFoldersOK creates a GetFoldersOK with default headers values
func NewGetFoldersOK() *GetFoldersOK {
	return &GetFoldersOK{}
}

/*
GetFoldersOK describes a response with status code 200, with default header values.

(empty)
*/
type GetFoldersOK struct {
	Payload []*models.FolderSearchHit
}

// IsSuccess returns true when this get folders Ok response has a 2xx status code
func (o *GetFoldersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get folders Ok response has a 3xx status code
func (o *GetFoldersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folders Ok response has a 4xx status code
func (o *GetFoldersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folders Ok response has a 5xx status code
func (o *GetFoldersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get folders Ok response a status code equal to that given
func (o *GetFoldersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get folders Ok response
func (o *GetFoldersOK) Code() int {
	return 200
}

func (o *GetFoldersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /folders][%d] getFoldersOk %s", 200, payload)
}

func (o *GetFoldersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /folders][%d] getFoldersOk %s", 200, payload)
}

func (o *GetFoldersOK) GetPayload() []*models.FolderSearchHit {
	return o.Payload
}

func (o *GetFoldersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersUnauthorized creates a GetFoldersUnauthorized with default headers values
func NewGetFoldersUnauthorized() *GetFoldersUnauthorized {
	return &GetFoldersUnauthorized{}
}

/*
GetFoldersUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetFoldersUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folders unauthorized response has a 2xx status code
func (o *GetFoldersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folders unauthorized response has a 3xx status code
func (o *GetFoldersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folders unauthorized response has a 4xx status code
func (o *GetFoldersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folders unauthorized response has a 5xx status code
func (o *GetFoldersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get folders unauthorized response a status code equal to that given
func (o *GetFoldersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get folders unauthorized response
func (o *GetFoldersUnauthorized) Code() int {
	return 401
}

func (o *GetFoldersUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /folders][%d] getFoldersUnauthorized %s", 401, payload)
}

func (o *GetFoldersUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /folders][%d] getFoldersUnauthorized %s", 401, payload)
}

func (o *GetFoldersUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFoldersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersForbidden creates a GetFoldersForbidden with default headers values
func NewGetFoldersForbidden() *GetFoldersForbidden {
	return &GetFoldersForbidden{}
}

/*
GetFoldersForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetFoldersForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folders forbidden response has a 2xx status code
func (o *GetFoldersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folders forbidden response has a 3xx status code
func (o *GetFoldersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folders forbidden response has a 4xx status code
func (o *GetFoldersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folders forbidden response has a 5xx status code
func (o *GetFoldersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get folders forbidden response a status code equal to that given
func (o *GetFoldersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get folders forbidden response
func (o *GetFoldersForbidden) Code() int {
	return 403
}

func (o *GetFoldersForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /folders][%d] getFoldersForbidden %s", 403, payload)
}

func (o *GetFoldersForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /folders][%d] getFoldersForbidden %s", 403, payload)
}

func (o *GetFoldersForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFoldersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersInternalServerError creates a GetFoldersInternalServerError with default headers values
func NewGetFoldersInternalServerError() *GetFoldersInternalServerError {
	return &GetFoldersInternalServerError{}
}

/*
GetFoldersInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetFoldersInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folders internal server error response has a 2xx status code
func (o *GetFoldersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folders internal server error response has a 3xx status code
func (o *GetFoldersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folders internal server error response has a 4xx status code
func (o *GetFoldersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folders internal server error response has a 5xx status code
func (o *GetFoldersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get folders internal server error response a status code equal to that given
func (o *GetFoldersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get folders internal server error response
func (o *GetFoldersInternalServerError) Code() int {
	return 500
}

func (o *GetFoldersInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /folders][%d] getFoldersInternalServerError %s", 500, payload)
}

func (o *GetFoldersInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /folders][%d] getFoldersInternalServerError %s", 500, payload)
}

func (o *GetFoldersInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFoldersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
