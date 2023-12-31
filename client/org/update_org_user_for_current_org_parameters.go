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

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewUpdateOrgUserForCurrentOrgParams creates a new UpdateOrgUserForCurrentOrgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrgUserForCurrentOrgParams() *UpdateOrgUserForCurrentOrgParams {
	return &UpdateOrgUserForCurrentOrgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrgUserForCurrentOrgParamsWithTimeout creates a new UpdateOrgUserForCurrentOrgParams object
// with the ability to set a timeout on a request.
func NewUpdateOrgUserForCurrentOrgParamsWithTimeout(timeout time.Duration) *UpdateOrgUserForCurrentOrgParams {
	return &UpdateOrgUserForCurrentOrgParams{
		timeout: timeout,
	}
}

// NewUpdateOrgUserForCurrentOrgParamsWithContext creates a new UpdateOrgUserForCurrentOrgParams object
// with the ability to set a context for a request.
func NewUpdateOrgUserForCurrentOrgParamsWithContext(ctx context.Context) *UpdateOrgUserForCurrentOrgParams {
	return &UpdateOrgUserForCurrentOrgParams{
		Context: ctx,
	}
}

// NewUpdateOrgUserForCurrentOrgParamsWithHTTPClient creates a new UpdateOrgUserForCurrentOrgParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrgUserForCurrentOrgParamsWithHTTPClient(client *http.Client) *UpdateOrgUserForCurrentOrgParams {
	return &UpdateOrgUserForCurrentOrgParams{
		HTTPClient: client,
	}
}

/*
UpdateOrgUserForCurrentOrgParams contains all the parameters to send to the API endpoint

	for the update org user for current org operation.

	Typically these are written to a http.Request.
*/
type UpdateOrgUserForCurrentOrgParams struct {

	// Body.
	Body *models.UpdateOrgUserCommand

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update org user for current org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrgUserForCurrentOrgParams) WithDefaults() *UpdateOrgUserForCurrentOrgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update org user for current org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrgUserForCurrentOrgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) WithTimeout(timeout time.Duration) *UpdateOrgUserForCurrentOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) WithContext(ctx context.Context) *UpdateOrgUserForCurrentOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) WithHTTPClient(client *http.Client) *UpdateOrgUserForCurrentOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) WithBody(body *models.UpdateOrgUserCommand) *UpdateOrgUserForCurrentOrgParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) SetBody(body *models.UpdateOrgUserCommand) {
	o.Body = body
}

// WithUserID adds the userID to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) WithUserID(userID int64) *UpdateOrgUserForCurrentOrgParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update org user for current org params
func (o *UpdateOrgUserForCurrentOrgParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrgUserForCurrentOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
