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

// RefreshLicenseStatsReader is a Reader for the RefreshLicenseStats structure.
type RefreshLicenseStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshLicenseStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshLicenseStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRefreshLicenseStatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /licensing/refresh-stats] refreshLicenseStats", response, response.Code())
	}
}

// NewRefreshLicenseStatsOK creates a RefreshLicenseStatsOK with default headers values
func NewRefreshLicenseStatsOK() *RefreshLicenseStatsOK {
	return &RefreshLicenseStatsOK{}
}

/*
RefreshLicenseStatsOK describes a response with status code 200, with default header values.

(empty)
*/
type RefreshLicenseStatsOK struct {
	Payload *models.ActiveUserStats
}

// IsSuccess returns true when this refresh license stats Ok response has a 2xx status code
func (o *RefreshLicenseStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this refresh license stats Ok response has a 3xx status code
func (o *RefreshLicenseStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh license stats Ok response has a 4xx status code
func (o *RefreshLicenseStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh license stats Ok response has a 5xx status code
func (o *RefreshLicenseStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh license stats Ok response a status code equal to that given
func (o *RefreshLicenseStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the refresh license stats Ok response
func (o *RefreshLicenseStatsOK) Code() int {
	return 200
}

func (o *RefreshLicenseStatsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /licensing/refresh-stats][%d] refreshLicenseStatsOk %s", 200, payload)
}

func (o *RefreshLicenseStatsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /licensing/refresh-stats][%d] refreshLicenseStatsOk %s", 200, payload)
}

func (o *RefreshLicenseStatsOK) GetPayload() *models.ActiveUserStats {
	return o.Payload
}

func (o *RefreshLicenseStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveUserStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshLicenseStatsInternalServerError creates a RefreshLicenseStatsInternalServerError with default headers values
func NewRefreshLicenseStatsInternalServerError() *RefreshLicenseStatsInternalServerError {
	return &RefreshLicenseStatsInternalServerError{}
}

/*
RefreshLicenseStatsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RefreshLicenseStatsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this refresh license stats internal server error response has a 2xx status code
func (o *RefreshLicenseStatsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh license stats internal server error response has a 3xx status code
func (o *RefreshLicenseStatsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh license stats internal server error response has a 4xx status code
func (o *RefreshLicenseStatsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh license stats internal server error response has a 5xx status code
func (o *RefreshLicenseStatsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this refresh license stats internal server error response a status code equal to that given
func (o *RefreshLicenseStatsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the refresh license stats internal server error response
func (o *RefreshLicenseStatsInternalServerError) Code() int {
	return 500
}

func (o *RefreshLicenseStatsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /licensing/refresh-stats][%d] refreshLicenseStatsInternalServerError %s", 500, payload)
}

func (o *RefreshLicenseStatsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /licensing/refresh-stats][%d] refreshLicenseStatsInternalServerError %s", 500, payload)
}

func (o *RefreshLicenseStatsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RefreshLicenseStatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
