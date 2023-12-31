// Code generated by go-swagger; DO NOT EDIT.

package admin_ldap

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

// NewPostSyncUserWithLDAPParams creates a new PostSyncUserWithLDAPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSyncUserWithLDAPParams() *PostSyncUserWithLDAPParams {
	return &PostSyncUserWithLDAPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSyncUserWithLDAPParamsWithTimeout creates a new PostSyncUserWithLDAPParams object
// with the ability to set a timeout on a request.
func NewPostSyncUserWithLDAPParamsWithTimeout(timeout time.Duration) *PostSyncUserWithLDAPParams {
	return &PostSyncUserWithLDAPParams{
		timeout: timeout,
	}
}

// NewPostSyncUserWithLDAPParamsWithContext creates a new PostSyncUserWithLDAPParams object
// with the ability to set a context for a request.
func NewPostSyncUserWithLDAPParamsWithContext(ctx context.Context) *PostSyncUserWithLDAPParams {
	return &PostSyncUserWithLDAPParams{
		Context: ctx,
	}
}

// NewPostSyncUserWithLDAPParamsWithHTTPClient creates a new PostSyncUserWithLDAPParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSyncUserWithLDAPParamsWithHTTPClient(client *http.Client) *PostSyncUserWithLDAPParams {
	return &PostSyncUserWithLDAPParams{
		HTTPClient: client,
	}
}

/*
PostSyncUserWithLDAPParams contains all the parameters to send to the API endpoint

	for the post sync user with LDAP operation.

	Typically these are written to a http.Request.
*/
type PostSyncUserWithLDAPParams struct {

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post sync user with LDAP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSyncUserWithLDAPParams) WithDefaults() *PostSyncUserWithLDAPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post sync user with LDAP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSyncUserWithLDAPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post sync user with LDAP params
func (o *PostSyncUserWithLDAPParams) WithTimeout(timeout time.Duration) *PostSyncUserWithLDAPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sync user with LDAP params
func (o *PostSyncUserWithLDAPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sync user with LDAP params
func (o *PostSyncUserWithLDAPParams) WithContext(ctx context.Context) *PostSyncUserWithLDAPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sync user with LDAP params
func (o *PostSyncUserWithLDAPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sync user with LDAP params
func (o *PostSyncUserWithLDAPParams) WithHTTPClient(client *http.Client) *PostSyncUserWithLDAPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sync user with LDAP params
func (o *PostSyncUserWithLDAPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the post sync user with LDAP params
func (o *PostSyncUserWithLDAPParams) WithUserID(userID int64) *PostSyncUserWithLDAPParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post sync user with LDAP params
func (o *PostSyncUserWithLDAPParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSyncUserWithLDAPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
