// Code generated by go-swagger; DO NOT EDIT.

package licensing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetCustomPermissionsReportReader is a Reader for the GetCustomPermissionsReport structure.
type GetCustomPermissionsReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomPermissionsReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomPermissionsReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetCustomPermissionsReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /licensing/custom-permissions] getCustomPermissionsReport", response, response.Code())
	}
}

// NewGetCustomPermissionsReportOK creates a GetCustomPermissionsReportOK with default headers values
func NewGetCustomPermissionsReportOK() *GetCustomPermissionsReportOK {
	return &GetCustomPermissionsReportOK{}
}

/*
GetCustomPermissionsReportOK describes a response with status code 200, with default header values.

(empty)
*/
type GetCustomPermissionsReportOK struct {
	Payload []*models.CustomPermissionsRecordDTO
}

// IsSuccess returns true when this get custom permissions report Ok response has a 2xx status code
func (o *GetCustomPermissionsReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get custom permissions report Ok response has a 3xx status code
func (o *GetCustomPermissionsReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom permissions report Ok response has a 4xx status code
func (o *GetCustomPermissionsReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get custom permissions report Ok response has a 5xx status code
func (o *GetCustomPermissionsReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom permissions report Ok response a status code equal to that given
func (o *GetCustomPermissionsReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get custom permissions report Ok response
func (o *GetCustomPermissionsReportOK) Code() int {
	return 200
}

func (o *GetCustomPermissionsReportOK) Error() string {
	return fmt.Sprintf("[GET /licensing/custom-permissions][%d] getCustomPermissionsReportOk  %+v", 200, o.Payload)
}

func (o *GetCustomPermissionsReportOK) String() string {
	return fmt.Sprintf("[GET /licensing/custom-permissions][%d] getCustomPermissionsReportOk  %+v", 200, o.Payload)
}

func (o *GetCustomPermissionsReportOK) GetPayload() []*models.CustomPermissionsRecordDTO {
	return o.Payload
}

func (o *GetCustomPermissionsReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomPermissionsReportInternalServerError creates a GetCustomPermissionsReportInternalServerError with default headers values
func NewGetCustomPermissionsReportInternalServerError() *GetCustomPermissionsReportInternalServerError {
	return &GetCustomPermissionsReportInternalServerError{}
}

/*
GetCustomPermissionsReportInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCustomPermissionsReportInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get custom permissions report internal server error response has a 2xx status code
func (o *GetCustomPermissionsReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom permissions report internal server error response has a 3xx status code
func (o *GetCustomPermissionsReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom permissions report internal server error response has a 4xx status code
func (o *GetCustomPermissionsReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get custom permissions report internal server error response has a 5xx status code
func (o *GetCustomPermissionsReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get custom permissions report internal server error response a status code equal to that given
func (o *GetCustomPermissionsReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get custom permissions report internal server error response
func (o *GetCustomPermissionsReportInternalServerError) Code() int {
	return 500
}

func (o *GetCustomPermissionsReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /licensing/custom-permissions][%d] getCustomPermissionsReportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCustomPermissionsReportInternalServerError) String() string {
	return fmt.Sprintf("[GET /licensing/custom-permissions][%d] getCustomPermissionsReportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCustomPermissionsReportInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCustomPermissionsReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
