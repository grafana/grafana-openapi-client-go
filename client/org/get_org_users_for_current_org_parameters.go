// Code generated by go-swagger; DO NOT EDIT.

package org

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

// NewGetOrgUsersForCurrentOrgParams creates a new GetOrgUsersForCurrentOrgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgUsersForCurrentOrgParams() *GetOrgUsersForCurrentOrgParams {
	return &GetOrgUsersForCurrentOrgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgUsersForCurrentOrgParamsWithTimeout creates a new GetOrgUsersForCurrentOrgParams object
// with the ability to set a timeout on a request.
func NewGetOrgUsersForCurrentOrgParamsWithTimeout(timeout time.Duration) *GetOrgUsersForCurrentOrgParams {
	return &GetOrgUsersForCurrentOrgParams{
		timeout: timeout,
	}
}

// NewGetOrgUsersForCurrentOrgParamsWithContext creates a new GetOrgUsersForCurrentOrgParams object
// with the ability to set a context for a request.
func NewGetOrgUsersForCurrentOrgParamsWithContext(ctx context.Context) *GetOrgUsersForCurrentOrgParams {
	return &GetOrgUsersForCurrentOrgParams{
		Context: ctx,
	}
}

// NewGetOrgUsersForCurrentOrgParamsWithHTTPClient creates a new GetOrgUsersForCurrentOrgParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgUsersForCurrentOrgParamsWithHTTPClient(client *http.Client) *GetOrgUsersForCurrentOrgParams {
	return &GetOrgUsersForCurrentOrgParams{
		HTTPClient: client,
	}
}

/*
GetOrgUsersForCurrentOrgParams contains all the parameters to send to the API endpoint

	for the get org users for current org operation.

	Typically these are written to a http.Request.
*/
type GetOrgUsersForCurrentOrgParams struct {

	// Limit.
	//
	// Format: int64
	Limit *int64

	// Query.
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get org users for current org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgUsersForCurrentOrgParams) WithDefaults() *GetOrgUsersForCurrentOrgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get org users for current org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgUsersForCurrentOrgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) WithTimeout(timeout time.Duration) *GetOrgUsersForCurrentOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) WithContext(ctx context.Context) *GetOrgUsersForCurrentOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) WithHTTPClient(client *http.Client) *GetOrgUsersForCurrentOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) WithLimit(limit *int64) *GetOrgUsersForCurrentOrgParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithQuery adds the query to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) WithQuery(query *string) *GetOrgUsersForCurrentOrgParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get org users for current org params
func (o *GetOrgUsersForCurrentOrgParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgUsersForCurrentOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
