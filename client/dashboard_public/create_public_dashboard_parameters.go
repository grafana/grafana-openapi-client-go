// Code generated by go-swagger; DO NOT EDIT.

package dashboard_public

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

// NewCreatePublicDashboardParams creates a new CreatePublicDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePublicDashboardParams() *CreatePublicDashboardParams {
	return &CreatePublicDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePublicDashboardParamsWithTimeout creates a new CreatePublicDashboardParams object
// with the ability to set a timeout on a request.
func NewCreatePublicDashboardParamsWithTimeout(timeout time.Duration) *CreatePublicDashboardParams {
	return &CreatePublicDashboardParams{
		timeout: timeout,
	}
}

// NewCreatePublicDashboardParamsWithContext creates a new CreatePublicDashboardParams object
// with the ability to set a context for a request.
func NewCreatePublicDashboardParamsWithContext(ctx context.Context) *CreatePublicDashboardParams {
	return &CreatePublicDashboardParams{
		Context: ctx,
	}
}

// NewCreatePublicDashboardParamsWithHTTPClient creates a new CreatePublicDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePublicDashboardParamsWithHTTPClient(client *http.Client) *CreatePublicDashboardParams {
	return &CreatePublicDashboardParams{
		HTTPClient: client,
	}
}

/*
CreatePublicDashboardParams contains all the parameters to send to the API endpoint

	for the create public dashboard operation.

	Typically these are written to a http.Request.
*/
type CreatePublicDashboardParams struct {

	// Body.
	Body *models.PublicDashboardDTO

	// DashboardUID.
	DashboardUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create public dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePublicDashboardParams) WithDefaults() *CreatePublicDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create public dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePublicDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create public dashboard params
func (o *CreatePublicDashboardParams) WithTimeout(timeout time.Duration) *CreatePublicDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create public dashboard params
func (o *CreatePublicDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create public dashboard params
func (o *CreatePublicDashboardParams) WithContext(ctx context.Context) *CreatePublicDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create public dashboard params
func (o *CreatePublicDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create public dashboard params
func (o *CreatePublicDashboardParams) WithHTTPClient(client *http.Client) *CreatePublicDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create public dashboard params
func (o *CreatePublicDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create public dashboard params
func (o *CreatePublicDashboardParams) WithBody(body *models.PublicDashboardDTO) *CreatePublicDashboardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create public dashboard params
func (o *CreatePublicDashboardParams) SetBody(body *models.PublicDashboardDTO) {
	o.Body = body
}

// WithDashboardUID adds the dashboardUID to the create public dashboard params
func (o *CreatePublicDashboardParams) WithDashboardUID(dashboardUID string) *CreatePublicDashboardParams {
	o.SetDashboardUID(dashboardUID)
	return o
}

// SetDashboardUID adds the dashboardUid to the create public dashboard params
func (o *CreatePublicDashboardParams) SetDashboardUID(dashboardUID string) {
	o.DashboardUID = dashboardUID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePublicDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dashboardUid
	if err := r.SetPathParam("dashboardUid", o.DashboardUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
