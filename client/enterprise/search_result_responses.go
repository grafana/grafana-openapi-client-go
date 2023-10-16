// Code generated by go-swagger; DO NOT EDIT.

package enterprise

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// SearchResultReader is a Reader for the SearchResult structure.
type SearchResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchResultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchResultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchResultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /access-control/assignments/search] searchResult", response, response.Code())
	}
}

// NewSearchResultOK creates a SearchResultOK with default headers values
func NewSearchResultOK() *SearchResultOK {
	return &SearchResultOK{}
}

/*
SearchResultOK describes a response with status code 200, with default header values.

(empty)
*/
type SearchResultOK struct {
	Payload *models.SearchResult
}

// IsSuccess returns true when this search result o k response has a 2xx status code
func (o *SearchResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search result o k response has a 3xx status code
func (o *SearchResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search result o k response has a 4xx status code
func (o *SearchResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search result o k response has a 5xx status code
func (o *SearchResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search result o k response a status code equal to that given
func (o *SearchResultOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search result o k response
func (o *SearchResultOK) Code() int {
	return 200
}

func (o *SearchResultOK) Error() string {
	return fmt.Sprintf("[POST /access-control/assignments/search][%d] searchResultOK  %+v", 200, o.Payload)
}

func (o *SearchResultOK) String() string {
	return fmt.Sprintf("[POST /access-control/assignments/search][%d] searchResultOK  %+v", 200, o.Payload)
}

func (o *SearchResultOK) GetPayload() *models.SearchResult {
	return o.Payload
}

func (o *SearchResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchResultBadRequest creates a SearchResultBadRequest with default headers values
func NewSearchResultBadRequest() *SearchResultBadRequest {
	return &SearchResultBadRequest{}
}

/*
SearchResultBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type SearchResultBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search result bad request response has a 2xx status code
func (o *SearchResultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search result bad request response has a 3xx status code
func (o *SearchResultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search result bad request response has a 4xx status code
func (o *SearchResultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search result bad request response has a 5xx status code
func (o *SearchResultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search result bad request response a status code equal to that given
func (o *SearchResultBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the search result bad request response
func (o *SearchResultBadRequest) Code() int {
	return 400
}

func (o *SearchResultBadRequest) Error() string {
	return fmt.Sprintf("[POST /access-control/assignments/search][%d] searchResultBadRequest  %+v", 400, o.Payload)
}

func (o *SearchResultBadRequest) String() string {
	return fmt.Sprintf("[POST /access-control/assignments/search][%d] searchResultBadRequest  %+v", 400, o.Payload)
}

func (o *SearchResultBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchResultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchResultForbidden creates a SearchResultForbidden with default headers values
func NewSearchResultForbidden() *SearchResultForbidden {
	return &SearchResultForbidden{}
}

/*
SearchResultForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SearchResultForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search result forbidden response has a 2xx status code
func (o *SearchResultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search result forbidden response has a 3xx status code
func (o *SearchResultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search result forbidden response has a 4xx status code
func (o *SearchResultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search result forbidden response has a 5xx status code
func (o *SearchResultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search result forbidden response a status code equal to that given
func (o *SearchResultForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search result forbidden response
func (o *SearchResultForbidden) Code() int {
	return 403
}

func (o *SearchResultForbidden) Error() string {
	return fmt.Sprintf("[POST /access-control/assignments/search][%d] searchResultForbidden  %+v", 403, o.Payload)
}

func (o *SearchResultForbidden) String() string {
	return fmt.Sprintf("[POST /access-control/assignments/search][%d] searchResultForbidden  %+v", 403, o.Payload)
}

func (o *SearchResultForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchResultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchResultInternalServerError creates a SearchResultInternalServerError with default headers values
func NewSearchResultInternalServerError() *SearchResultInternalServerError {
	return &SearchResultInternalServerError{}
}

/*
SearchResultInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SearchResultInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search result internal server error response has a 2xx status code
func (o *SearchResultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search result internal server error response has a 3xx status code
func (o *SearchResultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search result internal server error response has a 4xx status code
func (o *SearchResultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search result internal server error response has a 5xx status code
func (o *SearchResultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search result internal server error response a status code equal to that given
func (o *SearchResultInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the search result internal server error response
func (o *SearchResultInternalServerError) Code() int {
	return 500
}

func (o *SearchResultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access-control/assignments/search][%d] searchResultInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchResultInternalServerError) String() string {
	return fmt.Sprintf("[POST /access-control/assignments/search][%d] searchResultInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchResultInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchResultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}