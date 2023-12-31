// Code generated by go-swagger; DO NOT EDIT.

package reports

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

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewSaveReportSettingsParams creates a new SaveReportSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSaveReportSettingsParams() *SaveReportSettingsParams {
	return &SaveReportSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSaveReportSettingsParamsWithTimeout creates a new SaveReportSettingsParams object
// with the ability to set a timeout on a request.
func NewSaveReportSettingsParamsWithTimeout(timeout time.Duration) *SaveReportSettingsParams {
	return &SaveReportSettingsParams{
		timeout: timeout,
	}
}

// NewSaveReportSettingsParamsWithContext creates a new SaveReportSettingsParams object
// with the ability to set a context for a request.
func NewSaveReportSettingsParamsWithContext(ctx context.Context) *SaveReportSettingsParams {
	return &SaveReportSettingsParams{
		Context: ctx,
	}
}

// NewSaveReportSettingsParamsWithHTTPClient creates a new SaveReportSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSaveReportSettingsParamsWithHTTPClient(client *http.Client) *SaveReportSettingsParams {
	return &SaveReportSettingsParams{
		HTTPClient: client,
	}
}

/*
SaveReportSettingsParams contains all the parameters to send to the API endpoint

	for the save report settings operation.

	Typically these are written to a http.Request.
*/
type SaveReportSettingsParams struct {

	// Body.
	Body *models.ReportSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the save report settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveReportSettingsParams) WithDefaults() *SaveReportSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the save report settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveReportSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the save report settings params
func (o *SaveReportSettingsParams) WithTimeout(timeout time.Duration) *SaveReportSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save report settings params
func (o *SaveReportSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save report settings params
func (o *SaveReportSettingsParams) WithContext(ctx context.Context) *SaveReportSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save report settings params
func (o *SaveReportSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save report settings params
func (o *SaveReportSettingsParams) WithHTTPClient(client *http.Client) *SaveReportSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save report settings params
func (o *SaveReportSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the save report settings params
func (o *SaveReportSettingsParams) WithBody(body *models.ReportSettings) *SaveReportSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the save report settings params
func (o *SaveReportSettingsParams) SetBody(body *models.ReportSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SaveReportSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
