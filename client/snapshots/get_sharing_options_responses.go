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

// GetSharingOptionsReader is a Reader for the GetSharingOptions structure.
type GetSharingOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSharingOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSharingOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSharingOptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /snapshot/shared-options] getSharingOptions", response, response.Code())
	}
}

// NewGetSharingOptionsOK creates a GetSharingOptionsOK with default headers values
func NewGetSharingOptionsOK() *GetSharingOptionsOK {
	return &GetSharingOptionsOK{}
}

/*
GetSharingOptionsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetSharingOptionsOK struct {
	Payload *models.GetSharingOptionsOKBody
}

// IsSuccess returns true when this get sharing options Ok response has a 2xx status code
func (o *GetSharingOptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sharing options Ok response has a 3xx status code
func (o *GetSharingOptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sharing options Ok response has a 4xx status code
func (o *GetSharingOptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sharing options Ok response has a 5xx status code
func (o *GetSharingOptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sharing options Ok response a status code equal to that given
func (o *GetSharingOptionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sharing options Ok response
func (o *GetSharingOptionsOK) Code() int {
	return 200
}

func (o *GetSharingOptionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /snapshot/shared-options][%d] getSharingOptionsOk %s", 200, payload)
}

func (o *GetSharingOptionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /snapshot/shared-options][%d] getSharingOptionsOk %s", 200, payload)
}

func (o *GetSharingOptionsOK) GetPayload() *models.GetSharingOptionsOKBody {
	return o.Payload
}

func (o *GetSharingOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSharingOptionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSharingOptionsUnauthorized creates a GetSharingOptionsUnauthorized with default headers values
func NewGetSharingOptionsUnauthorized() *GetSharingOptionsUnauthorized {
	return &GetSharingOptionsUnauthorized{}
}

/*
GetSharingOptionsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetSharingOptionsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get sharing options unauthorized response has a 2xx status code
func (o *GetSharingOptionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sharing options unauthorized response has a 3xx status code
func (o *GetSharingOptionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sharing options unauthorized response has a 4xx status code
func (o *GetSharingOptionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sharing options unauthorized response has a 5xx status code
func (o *GetSharingOptionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get sharing options unauthorized response a status code equal to that given
func (o *GetSharingOptionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get sharing options unauthorized response
func (o *GetSharingOptionsUnauthorized) Code() int {
	return 401
}

func (o *GetSharingOptionsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /snapshot/shared-options][%d] getSharingOptionsUnauthorized %s", 401, payload)
}

func (o *GetSharingOptionsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /snapshot/shared-options][%d] getSharingOptionsUnauthorized %s", 401, payload)
}

func (o *GetSharingOptionsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSharingOptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
