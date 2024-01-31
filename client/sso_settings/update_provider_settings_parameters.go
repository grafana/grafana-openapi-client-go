// Code generated by go-swagger; DO NOT EDIT.

package sso_settings

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

// NewUpdateProviderSettingsParams creates a new UpdateProviderSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProviderSettingsParams() *UpdateProviderSettingsParams {
	return &UpdateProviderSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProviderSettingsParamsWithTimeout creates a new UpdateProviderSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateProviderSettingsParamsWithTimeout(timeout time.Duration) *UpdateProviderSettingsParams {
	return &UpdateProviderSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateProviderSettingsParamsWithContext creates a new UpdateProviderSettingsParams object
// with the ability to set a context for a request.
func NewUpdateProviderSettingsParamsWithContext(ctx context.Context) *UpdateProviderSettingsParams {
	return &UpdateProviderSettingsParams{
		Context: ctx,
	}
}

// NewUpdateProviderSettingsParamsWithHTTPClient creates a new UpdateProviderSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProviderSettingsParamsWithHTTPClient(client *http.Client) *UpdateProviderSettingsParams {
	return &UpdateProviderSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateProviderSettingsParams contains all the parameters to send to the API endpoint

	for the update provider settings operation.

	Typically these are written to a http.Request.
*/
type UpdateProviderSettingsParams struct {

	// Body.
	Body *models.UpdateProviderSettingsParamsBody

	// Key.
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update provider settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProviderSettingsParams) WithDefaults() *UpdateProviderSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update provider settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProviderSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update provider settings params
func (o *UpdateProviderSettingsParams) WithTimeout(timeout time.Duration) *UpdateProviderSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update provider settings params
func (o *UpdateProviderSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update provider settings params
func (o *UpdateProviderSettingsParams) WithContext(ctx context.Context) *UpdateProviderSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update provider settings params
func (o *UpdateProviderSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update provider settings params
func (o *UpdateProviderSettingsParams) WithHTTPClient(client *http.Client) *UpdateProviderSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update provider settings params
func (o *UpdateProviderSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update provider settings params
func (o *UpdateProviderSettingsParams) WithBody(body *models.UpdateProviderSettingsParamsBody) *UpdateProviderSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update provider settings params
func (o *UpdateProviderSettingsParams) SetBody(body *models.UpdateProviderSettingsParamsBody) {
	o.Body = body
}

// WithKey adds the key to the update provider settings params
func (o *UpdateProviderSettingsParams) WithKey(key string) *UpdateProviderSettingsParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the update provider settings params
func (o *UpdateProviderSettingsParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProviderSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
