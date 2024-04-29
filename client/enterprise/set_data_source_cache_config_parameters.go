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

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewSetDataSourceCacheConfigParams creates a new SetDataSourceCacheConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDataSourceCacheConfigParams() *SetDataSourceCacheConfigParams {
	return &SetDataSourceCacheConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDataSourceCacheConfigParamsWithTimeout creates a new SetDataSourceCacheConfigParams object
// with the ability to set a timeout on a request.
func NewSetDataSourceCacheConfigParamsWithTimeout(timeout time.Duration) *SetDataSourceCacheConfigParams {
	return &SetDataSourceCacheConfigParams{
		timeout: timeout,
	}
}

// NewSetDataSourceCacheConfigParamsWithContext creates a new SetDataSourceCacheConfigParams object
// with the ability to set a context for a request.
func NewSetDataSourceCacheConfigParamsWithContext(ctx context.Context) *SetDataSourceCacheConfigParams {
	return &SetDataSourceCacheConfigParams{
		Context: ctx,
	}
}

// NewSetDataSourceCacheConfigParamsWithHTTPClient creates a new SetDataSourceCacheConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDataSourceCacheConfigParamsWithHTTPClient(client *http.Client) *SetDataSourceCacheConfigParams {
	return &SetDataSourceCacheConfigParams{
		HTTPClient: client,
	}
}

/*
SetDataSourceCacheConfigParams contains all the parameters to send to the API endpoint

	for the set data source cache config operation.

	Typically these are written to a http.Request.
*/
type SetDataSourceCacheConfigParams struct {

	// Body.
	Body *models.CacheConfigSetter

	// DataSourceUID.
	DataSourceUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set data source cache config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDataSourceCacheConfigParams) WithDefaults() *SetDataSourceCacheConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set data source cache config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDataSourceCacheConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) WithTimeout(timeout time.Duration) *SetDataSourceCacheConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) WithContext(ctx context.Context) *SetDataSourceCacheConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) WithHTTPClient(client *http.Client) *SetDataSourceCacheConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) WithBody(body *models.CacheConfigSetter) *SetDataSourceCacheConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) SetBody(body *models.CacheConfigSetter) {
	o.Body = body
}

// WithDataSourceUID adds the dataSourceUID to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) WithDataSourceUID(dataSourceUID string) *SetDataSourceCacheConfigParams {
	o.SetDataSourceUID(dataSourceUID)
	return o
}

// SetDataSourceUID adds the dataSourceUid to the set data source cache config params
func (o *SetDataSourceCacheConfigParams) SetDataSourceUID(dataSourceUID string) {
	o.DataSourceUID = dataSourceUID
}

// WriteToRequest writes these params to a swagger request
func (o *SetDataSourceCacheConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dataSourceUID
	if err := r.SetPathParam("dataSourceUID", o.DataSourceUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
