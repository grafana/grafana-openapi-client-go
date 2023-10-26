// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// RouteGetAlertRuleGroupExportReader is a Reader for the RouteGetAlertRuleGroupExport structure.
type RouteGetAlertRuleGroupExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetAlertRuleGroupExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetAlertRuleGroupExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRouteGetAlertRuleGroupExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}/export] RouteGetAlertRuleGroupExport", response, response.Code())
	}
}

// NewRouteGetAlertRuleGroupExportOK creates a RouteGetAlertRuleGroupExportOK with default headers values
func NewRouteGetAlertRuleGroupExportOK() *RouteGetAlertRuleGroupExportOK {
	return &RouteGetAlertRuleGroupExportOK{}
}

/*
RouteGetAlertRuleGroupExportOK describes a response with status code 200, with default header values.

AlertingFileExport
*/
type RouteGetAlertRuleGroupExportOK struct {
	Payload *models.AlertingFileExport
}

// IsSuccess returns true when this route get alert rule group export Ok response has a 2xx status code
func (o *RouteGetAlertRuleGroupExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this route get alert rule group export Ok response has a 3xx status code
func (o *RouteGetAlertRuleGroupExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route get alert rule group export Ok response has a 4xx status code
func (o *RouteGetAlertRuleGroupExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this route get alert rule group export Ok response has a 5xx status code
func (o *RouteGetAlertRuleGroupExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this route get alert rule group export Ok response a status code equal to that given
func (o *RouteGetAlertRuleGroupExportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the route get alert rule group export Ok response
func (o *RouteGetAlertRuleGroupExportOK) Code() int {
	return 200
}

func (o *RouteGetAlertRuleGroupExportOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}/export][%d] routeGetAlertRuleGroupExportOk  %+v", 200, o.Payload)
}

func (o *RouteGetAlertRuleGroupExportOK) String() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}/export][%d] routeGetAlertRuleGroupExportOk  %+v", 200, o.Payload)
}

func (o *RouteGetAlertRuleGroupExportOK) GetPayload() *models.AlertingFileExport {
	return o.Payload
}

func (o *RouteGetAlertRuleGroupExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingFileExport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteGetAlertRuleGroupExportNotFound creates a RouteGetAlertRuleGroupExportNotFound with default headers values
func NewRouteGetAlertRuleGroupExportNotFound() *RouteGetAlertRuleGroupExportNotFound {
	return &RouteGetAlertRuleGroupExportNotFound{}
}

/*
RouteGetAlertRuleGroupExportNotFound describes a response with status code 404, with default header values.

	Not found.
*/
type RouteGetAlertRuleGroupExportNotFound struct {
}

// IsSuccess returns true when this route get alert rule group export not found response has a 2xx status code
func (o *RouteGetAlertRuleGroupExportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this route get alert rule group export not found response has a 3xx status code
func (o *RouteGetAlertRuleGroupExportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route get alert rule group export not found response has a 4xx status code
func (o *RouteGetAlertRuleGroupExportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this route get alert rule group export not found response has a 5xx status code
func (o *RouteGetAlertRuleGroupExportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this route get alert rule group export not found response a status code equal to that given
func (o *RouteGetAlertRuleGroupExportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the route get alert rule group export not found response
func (o *RouteGetAlertRuleGroupExportNotFound) Code() int {
	return 404
}

func (o *RouteGetAlertRuleGroupExportNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}/export][%d] routeGetAlertRuleGroupExportNotFound ", 404)
}

func (o *RouteGetAlertRuleGroupExportNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}/export][%d] routeGetAlertRuleGroupExportNotFound ", 404)
}

func (o *RouteGetAlertRuleGroupExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
