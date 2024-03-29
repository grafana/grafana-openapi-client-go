// Code generated by go-swagger; DO NOT EDIT.

package correlations

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

// UpdateCorrelationReader is a Reader for the UpdateCorrelation structure.
type UpdateCorrelationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCorrelationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCorrelationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCorrelationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCorrelationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCorrelationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCorrelationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCorrelationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}] updateCorrelation", response, response.Code())
	}
}

// NewUpdateCorrelationOK creates a UpdateCorrelationOK with default headers values
func NewUpdateCorrelationOK() *UpdateCorrelationOK {
	return &UpdateCorrelationOK{}
}

/*
UpdateCorrelationOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateCorrelationOK struct {
	Payload *models.UpdateCorrelationResponseBody
}

// IsSuccess returns true when this update correlation Ok response has a 2xx status code
func (o *UpdateCorrelationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update correlation Ok response has a 3xx status code
func (o *UpdateCorrelationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update correlation Ok response has a 4xx status code
func (o *UpdateCorrelationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update correlation Ok response has a 5xx status code
func (o *UpdateCorrelationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update correlation Ok response a status code equal to that given
func (o *UpdateCorrelationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update correlation Ok response
func (o *UpdateCorrelationOK) Code() int {
	return 200
}

func (o *UpdateCorrelationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationOk %s", 200, payload)
}

func (o *UpdateCorrelationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationOk %s", 200, payload)
}

func (o *UpdateCorrelationOK) GetPayload() *models.UpdateCorrelationResponseBody {
	return o.Payload
}

func (o *UpdateCorrelationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateCorrelationResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCorrelationBadRequest creates a UpdateCorrelationBadRequest with default headers values
func NewUpdateCorrelationBadRequest() *UpdateCorrelationBadRequest {
	return &UpdateCorrelationBadRequest{}
}

/*
UpdateCorrelationBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateCorrelationBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update correlation bad request response has a 2xx status code
func (o *UpdateCorrelationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update correlation bad request response has a 3xx status code
func (o *UpdateCorrelationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update correlation bad request response has a 4xx status code
func (o *UpdateCorrelationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update correlation bad request response has a 5xx status code
func (o *UpdateCorrelationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update correlation bad request response a status code equal to that given
func (o *UpdateCorrelationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update correlation bad request response
func (o *UpdateCorrelationBadRequest) Code() int {
	return 400
}

func (o *UpdateCorrelationBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationBadRequest %s", 400, payload)
}

func (o *UpdateCorrelationBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationBadRequest %s", 400, payload)
}

func (o *UpdateCorrelationBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCorrelationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCorrelationUnauthorized creates a UpdateCorrelationUnauthorized with default headers values
func NewUpdateCorrelationUnauthorized() *UpdateCorrelationUnauthorized {
	return &UpdateCorrelationUnauthorized{}
}

/*
UpdateCorrelationUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateCorrelationUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update correlation unauthorized response has a 2xx status code
func (o *UpdateCorrelationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update correlation unauthorized response has a 3xx status code
func (o *UpdateCorrelationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update correlation unauthorized response has a 4xx status code
func (o *UpdateCorrelationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update correlation unauthorized response has a 5xx status code
func (o *UpdateCorrelationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update correlation unauthorized response a status code equal to that given
func (o *UpdateCorrelationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update correlation unauthorized response
func (o *UpdateCorrelationUnauthorized) Code() int {
	return 401
}

func (o *UpdateCorrelationUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationUnauthorized %s", 401, payload)
}

func (o *UpdateCorrelationUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationUnauthorized %s", 401, payload)
}

func (o *UpdateCorrelationUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCorrelationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCorrelationForbidden creates a UpdateCorrelationForbidden with default headers values
func NewUpdateCorrelationForbidden() *UpdateCorrelationForbidden {
	return &UpdateCorrelationForbidden{}
}

/*
UpdateCorrelationForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateCorrelationForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update correlation forbidden response has a 2xx status code
func (o *UpdateCorrelationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update correlation forbidden response has a 3xx status code
func (o *UpdateCorrelationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update correlation forbidden response has a 4xx status code
func (o *UpdateCorrelationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update correlation forbidden response has a 5xx status code
func (o *UpdateCorrelationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update correlation forbidden response a status code equal to that given
func (o *UpdateCorrelationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update correlation forbidden response
func (o *UpdateCorrelationForbidden) Code() int {
	return 403
}

func (o *UpdateCorrelationForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationForbidden %s", 403, payload)
}

func (o *UpdateCorrelationForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationForbidden %s", 403, payload)
}

func (o *UpdateCorrelationForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCorrelationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCorrelationNotFound creates a UpdateCorrelationNotFound with default headers values
func NewUpdateCorrelationNotFound() *UpdateCorrelationNotFound {
	return &UpdateCorrelationNotFound{}
}

/*
UpdateCorrelationNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateCorrelationNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update correlation not found response has a 2xx status code
func (o *UpdateCorrelationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update correlation not found response has a 3xx status code
func (o *UpdateCorrelationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update correlation not found response has a 4xx status code
func (o *UpdateCorrelationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update correlation not found response has a 5xx status code
func (o *UpdateCorrelationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update correlation not found response a status code equal to that given
func (o *UpdateCorrelationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update correlation not found response
func (o *UpdateCorrelationNotFound) Code() int {
	return 404
}

func (o *UpdateCorrelationNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationNotFound %s", 404, payload)
}

func (o *UpdateCorrelationNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationNotFound %s", 404, payload)
}

func (o *UpdateCorrelationNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCorrelationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCorrelationInternalServerError creates a UpdateCorrelationInternalServerError with default headers values
func NewUpdateCorrelationInternalServerError() *UpdateCorrelationInternalServerError {
	return &UpdateCorrelationInternalServerError{}
}

/*
UpdateCorrelationInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateCorrelationInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update correlation internal server error response has a 2xx status code
func (o *UpdateCorrelationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update correlation internal server error response has a 3xx status code
func (o *UpdateCorrelationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update correlation internal server error response has a 4xx status code
func (o *UpdateCorrelationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update correlation internal server error response has a 5xx status code
func (o *UpdateCorrelationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update correlation internal server error response a status code equal to that given
func (o *UpdateCorrelationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update correlation internal server error response
func (o *UpdateCorrelationInternalServerError) Code() int {
	return 500
}

func (o *UpdateCorrelationInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationInternalServerError %s", 500, payload)
}

func (o *UpdateCorrelationInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] updateCorrelationInternalServerError %s", 500, payload)
}

func (o *UpdateCorrelationInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCorrelationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
