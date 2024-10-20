// Code generated by go-swagger; DO NOT EDIT.

package group_attribute_sync

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

// NewCreateGroupMappingParams creates a new CreateGroupMappingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGroupMappingParams() *CreateGroupMappingParams {
	return &CreateGroupMappingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGroupMappingParamsWithTimeout creates a new CreateGroupMappingParams object
// with the ability to set a timeout on a request.
func NewCreateGroupMappingParamsWithTimeout(timeout time.Duration) *CreateGroupMappingParams {
	return &CreateGroupMappingParams{
		timeout: timeout,
	}
}

// NewCreateGroupMappingParamsWithContext creates a new CreateGroupMappingParams object
// with the ability to set a context for a request.
func NewCreateGroupMappingParamsWithContext(ctx context.Context) *CreateGroupMappingParams {
	return &CreateGroupMappingParams{
		Context: ctx,
	}
}

// NewCreateGroupMappingParamsWithHTTPClient creates a new CreateGroupMappingParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGroupMappingParamsWithHTTPClient(client *http.Client) *CreateGroupMappingParams {
	return &CreateGroupMappingParams{
		HTTPClient: client,
	}
}

/*
CreateGroupMappingParams contains all the parameters to send to the API endpoint

	for the create group mapping operation.

	Typically these are written to a http.Request.
*/
type CreateGroupMappingParams struct {

	// Body.
	Body *models.GroupMappingRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create group mapping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGroupMappingParams) WithDefaults() *CreateGroupMappingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create group mapping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGroupMappingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create group mapping params
func (o *CreateGroupMappingParams) WithTimeout(timeout time.Duration) *CreateGroupMappingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create group mapping params
func (o *CreateGroupMappingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create group mapping params
func (o *CreateGroupMappingParams) WithContext(ctx context.Context) *CreateGroupMappingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create group mapping params
func (o *CreateGroupMappingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create group mapping params
func (o *CreateGroupMappingParams) WithHTTPClient(client *http.Client) *CreateGroupMappingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create group mapping params
func (o *CreateGroupMappingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create group mapping params
func (o *CreateGroupMappingParams) WithBody(body *models.GroupMappingRequestBody) *CreateGroupMappingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create group mapping params
func (o *CreateGroupMappingParams) SetBody(body *models.GroupMappingRequestBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGroupMappingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
