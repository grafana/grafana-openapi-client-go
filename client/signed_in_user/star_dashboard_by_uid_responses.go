// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

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

// StarDashboardByUIDReader is a Reader for the StarDashboardByUID structure.
type StarDashboardByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StarDashboardByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStarDashboardByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStarDashboardByUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStarDashboardByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStarDashboardByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStarDashboardByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /user/stars/dashboard/uid/{dashboard_uid}] starDashboardByUID", response, response.Code())
	}
}

// NewStarDashboardByUIDOK creates a StarDashboardByUIDOK with default headers values
func NewStarDashboardByUIDOK() *StarDashboardByUIDOK {
	return &StarDashboardByUIDOK{}
}

/*
StarDashboardByUIDOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type StarDashboardByUIDOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this star dashboard by Uid Ok response has a 2xx status code
func (o *StarDashboardByUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this star dashboard by Uid Ok response has a 3xx status code
func (o *StarDashboardByUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this star dashboard by Uid Ok response has a 4xx status code
func (o *StarDashboardByUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this star dashboard by Uid Ok response has a 5xx status code
func (o *StarDashboardByUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this star dashboard by Uid Ok response a status code equal to that given
func (o *StarDashboardByUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the star dashboard by Uid Ok response
func (o *StarDashboardByUIDOK) Code() int {
	return 200
}

func (o *StarDashboardByUIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidOk %s", 200, payload)
}

func (o *StarDashboardByUIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidOk %s", 200, payload)
}

func (o *StarDashboardByUIDOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarDashboardByUIDBadRequest creates a StarDashboardByUIDBadRequest with default headers values
func NewStarDashboardByUIDBadRequest() *StarDashboardByUIDBadRequest {
	return &StarDashboardByUIDBadRequest{}
}

/*
StarDashboardByUIDBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type StarDashboardByUIDBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this star dashboard by Uid bad request response has a 2xx status code
func (o *StarDashboardByUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this star dashboard by Uid bad request response has a 3xx status code
func (o *StarDashboardByUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this star dashboard by Uid bad request response has a 4xx status code
func (o *StarDashboardByUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this star dashboard by Uid bad request response has a 5xx status code
func (o *StarDashboardByUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this star dashboard by Uid bad request response a status code equal to that given
func (o *StarDashboardByUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the star dashboard by Uid bad request response
func (o *StarDashboardByUIDBadRequest) Code() int {
	return 400
}

func (o *StarDashboardByUIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidBadRequest %s", 400, payload)
}

func (o *StarDashboardByUIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidBadRequest %s", 400, payload)
}

func (o *StarDashboardByUIDBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarDashboardByUIDUnauthorized creates a StarDashboardByUIDUnauthorized with default headers values
func NewStarDashboardByUIDUnauthorized() *StarDashboardByUIDUnauthorized {
	return &StarDashboardByUIDUnauthorized{}
}

/*
StarDashboardByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type StarDashboardByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this star dashboard by Uid unauthorized response has a 2xx status code
func (o *StarDashboardByUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this star dashboard by Uid unauthorized response has a 3xx status code
func (o *StarDashboardByUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this star dashboard by Uid unauthorized response has a 4xx status code
func (o *StarDashboardByUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this star dashboard by Uid unauthorized response has a 5xx status code
func (o *StarDashboardByUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this star dashboard by Uid unauthorized response a status code equal to that given
func (o *StarDashboardByUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the star dashboard by Uid unauthorized response
func (o *StarDashboardByUIDUnauthorized) Code() int {
	return 401
}

func (o *StarDashboardByUIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidUnauthorized %s", 401, payload)
}

func (o *StarDashboardByUIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidUnauthorized %s", 401, payload)
}

func (o *StarDashboardByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarDashboardByUIDForbidden creates a StarDashboardByUIDForbidden with default headers values
func NewStarDashboardByUIDForbidden() *StarDashboardByUIDForbidden {
	return &StarDashboardByUIDForbidden{}
}

/*
StarDashboardByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type StarDashboardByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this star dashboard by Uid forbidden response has a 2xx status code
func (o *StarDashboardByUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this star dashboard by Uid forbidden response has a 3xx status code
func (o *StarDashboardByUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this star dashboard by Uid forbidden response has a 4xx status code
func (o *StarDashboardByUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this star dashboard by Uid forbidden response has a 5xx status code
func (o *StarDashboardByUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this star dashboard by Uid forbidden response a status code equal to that given
func (o *StarDashboardByUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the star dashboard by Uid forbidden response
func (o *StarDashboardByUIDForbidden) Code() int {
	return 403
}

func (o *StarDashboardByUIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidForbidden %s", 403, payload)
}

func (o *StarDashboardByUIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidForbidden %s", 403, payload)
}

func (o *StarDashboardByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarDashboardByUIDInternalServerError creates a StarDashboardByUIDInternalServerError with default headers values
func NewStarDashboardByUIDInternalServerError() *StarDashboardByUIDInternalServerError {
	return &StarDashboardByUIDInternalServerError{}
}

/*
StarDashboardByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type StarDashboardByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this star dashboard by Uid internal server error response has a 2xx status code
func (o *StarDashboardByUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this star dashboard by Uid internal server error response has a 3xx status code
func (o *StarDashboardByUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this star dashboard by Uid internal server error response has a 4xx status code
func (o *StarDashboardByUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this star dashboard by Uid internal server error response has a 5xx status code
func (o *StarDashboardByUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this star dashboard by Uid internal server error response a status code equal to that given
func (o *StarDashboardByUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the star dashboard by Uid internal server error response
func (o *StarDashboardByUIDInternalServerError) Code() int {
	return 500
}

func (o *StarDashboardByUIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidInternalServerError %s", 500, payload)
}

func (o *StarDashboardByUIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidInternalServerError %s", 500, payload)
}

func (o *StarDashboardByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
