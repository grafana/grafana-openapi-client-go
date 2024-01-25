// Code generated by go-swagger; DO NOT EDIT.

package provisioning

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

// GetMuteTimingReader is a Reader for the GetMuteTiming structure.
type GetMuteTimingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMuteTimingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMuteTimingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetMuteTimingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/provisioning/mute-timings/{name}] GetMuteTiming", response, response.Code())
	}
}

// NewGetMuteTimingOK creates a GetMuteTimingOK with default headers values
func NewGetMuteTimingOK() *GetMuteTimingOK {
	return &GetMuteTimingOK{}
}

/*
GetMuteTimingOK describes a response with status code 200, with default header values.

MuteTimeInterval
*/
type GetMuteTimingOK struct {
	Payload *models.MuteTimeInterval
}

// IsSuccess returns true when this get mute timing Ok response has a 2xx status code
func (o *GetMuteTimingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mute timing Ok response has a 3xx status code
func (o *GetMuteTimingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mute timing Ok response has a 4xx status code
func (o *GetMuteTimingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mute timing Ok response has a 5xx status code
func (o *GetMuteTimingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mute timing Ok response a status code equal to that given
func (o *GetMuteTimingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mute timing Ok response
func (o *GetMuteTimingOK) Code() int {
	return 200
}

func (o *GetMuteTimingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/mute-timings/{name}][%d] getMuteTimingOk %s", 200, payload)
}

func (o *GetMuteTimingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/mute-timings/{name}][%d] getMuteTimingOk %s", 200, payload)
}

func (o *GetMuteTimingOK) GetPayload() *models.MuteTimeInterval {
	return o.Payload
}

func (o *GetMuteTimingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MuteTimeInterval)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMuteTimingNotFound creates a GetMuteTimingNotFound with default headers values
func NewGetMuteTimingNotFound() *GetMuteTimingNotFound {
	return &GetMuteTimingNotFound{}
}

/*
GetMuteTimingNotFound describes a response with status code 404, with default header values.

	Not found.
*/
type GetMuteTimingNotFound struct {
}

// IsSuccess returns true when this get mute timing not found response has a 2xx status code
func (o *GetMuteTimingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mute timing not found response has a 3xx status code
func (o *GetMuteTimingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mute timing not found response has a 4xx status code
func (o *GetMuteTimingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mute timing not found response has a 5xx status code
func (o *GetMuteTimingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get mute timing not found response a status code equal to that given
func (o *GetMuteTimingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get mute timing not found response
func (o *GetMuteTimingNotFound) Code() int {
	return 404
}

func (o *GetMuteTimingNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/provisioning/mute-timings/{name}][%d] getMuteTimingNotFound", 404)
}

func (o *GetMuteTimingNotFound) String() string {
	return fmt.Sprintf("[GET /v1/provisioning/mute-timings/{name}][%d] getMuteTimingNotFound", 404)
}

func (o *GetMuteTimingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
