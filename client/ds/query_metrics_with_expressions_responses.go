// Code generated by go-swagger; DO NOT EDIT.

package ds

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

// QueryMetricsWithExpressionsReader is a Reader for the QueryMetricsWithExpressions structure.
type QueryMetricsWithExpressionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryMetricsWithExpressionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryMetricsWithExpressionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewQueryMetricsWithExpressionsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryMetricsWithExpressionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryMetricsWithExpressionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryMetricsWithExpressionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryMetricsWithExpressionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /ds/query] queryMetricsWithExpressions", response, response.Code())
	}
}

// NewQueryMetricsWithExpressionsOK creates a QueryMetricsWithExpressionsOK with default headers values
func NewQueryMetricsWithExpressionsOK() *QueryMetricsWithExpressionsOK {
	return &QueryMetricsWithExpressionsOK{}
}

/*
QueryMetricsWithExpressionsOK describes a response with status code 200, with default header values.

(empty)
*/
type QueryMetricsWithExpressionsOK struct {
	Payload *models.QueryDataResponse
}

// IsSuccess returns true when this query metrics with expressions Ok response has a 2xx status code
func (o *QueryMetricsWithExpressionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query metrics with expressions Ok response has a 3xx status code
func (o *QueryMetricsWithExpressionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query metrics with expressions Ok response has a 4xx status code
func (o *QueryMetricsWithExpressionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query metrics with expressions Ok response has a 5xx status code
func (o *QueryMetricsWithExpressionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query metrics with expressions Ok response a status code equal to that given
func (o *QueryMetricsWithExpressionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query metrics with expressions Ok response
func (o *QueryMetricsWithExpressionsOK) Code() int {
	return 200
}

func (o *QueryMetricsWithExpressionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsOk %s", 200, payload)
}

func (o *QueryMetricsWithExpressionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsOk %s", 200, payload)
}

func (o *QueryMetricsWithExpressionsOK) GetPayload() *models.QueryDataResponse {
	return o.Payload
}

func (o *QueryMetricsWithExpressionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryDataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryMetricsWithExpressionsMultiStatus creates a QueryMetricsWithExpressionsMultiStatus with default headers values
func NewQueryMetricsWithExpressionsMultiStatus() *QueryMetricsWithExpressionsMultiStatus {
	return &QueryMetricsWithExpressionsMultiStatus{}
}

/*
QueryMetricsWithExpressionsMultiStatus describes a response with status code 207, with default header values.

(empty)
*/
type QueryMetricsWithExpressionsMultiStatus struct {
	Payload *models.QueryDataResponse
}

// IsSuccess returns true when this query metrics with expressions multi status response has a 2xx status code
func (o *QueryMetricsWithExpressionsMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query metrics with expressions multi status response has a 3xx status code
func (o *QueryMetricsWithExpressionsMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query metrics with expressions multi status response has a 4xx status code
func (o *QueryMetricsWithExpressionsMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this query metrics with expressions multi status response has a 5xx status code
func (o *QueryMetricsWithExpressionsMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this query metrics with expressions multi status response a status code equal to that given
func (o *QueryMetricsWithExpressionsMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the query metrics with expressions multi status response
func (o *QueryMetricsWithExpressionsMultiStatus) Code() int {
	return 207
}

func (o *QueryMetricsWithExpressionsMultiStatus) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsMultiStatus %s", 207, payload)
}

func (o *QueryMetricsWithExpressionsMultiStatus) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsMultiStatus %s", 207, payload)
}

func (o *QueryMetricsWithExpressionsMultiStatus) GetPayload() *models.QueryDataResponse {
	return o.Payload
}

func (o *QueryMetricsWithExpressionsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryDataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryMetricsWithExpressionsBadRequest creates a QueryMetricsWithExpressionsBadRequest with default headers values
func NewQueryMetricsWithExpressionsBadRequest() *QueryMetricsWithExpressionsBadRequest {
	return &QueryMetricsWithExpressionsBadRequest{}
}

/*
QueryMetricsWithExpressionsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type QueryMetricsWithExpressionsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this query metrics with expressions bad request response has a 2xx status code
func (o *QueryMetricsWithExpressionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query metrics with expressions bad request response has a 3xx status code
func (o *QueryMetricsWithExpressionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query metrics with expressions bad request response has a 4xx status code
func (o *QueryMetricsWithExpressionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query metrics with expressions bad request response has a 5xx status code
func (o *QueryMetricsWithExpressionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query metrics with expressions bad request response a status code equal to that given
func (o *QueryMetricsWithExpressionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query metrics with expressions bad request response
func (o *QueryMetricsWithExpressionsBadRequest) Code() int {
	return 400
}

func (o *QueryMetricsWithExpressionsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsBadRequest %s", 400, payload)
}

func (o *QueryMetricsWithExpressionsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsBadRequest %s", 400, payload)
}

func (o *QueryMetricsWithExpressionsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *QueryMetricsWithExpressionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryMetricsWithExpressionsUnauthorized creates a QueryMetricsWithExpressionsUnauthorized with default headers values
func NewQueryMetricsWithExpressionsUnauthorized() *QueryMetricsWithExpressionsUnauthorized {
	return &QueryMetricsWithExpressionsUnauthorized{}
}

/*
QueryMetricsWithExpressionsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type QueryMetricsWithExpressionsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this query metrics with expressions unauthorized response has a 2xx status code
func (o *QueryMetricsWithExpressionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query metrics with expressions unauthorized response has a 3xx status code
func (o *QueryMetricsWithExpressionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query metrics with expressions unauthorized response has a 4xx status code
func (o *QueryMetricsWithExpressionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this query metrics with expressions unauthorized response has a 5xx status code
func (o *QueryMetricsWithExpressionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this query metrics with expressions unauthorized response a status code equal to that given
func (o *QueryMetricsWithExpressionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the query metrics with expressions unauthorized response
func (o *QueryMetricsWithExpressionsUnauthorized) Code() int {
	return 401
}

func (o *QueryMetricsWithExpressionsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsUnauthorized %s", 401, payload)
}

func (o *QueryMetricsWithExpressionsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsUnauthorized %s", 401, payload)
}

func (o *QueryMetricsWithExpressionsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *QueryMetricsWithExpressionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryMetricsWithExpressionsForbidden creates a QueryMetricsWithExpressionsForbidden with default headers values
func NewQueryMetricsWithExpressionsForbidden() *QueryMetricsWithExpressionsForbidden {
	return &QueryMetricsWithExpressionsForbidden{}
}

/*
QueryMetricsWithExpressionsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type QueryMetricsWithExpressionsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this query metrics with expressions forbidden response has a 2xx status code
func (o *QueryMetricsWithExpressionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query metrics with expressions forbidden response has a 3xx status code
func (o *QueryMetricsWithExpressionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query metrics with expressions forbidden response has a 4xx status code
func (o *QueryMetricsWithExpressionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query metrics with expressions forbidden response has a 5xx status code
func (o *QueryMetricsWithExpressionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query metrics with expressions forbidden response a status code equal to that given
func (o *QueryMetricsWithExpressionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query metrics with expressions forbidden response
func (o *QueryMetricsWithExpressionsForbidden) Code() int {
	return 403
}

func (o *QueryMetricsWithExpressionsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsForbidden %s", 403, payload)
}

func (o *QueryMetricsWithExpressionsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsForbidden %s", 403, payload)
}

func (o *QueryMetricsWithExpressionsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *QueryMetricsWithExpressionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryMetricsWithExpressionsInternalServerError creates a QueryMetricsWithExpressionsInternalServerError with default headers values
func NewQueryMetricsWithExpressionsInternalServerError() *QueryMetricsWithExpressionsInternalServerError {
	return &QueryMetricsWithExpressionsInternalServerError{}
}

/*
QueryMetricsWithExpressionsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type QueryMetricsWithExpressionsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this query metrics with expressions internal server error response has a 2xx status code
func (o *QueryMetricsWithExpressionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query metrics with expressions internal server error response has a 3xx status code
func (o *QueryMetricsWithExpressionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query metrics with expressions internal server error response has a 4xx status code
func (o *QueryMetricsWithExpressionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query metrics with expressions internal server error response has a 5xx status code
func (o *QueryMetricsWithExpressionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query metrics with expressions internal server error response a status code equal to that given
func (o *QueryMetricsWithExpressionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query metrics with expressions internal server error response
func (o *QueryMetricsWithExpressionsInternalServerError) Code() int {
	return 500
}

func (o *QueryMetricsWithExpressionsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsInternalServerError %s", 500, payload)
}

func (o *QueryMetricsWithExpressionsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ds/query][%d] queryMetricsWithExpressionsInternalServerError %s", 500, payload)
}

func (o *QueryMetricsWithExpressionsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *QueryMetricsWithExpressionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
