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

// PostLicenseTokenReader is a Reader for the PostLicenseToken structure.
type PostLicenseTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLicenseTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLicenseTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLicenseTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /licensing/token] postLicenseToken", response, response.Code())
	}
}

// NewPostLicenseTokenOK creates a PostLicenseTokenOK with default headers values
func NewPostLicenseTokenOK() *PostLicenseTokenOK {
	return &PostLicenseTokenOK{}
}

/*
PostLicenseTokenOK describes a response with status code 200, with default header values.

(empty)
*/
type PostLicenseTokenOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this post license token Ok response has a 2xx status code
func (o *PostLicenseTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post license token Ok response has a 3xx status code
func (o *PostLicenseTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license token Ok response has a 4xx status code
func (o *PostLicenseTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license token Ok response has a 5xx status code
func (o *PostLicenseTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post license token Ok response a status code equal to that given
func (o *PostLicenseTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post license token Ok response
func (o *PostLicenseTokenOK) Code() int {
	return 200
}

func (o *PostLicenseTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /licensing/token][%d] postLicenseTokenOk %s", 200, payload)
}

func (o *PostLicenseTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /licensing/token][%d] postLicenseTokenOk %s", 200, payload)
}

func (o *PostLicenseTokenOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *PostLicenseTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseTokenBadRequest creates a PostLicenseTokenBadRequest with default headers values
func NewPostLicenseTokenBadRequest() *PostLicenseTokenBadRequest {
	return &PostLicenseTokenBadRequest{}
}

/*
PostLicenseTokenBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type PostLicenseTokenBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this post license token bad request response has a 2xx status code
func (o *PostLicenseTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license token bad request response has a 3xx status code
func (o *PostLicenseTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license token bad request response has a 4xx status code
func (o *PostLicenseTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license token bad request response has a 5xx status code
func (o *PostLicenseTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post license token bad request response a status code equal to that given
func (o *PostLicenseTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post license token bad request response
func (o *PostLicenseTokenBadRequest) Code() int {
	return 400
}

func (o *PostLicenseTokenBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /licensing/token][%d] postLicenseTokenBadRequest %s", 400, payload)
}

func (o *PostLicenseTokenBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /licensing/token][%d] postLicenseTokenBadRequest %s", 400, payload)
}

func (o *PostLicenseTokenBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostLicenseTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
