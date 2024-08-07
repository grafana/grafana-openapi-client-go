// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// RenderReportCSVsReader is a Reader for the RenderReportCSVs structure.
type RenderReportCSVsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenderReportCSVsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenderReportCSVsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRenderReportCSVsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenderReportCSVsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRenderReportCSVsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRenderReportCSVsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /reports/render/csvs] renderReportCSVs", response, response.Code())
	}
}

// NewRenderReportCSVsOK creates a RenderReportCSVsOK with default headers values
func NewRenderReportCSVsOK() *RenderReportCSVsOK {
	return &RenderReportCSVsOK{}
}

/*
RenderReportCSVsOK describes a response with status code 200, with default header values.

(empty)
*/
type RenderReportCSVsOK struct {
	Payload []uint8
}

// IsSuccess returns true when this render report c s vs Ok response has a 2xx status code
func (o *RenderReportCSVsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this render report c s vs Ok response has a 3xx status code
func (o *RenderReportCSVsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report c s vs Ok response has a 4xx status code
func (o *RenderReportCSVsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this render report c s vs Ok response has a 5xx status code
func (o *RenderReportCSVsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this render report c s vs Ok response a status code equal to that given
func (o *RenderReportCSVsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the render report c s vs Ok response
func (o *RenderReportCSVsOK) Code() int {
	return 200
}

func (o *RenderReportCSVsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsOk %s", 200, payload)
}

func (o *RenderReportCSVsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsOk %s", 200, payload)
}

func (o *RenderReportCSVsOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *RenderReportCSVsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportCSVsNoContent creates a RenderReportCSVsNoContent with default headers values
func NewRenderReportCSVsNoContent() *RenderReportCSVsNoContent {
	return &RenderReportCSVsNoContent{}
}

/*
RenderReportCSVsNoContent describes a response with status code 204, with default header values.

(empty)
*/
type RenderReportCSVsNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this render report c s vs no content response has a 2xx status code
func (o *RenderReportCSVsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this render report c s vs no content response has a 3xx status code
func (o *RenderReportCSVsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report c s vs no content response has a 4xx status code
func (o *RenderReportCSVsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this render report c s vs no content response has a 5xx status code
func (o *RenderReportCSVsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this render report c s vs no content response a status code equal to that given
func (o *RenderReportCSVsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the render report c s vs no content response
func (o *RenderReportCSVsNoContent) Code() int {
	return 204
}

func (o *RenderReportCSVsNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsNoContent %s", 204, payload)
}

func (o *RenderReportCSVsNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsNoContent %s", 204, payload)
}

func (o *RenderReportCSVsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RenderReportCSVsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportCSVsBadRequest creates a RenderReportCSVsBadRequest with default headers values
func NewRenderReportCSVsBadRequest() *RenderReportCSVsBadRequest {
	return &RenderReportCSVsBadRequest{}
}

/*
RenderReportCSVsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RenderReportCSVsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this render report c s vs bad request response has a 2xx status code
func (o *RenderReportCSVsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render report c s vs bad request response has a 3xx status code
func (o *RenderReportCSVsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report c s vs bad request response has a 4xx status code
func (o *RenderReportCSVsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this render report c s vs bad request response has a 5xx status code
func (o *RenderReportCSVsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this render report c s vs bad request response a status code equal to that given
func (o *RenderReportCSVsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the render report c s vs bad request response
func (o *RenderReportCSVsBadRequest) Code() int {
	return 400
}

func (o *RenderReportCSVsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsBadRequest %s", 400, payload)
}

func (o *RenderReportCSVsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsBadRequest %s", 400, payload)
}

func (o *RenderReportCSVsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportCSVsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportCSVsUnauthorized creates a RenderReportCSVsUnauthorized with default headers values
func NewRenderReportCSVsUnauthorized() *RenderReportCSVsUnauthorized {
	return &RenderReportCSVsUnauthorized{}
}

/*
RenderReportCSVsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RenderReportCSVsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this render report c s vs unauthorized response has a 2xx status code
func (o *RenderReportCSVsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render report c s vs unauthorized response has a 3xx status code
func (o *RenderReportCSVsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report c s vs unauthorized response has a 4xx status code
func (o *RenderReportCSVsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this render report c s vs unauthorized response has a 5xx status code
func (o *RenderReportCSVsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this render report c s vs unauthorized response a status code equal to that given
func (o *RenderReportCSVsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the render report c s vs unauthorized response
func (o *RenderReportCSVsUnauthorized) Code() int {
	return 401
}

func (o *RenderReportCSVsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsUnauthorized %s", 401, payload)
}

func (o *RenderReportCSVsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsUnauthorized %s", 401, payload)
}

func (o *RenderReportCSVsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportCSVsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportCSVsInternalServerError creates a RenderReportCSVsInternalServerError with default headers values
func NewRenderReportCSVsInternalServerError() *RenderReportCSVsInternalServerError {
	return &RenderReportCSVsInternalServerError{}
}

/*
RenderReportCSVsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RenderReportCSVsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this render report c s vs internal server error response has a 2xx status code
func (o *RenderReportCSVsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render report c s vs internal server error response has a 3xx status code
func (o *RenderReportCSVsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render report c s vs internal server error response has a 4xx status code
func (o *RenderReportCSVsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this render report c s vs internal server error response has a 5xx status code
func (o *RenderReportCSVsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this render report c s vs internal server error response a status code equal to that given
func (o *RenderReportCSVsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the render report c s vs internal server error response
func (o *RenderReportCSVsInternalServerError) Code() int {
	return 500
}

func (o *RenderReportCSVsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsInternalServerError %s", 500, payload)
}

func (o *RenderReportCSVsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports/render/csvs][%d] renderReportCSVsInternalServerError %s", 500, payload)
}

func (o *RenderReportCSVsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportCSVsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
