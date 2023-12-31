// Code generated by go-swagger; DO NOT EDIT.

package ds

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

// NewQueryMetricsWithExpressionsParams creates a new QueryMetricsWithExpressionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryMetricsWithExpressionsParams() *QueryMetricsWithExpressionsParams {
	return &QueryMetricsWithExpressionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryMetricsWithExpressionsParamsWithTimeout creates a new QueryMetricsWithExpressionsParams object
// with the ability to set a timeout on a request.
func NewQueryMetricsWithExpressionsParamsWithTimeout(timeout time.Duration) *QueryMetricsWithExpressionsParams {
	return &QueryMetricsWithExpressionsParams{
		timeout: timeout,
	}
}

// NewQueryMetricsWithExpressionsParamsWithContext creates a new QueryMetricsWithExpressionsParams object
// with the ability to set a context for a request.
func NewQueryMetricsWithExpressionsParamsWithContext(ctx context.Context) *QueryMetricsWithExpressionsParams {
	return &QueryMetricsWithExpressionsParams{
		Context: ctx,
	}
}

// NewQueryMetricsWithExpressionsParamsWithHTTPClient creates a new QueryMetricsWithExpressionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryMetricsWithExpressionsParamsWithHTTPClient(client *http.Client) *QueryMetricsWithExpressionsParams {
	return &QueryMetricsWithExpressionsParams{
		HTTPClient: client,
	}
}

/*
QueryMetricsWithExpressionsParams contains all the parameters to send to the API endpoint

	for the query metrics with expressions operation.

	Typically these are written to a http.Request.
*/
type QueryMetricsWithExpressionsParams struct {

	// Body.
	Body *models.MetricRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query metrics with expressions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryMetricsWithExpressionsParams) WithDefaults() *QueryMetricsWithExpressionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query metrics with expressions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryMetricsWithExpressionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query metrics with expressions params
func (o *QueryMetricsWithExpressionsParams) WithTimeout(timeout time.Duration) *QueryMetricsWithExpressionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query metrics with expressions params
func (o *QueryMetricsWithExpressionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query metrics with expressions params
func (o *QueryMetricsWithExpressionsParams) WithContext(ctx context.Context) *QueryMetricsWithExpressionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query metrics with expressions params
func (o *QueryMetricsWithExpressionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query metrics with expressions params
func (o *QueryMetricsWithExpressionsParams) WithHTTPClient(client *http.Client) *QueryMetricsWithExpressionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query metrics with expressions params
func (o *QueryMetricsWithExpressionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the query metrics with expressions params
func (o *QueryMetricsWithExpressionsParams) WithBody(body *models.MetricRequest) *QueryMetricsWithExpressionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the query metrics with expressions params
func (o *QueryMetricsWithExpressionsParams) SetBody(body *models.MetricRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *QueryMetricsWithExpressionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
