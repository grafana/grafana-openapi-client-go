// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// DeleteDashboardSnapshotReader is a Reader for the DeleteDashboardSnapshot structure.
type DeleteDashboardSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDashboardSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteDashboardSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDashboardSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDashboardSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /snapshots/{key}] deleteDashboardSnapshot", response, response.Code())
	}
}

// NewDeleteDashboardSnapshotOK creates a DeleteDashboardSnapshotOK with default headers values
func NewDeleteDashboardSnapshotOK() *DeleteDashboardSnapshotOK {
	return &DeleteDashboardSnapshotOK{}
}

/*
DeleteDashboardSnapshotOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteDashboardSnapshotOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete dashboard snapshot Ok response has a 2xx status code
func (o *DeleteDashboardSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete dashboard snapshot Ok response has a 3xx status code
func (o *DeleteDashboardSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard snapshot Ok response has a 4xx status code
func (o *DeleteDashboardSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete dashboard snapshot Ok response has a 5xx status code
func (o *DeleteDashboardSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard snapshot Ok response a status code equal to that given
func (o *DeleteDashboardSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete dashboard snapshot Ok response
func (o *DeleteDashboardSnapshotOK) Code() int {
	return 200
}

func (o *DeleteDashboardSnapshotOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapshots/{key}][%d] deleteDashboardSnapshotOk %s", 200, payload)
}

func (o *DeleteDashboardSnapshotOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapshots/{key}][%d] deleteDashboardSnapshotOk %s", 200, payload)
}

func (o *DeleteDashboardSnapshotOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardSnapshotForbidden creates a DeleteDashboardSnapshotForbidden with default headers values
func NewDeleteDashboardSnapshotForbidden() *DeleteDashboardSnapshotForbidden {
	return &DeleteDashboardSnapshotForbidden{}
}

/*
DeleteDashboardSnapshotForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteDashboardSnapshotForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete dashboard snapshot forbidden response has a 2xx status code
func (o *DeleteDashboardSnapshotForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard snapshot forbidden response has a 3xx status code
func (o *DeleteDashboardSnapshotForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard snapshot forbidden response has a 4xx status code
func (o *DeleteDashboardSnapshotForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete dashboard snapshot forbidden response has a 5xx status code
func (o *DeleteDashboardSnapshotForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard snapshot forbidden response a status code equal to that given
func (o *DeleteDashboardSnapshotForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete dashboard snapshot forbidden response
func (o *DeleteDashboardSnapshotForbidden) Code() int {
	return 403
}

func (o *DeleteDashboardSnapshotForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapshots/{key}][%d] deleteDashboardSnapshotForbidden %s", 403, payload)
}

func (o *DeleteDashboardSnapshotForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapshots/{key}][%d] deleteDashboardSnapshotForbidden %s", 403, payload)
}

func (o *DeleteDashboardSnapshotForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardSnapshotNotFound creates a DeleteDashboardSnapshotNotFound with default headers values
func NewDeleteDashboardSnapshotNotFound() *DeleteDashboardSnapshotNotFound {
	return &DeleteDashboardSnapshotNotFound{}
}

/*
DeleteDashboardSnapshotNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteDashboardSnapshotNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete dashboard snapshot not found response has a 2xx status code
func (o *DeleteDashboardSnapshotNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard snapshot not found response has a 3xx status code
func (o *DeleteDashboardSnapshotNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard snapshot not found response has a 4xx status code
func (o *DeleteDashboardSnapshotNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete dashboard snapshot not found response has a 5xx status code
func (o *DeleteDashboardSnapshotNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete dashboard snapshot not found response a status code equal to that given
func (o *DeleteDashboardSnapshotNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete dashboard snapshot not found response
func (o *DeleteDashboardSnapshotNotFound) Code() int {
	return 404
}

func (o *DeleteDashboardSnapshotNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapshots/{key}][%d] deleteDashboardSnapshotNotFound %s", 404, payload)
}

func (o *DeleteDashboardSnapshotNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapshots/{key}][%d] deleteDashboardSnapshotNotFound %s", 404, payload)
}

func (o *DeleteDashboardSnapshotNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardSnapshotInternalServerError creates a DeleteDashboardSnapshotInternalServerError with default headers values
func NewDeleteDashboardSnapshotInternalServerError() *DeleteDashboardSnapshotInternalServerError {
	return &DeleteDashboardSnapshotInternalServerError{}
}

/*
DeleteDashboardSnapshotInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteDashboardSnapshotInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete dashboard snapshot internal server error response has a 2xx status code
func (o *DeleteDashboardSnapshotInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete dashboard snapshot internal server error response has a 3xx status code
func (o *DeleteDashboardSnapshotInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete dashboard snapshot internal server error response has a 4xx status code
func (o *DeleteDashboardSnapshotInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete dashboard snapshot internal server error response has a 5xx status code
func (o *DeleteDashboardSnapshotInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete dashboard snapshot internal server error response a status code equal to that given
func (o *DeleteDashboardSnapshotInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete dashboard snapshot internal server error response
func (o *DeleteDashboardSnapshotInternalServerError) Code() int {
	return 500
}

func (o *DeleteDashboardSnapshotInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapshots/{key}][%d] deleteDashboardSnapshotInternalServerError %s", 500, payload)
}

func (o *DeleteDashboardSnapshotInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapshots/{key}][%d] deleteDashboardSnapshotInternalServerError %s", 500, payload)
}

func (o *DeleteDashboardSnapshotInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
