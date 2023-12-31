// Code generated by go-swagger; DO NOT EDIT.

package provisioning

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

// NewDeleteContactpointsParams creates a new DeleteContactpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContactpointsParams() *DeleteContactpointsParams {
	return &DeleteContactpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContactpointsParamsWithTimeout creates a new DeleteContactpointsParams object
// with the ability to set a timeout on a request.
func NewDeleteContactpointsParamsWithTimeout(timeout time.Duration) *DeleteContactpointsParams {
	return &DeleteContactpointsParams{
		timeout: timeout,
	}
}

// NewDeleteContactpointsParamsWithContext creates a new DeleteContactpointsParams object
// with the ability to set a context for a request.
func NewDeleteContactpointsParamsWithContext(ctx context.Context) *DeleteContactpointsParams {
	return &DeleteContactpointsParams{
		Context: ctx,
	}
}

// NewDeleteContactpointsParamsWithHTTPClient creates a new DeleteContactpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContactpointsParamsWithHTTPClient(client *http.Client) *DeleteContactpointsParams {
	return &DeleteContactpointsParams{
		HTTPClient: client,
	}
}

/*
DeleteContactpointsParams contains all the parameters to send to the API endpoint

	for the delete contactpoints operation.

	Typically these are written to a http.Request.
*/
type DeleteContactpointsParams struct {

	/* UID.

	   UID is the contact point unique identifier
	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete contactpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContactpointsParams) WithDefaults() *DeleteContactpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete contactpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContactpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete contactpoints params
func (o *DeleteContactpointsParams) WithTimeout(timeout time.Duration) *DeleteContactpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete contactpoints params
func (o *DeleteContactpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete contactpoints params
func (o *DeleteContactpointsParams) WithContext(ctx context.Context) *DeleteContactpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete contactpoints params
func (o *DeleteContactpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete contactpoints params
func (o *DeleteContactpointsParams) WithHTTPClient(client *http.Client) *DeleteContactpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete contactpoints params
func (o *DeleteContactpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the delete contactpoints params
func (o *DeleteContactpointsParams) WithUID(uid string) *DeleteContactpointsParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the delete contactpoints params
func (o *DeleteContactpointsParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContactpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param UID
	if err := r.SetPathParam("UID", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
