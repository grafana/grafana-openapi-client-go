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

// ResetPolicyTreeReader is a Reader for the ResetPolicyTree structure.
type ResetPolicyTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetPolicyTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewResetPolicyTreeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/provisioning/policies] ResetPolicyTree", response, response.Code())
	}
}

// NewResetPolicyTreeAccepted creates a ResetPolicyTreeAccepted with default headers values
func NewResetPolicyTreeAccepted() *ResetPolicyTreeAccepted {
	return &ResetPolicyTreeAccepted{}
}

/*
ResetPolicyTreeAccepted describes a response with status code 202, with default header values.

Ack
*/
type ResetPolicyTreeAccepted struct {
	Payload models.Ack
}

// IsSuccess returns true when this reset policy tree accepted response has a 2xx status code
func (o *ResetPolicyTreeAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reset policy tree accepted response has a 3xx status code
func (o *ResetPolicyTreeAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reset policy tree accepted response has a 4xx status code
func (o *ResetPolicyTreeAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this reset policy tree accepted response has a 5xx status code
func (o *ResetPolicyTreeAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this reset policy tree accepted response a status code equal to that given
func (o *ResetPolicyTreeAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the reset policy tree accepted response
func (o *ResetPolicyTreeAccepted) Code() int {
	return 202
}

func (o *ResetPolicyTreeAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/provisioning/policies][%d] resetPolicyTreeAccepted %s", 202, payload)
}

func (o *ResetPolicyTreeAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/provisioning/policies][%d] resetPolicyTreeAccepted %s", 202, payload)
}

func (o *ResetPolicyTreeAccepted) GetPayload() models.Ack {
	return o.Payload
}

func (o *ResetPolicyTreeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
