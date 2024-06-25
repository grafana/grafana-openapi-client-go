// Code generated by go-swagger; DO NOT EDIT.

package enterprise

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

// NewGetTeamLBACRulesAPIParams creates a new GetTeamLBACRulesAPIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTeamLBACRulesAPIParams() *GetTeamLBACRulesAPIParams {
	return &GetTeamLBACRulesAPIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamLBACRulesAPIParamsWithTimeout creates a new GetTeamLBACRulesAPIParams object
// with the ability to set a timeout on a request.
func NewGetTeamLBACRulesAPIParamsWithTimeout(timeout time.Duration) *GetTeamLBACRulesAPIParams {
	return &GetTeamLBACRulesAPIParams{
		timeout: timeout,
	}
}

// NewGetTeamLBACRulesAPIParamsWithContext creates a new GetTeamLBACRulesAPIParams object
// with the ability to set a context for a request.
func NewGetTeamLBACRulesAPIParamsWithContext(ctx context.Context) *GetTeamLBACRulesAPIParams {
	return &GetTeamLBACRulesAPIParams{
		Context: ctx,
	}
}

// NewGetTeamLBACRulesAPIParamsWithHTTPClient creates a new GetTeamLBACRulesAPIParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTeamLBACRulesAPIParamsWithHTTPClient(client *http.Client) *GetTeamLBACRulesAPIParams {
	return &GetTeamLBACRulesAPIParams{
		HTTPClient: client,
	}
}

/*
GetTeamLBACRulesAPIParams contains all the parameters to send to the API endpoint

	for the get team l b a c rules Api operation.

	Typically these are written to a http.Request.
*/
type GetTeamLBACRulesAPIParams struct {

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get team l b a c rules Api params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamLBACRulesAPIParams) WithDefaults() *GetTeamLBACRulesAPIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get team l b a c rules Api params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamLBACRulesAPIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get team l b a c rules Api params
func (o *GetTeamLBACRulesAPIParams) WithTimeout(timeout time.Duration) *GetTeamLBACRulesAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team l b a c rules Api params
func (o *GetTeamLBACRulesAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team l b a c rules Api params
func (o *GetTeamLBACRulesAPIParams) WithContext(ctx context.Context) *GetTeamLBACRulesAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team l b a c rules Api params
func (o *GetTeamLBACRulesAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team l b a c rules Api params
func (o *GetTeamLBACRulesAPIParams) WithHTTPClient(client *http.Client) *GetTeamLBACRulesAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team l b a c rules Api params
func (o *GetTeamLBACRulesAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the get team l b a c rules Api params
func (o *GetTeamLBACRulesAPIParams) WithUID(uid string) *GetTeamLBACRulesAPIParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the get team l b a c rules Api params
func (o *GetTeamLBACRulesAPIParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamLBACRulesAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
