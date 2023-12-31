// Code generated by go-swagger; DO NOT EDIT.

package org_preferences

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

// NewUpdateOrgPreferencesParams creates a new UpdateOrgPreferencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrgPreferencesParams() *UpdateOrgPreferencesParams {
	return &UpdateOrgPreferencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrgPreferencesParamsWithTimeout creates a new UpdateOrgPreferencesParams object
// with the ability to set a timeout on a request.
func NewUpdateOrgPreferencesParamsWithTimeout(timeout time.Duration) *UpdateOrgPreferencesParams {
	return &UpdateOrgPreferencesParams{
		timeout: timeout,
	}
}

// NewUpdateOrgPreferencesParamsWithContext creates a new UpdateOrgPreferencesParams object
// with the ability to set a context for a request.
func NewUpdateOrgPreferencesParamsWithContext(ctx context.Context) *UpdateOrgPreferencesParams {
	return &UpdateOrgPreferencesParams{
		Context: ctx,
	}
}

// NewUpdateOrgPreferencesParamsWithHTTPClient creates a new UpdateOrgPreferencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrgPreferencesParamsWithHTTPClient(client *http.Client) *UpdateOrgPreferencesParams {
	return &UpdateOrgPreferencesParams{
		HTTPClient: client,
	}
}

/*
UpdateOrgPreferencesParams contains all the parameters to send to the API endpoint

	for the update org preferences operation.

	Typically these are written to a http.Request.
*/
type UpdateOrgPreferencesParams struct {

	// Body.
	Body *models.UpdatePrefsCmd

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update org preferences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrgPreferencesParams) WithDefaults() *UpdateOrgPreferencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update org preferences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrgPreferencesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update org preferences params
func (o *UpdateOrgPreferencesParams) WithTimeout(timeout time.Duration) *UpdateOrgPreferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update org preferences params
func (o *UpdateOrgPreferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update org preferences params
func (o *UpdateOrgPreferencesParams) WithContext(ctx context.Context) *UpdateOrgPreferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update org preferences params
func (o *UpdateOrgPreferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update org preferences params
func (o *UpdateOrgPreferencesParams) WithHTTPClient(client *http.Client) *UpdateOrgPreferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update org preferences params
func (o *UpdateOrgPreferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update org preferences params
func (o *UpdateOrgPreferencesParams) WithBody(body *models.UpdatePrefsCmd) *UpdateOrgPreferencesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update org preferences params
func (o *UpdateOrgPreferencesParams) SetBody(body *models.UpdatePrefsCmd) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrgPreferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
