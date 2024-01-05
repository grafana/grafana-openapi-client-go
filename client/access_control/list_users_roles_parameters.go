// Code generated by go-swagger; DO NOT EDIT.

package access_control

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

// NewListUsersRolesParams creates a new ListUsersRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListUsersRolesParams() *ListUsersRolesParams {
	return &ListUsersRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListUsersRolesParamsWithTimeout creates a new ListUsersRolesParams object
// with the ability to set a timeout on a request.
func NewListUsersRolesParamsWithTimeout(timeout time.Duration) *ListUsersRolesParams {
	return &ListUsersRolesParams{
		timeout: timeout,
	}
}

// NewListUsersRolesParamsWithContext creates a new ListUsersRolesParams object
// with the ability to set a context for a request.
func NewListUsersRolesParamsWithContext(ctx context.Context) *ListUsersRolesParams {
	return &ListUsersRolesParams{
		Context: ctx,
	}
}

// NewListUsersRolesParamsWithHTTPClient creates a new ListUsersRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListUsersRolesParamsWithHTTPClient(client *http.Client) *ListUsersRolesParams {
	return &ListUsersRolesParams{
		HTTPClient: client,
	}
}

/*
ListUsersRolesParams contains all the parameters to send to the API endpoint

	for the list users roles operation.

	Typically these are written to a http.Request.
*/
type ListUsersRolesParams struct {

	// Body.
	Body *models.RolesSearchQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list users roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUsersRolesParams) WithDefaults() *ListUsersRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list users roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUsersRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list users roles params
func (o *ListUsersRolesParams) WithTimeout(timeout time.Duration) *ListUsersRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list users roles params
func (o *ListUsersRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list users roles params
func (o *ListUsersRolesParams) WithContext(ctx context.Context) *ListUsersRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list users roles params
func (o *ListUsersRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list users roles params
func (o *ListUsersRolesParams) WithHTTPClient(client *http.Client) *ListUsersRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list users roles params
func (o *ListUsersRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list users roles params
func (o *ListUsersRolesParams) WithBody(body *models.RolesSearchQuery) *ListUsersRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list users roles params
func (o *ListUsersRolesParams) SetBody(body *models.RolesSearchQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListUsersRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
