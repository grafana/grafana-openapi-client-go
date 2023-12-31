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

// NewAdminRevokeUserAuthTokenParams creates a new AdminRevokeUserAuthTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminRevokeUserAuthTokenParams() *AdminRevokeUserAuthTokenParams {
	return &AdminRevokeUserAuthTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminRevokeUserAuthTokenParamsWithTimeout creates a new AdminRevokeUserAuthTokenParams object
// with the ability to set a timeout on a request.
func NewAdminRevokeUserAuthTokenParamsWithTimeout(timeout time.Duration) *AdminRevokeUserAuthTokenParams {
	return &AdminRevokeUserAuthTokenParams{
		timeout: timeout,
	}
}

// NewAdminRevokeUserAuthTokenParamsWithContext creates a new AdminRevokeUserAuthTokenParams object
// with the ability to set a context for a request.
func NewAdminRevokeUserAuthTokenParamsWithContext(ctx context.Context) *AdminRevokeUserAuthTokenParams {
	return &AdminRevokeUserAuthTokenParams{
		Context: ctx,
	}
}

// NewAdminRevokeUserAuthTokenParamsWithHTTPClient creates a new AdminRevokeUserAuthTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminRevokeUserAuthTokenParamsWithHTTPClient(client *http.Client) *AdminRevokeUserAuthTokenParams {
	return &AdminRevokeUserAuthTokenParams{
		HTTPClient: client,
	}
}

/*
AdminRevokeUserAuthTokenParams contains all the parameters to send to the API endpoint

	for the admin revoke user auth token operation.

	Typically these are written to a http.Request.
*/
type AdminRevokeUserAuthTokenParams struct {

	// Body.
	Body *models.RevokeAuthTokenCmd

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin revoke user auth token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminRevokeUserAuthTokenParams) WithDefaults() *AdminRevokeUserAuthTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin revoke user auth token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminRevokeUserAuthTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) WithTimeout(timeout time.Duration) *AdminRevokeUserAuthTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) WithContext(ctx context.Context) *AdminRevokeUserAuthTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) WithHTTPClient(client *http.Client) *AdminRevokeUserAuthTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) WithBody(body *models.RevokeAuthTokenCmd) *AdminRevokeUserAuthTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) SetBody(body *models.RevokeAuthTokenCmd) {
	o.Body = body
}

// WithUserID adds the userID to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) WithUserID(userID int64) *AdminRevokeUserAuthTokenParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin revoke user auth token params
func (o *AdminRevokeUserAuthTokenParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminRevokeUserAuthTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
