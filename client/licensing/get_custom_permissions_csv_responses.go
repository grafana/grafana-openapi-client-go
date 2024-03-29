// Code generated by go-swagger; DO NOT EDIT.

package licensing

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

// GetCustomPermissionsCSVReader is a Reader for the GetCustomPermissionsCSV structure.
type GetCustomPermissionsCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomPermissionsCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 500:
		result := NewGetCustomPermissionsCSVInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /licensing/custom-permissions-csv] getCustomPermissionsCSV", response, response.Code())
	}
}

// NewGetCustomPermissionsCSVInternalServerError creates a GetCustomPermissionsCSVInternalServerError with default headers values
func NewGetCustomPermissionsCSVInternalServerError() *GetCustomPermissionsCSVInternalServerError {
	return &GetCustomPermissionsCSVInternalServerError{}
}

/*
GetCustomPermissionsCSVInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCustomPermissionsCSVInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get custom permissions Csv internal server error response has a 2xx status code
func (o *GetCustomPermissionsCSVInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom permissions Csv internal server error response has a 3xx status code
func (o *GetCustomPermissionsCSVInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom permissions Csv internal server error response has a 4xx status code
func (o *GetCustomPermissionsCSVInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get custom permissions Csv internal server error response has a 5xx status code
func (o *GetCustomPermissionsCSVInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get custom permissions Csv internal server error response a status code equal to that given
func (o *GetCustomPermissionsCSVInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get custom permissions Csv internal server error response
func (o *GetCustomPermissionsCSVInternalServerError) Code() int {
	return 500
}

func (o *GetCustomPermissionsCSVInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /licensing/custom-permissions-csv][%d] getCustomPermissionsCsvInternalServerError %s", 500, payload)
}

func (o *GetCustomPermissionsCSVInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /licensing/custom-permissions-csv][%d] getCustomPermissionsCsvInternalServerError %s", 500, payload)
}

func (o *GetCustomPermissionsCSVInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCustomPermissionsCSVInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
