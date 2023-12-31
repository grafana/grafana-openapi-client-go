// Code generated by go-swagger; DO NOT EDIT.

package dashboard_permissions

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

// NewUpdateDashboardPermissionsByUIDParams creates a new UpdateDashboardPermissionsByUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDashboardPermissionsByUIDParams() *UpdateDashboardPermissionsByUIDParams {
	return &UpdateDashboardPermissionsByUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDashboardPermissionsByUIDParamsWithTimeout creates a new UpdateDashboardPermissionsByUIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDashboardPermissionsByUIDParamsWithTimeout(timeout time.Duration) *UpdateDashboardPermissionsByUIDParams {
	return &UpdateDashboardPermissionsByUIDParams{
		timeout: timeout,
	}
}

// NewUpdateDashboardPermissionsByUIDParamsWithContext creates a new UpdateDashboardPermissionsByUIDParams object
// with the ability to set a context for a request.
func NewUpdateDashboardPermissionsByUIDParamsWithContext(ctx context.Context) *UpdateDashboardPermissionsByUIDParams {
	return &UpdateDashboardPermissionsByUIDParams{
		Context: ctx,
	}
}

// NewUpdateDashboardPermissionsByUIDParamsWithHTTPClient creates a new UpdateDashboardPermissionsByUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDashboardPermissionsByUIDParamsWithHTTPClient(client *http.Client) *UpdateDashboardPermissionsByUIDParams {
	return &UpdateDashboardPermissionsByUIDParams{
		HTTPClient: client,
	}
}

/*
UpdateDashboardPermissionsByUIDParams contains all the parameters to send to the API endpoint

	for the update dashboard permissions by UID operation.

	Typically these are written to a http.Request.
*/
type UpdateDashboardPermissionsByUIDParams struct {

	// Body.
	Body *models.UpdateDashboardACLCommand

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update dashboard permissions by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardPermissionsByUIDParams) WithDefaults() *UpdateDashboardPermissionsByUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update dashboard permissions by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardPermissionsByUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) WithTimeout(timeout time.Duration) *UpdateDashboardPermissionsByUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) WithContext(ctx context.Context) *UpdateDashboardPermissionsByUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) WithHTTPClient(client *http.Client) *UpdateDashboardPermissionsByUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) WithBody(body *models.UpdateDashboardACLCommand) *UpdateDashboardPermissionsByUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) SetBody(body *models.UpdateDashboardACLCommand) {
	o.Body = body
}

// WithUID adds the uid to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) WithUID(uid string) *UpdateDashboardPermissionsByUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the update dashboard permissions by UID params
func (o *UpdateDashboardPermissionsByUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDashboardPermissionsByUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
