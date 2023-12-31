// Code generated by go-swagger; DO NOT EDIT.

package correlations

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

// NewGetCorrelationsParams creates a new GetCorrelationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCorrelationsParams() *GetCorrelationsParams {
	return &GetCorrelationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorrelationsParamsWithTimeout creates a new GetCorrelationsParams object
// with the ability to set a timeout on a request.
func NewGetCorrelationsParamsWithTimeout(timeout time.Duration) *GetCorrelationsParams {
	return &GetCorrelationsParams{
		timeout: timeout,
	}
}

// NewGetCorrelationsParamsWithContext creates a new GetCorrelationsParams object
// with the ability to set a context for a request.
func NewGetCorrelationsParamsWithContext(ctx context.Context) *GetCorrelationsParams {
	return &GetCorrelationsParams{
		Context: ctx,
	}
}

// NewGetCorrelationsParamsWithHTTPClient creates a new GetCorrelationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCorrelationsParamsWithHTTPClient(client *http.Client) *GetCorrelationsParams {
	return &GetCorrelationsParams{
		HTTPClient: client,
	}
}

/*
GetCorrelationsParams contains all the parameters to send to the API endpoint

	for the get correlations operation.

	Typically these are written to a http.Request.
*/
type GetCorrelationsParams struct {

	/* Limit.

	   Limit the maximum number of correlations to return per page

	   Format: int64
	   Default: 100
	*/
	Limit *int64

	/* Page.

	   Page index for starting fetching correlations

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* SourceUID.

	   Source datasource UID filter to be applied to correlations
	*/
	SourceUID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get correlations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorrelationsParams) WithDefaults() *GetCorrelationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get correlations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorrelationsParams) SetDefaults() {
	var (
		limitDefault = int64(100)

		pageDefault = int64(1)
	)

	val := GetCorrelationsParams{
		Limit: &limitDefault,
		Page:  &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get correlations params
func (o *GetCorrelationsParams) WithTimeout(timeout time.Duration) *GetCorrelationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get correlations params
func (o *GetCorrelationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get correlations params
func (o *GetCorrelationsParams) WithContext(ctx context.Context) *GetCorrelationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get correlations params
func (o *GetCorrelationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get correlations params
func (o *GetCorrelationsParams) WithHTTPClient(client *http.Client) *GetCorrelationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get correlations params
func (o *GetCorrelationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get correlations params
func (o *GetCorrelationsParams) WithLimit(limit *int64) *GetCorrelationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get correlations params
func (o *GetCorrelationsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get correlations params
func (o *GetCorrelationsParams) WithPage(page *int64) *GetCorrelationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get correlations params
func (o *GetCorrelationsParams) SetPage(page *int64) {
	o.Page = page
}

// WithSourceUID adds the sourceUID to the get correlations params
func (o *GetCorrelationsParams) WithSourceUID(sourceUID []string) *GetCorrelationsParams {
	o.SetSourceUID(sourceUID)
	return o
}

// SetSourceUID adds the sourceUid to the get correlations params
func (o *GetCorrelationsParams) SetSourceUID(sourceUID []string) {
	o.SourceUID = sourceUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorrelationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.SourceUID != nil {

		// binding items for sourceUID
		joinedSourceUID := o.bindParamSourceUID(reg)

		// query array param sourceUID
		if err := r.SetQueryParam("sourceUID", joinedSourceUID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetCorrelations binds the parameter sourceUID
func (o *GetCorrelationsParams) bindParamSourceUID(formats strfmt.Registry) []string {
	sourceUIDIR := o.SourceUID

	var sourceUIDIC []string
	for _, sourceUIDIIR := range sourceUIDIR { // explode []string

		sourceUIDIIV := sourceUIDIIR // string as string
		sourceUIDIC = append(sourceUIDIC, sourceUIDIIV)
	}

	// items.CollectionFormat: "multi"
	sourceUIDIS := swag.JoinByFormat(sourceUIDIC, "multi")

	return sourceUIDIS
}
