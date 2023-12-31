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

// NewDeleteAlertRuleParams creates a new DeleteAlertRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAlertRuleParams() *DeleteAlertRuleParams {
	return &DeleteAlertRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAlertRuleParamsWithTimeout creates a new DeleteAlertRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteAlertRuleParamsWithTimeout(timeout time.Duration) *DeleteAlertRuleParams {
	return &DeleteAlertRuleParams{
		timeout: timeout,
	}
}

// NewDeleteAlertRuleParamsWithContext creates a new DeleteAlertRuleParams object
// with the ability to set a context for a request.
func NewDeleteAlertRuleParamsWithContext(ctx context.Context) *DeleteAlertRuleParams {
	return &DeleteAlertRuleParams{
		Context: ctx,
	}
}

// NewDeleteAlertRuleParamsWithHTTPClient creates a new DeleteAlertRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAlertRuleParamsWithHTTPClient(client *http.Client) *DeleteAlertRuleParams {
	return &DeleteAlertRuleParams{
		HTTPClient: client,
	}
}

/*
DeleteAlertRuleParams contains all the parameters to send to the API endpoint

	for the delete alert rule operation.

	Typically these are written to a http.Request.
*/
type DeleteAlertRuleParams struct {

	/* UID.

	   Alert rule UID
	*/
	UID string

	// XDisableProvenance.
	XDisableProvenance *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAlertRuleParams) WithDefaults() *DeleteAlertRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAlertRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete alert rule params
func (o *DeleteAlertRuleParams) WithTimeout(timeout time.Duration) *DeleteAlertRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete alert rule params
func (o *DeleteAlertRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete alert rule params
func (o *DeleteAlertRuleParams) WithContext(ctx context.Context) *DeleteAlertRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete alert rule params
func (o *DeleteAlertRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete alert rule params
func (o *DeleteAlertRuleParams) WithHTTPClient(client *http.Client) *DeleteAlertRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete alert rule params
func (o *DeleteAlertRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the delete alert rule params
func (o *DeleteAlertRuleParams) WithUID(uid string) *DeleteAlertRuleParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the delete alert rule params
func (o *DeleteAlertRuleParams) SetUID(uid string) {
	o.UID = uid
}

// WithXDisableProvenance adds the xDisableProvenance to the delete alert rule params
func (o *DeleteAlertRuleParams) WithXDisableProvenance(xDisableProvenance *string) *DeleteAlertRuleParams {
	o.SetXDisableProvenance(xDisableProvenance)
	return o
}

// SetXDisableProvenance adds the xDisableProvenance to the delete alert rule params
func (o *DeleteAlertRuleParams) SetXDisableProvenance(xDisableProvenance *string) {
	o.XDisableProvenance = xDisableProvenance
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAlertRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param UID
	if err := r.SetPathParam("UID", o.UID); err != nil {
		return err
	}

	if o.XDisableProvenance != nil {

		// header param X-Disable-Provenance
		if err := r.SetHeaderParam("X-Disable-Provenance", *o.XDisableProvenance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
