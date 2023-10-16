// Code generated by go-swagger; DO NOT EDIT.

package service_accounts

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

// NewRetrieveServiceAccountParams creates a new RetrieveServiceAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveServiceAccountParams() *RetrieveServiceAccountParams {
	return &RetrieveServiceAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveServiceAccountParamsWithTimeout creates a new RetrieveServiceAccountParams object
// with the ability to set a timeout on a request.
func NewRetrieveServiceAccountParamsWithTimeout(timeout time.Duration) *RetrieveServiceAccountParams {
	return &RetrieveServiceAccountParams{
		timeout: timeout,
	}
}

// NewRetrieveServiceAccountParamsWithContext creates a new RetrieveServiceAccountParams object
// with the ability to set a context for a request.
func NewRetrieveServiceAccountParamsWithContext(ctx context.Context) *RetrieveServiceAccountParams {
	return &RetrieveServiceAccountParams{
		Context: ctx,
	}
}

// NewRetrieveServiceAccountParamsWithHTTPClient creates a new RetrieveServiceAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveServiceAccountParamsWithHTTPClient(client *http.Client) *RetrieveServiceAccountParams {
	return &RetrieveServiceAccountParams{
		HTTPClient: client,
	}
}

/*
RetrieveServiceAccountParams contains all the parameters to send to the API endpoint

	for the retrieve service account operation.

	Typically these are written to a http.Request.
*/
type RetrieveServiceAccountParams struct {

	// ServiceAccountID.
	//
	// Format: int64
	ServiceAccountID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveServiceAccountParams) WithDefaults() *RetrieveServiceAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveServiceAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve service account params
func (o *RetrieveServiceAccountParams) WithTimeout(timeout time.Duration) *RetrieveServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve service account params
func (o *RetrieveServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve service account params
func (o *RetrieveServiceAccountParams) WithContext(ctx context.Context) *RetrieveServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve service account params
func (o *RetrieveServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve service account params
func (o *RetrieveServiceAccountParams) WithHTTPClient(client *http.Client) *RetrieveServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve service account params
func (o *RetrieveServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceAccountID adds the serviceAccountID to the retrieve service account params
func (o *RetrieveServiceAccountParams) WithServiceAccountID(serviceAccountID int64) *RetrieveServiceAccountParams {
	o.SetServiceAccountID(serviceAccountID)
	return o
}

// SetServiceAccountID adds the serviceAccountId to the retrieve service account params
func (o *RetrieveServiceAccountParams) SetServiceAccountID(serviceAccountID int64) {
	o.ServiceAccountID = serviceAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceAccountId
	if err := r.SetPathParam("serviceAccountId", swag.FormatInt64(o.ServiceAccountID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}