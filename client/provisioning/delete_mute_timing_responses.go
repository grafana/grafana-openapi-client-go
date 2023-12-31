// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMuteTimingReader is a Reader for the DeleteMuteTiming structure.
type DeleteMuteTimingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMuteTimingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMuteTimingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/provisioning/mute-timings/{name}] DeleteMuteTiming", response, response.Code())
	}
}

// NewDeleteMuteTimingNoContent creates a DeleteMuteTimingNoContent with default headers values
func NewDeleteMuteTimingNoContent() *DeleteMuteTimingNoContent {
	return &DeleteMuteTimingNoContent{}
}

/*
DeleteMuteTimingNoContent describes a response with status code 204, with default header values.

	The mute timing was deleted successfully.
*/
type DeleteMuteTimingNoContent struct {
}

// IsSuccess returns true when this delete mute timing no content response has a 2xx status code
func (o *DeleteMuteTimingNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mute timing no content response has a 3xx status code
func (o *DeleteMuteTimingNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mute timing no content response has a 4xx status code
func (o *DeleteMuteTimingNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mute timing no content response has a 5xx status code
func (o *DeleteMuteTimingNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mute timing no content response a status code equal to that given
func (o *DeleteMuteTimingNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete mute timing no content response
func (o *DeleteMuteTimingNoContent) Code() int {
	return 204
}

func (o *DeleteMuteTimingNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/provisioning/mute-timings/{name}][%d] deleteMuteTimingNoContent ", 204)
}

func (o *DeleteMuteTimingNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/provisioning/mute-timings/{name}][%d] deleteMuteTimingNoContent ", 204)
}

func (o *DeleteMuteTimingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
