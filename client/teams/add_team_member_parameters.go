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

// NewAddTeamMemberParams creates a new AddTeamMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddTeamMemberParams() *AddTeamMemberParams {
	return &AddTeamMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddTeamMemberParamsWithTimeout creates a new AddTeamMemberParams object
// with the ability to set a timeout on a request.
func NewAddTeamMemberParamsWithTimeout(timeout time.Duration) *AddTeamMemberParams {
	return &AddTeamMemberParams{
		timeout: timeout,
	}
}

// NewAddTeamMemberParamsWithContext creates a new AddTeamMemberParams object
// with the ability to set a context for a request.
func NewAddTeamMemberParamsWithContext(ctx context.Context) *AddTeamMemberParams {
	return &AddTeamMemberParams{
		Context: ctx,
	}
}

// NewAddTeamMemberParamsWithHTTPClient creates a new AddTeamMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddTeamMemberParamsWithHTTPClient(client *http.Client) *AddTeamMemberParams {
	return &AddTeamMemberParams{
		HTTPClient: client,
	}
}

/*
AddTeamMemberParams contains all the parameters to send to the API endpoint

	for the add team member operation.

	Typically these are written to a http.Request.
*/
type AddTeamMemberParams struct {

	// Body.
	Body *models.AddTeamMemberCommand

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add team member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTeamMemberParams) WithDefaults() *AddTeamMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add team member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTeamMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add team member params
func (o *AddTeamMemberParams) WithTimeout(timeout time.Duration) *AddTeamMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add team member params
func (o *AddTeamMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add team member params
func (o *AddTeamMemberParams) WithContext(ctx context.Context) *AddTeamMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add team member params
func (o *AddTeamMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add team member params
func (o *AddTeamMemberParams) WithHTTPClient(client *http.Client) *AddTeamMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add team member params
func (o *AddTeamMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add team member params
func (o *AddTeamMemberParams) WithBody(body *models.AddTeamMemberCommand) *AddTeamMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add team member params
func (o *AddTeamMemberParams) SetBody(body *models.AddTeamMemberCommand) {
	o.Body = body
}

// WithTeamID adds the teamID to the add team member params
func (o *AddTeamMemberParams) WithTeamID(teamID string) *AddTeamMemberParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the add team member params
func (o *AddTeamMemberParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *AddTeamMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
