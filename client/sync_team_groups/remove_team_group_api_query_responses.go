// Code generated by go-swagger; DO NOT EDIT.

package sync_team_groups

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

// RemoveTeamGroupAPIQueryReader is a Reader for the RemoveTeamGroupAPIQuery structure.
type RemoveTeamGroupAPIQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveTeamGroupAPIQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveTeamGroupAPIQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveTeamGroupAPIQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveTeamGroupAPIQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveTeamGroupAPIQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveTeamGroupAPIQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveTeamGroupAPIQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /teams/{teamId}/groups] removeTeamGroupApiQuery", response, response.Code())
	}
}

// NewRemoveTeamGroupAPIQueryOK creates a RemoveTeamGroupAPIQueryOK with default headers values
func NewRemoveTeamGroupAPIQueryOK() *RemoveTeamGroupAPIQueryOK {
	return &RemoveTeamGroupAPIQueryOK{}
}

/*
RemoveTeamGroupAPIQueryOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type RemoveTeamGroupAPIQueryOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this remove team group Api query Ok response has a 2xx status code
func (o *RemoveTeamGroupAPIQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove team group Api query Ok response has a 3xx status code
func (o *RemoveTeamGroupAPIQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team group Api query Ok response has a 4xx status code
func (o *RemoveTeamGroupAPIQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove team group Api query Ok response has a 5xx status code
func (o *RemoveTeamGroupAPIQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team group Api query Ok response a status code equal to that given
func (o *RemoveTeamGroupAPIQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove team group Api query Ok response
func (o *RemoveTeamGroupAPIQueryOK) Code() int {
	return 200
}

func (o *RemoveTeamGroupAPIQueryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryOk %s", 200, payload)
}

func (o *RemoveTeamGroupAPIQueryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryOk %s", 200, payload)
}

func (o *RemoveTeamGroupAPIQueryOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *RemoveTeamGroupAPIQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamGroupAPIQueryBadRequest creates a RemoveTeamGroupAPIQueryBadRequest with default headers values
func NewRemoveTeamGroupAPIQueryBadRequest() *RemoveTeamGroupAPIQueryBadRequest {
	return &RemoveTeamGroupAPIQueryBadRequest{}
}

/*
RemoveTeamGroupAPIQueryBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RemoveTeamGroupAPIQueryBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team group Api query bad request response has a 2xx status code
func (o *RemoveTeamGroupAPIQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team group Api query bad request response has a 3xx status code
func (o *RemoveTeamGroupAPIQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team group Api query bad request response has a 4xx status code
func (o *RemoveTeamGroupAPIQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove team group Api query bad request response has a 5xx status code
func (o *RemoveTeamGroupAPIQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team group Api query bad request response a status code equal to that given
func (o *RemoveTeamGroupAPIQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove team group Api query bad request response
func (o *RemoveTeamGroupAPIQueryBadRequest) Code() int {
	return 400
}

func (o *RemoveTeamGroupAPIQueryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryBadRequest %s", 400, payload)
}

func (o *RemoveTeamGroupAPIQueryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryBadRequest %s", 400, payload)
}

func (o *RemoveTeamGroupAPIQueryBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamGroupAPIQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamGroupAPIQueryUnauthorized creates a RemoveTeamGroupAPIQueryUnauthorized with default headers values
func NewRemoveTeamGroupAPIQueryUnauthorized() *RemoveTeamGroupAPIQueryUnauthorized {
	return &RemoveTeamGroupAPIQueryUnauthorized{}
}

/*
RemoveTeamGroupAPIQueryUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RemoveTeamGroupAPIQueryUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team group Api query unauthorized response has a 2xx status code
func (o *RemoveTeamGroupAPIQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team group Api query unauthorized response has a 3xx status code
func (o *RemoveTeamGroupAPIQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team group Api query unauthorized response has a 4xx status code
func (o *RemoveTeamGroupAPIQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove team group Api query unauthorized response has a 5xx status code
func (o *RemoveTeamGroupAPIQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team group Api query unauthorized response a status code equal to that given
func (o *RemoveTeamGroupAPIQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove team group Api query unauthorized response
func (o *RemoveTeamGroupAPIQueryUnauthorized) Code() int {
	return 401
}

func (o *RemoveTeamGroupAPIQueryUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryUnauthorized %s", 401, payload)
}

func (o *RemoveTeamGroupAPIQueryUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryUnauthorized %s", 401, payload)
}

func (o *RemoveTeamGroupAPIQueryUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamGroupAPIQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamGroupAPIQueryForbidden creates a RemoveTeamGroupAPIQueryForbidden with default headers values
func NewRemoveTeamGroupAPIQueryForbidden() *RemoveTeamGroupAPIQueryForbidden {
	return &RemoveTeamGroupAPIQueryForbidden{}
}

/*
RemoveTeamGroupAPIQueryForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RemoveTeamGroupAPIQueryForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team group Api query forbidden response has a 2xx status code
func (o *RemoveTeamGroupAPIQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team group Api query forbidden response has a 3xx status code
func (o *RemoveTeamGroupAPIQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team group Api query forbidden response has a 4xx status code
func (o *RemoveTeamGroupAPIQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove team group Api query forbidden response has a 5xx status code
func (o *RemoveTeamGroupAPIQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team group Api query forbidden response a status code equal to that given
func (o *RemoveTeamGroupAPIQueryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove team group Api query forbidden response
func (o *RemoveTeamGroupAPIQueryForbidden) Code() int {
	return 403
}

func (o *RemoveTeamGroupAPIQueryForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryForbidden %s", 403, payload)
}

func (o *RemoveTeamGroupAPIQueryForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryForbidden %s", 403, payload)
}

func (o *RemoveTeamGroupAPIQueryForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamGroupAPIQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamGroupAPIQueryNotFound creates a RemoveTeamGroupAPIQueryNotFound with default headers values
func NewRemoveTeamGroupAPIQueryNotFound() *RemoveTeamGroupAPIQueryNotFound {
	return &RemoveTeamGroupAPIQueryNotFound{}
}

/*
RemoveTeamGroupAPIQueryNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RemoveTeamGroupAPIQueryNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team group Api query not found response has a 2xx status code
func (o *RemoveTeamGroupAPIQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team group Api query not found response has a 3xx status code
func (o *RemoveTeamGroupAPIQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team group Api query not found response has a 4xx status code
func (o *RemoveTeamGroupAPIQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove team group Api query not found response has a 5xx status code
func (o *RemoveTeamGroupAPIQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove team group Api query not found response a status code equal to that given
func (o *RemoveTeamGroupAPIQueryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove team group Api query not found response
func (o *RemoveTeamGroupAPIQueryNotFound) Code() int {
	return 404
}

func (o *RemoveTeamGroupAPIQueryNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryNotFound %s", 404, payload)
}

func (o *RemoveTeamGroupAPIQueryNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryNotFound %s", 404, payload)
}

func (o *RemoveTeamGroupAPIQueryNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamGroupAPIQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamGroupAPIQueryInternalServerError creates a RemoveTeamGroupAPIQueryInternalServerError with default headers values
func NewRemoveTeamGroupAPIQueryInternalServerError() *RemoveTeamGroupAPIQueryInternalServerError {
	return &RemoveTeamGroupAPIQueryInternalServerError{}
}

/*
RemoveTeamGroupAPIQueryInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RemoveTeamGroupAPIQueryInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove team group Api query internal server error response has a 2xx status code
func (o *RemoveTeamGroupAPIQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove team group Api query internal server error response has a 3xx status code
func (o *RemoveTeamGroupAPIQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove team group Api query internal server error response has a 4xx status code
func (o *RemoveTeamGroupAPIQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove team group Api query internal server error response has a 5xx status code
func (o *RemoveTeamGroupAPIQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove team group Api query internal server error response a status code equal to that given
func (o *RemoveTeamGroupAPIQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove team group Api query internal server error response
func (o *RemoveTeamGroupAPIQueryInternalServerError) Code() int {
	return 500
}

func (o *RemoveTeamGroupAPIQueryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryInternalServerError %s", 500, payload)
}

func (o *RemoveTeamGroupAPIQueryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /teams/{teamId}/groups][%d] removeTeamGroupApiQueryInternalServerError %s", 500, payload)
}

func (o *RemoveTeamGroupAPIQueryInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamGroupAPIQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
