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

// GetPolicyTreeExportReader is a Reader for the GetPolicyTreeExport structure.
type GetPolicyTreeExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyTreeExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyTreeExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPolicyTreeExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/provisioning/policies/export] GetPolicyTreeExport", response, response.Code())
	}
}

// NewGetPolicyTreeExportOK creates a GetPolicyTreeExportOK with default headers values
func NewGetPolicyTreeExportOK() *GetPolicyTreeExportOK {
	return &GetPolicyTreeExportOK{}
}

/*
GetPolicyTreeExportOK describes a response with status code 200, with default header values.

AlertingFileExport
*/
type GetPolicyTreeExportOK struct {
	Payload *models.AlertingFileExport
}

// IsSuccess returns true when this get policy tree export Ok response has a 2xx status code
func (o *GetPolicyTreeExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get policy tree export Ok response has a 3xx status code
func (o *GetPolicyTreeExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy tree export Ok response has a 4xx status code
func (o *GetPolicyTreeExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get policy tree export Ok response has a 5xx status code
func (o *GetPolicyTreeExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy tree export Ok response a status code equal to that given
func (o *GetPolicyTreeExportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get policy tree export Ok response
func (o *GetPolicyTreeExportOK) Code() int {
	return 200
}

func (o *GetPolicyTreeExportOK) Error() string {
	return fmt.Sprintf("[GET /v1/provisioning/policies/export][%d] getPolicyTreeExportOk  %+v", 200, o.Payload)
}

func (o *GetPolicyTreeExportOK) String() string {
	return fmt.Sprintf("[GET /v1/provisioning/policies/export][%d] getPolicyTreeExportOk  %+v", 200, o.Payload)
}

func (o *GetPolicyTreeExportOK) GetPayload() *models.AlertingFileExport {
	return o.Payload
}

func (o *GetPolicyTreeExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingFileExport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyTreeExportNotFound creates a GetPolicyTreeExportNotFound with default headers values
func NewGetPolicyTreeExportNotFound() *GetPolicyTreeExportNotFound {
	return &GetPolicyTreeExportNotFound{}
}

/*
GetPolicyTreeExportNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetPolicyTreeExportNotFound struct {
	Payload models.NotFound
}

// IsSuccess returns true when this get policy tree export not found response has a 2xx status code
func (o *GetPolicyTreeExportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy tree export not found response has a 3xx status code
func (o *GetPolicyTreeExportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy tree export not found response has a 4xx status code
func (o *GetPolicyTreeExportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy tree export not found response has a 5xx status code
func (o *GetPolicyTreeExportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy tree export not found response a status code equal to that given
func (o *GetPolicyTreeExportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get policy tree export not found response
func (o *GetPolicyTreeExportNotFound) Code() int {
	return 404
}

func (o *GetPolicyTreeExportNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/provisioning/policies/export][%d] getPolicyTreeExportNotFound  %+v", 404, o.Payload)
}

func (o *GetPolicyTreeExportNotFound) String() string {
	return fmt.Sprintf("[GET /v1/provisioning/policies/export][%d] getPolicyTreeExportNotFound  %+v", 404, o.Payload)
}

func (o *GetPolicyTreeExportNotFound) GetPayload() models.NotFound {
	return o.Payload
}

func (o *GetPolicyTreeExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
