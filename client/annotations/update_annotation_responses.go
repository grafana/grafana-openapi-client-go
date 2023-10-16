// Code generated by go-swagger; DO NOT EDIT.

package annotations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// UpdateAnnotationReader is a Reader for the UpdateAnnotation structure.
type UpdateAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAnnotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAnnotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAnnotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAnnotationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /annotations/{annotation_id}] updateAnnotation", response, response.Code())
	}
}

// NewUpdateAnnotationOK creates a UpdateAnnotationOK with default headers values
func NewUpdateAnnotationOK() *UpdateAnnotationOK {
	return &UpdateAnnotationOK{}
}

/*
UpdateAnnotationOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateAnnotationOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this update annotation o k response has a 2xx status code
func (o *UpdateAnnotationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update annotation o k response has a 3xx status code
func (o *UpdateAnnotationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update annotation o k response has a 4xx status code
func (o *UpdateAnnotationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update annotation o k response has a 5xx status code
func (o *UpdateAnnotationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update annotation o k response a status code equal to that given
func (o *UpdateAnnotationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update annotation o k response
func (o *UpdateAnnotationOK) Code() int {
	return 200
}

func (o *UpdateAnnotationOK) Error() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationOK  %+v", 200, o.Payload)
}

func (o *UpdateAnnotationOK) String() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationOK  %+v", 200, o.Payload)
}

func (o *UpdateAnnotationOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAnnotationBadRequest creates a UpdateAnnotationBadRequest with default headers values
func NewUpdateAnnotationBadRequest() *UpdateAnnotationBadRequest {
	return &UpdateAnnotationBadRequest{}
}

/*
UpdateAnnotationBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateAnnotationBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update annotation bad request response has a 2xx status code
func (o *UpdateAnnotationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update annotation bad request response has a 3xx status code
func (o *UpdateAnnotationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update annotation bad request response has a 4xx status code
func (o *UpdateAnnotationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update annotation bad request response has a 5xx status code
func (o *UpdateAnnotationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update annotation bad request response a status code equal to that given
func (o *UpdateAnnotationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update annotation bad request response
func (o *UpdateAnnotationBadRequest) Code() int {
	return 400
}

func (o *UpdateAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAnnotationBadRequest) String() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAnnotationBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAnnotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAnnotationUnauthorized creates a UpdateAnnotationUnauthorized with default headers values
func NewUpdateAnnotationUnauthorized() *UpdateAnnotationUnauthorized {
	return &UpdateAnnotationUnauthorized{}
}

/*
UpdateAnnotationUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateAnnotationUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update annotation unauthorized response has a 2xx status code
func (o *UpdateAnnotationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update annotation unauthorized response has a 3xx status code
func (o *UpdateAnnotationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update annotation unauthorized response has a 4xx status code
func (o *UpdateAnnotationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update annotation unauthorized response has a 5xx status code
func (o *UpdateAnnotationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update annotation unauthorized response a status code equal to that given
func (o *UpdateAnnotationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update annotation unauthorized response
func (o *UpdateAnnotationUnauthorized) Code() int {
	return 401
}

func (o *UpdateAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAnnotationUnauthorized) String() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAnnotationUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAnnotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAnnotationForbidden creates a UpdateAnnotationForbidden with default headers values
func NewUpdateAnnotationForbidden() *UpdateAnnotationForbidden {
	return &UpdateAnnotationForbidden{}
}

/*
UpdateAnnotationForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateAnnotationForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update annotation forbidden response has a 2xx status code
func (o *UpdateAnnotationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update annotation forbidden response has a 3xx status code
func (o *UpdateAnnotationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update annotation forbidden response has a 4xx status code
func (o *UpdateAnnotationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update annotation forbidden response has a 5xx status code
func (o *UpdateAnnotationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update annotation forbidden response a status code equal to that given
func (o *UpdateAnnotationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update annotation forbidden response
func (o *UpdateAnnotationForbidden) Code() int {
	return 403
}

func (o *UpdateAnnotationForbidden) Error() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAnnotationForbidden) String() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAnnotationForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAnnotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAnnotationInternalServerError creates a UpdateAnnotationInternalServerError with default headers values
func NewUpdateAnnotationInternalServerError() *UpdateAnnotationInternalServerError {
	return &UpdateAnnotationInternalServerError{}
}

/*
UpdateAnnotationInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateAnnotationInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update annotation internal server error response has a 2xx status code
func (o *UpdateAnnotationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update annotation internal server error response has a 3xx status code
func (o *UpdateAnnotationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update annotation internal server error response has a 4xx status code
func (o *UpdateAnnotationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update annotation internal server error response has a 5xx status code
func (o *UpdateAnnotationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update annotation internal server error response a status code equal to that given
func (o *UpdateAnnotationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update annotation internal server error response
func (o *UpdateAnnotationInternalServerError) Code() int {
	return 500
}

func (o *UpdateAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateAnnotationInternalServerError) String() string {
	return fmt.Sprintf("[PUT /annotations/{annotation_id}][%d] updateAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateAnnotationInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAnnotationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}