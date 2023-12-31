// Code generated by go-swagger; DO NOT EDIT.

package admin_users

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

// NewAdminUpdateUserPermissionsParams creates a new AdminUpdateUserPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminUpdateUserPermissionsParams() *AdminUpdateUserPermissionsParams {
	return &AdminUpdateUserPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminUpdateUserPermissionsParamsWithTimeout creates a new AdminUpdateUserPermissionsParams object
// with the ability to set a timeout on a request.
func NewAdminUpdateUserPermissionsParamsWithTimeout(timeout time.Duration) *AdminUpdateUserPermissionsParams {
	return &AdminUpdateUserPermissionsParams{
		timeout: timeout,
	}
}

// NewAdminUpdateUserPermissionsParamsWithContext creates a new AdminUpdateUserPermissionsParams object
// with the ability to set a context for a request.
func NewAdminUpdateUserPermissionsParamsWithContext(ctx context.Context) *AdminUpdateUserPermissionsParams {
	return &AdminUpdateUserPermissionsParams{
		Context: ctx,
	}
}

// NewAdminUpdateUserPermissionsParamsWithHTTPClient creates a new AdminUpdateUserPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminUpdateUserPermissionsParamsWithHTTPClient(client *http.Client) *AdminUpdateUserPermissionsParams {
	return &AdminUpdateUserPermissionsParams{
		HTTPClient: client,
	}
}

/*
AdminUpdateUserPermissionsParams contains all the parameters to send to the API endpoint

	for the admin update user permissions operation.

	Typically these are written to a http.Request.
*/
type AdminUpdateUserPermissionsParams struct {

	// Body.
	Body *models.AdminUpdateUserPermissionsForm

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin update user permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUpdateUserPermissionsParams) WithDefaults() *AdminUpdateUserPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin update user permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUpdateUserPermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) WithTimeout(timeout time.Duration) *AdminUpdateUserPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) WithContext(ctx context.Context) *AdminUpdateUserPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) WithHTTPClient(client *http.Client) *AdminUpdateUserPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) WithBody(body *models.AdminUpdateUserPermissionsForm) *AdminUpdateUserPermissionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) SetBody(body *models.AdminUpdateUserPermissionsForm) {
	o.Body = body
}

// WithUserID adds the userID to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) WithUserID(userID int64) *AdminUpdateUserPermissionsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin update user permissions params
func (o *AdminUpdateUserPermissionsParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminUpdateUserPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
