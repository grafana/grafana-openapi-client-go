// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRouteGetAlertRulesExportParams creates a new RouteGetAlertRulesExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteGetAlertRulesExportParams() *RouteGetAlertRulesExportParams {
	return &RouteGetAlertRulesExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteGetAlertRulesExportParamsWithTimeout creates a new RouteGetAlertRulesExportParams object
// with the ability to set a timeout on a request.
func NewRouteGetAlertRulesExportParamsWithTimeout(timeout time.Duration) *RouteGetAlertRulesExportParams {
	return &RouteGetAlertRulesExportParams{
		timeout: timeout,
	}
}

// NewRouteGetAlertRulesExportParamsWithContext creates a new RouteGetAlertRulesExportParams object
// with the ability to set a context for a request.
func NewRouteGetAlertRulesExportParamsWithContext(ctx context.Context) *RouteGetAlertRulesExportParams {
	return &RouteGetAlertRulesExportParams{
		Context: ctx,
	}
}

// NewRouteGetAlertRulesExportParamsWithHTTPClient creates a new RouteGetAlertRulesExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteGetAlertRulesExportParamsWithHTTPClient(client *http.Client) *RouteGetAlertRulesExportParams {
	return &RouteGetAlertRulesExportParams{
		HTTPClient: client,
	}
}

/*
RouteGetAlertRulesExportParams contains all the parameters to send to the API endpoint

	for the route get alert rules export operation.

	Typically these are written to a http.Request.
*/
type RouteGetAlertRulesExportParams struct {

	/* Download.

	   Whether to initiate a download of the file or not.
	*/
	Download *bool

	/* Format.

	   Format of the downloaded file, either yaml or json. Accept header can also be used, but the query parameter will take precedence.

	   Default: "yaml"
	*/
	Format *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route get alert rules export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetAlertRulesExportParams) WithDefaults() *RouteGetAlertRulesExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route get alert rules export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetAlertRulesExportParams) SetDefaults() {
	var (
		downloadDefault = bool(false)

		formatDefault = string("yaml")
	)

	val := RouteGetAlertRulesExportParams{
		Download: &downloadDefault,
		Format:   &formatDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) WithTimeout(timeout time.Duration) *RouteGetAlertRulesExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) WithContext(ctx context.Context) *RouteGetAlertRulesExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) WithHTTPClient(client *http.Client) *RouteGetAlertRulesExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDownload adds the download to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) WithDownload(download *bool) *RouteGetAlertRulesExportParams {
	o.SetDownload(download)
	return o
}

// SetDownload adds the download to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) SetDownload(download *bool) {
	o.Download = download
}

// WithFormat adds the format to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) WithFormat(format *string) *RouteGetAlertRulesExportParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the route get alert rules export params
func (o *RouteGetAlertRulesExportParams) SetFormat(format *string) {
	o.Format = format
}

// WriteToRequest writes these params to a swagger request
func (o *RouteGetAlertRulesExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Download != nil {

		// query param download
		var qrDownload bool

		if o.Download != nil {
			qrDownload = *o.Download
		}
		qDownload := swag.FormatBool(qrDownload)
		if qDownload != "" {

			if err := r.SetQueryParam("download", qDownload); err != nil {
				return err
			}
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}