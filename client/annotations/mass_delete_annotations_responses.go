// Code generated by go-swagger; DO NOT EDIT.

package annotations

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

// MassDeleteAnnotationsReader is a Reader for the MassDeleteAnnotations structure.
type MassDeleteAnnotationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MassDeleteAnnotationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMassDeleteAnnotationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewMassDeleteAnnotationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMassDeleteAnnotationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /annotations/mass-delete] massDeleteAnnotations", response, response.Code())
	}
}

// NewMassDeleteAnnotationsOK creates a MassDeleteAnnotationsOK with default headers values
func NewMassDeleteAnnotationsOK() *MassDeleteAnnotationsOK {
	return &MassDeleteAnnotationsOK{}
}

/*
MassDeleteAnnotationsOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type MassDeleteAnnotationsOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this mass delete annotations Ok response has a 2xx status code
func (o *MassDeleteAnnotationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mass delete annotations Ok response has a 3xx status code
func (o *MassDeleteAnnotationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mass delete annotations Ok response has a 4xx status code
func (o *MassDeleteAnnotationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mass delete annotations Ok response has a 5xx status code
func (o *MassDeleteAnnotationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mass delete annotations Ok response a status code equal to that given
func (o *MassDeleteAnnotationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mass delete annotations Ok response
func (o *MassDeleteAnnotationsOK) Code() int {
	return 200
}

func (o *MassDeleteAnnotationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /annotations/mass-delete][%d] massDeleteAnnotationsOk %s", 200, payload)
}

func (o *MassDeleteAnnotationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /annotations/mass-delete][%d] massDeleteAnnotationsOk %s", 200, payload)
}

func (o *MassDeleteAnnotationsOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *MassDeleteAnnotationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMassDeleteAnnotationsUnauthorized creates a MassDeleteAnnotationsUnauthorized with default headers values
func NewMassDeleteAnnotationsUnauthorized() *MassDeleteAnnotationsUnauthorized {
	return &MassDeleteAnnotationsUnauthorized{}
}

/*
MassDeleteAnnotationsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type MassDeleteAnnotationsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this mass delete annotations unauthorized response has a 2xx status code
func (o *MassDeleteAnnotationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this mass delete annotations unauthorized response has a 3xx status code
func (o *MassDeleteAnnotationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mass delete annotations unauthorized response has a 4xx status code
func (o *MassDeleteAnnotationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this mass delete annotations unauthorized response has a 5xx status code
func (o *MassDeleteAnnotationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this mass delete annotations unauthorized response a status code equal to that given
func (o *MassDeleteAnnotationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the mass delete annotations unauthorized response
func (o *MassDeleteAnnotationsUnauthorized) Code() int {
	return 401
}

func (o *MassDeleteAnnotationsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /annotations/mass-delete][%d] massDeleteAnnotationsUnauthorized %s", 401, payload)
}

func (o *MassDeleteAnnotationsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /annotations/mass-delete][%d] massDeleteAnnotationsUnauthorized %s", 401, payload)
}

func (o *MassDeleteAnnotationsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MassDeleteAnnotationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMassDeleteAnnotationsInternalServerError creates a MassDeleteAnnotationsInternalServerError with default headers values
func NewMassDeleteAnnotationsInternalServerError() *MassDeleteAnnotationsInternalServerError {
	return &MassDeleteAnnotationsInternalServerError{}
}

/*
MassDeleteAnnotationsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type MassDeleteAnnotationsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this mass delete annotations internal server error response has a 2xx status code
func (o *MassDeleteAnnotationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this mass delete annotations internal server error response has a 3xx status code
func (o *MassDeleteAnnotationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mass delete annotations internal server error response has a 4xx status code
func (o *MassDeleteAnnotationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this mass delete annotations internal server error response has a 5xx status code
func (o *MassDeleteAnnotationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this mass delete annotations internal server error response a status code equal to that given
func (o *MassDeleteAnnotationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the mass delete annotations internal server error response
func (o *MassDeleteAnnotationsInternalServerError) Code() int {
	return 500
}

func (o *MassDeleteAnnotationsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /annotations/mass-delete][%d] massDeleteAnnotationsInternalServerError %s", 500, payload)
}

func (o *MassDeleteAnnotationsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /annotations/mass-delete][%d] massDeleteAnnotationsInternalServerError %s", 500, payload)
}

func (o *MassDeleteAnnotationsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MassDeleteAnnotationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
