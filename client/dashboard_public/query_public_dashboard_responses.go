// Code generated by go-swagger; DO NOT EDIT.

package dashboard_public

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

// QueryPublicDashboardReader is a Reader for the QueryPublicDashboard structure.
type QueryPublicDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryPublicDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryPublicDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryPublicDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryPublicDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryPublicDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryPublicDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryPublicDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /public/dashboards/{accessToken}/panels/{panelId}/query] queryPublicDashboard", response, response.Code())
	}
}

// NewQueryPublicDashboardOK creates a QueryPublicDashboardOK with default headers values
func NewQueryPublicDashboardOK() *QueryPublicDashboardOK {
	return &QueryPublicDashboardOK{}
}

/*
QueryPublicDashboardOK describes a response with status code 200, with default header values.

(empty)
*/
type QueryPublicDashboardOK struct {
	Payload *models.QueryDataResponse
}

// IsSuccess returns true when this query public dashboard Ok response has a 2xx status code
func (o *QueryPublicDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query public dashboard Ok response has a 3xx status code
func (o *QueryPublicDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query public dashboard Ok response has a 4xx status code
func (o *QueryPublicDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query public dashboard Ok response has a 5xx status code
func (o *QueryPublicDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query public dashboard Ok response a status code equal to that given
func (o *QueryPublicDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query public dashboard Ok response
func (o *QueryPublicDashboardOK) Code() int {
	return 200
}

func (o *QueryPublicDashboardOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardOk %s", 200, payload)
}

func (o *QueryPublicDashboardOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardOk %s", 200, payload)
}

func (o *QueryPublicDashboardOK) GetPayload() *models.QueryDataResponse {
	return o.Payload
}

func (o *QueryPublicDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryDataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPublicDashboardBadRequest creates a QueryPublicDashboardBadRequest with default headers values
func NewQueryPublicDashboardBadRequest() *QueryPublicDashboardBadRequest {
	return &QueryPublicDashboardBadRequest{}
}

/*
QueryPublicDashboardBadRequest describes a response with status code 400, with default header values.

BadRequestPublicError is returned when the request is invalid and it cannot be processed.
*/
type QueryPublicDashboardBadRequest struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this query public dashboard bad request response has a 2xx status code
func (o *QueryPublicDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query public dashboard bad request response has a 3xx status code
func (o *QueryPublicDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query public dashboard bad request response has a 4xx status code
func (o *QueryPublicDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query public dashboard bad request response has a 5xx status code
func (o *QueryPublicDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query public dashboard bad request response a status code equal to that given
func (o *QueryPublicDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query public dashboard bad request response
func (o *QueryPublicDashboardBadRequest) Code() int {
	return 400
}

func (o *QueryPublicDashboardBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardBadRequest %s", 400, payload)
}

func (o *QueryPublicDashboardBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardBadRequest %s", 400, payload)
}

func (o *QueryPublicDashboardBadRequest) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *QueryPublicDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPublicDashboardUnauthorized creates a QueryPublicDashboardUnauthorized with default headers values
func NewQueryPublicDashboardUnauthorized() *QueryPublicDashboardUnauthorized {
	return &QueryPublicDashboardUnauthorized{}
}

/*
QueryPublicDashboardUnauthorized describes a response with status code 401, with default header values.

UnauthorisedPublicError is returned when the request is not authenticated.
*/
type QueryPublicDashboardUnauthorized struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this query public dashboard unauthorized response has a 2xx status code
func (o *QueryPublicDashboardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query public dashboard unauthorized response has a 3xx status code
func (o *QueryPublicDashboardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query public dashboard unauthorized response has a 4xx status code
func (o *QueryPublicDashboardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this query public dashboard unauthorized response has a 5xx status code
func (o *QueryPublicDashboardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this query public dashboard unauthorized response a status code equal to that given
func (o *QueryPublicDashboardUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the query public dashboard unauthorized response
func (o *QueryPublicDashboardUnauthorized) Code() int {
	return 401
}

func (o *QueryPublicDashboardUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardUnauthorized %s", 401, payload)
}

func (o *QueryPublicDashboardUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardUnauthorized %s", 401, payload)
}

func (o *QueryPublicDashboardUnauthorized) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *QueryPublicDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPublicDashboardForbidden creates a QueryPublicDashboardForbidden with default headers values
func NewQueryPublicDashboardForbidden() *QueryPublicDashboardForbidden {
	return &QueryPublicDashboardForbidden{}
}

/*
QueryPublicDashboardForbidden describes a response with status code 403, with default header values.

ForbiddenPublicError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type QueryPublicDashboardForbidden struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this query public dashboard forbidden response has a 2xx status code
func (o *QueryPublicDashboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query public dashboard forbidden response has a 3xx status code
func (o *QueryPublicDashboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query public dashboard forbidden response has a 4xx status code
func (o *QueryPublicDashboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query public dashboard forbidden response has a 5xx status code
func (o *QueryPublicDashboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query public dashboard forbidden response a status code equal to that given
func (o *QueryPublicDashboardForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query public dashboard forbidden response
func (o *QueryPublicDashboardForbidden) Code() int {
	return 403
}

func (o *QueryPublicDashboardForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardForbidden %s", 403, payload)
}

func (o *QueryPublicDashboardForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardForbidden %s", 403, payload)
}

func (o *QueryPublicDashboardForbidden) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *QueryPublicDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPublicDashboardNotFound creates a QueryPublicDashboardNotFound with default headers values
func NewQueryPublicDashboardNotFound() *QueryPublicDashboardNotFound {
	return &QueryPublicDashboardNotFound{}
}

/*
QueryPublicDashboardNotFound describes a response with status code 404, with default header values.

NotFoundPublicError is returned when the requested resource was not found.
*/
type QueryPublicDashboardNotFound struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this query public dashboard not found response has a 2xx status code
func (o *QueryPublicDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query public dashboard not found response has a 3xx status code
func (o *QueryPublicDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query public dashboard not found response has a 4xx status code
func (o *QueryPublicDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query public dashboard not found response has a 5xx status code
func (o *QueryPublicDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query public dashboard not found response a status code equal to that given
func (o *QueryPublicDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query public dashboard not found response
func (o *QueryPublicDashboardNotFound) Code() int {
	return 404
}

func (o *QueryPublicDashboardNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardNotFound %s", 404, payload)
}

func (o *QueryPublicDashboardNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardNotFound %s", 404, payload)
}

func (o *QueryPublicDashboardNotFound) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *QueryPublicDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPublicDashboardInternalServerError creates a QueryPublicDashboardInternalServerError with default headers values
func NewQueryPublicDashboardInternalServerError() *QueryPublicDashboardInternalServerError {
	return &QueryPublicDashboardInternalServerError{}
}

/*
QueryPublicDashboardInternalServerError describes a response with status code 500, with default header values.

InternalServerPublicError is a general error indicating something went wrong internally.
*/
type QueryPublicDashboardInternalServerError struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this query public dashboard internal server error response has a 2xx status code
func (o *QueryPublicDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query public dashboard internal server error response has a 3xx status code
func (o *QueryPublicDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query public dashboard internal server error response has a 4xx status code
func (o *QueryPublicDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query public dashboard internal server error response has a 5xx status code
func (o *QueryPublicDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query public dashboard internal server error response a status code equal to that given
func (o *QueryPublicDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query public dashboard internal server error response
func (o *QueryPublicDashboardInternalServerError) Code() int {
	return 500
}

func (o *QueryPublicDashboardInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardInternalServerError %s", 500, payload)
}

func (o *QueryPublicDashboardInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/dashboards/{accessToken}/panels/{panelId}/query][%d] queryPublicDashboardInternalServerError %s", 500, payload)
}

func (o *QueryPublicDashboardInternalServerError) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *QueryPublicDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
