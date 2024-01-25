// Code generated by go-swagger; DO NOT EDIT.

package dashboard_permissions

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

// UpdateDashboardPermissionsByIDReader is a Reader for the UpdateDashboardPermissionsByID structure.
type UpdateDashboardPermissionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardPermissionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDashboardPermissionsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDashboardPermissionsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateDashboardPermissionsByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDashboardPermissionsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDashboardPermissionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDashboardPermissionsByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /dashboards/id/{DashboardID}/permissions] updateDashboardPermissionsByID", response, response.Code())
	}
}

// NewUpdateDashboardPermissionsByIDOK creates a UpdateDashboardPermissionsByIDOK with default headers values
func NewUpdateDashboardPermissionsByIDOK() *UpdateDashboardPermissionsByIDOK {
	return &UpdateDashboardPermissionsByIDOK{}
}

/*
UpdateDashboardPermissionsByIDOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateDashboardPermissionsByIDOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this update dashboard permissions by Id Ok response has a 2xx status code
func (o *UpdateDashboardPermissionsByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update dashboard permissions by Id Ok response has a 3xx status code
func (o *UpdateDashboardPermissionsByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard permissions by Id Ok response has a 4xx status code
func (o *UpdateDashboardPermissionsByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update dashboard permissions by Id Ok response has a 5xx status code
func (o *UpdateDashboardPermissionsByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard permissions by Id Ok response a status code equal to that given
func (o *UpdateDashboardPermissionsByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update dashboard permissions by Id Ok response
func (o *UpdateDashboardPermissionsByIDOK) Code() int {
	return 200
}

func (o *UpdateDashboardPermissionsByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdOk %s", 200, payload)
}

func (o *UpdateDashboardPermissionsByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdOk %s", 200, payload)
}

func (o *UpdateDashboardPermissionsByIDOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateDashboardPermissionsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardPermissionsByIDBadRequest creates a UpdateDashboardPermissionsByIDBadRequest with default headers values
func NewUpdateDashboardPermissionsByIDBadRequest() *UpdateDashboardPermissionsByIDBadRequest {
	return &UpdateDashboardPermissionsByIDBadRequest{}
}

/*
UpdateDashboardPermissionsByIDBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateDashboardPermissionsByIDBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update dashboard permissions by Id bad request response has a 2xx status code
func (o *UpdateDashboardPermissionsByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard permissions by Id bad request response has a 3xx status code
func (o *UpdateDashboardPermissionsByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard permissions by Id bad request response has a 4xx status code
func (o *UpdateDashboardPermissionsByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard permissions by Id bad request response has a 5xx status code
func (o *UpdateDashboardPermissionsByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard permissions by Id bad request response a status code equal to that given
func (o *UpdateDashboardPermissionsByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update dashboard permissions by Id bad request response
func (o *UpdateDashboardPermissionsByIDBadRequest) Code() int {
	return 400
}

func (o *UpdateDashboardPermissionsByIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdBadRequest %s", 400, payload)
}

func (o *UpdateDashboardPermissionsByIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdBadRequest %s", 400, payload)
}

func (o *UpdateDashboardPermissionsByIDBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDashboardPermissionsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardPermissionsByIDUnauthorized creates a UpdateDashboardPermissionsByIDUnauthorized with default headers values
func NewUpdateDashboardPermissionsByIDUnauthorized() *UpdateDashboardPermissionsByIDUnauthorized {
	return &UpdateDashboardPermissionsByIDUnauthorized{}
}

/*
UpdateDashboardPermissionsByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateDashboardPermissionsByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update dashboard permissions by Id unauthorized response has a 2xx status code
func (o *UpdateDashboardPermissionsByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard permissions by Id unauthorized response has a 3xx status code
func (o *UpdateDashboardPermissionsByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard permissions by Id unauthorized response has a 4xx status code
func (o *UpdateDashboardPermissionsByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard permissions by Id unauthorized response has a 5xx status code
func (o *UpdateDashboardPermissionsByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard permissions by Id unauthorized response a status code equal to that given
func (o *UpdateDashboardPermissionsByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update dashboard permissions by Id unauthorized response
func (o *UpdateDashboardPermissionsByIDUnauthorized) Code() int {
	return 401
}

func (o *UpdateDashboardPermissionsByIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdUnauthorized %s", 401, payload)
}

func (o *UpdateDashboardPermissionsByIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdUnauthorized %s", 401, payload)
}

func (o *UpdateDashboardPermissionsByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDashboardPermissionsByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardPermissionsByIDForbidden creates a UpdateDashboardPermissionsByIDForbidden with default headers values
func NewUpdateDashboardPermissionsByIDForbidden() *UpdateDashboardPermissionsByIDForbidden {
	return &UpdateDashboardPermissionsByIDForbidden{}
}

/*
UpdateDashboardPermissionsByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateDashboardPermissionsByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update dashboard permissions by Id forbidden response has a 2xx status code
func (o *UpdateDashboardPermissionsByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard permissions by Id forbidden response has a 3xx status code
func (o *UpdateDashboardPermissionsByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard permissions by Id forbidden response has a 4xx status code
func (o *UpdateDashboardPermissionsByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard permissions by Id forbidden response has a 5xx status code
func (o *UpdateDashboardPermissionsByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard permissions by Id forbidden response a status code equal to that given
func (o *UpdateDashboardPermissionsByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update dashboard permissions by Id forbidden response
func (o *UpdateDashboardPermissionsByIDForbidden) Code() int {
	return 403
}

func (o *UpdateDashboardPermissionsByIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdForbidden %s", 403, payload)
}

func (o *UpdateDashboardPermissionsByIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdForbidden %s", 403, payload)
}

func (o *UpdateDashboardPermissionsByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDashboardPermissionsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardPermissionsByIDNotFound creates a UpdateDashboardPermissionsByIDNotFound with default headers values
func NewUpdateDashboardPermissionsByIDNotFound() *UpdateDashboardPermissionsByIDNotFound {
	return &UpdateDashboardPermissionsByIDNotFound{}
}

/*
UpdateDashboardPermissionsByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateDashboardPermissionsByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update dashboard permissions by Id not found response has a 2xx status code
func (o *UpdateDashboardPermissionsByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard permissions by Id not found response has a 3xx status code
func (o *UpdateDashboardPermissionsByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard permissions by Id not found response has a 4xx status code
func (o *UpdateDashboardPermissionsByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dashboard permissions by Id not found response has a 5xx status code
func (o *UpdateDashboardPermissionsByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update dashboard permissions by Id not found response a status code equal to that given
func (o *UpdateDashboardPermissionsByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update dashboard permissions by Id not found response
func (o *UpdateDashboardPermissionsByIDNotFound) Code() int {
	return 404
}

func (o *UpdateDashboardPermissionsByIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdNotFound %s", 404, payload)
}

func (o *UpdateDashboardPermissionsByIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdNotFound %s", 404, payload)
}

func (o *UpdateDashboardPermissionsByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDashboardPermissionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardPermissionsByIDInternalServerError creates a UpdateDashboardPermissionsByIDInternalServerError with default headers values
func NewUpdateDashboardPermissionsByIDInternalServerError() *UpdateDashboardPermissionsByIDInternalServerError {
	return &UpdateDashboardPermissionsByIDInternalServerError{}
}

/*
UpdateDashboardPermissionsByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateDashboardPermissionsByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update dashboard permissions by Id internal server error response has a 2xx status code
func (o *UpdateDashboardPermissionsByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dashboard permissions by Id internal server error response has a 3xx status code
func (o *UpdateDashboardPermissionsByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dashboard permissions by Id internal server error response has a 4xx status code
func (o *UpdateDashboardPermissionsByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update dashboard permissions by Id internal server error response has a 5xx status code
func (o *UpdateDashboardPermissionsByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update dashboard permissions by Id internal server error response a status code equal to that given
func (o *UpdateDashboardPermissionsByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update dashboard permissions by Id internal server error response
func (o *UpdateDashboardPermissionsByIDInternalServerError) Code() int {
	return 500
}

func (o *UpdateDashboardPermissionsByIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdInternalServerError %s", 500, payload)
}

func (o *UpdateDashboardPermissionsByIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/permissions][%d] updateDashboardPermissionsByIdInternalServerError %s", 500, payload)
}

func (o *UpdateDashboardPermissionsByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDashboardPermissionsByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
