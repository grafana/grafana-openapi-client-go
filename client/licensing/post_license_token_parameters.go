// Code generated by go-swagger; DO NOT EDIT.

package licensing

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

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewPostLicenseTokenParams creates a new PostLicenseTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLicenseTokenParams() *PostLicenseTokenParams {
	return &PostLicenseTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLicenseTokenParamsWithTimeout creates a new PostLicenseTokenParams object
// with the ability to set a timeout on a request.
func NewPostLicenseTokenParamsWithTimeout(timeout time.Duration) *PostLicenseTokenParams {
	return &PostLicenseTokenParams{
		timeout: timeout,
	}
}

// NewPostLicenseTokenParamsWithContext creates a new PostLicenseTokenParams object
// with the ability to set a context for a request.
func NewPostLicenseTokenParamsWithContext(ctx context.Context) *PostLicenseTokenParams {
	return &PostLicenseTokenParams{
		Context: ctx,
	}
}

// NewPostLicenseTokenParamsWithHTTPClient creates a new PostLicenseTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLicenseTokenParamsWithHTTPClient(client *http.Client) *PostLicenseTokenParams {
	return &PostLicenseTokenParams{
		HTTPClient: client,
	}
}

/*
PostLicenseTokenParams contains all the parameters to send to the API endpoint

	for the post license token operation.

	Typically these are written to a http.Request.
*/
type PostLicenseTokenParams struct {

	// Body.
	Body *models.DeleteTokenCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post license token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLicenseTokenParams) WithDefaults() *PostLicenseTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post license token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLicenseTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post license token params
func (o *PostLicenseTokenParams) WithTimeout(timeout time.Duration) *PostLicenseTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post license token params
func (o *PostLicenseTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post license token params
func (o *PostLicenseTokenParams) WithContext(ctx context.Context) *PostLicenseTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post license token params
func (o *PostLicenseTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post license token params
func (o *PostLicenseTokenParams) WithHTTPClient(client *http.Client) *PostLicenseTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post license token params
func (o *PostLicenseTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post license token params
func (o *PostLicenseTokenParams) WithBody(body *models.DeleteTokenCommand) *PostLicenseTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post license token params
func (o *PostLicenseTokenParams) SetBody(body *models.DeleteTokenCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostLicenseTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
