// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

// NewSearchOrgUsersParams creates a new SearchOrgUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchOrgUsersParams() *SearchOrgUsersParams {
	return &SearchOrgUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchOrgUsersParamsWithTimeout creates a new SearchOrgUsersParams object
// with the ability to set a timeout on a request.
func NewSearchOrgUsersParamsWithTimeout(timeout time.Duration) *SearchOrgUsersParams {
	return &SearchOrgUsersParams{
		timeout: timeout,
	}
}

// NewSearchOrgUsersParamsWithContext creates a new SearchOrgUsersParams object
// with the ability to set a context for a request.
func NewSearchOrgUsersParamsWithContext(ctx context.Context) *SearchOrgUsersParams {
	return &SearchOrgUsersParams{
		Context: ctx,
	}
}

// NewSearchOrgUsersParamsWithHTTPClient creates a new SearchOrgUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchOrgUsersParamsWithHTTPClient(client *http.Client) *SearchOrgUsersParams {
	return &SearchOrgUsersParams{
		HTTPClient: client,
	}
}

/*
SearchOrgUsersParams contains all the parameters to send to the API endpoint

	for the search org users operation.

	Typically these are written to a http.Request.
*/
type SearchOrgUsersParams struct {

	// OrgID.
	//
	// Format: int64
	OrgID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search org users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchOrgUsersParams) WithDefaults() *SearchOrgUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search org users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchOrgUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search org users params
func (o *SearchOrgUsersParams) WithTimeout(timeout time.Duration) *SearchOrgUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search org users params
func (o *SearchOrgUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search org users params
func (o *SearchOrgUsersParams) WithContext(ctx context.Context) *SearchOrgUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search org users params
func (o *SearchOrgUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search org users params
func (o *SearchOrgUsersParams) WithHTTPClient(client *http.Client) *SearchOrgUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search org users params
func (o *SearchOrgUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgID adds the orgID to the search org users params
func (o *SearchOrgUsersParams) WithOrgID(orgID int64) *SearchOrgUsersParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the search org users params
func (o *SearchOrgUsersParams) SetOrgID(orgID int64) {
	o.OrgID = orgID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchOrgUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param org_id
	if err := r.SetPathParam("org_id", swag.FormatInt64(o.OrgID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
