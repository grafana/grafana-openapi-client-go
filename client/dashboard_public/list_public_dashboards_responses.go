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

// ListPublicDashboardsReader is a Reader for the ListPublicDashboards structure.
type ListPublicDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPublicDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPublicDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListPublicDashboardsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPublicDashboardsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPublicDashboardsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /dashboards/public-dashboards] listPublicDashboards", response, response.Code())
	}
}

// NewListPublicDashboardsOK creates a ListPublicDashboardsOK with default headers values
func NewListPublicDashboardsOK() *ListPublicDashboardsOK {
	return &ListPublicDashboardsOK{}
}

/*
ListPublicDashboardsOK describes a response with status code 200, with default header values.

(empty)
*/
type ListPublicDashboardsOK struct {
	Payload *models.PublicDashboardListResponseWithPagination
}

// IsSuccess returns true when this list public dashboards Ok response has a 2xx status code
func (o *ListPublicDashboardsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list public dashboards Ok response has a 3xx status code
func (o *ListPublicDashboardsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list public dashboards Ok response has a 4xx status code
func (o *ListPublicDashboardsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list public dashboards Ok response has a 5xx status code
func (o *ListPublicDashboardsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list public dashboards Ok response a status code equal to that given
func (o *ListPublicDashboardsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list public dashboards Ok response
func (o *ListPublicDashboardsOK) Code() int {
	return 200
}

func (o *ListPublicDashboardsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/public-dashboards][%d] listPublicDashboardsOk %s", 200, payload)
}

func (o *ListPublicDashboardsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/public-dashboards][%d] listPublicDashboardsOk %s", 200, payload)
}

func (o *ListPublicDashboardsOK) GetPayload() *models.PublicDashboardListResponseWithPagination {
	return o.Payload
}

func (o *ListPublicDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicDashboardListResponseWithPagination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPublicDashboardsUnauthorized creates a ListPublicDashboardsUnauthorized with default headers values
func NewListPublicDashboardsUnauthorized() *ListPublicDashboardsUnauthorized {
	return &ListPublicDashboardsUnauthorized{}
}

/*
ListPublicDashboardsUnauthorized describes a response with status code 401, with default header values.

UnauthorisedPublicError is returned when the request is not authenticated.
*/
type ListPublicDashboardsUnauthorized struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this list public dashboards unauthorized response has a 2xx status code
func (o *ListPublicDashboardsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list public dashboards unauthorized response has a 3xx status code
func (o *ListPublicDashboardsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list public dashboards unauthorized response has a 4xx status code
func (o *ListPublicDashboardsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list public dashboards unauthorized response has a 5xx status code
func (o *ListPublicDashboardsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list public dashboards unauthorized response a status code equal to that given
func (o *ListPublicDashboardsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list public dashboards unauthorized response
func (o *ListPublicDashboardsUnauthorized) Code() int {
	return 401
}

func (o *ListPublicDashboardsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/public-dashboards][%d] listPublicDashboardsUnauthorized %s", 401, payload)
}

func (o *ListPublicDashboardsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/public-dashboards][%d] listPublicDashboardsUnauthorized %s", 401, payload)
}

func (o *ListPublicDashboardsUnauthorized) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *ListPublicDashboardsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPublicDashboardsForbidden creates a ListPublicDashboardsForbidden with default headers values
func NewListPublicDashboardsForbidden() *ListPublicDashboardsForbidden {
	return &ListPublicDashboardsForbidden{}
}

/*
ListPublicDashboardsForbidden describes a response with status code 403, with default header values.

ForbiddenPublicError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ListPublicDashboardsForbidden struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this list public dashboards forbidden response has a 2xx status code
func (o *ListPublicDashboardsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list public dashboards forbidden response has a 3xx status code
func (o *ListPublicDashboardsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list public dashboards forbidden response has a 4xx status code
func (o *ListPublicDashboardsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list public dashboards forbidden response has a 5xx status code
func (o *ListPublicDashboardsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list public dashboards forbidden response a status code equal to that given
func (o *ListPublicDashboardsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list public dashboards forbidden response
func (o *ListPublicDashboardsForbidden) Code() int {
	return 403
}

func (o *ListPublicDashboardsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/public-dashboards][%d] listPublicDashboardsForbidden %s", 403, payload)
}

func (o *ListPublicDashboardsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/public-dashboards][%d] listPublicDashboardsForbidden %s", 403, payload)
}

func (o *ListPublicDashboardsForbidden) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *ListPublicDashboardsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPublicDashboardsInternalServerError creates a ListPublicDashboardsInternalServerError with default headers values
func NewListPublicDashboardsInternalServerError() *ListPublicDashboardsInternalServerError {
	return &ListPublicDashboardsInternalServerError{}
}

/*
ListPublicDashboardsInternalServerError describes a response with status code 500, with default header values.

InternalServerPublicError is a general error indicating something went wrong internally.
*/
type ListPublicDashboardsInternalServerError struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this list public dashboards internal server error response has a 2xx status code
func (o *ListPublicDashboardsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list public dashboards internal server error response has a 3xx status code
func (o *ListPublicDashboardsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list public dashboards internal server error response has a 4xx status code
func (o *ListPublicDashboardsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list public dashboards internal server error response has a 5xx status code
func (o *ListPublicDashboardsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list public dashboards internal server error response a status code equal to that given
func (o *ListPublicDashboardsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list public dashboards internal server error response
func (o *ListPublicDashboardsInternalServerError) Code() int {
	return 500
}

func (o *ListPublicDashboardsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/public-dashboards][%d] listPublicDashboardsInternalServerError %s", 500, payload)
}

func (o *ListPublicDashboardsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/public-dashboards][%d] listPublicDashboardsInternalServerError %s", 500, payload)
}

func (o *ListPublicDashboardsInternalServerError) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *ListPublicDashboardsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
