// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetDashboardStatesReader is a Reader for the GetDashboardStates structure.
type GetDashboardStatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardStatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardStatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDashboardStatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDashboardStatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /alerts/states-for-dashboard] getDashboardStates", response, response.Code())
	}
}

// NewGetDashboardStatesOK creates a GetDashboardStatesOK with default headers values
func NewGetDashboardStatesOK() *GetDashboardStatesOK {
	return &GetDashboardStatesOK{}
}

/*
GetDashboardStatesOK describes a response with status code 200, with default header values.

(empty)
*/
type GetDashboardStatesOK struct {
	Payload []*models.AlertStateInfoDTO
}

// IsSuccess returns true when this get dashboard states Ok response has a 2xx status code
func (o *GetDashboardStatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dashboard states Ok response has a 3xx status code
func (o *GetDashboardStatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard states Ok response has a 4xx status code
func (o *GetDashboardStatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard states Ok response has a 5xx status code
func (o *GetDashboardStatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard states Ok response a status code equal to that given
func (o *GetDashboardStatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get dashboard states Ok response
func (o *GetDashboardStatesOK) Code() int {
	return 200
}

func (o *GetDashboardStatesOK) Error() string {
	return fmt.Sprintf("[GET /alerts/states-for-dashboard][%d] getDashboardStatesOk  %+v", 200, o.Payload)
}

func (o *GetDashboardStatesOK) String() string {
	return fmt.Sprintf("[GET /alerts/states-for-dashboard][%d] getDashboardStatesOk  %+v", 200, o.Payload)
}

func (o *GetDashboardStatesOK) GetPayload() []*models.AlertStateInfoDTO {
	return o.Payload
}

func (o *GetDashboardStatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardStatesBadRequest creates a GetDashboardStatesBadRequest with default headers values
func NewGetDashboardStatesBadRequest() *GetDashboardStatesBadRequest {
	return &GetDashboardStatesBadRequest{}
}

/*
GetDashboardStatesBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type GetDashboardStatesBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard states bad request response has a 2xx status code
func (o *GetDashboardStatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard states bad request response has a 3xx status code
func (o *GetDashboardStatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard states bad request response has a 4xx status code
func (o *GetDashboardStatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dashboard states bad request response has a 5xx status code
func (o *GetDashboardStatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard states bad request response a status code equal to that given
func (o *GetDashboardStatesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get dashboard states bad request response
func (o *GetDashboardStatesBadRequest) Code() int {
	return 400
}

func (o *GetDashboardStatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /alerts/states-for-dashboard][%d] getDashboardStatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetDashboardStatesBadRequest) String() string {
	return fmt.Sprintf("[GET /alerts/states-for-dashboard][%d] getDashboardStatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetDashboardStatesBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardStatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardStatesInternalServerError creates a GetDashboardStatesInternalServerError with default headers values
func NewGetDashboardStatesInternalServerError() *GetDashboardStatesInternalServerError {
	return &GetDashboardStatesInternalServerError{}
}

/*
GetDashboardStatesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetDashboardStatesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get dashboard states internal server error response has a 2xx status code
func (o *GetDashboardStatesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dashboard states internal server error response has a 3xx status code
func (o *GetDashboardStatesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard states internal server error response has a 4xx status code
func (o *GetDashboardStatesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard states internal server error response has a 5xx status code
func (o *GetDashboardStatesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dashboard states internal server error response a status code equal to that given
func (o *GetDashboardStatesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get dashboard states internal server error response
func (o *GetDashboardStatesInternalServerError) Code() int {
	return 500
}

func (o *GetDashboardStatesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alerts/states-for-dashboard][%d] getDashboardStatesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDashboardStatesInternalServerError) String() string {
	return fmt.Sprintf("[GET /alerts/states-for-dashboard][%d] getDashboardStatesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDashboardStatesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardStatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
