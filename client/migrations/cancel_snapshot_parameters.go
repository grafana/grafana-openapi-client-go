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

// NewCancelSnapshotParams creates a new CancelSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelSnapshotParams() *CancelSnapshotParams {
	return &CancelSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelSnapshotParamsWithTimeout creates a new CancelSnapshotParams object
// with the ability to set a timeout on a request.
func NewCancelSnapshotParamsWithTimeout(timeout time.Duration) *CancelSnapshotParams {
	return &CancelSnapshotParams{
		timeout: timeout,
	}
}

// NewCancelSnapshotParamsWithContext creates a new CancelSnapshotParams object
// with the ability to set a context for a request.
func NewCancelSnapshotParamsWithContext(ctx context.Context) *CancelSnapshotParams {
	return &CancelSnapshotParams{
		Context: ctx,
	}
}

// NewCancelSnapshotParamsWithHTTPClient creates a new CancelSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelSnapshotParamsWithHTTPClient(client *http.Client) *CancelSnapshotParams {
	return &CancelSnapshotParams{
		HTTPClient: client,
	}
}

/*
CancelSnapshotParams contains all the parameters to send to the API endpoint

	for the cancel snapshot operation.

	Typically these are written to a http.Request.
*/
type CancelSnapshotParams struct {

	/* SnapshotUID.

	   UID of a snapshot
	*/
	SnapshotUID string

	/* UID.

	   Session UID of a session
	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelSnapshotParams) WithDefaults() *CancelSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel snapshot params
func (o *CancelSnapshotParams) WithTimeout(timeout time.Duration) *CancelSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel snapshot params
func (o *CancelSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel snapshot params
func (o *CancelSnapshotParams) WithContext(ctx context.Context) *CancelSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel snapshot params
func (o *CancelSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel snapshot params
func (o *CancelSnapshotParams) WithHTTPClient(client *http.Client) *CancelSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel snapshot params
func (o *CancelSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotUID adds the snapshotUID to the cancel snapshot params
func (o *CancelSnapshotParams) WithSnapshotUID(snapshotUID string) *CancelSnapshotParams {
	o.SetSnapshotUID(snapshotUID)
	return o
}

// SetSnapshotUID adds the snapshotUid to the cancel snapshot params
func (o *CancelSnapshotParams) SetSnapshotUID(snapshotUID string) {
	o.SnapshotUID = snapshotUID
}

// WithUID adds the uid to the cancel snapshot params
func (o *CancelSnapshotParams) WithUID(uid string) *CancelSnapshotParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the cancel snapshot params
func (o *CancelSnapshotParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *CancelSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param snapshotUid
	if err := r.SetPathParam("snapshotUid", o.SnapshotUID); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}