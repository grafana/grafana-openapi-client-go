// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewSetTeamMembershipsParams creates a new SetTeamMembershipsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetTeamMembershipsParams() *SetTeamMembershipsParams {
	return &SetTeamMembershipsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetTeamMembershipsParamsWithTimeout creates a new SetTeamMembershipsParams object
// with the ability to set a timeout on a request.
func NewSetTeamMembershipsParamsWithTimeout(timeout time.Duration) *SetTeamMembershipsParams {
	return &SetTeamMembershipsParams{
		timeout: timeout,
	}
}

// NewSetTeamMembershipsParamsWithContext creates a new SetTeamMembershipsParams object
// with the ability to set a context for a request.
func NewSetTeamMembershipsParamsWithContext(ctx context.Context) *SetTeamMembershipsParams {
	return &SetTeamMembershipsParams{
		Context: ctx,
	}
}

// NewSetTeamMembershipsParamsWithHTTPClient creates a new SetTeamMembershipsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetTeamMembershipsParamsWithHTTPClient(client *http.Client) *SetTeamMembershipsParams {
	return &SetTeamMembershipsParams{
		HTTPClient: client,
	}
}

/*
SetTeamMembershipsParams contains all the parameters to send to the API endpoint

	for the set team memberships operation.

	Typically these are written to a http.Request.
*/
type SetTeamMembershipsParams struct {

	// Body.
	Body *models.SetTeamMembershipsCommand

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set team memberships params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTeamMembershipsParams) WithDefaults() *SetTeamMembershipsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set team memberships params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTeamMembershipsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set team memberships params
func (o *SetTeamMembershipsParams) WithTimeout(timeout time.Duration) *SetTeamMembershipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set team memberships params
func (o *SetTeamMembershipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set team memberships params
func (o *SetTeamMembershipsParams) WithContext(ctx context.Context) *SetTeamMembershipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set team memberships params
func (o *SetTeamMembershipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set team memberships params
func (o *SetTeamMembershipsParams) WithHTTPClient(client *http.Client) *SetTeamMembershipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set team memberships params
func (o *SetTeamMembershipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set team memberships params
func (o *SetTeamMembershipsParams) WithBody(body *models.SetTeamMembershipsCommand) *SetTeamMembershipsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set team memberships params
func (o *SetTeamMembershipsParams) SetBody(body *models.SetTeamMembershipsCommand) {
	o.Body = body
}

// WithTeamID adds the teamID to the set team memberships params
func (o *SetTeamMembershipsParams) WithTeamID(teamID string) *SetTeamMembershipsParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the set team memberships params
func (o *SetTeamMembershipsParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *SetTeamMembershipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
