// Code generated by go-swagger; DO NOT EDIT.

package provisioning

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

// GetTemplateReader is a Reader for the GetTemplate structure.
type GetTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/provisioning/templates/{name}] GetTemplate", response, response.Code())
	}
}

// NewGetTemplateOK creates a GetTemplateOK with default headers values
func NewGetTemplateOK() *GetTemplateOK {
	return &GetTemplateOK{}
}

/*
GetTemplateOK describes a response with status code 200, with default header values.

NotificationTemplate
*/
type GetTemplateOK struct {
	Payload *models.NotificationTemplate
}

// IsSuccess returns true when this get template Ok response has a 2xx status code
func (o *GetTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get template Ok response has a 3xx status code
func (o *GetTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get template Ok response has a 4xx status code
func (o *GetTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get template Ok response has a 5xx status code
func (o *GetTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get template Ok response a status code equal to that given
func (o *GetTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get template Ok response
func (o *GetTemplateOK) Code() int {
	return 200
}

func (o *GetTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/templates/{name}][%d] getTemplateOk %s", 200, payload)
}

func (o *GetTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/templates/{name}][%d] getTemplateOk %s", 200, payload)
}

func (o *GetTemplateOK) GetPayload() *models.NotificationTemplate {
	return o.Payload
}

func (o *GetTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTemplateNotFound creates a GetTemplateNotFound with default headers values
func NewGetTemplateNotFound() *GetTemplateNotFound {
	return &GetTemplateNotFound{}
}

/*
GetTemplateNotFound describes a response with status code 404, with default header values.

PublicError
*/
type GetTemplateNotFound struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this get template not found response has a 2xx status code
func (o *GetTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get template not found response has a 3xx status code
func (o *GetTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get template not found response has a 4xx status code
func (o *GetTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get template not found response has a 5xx status code
func (o *GetTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get template not found response a status code equal to that given
func (o *GetTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get template not found response
func (o *GetTemplateNotFound) Code() int {
	return 404
}

func (o *GetTemplateNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/templates/{name}][%d] getTemplateNotFound %s", 404, payload)
}

func (o *GetTemplateNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/templates/{name}][%d] getTemplateNotFound %s", 404, payload)
}

func (o *GetTemplateNotFound) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *GetTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
