// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetAlertRuleReader is a Reader for the GetAlertRule structure.
type GetAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAlertRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/provisioning/alert-rules/{UID}] GetAlertRule", response, response.Code())
	}
}

// NewGetAlertRuleOK creates a GetAlertRuleOK with default headers values
func NewGetAlertRuleOK() *GetAlertRuleOK {
	return &GetAlertRuleOK{}
}

/*
GetAlertRuleOK describes a response with status code 200, with default header values.

ProvisionedAlertRule
*/
type GetAlertRuleOK struct {
	Payload *models.ProvisionedAlertRule
}

// IsSuccess returns true when this get alert rule Ok response has a 2xx status code
func (o *GetAlertRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert rule Ok response has a 3xx status code
func (o *GetAlertRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule Ok response has a 4xx status code
func (o *GetAlertRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert rule Ok response has a 5xx status code
func (o *GetAlertRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule Ok response a status code equal to that given
func (o *GetAlertRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert rule Ok response
func (o *GetAlertRuleOK) Code() int {
	return 200
}

func (o *GetAlertRuleOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/alert-rules/{UID}][%d] getAlertRuleOk  %+v", 200, o.Payload)
}

func (o *GetAlertRuleOK) String() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/alert-rules/{UID}][%d] getAlertRuleOk  %+v", 200, o.Payload)
}

func (o *GetAlertRuleOK) GetPayload() *models.ProvisionedAlertRule {
	return o.Payload
}

func (o *GetAlertRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvisionedAlertRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRuleNotFound creates a GetAlertRuleNotFound with default headers values
func NewGetAlertRuleNotFound() *GetAlertRuleNotFound {
	return &GetAlertRuleNotFound{}
}

/*
GetAlertRuleNotFound describes a response with status code 404, with default header values.

	Not found.
*/
type GetAlertRuleNotFound struct {
}

// IsSuccess returns true when this get alert rule not found response has a 2xx status code
func (o *GetAlertRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert rule not found response has a 3xx status code
func (o *GetAlertRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule not found response has a 4xx status code
func (o *GetAlertRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert rule not found response has a 5xx status code
func (o *GetAlertRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule not found response a status code equal to that given
func (o *GetAlertRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get alert rule not found response
func (o *GetAlertRuleNotFound) Code() int {
	return 404
}

func (o *GetAlertRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/alert-rules/{UID}][%d] getAlertRuleNotFound ", 404)
}

func (o *GetAlertRuleNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/alert-rules/{UID}][%d] getAlertRuleNotFound ", 404)
}

func (o *GetAlertRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
