// Code generated by go-swagger; DO NOT EDIT.

package migrations

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
)

// NewCreateCloudMigrationTokenParams creates a new CreateCloudMigrationTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCloudMigrationTokenParams() *CreateCloudMigrationTokenParams {
	return &CreateCloudMigrationTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCloudMigrationTokenParamsWithTimeout creates a new CreateCloudMigrationTokenParams object
// with the ability to set a timeout on a request.
func NewCreateCloudMigrationTokenParamsWithTimeout(timeout time.Duration) *CreateCloudMigrationTokenParams {
	return &CreateCloudMigrationTokenParams{
		timeout: timeout,
	}
}

// NewCreateCloudMigrationTokenParamsWithContext creates a new CreateCloudMigrationTokenParams object
// with the ability to set a context for a request.
func NewCreateCloudMigrationTokenParamsWithContext(ctx context.Context) *CreateCloudMigrationTokenParams {
	return &CreateCloudMigrationTokenParams{
		Context: ctx,
	}
}

// NewCreateCloudMigrationTokenParamsWithHTTPClient creates a new CreateCloudMigrationTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCloudMigrationTokenParamsWithHTTPClient(client *http.Client) *CreateCloudMigrationTokenParams {
	return &CreateCloudMigrationTokenParams{
		HTTPClient: client,
	}
}

/*
CreateCloudMigrationTokenParams contains all the parameters to send to the API endpoint

	for the create cloud migration token operation.

	Typically these are written to a http.Request.
*/
type CreateCloudMigrationTokenParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cloud migration token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudMigrationTokenParams) WithDefaults() *CreateCloudMigrationTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cloud migration token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudMigrationTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cloud migration token params
func (o *CreateCloudMigrationTokenParams) WithTimeout(timeout time.Duration) *CreateCloudMigrationTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cloud migration token params
func (o *CreateCloudMigrationTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cloud migration token params
func (o *CreateCloudMigrationTokenParams) WithContext(ctx context.Context) *CreateCloudMigrationTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cloud migration token params
func (o *CreateCloudMigrationTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cloud migration token params
func (o *CreateCloudMigrationTokenParams) WithHTTPClient(client *http.Client) *CreateCloudMigrationTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cloud migration token params
func (o *CreateCloudMigrationTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCloudMigrationTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
