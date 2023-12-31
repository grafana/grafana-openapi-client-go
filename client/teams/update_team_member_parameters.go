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
	"github.com/go-openapi/swag"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewUpdateTeamMemberParams creates a new UpdateTeamMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTeamMemberParams() *UpdateTeamMemberParams {
	return &UpdateTeamMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTeamMemberParamsWithTimeout creates a new UpdateTeamMemberParams object
// with the ability to set a timeout on a request.
func NewUpdateTeamMemberParamsWithTimeout(timeout time.Duration) *UpdateTeamMemberParams {
	return &UpdateTeamMemberParams{
		timeout: timeout,
	}
}

// NewUpdateTeamMemberParamsWithContext creates a new UpdateTeamMemberParams object
// with the ability to set a context for a request.
func NewUpdateTeamMemberParamsWithContext(ctx context.Context) *UpdateTeamMemberParams {
	return &UpdateTeamMemberParams{
		Context: ctx,
	}
}

// NewUpdateTeamMemberParamsWithHTTPClient creates a new UpdateTeamMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTeamMemberParamsWithHTTPClient(client *http.Client) *UpdateTeamMemberParams {
	return &UpdateTeamMemberParams{
		HTTPClient: client,
	}
}

/*
UpdateTeamMemberParams contains all the parameters to send to the API endpoint

	for the update team member operation.

	Typically these are written to a http.Request.
*/
type UpdateTeamMemberParams struct {

	// Body.
	Body *models.UpdateTeamMemberCommand

	// TeamID.
	TeamID string

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update team member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTeamMemberParams) WithDefaults() *UpdateTeamMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update team member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTeamMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update team member params
func (o *UpdateTeamMemberParams) WithTimeout(timeout time.Duration) *UpdateTeamMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update team member params
func (o *UpdateTeamMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update team member params
func (o *UpdateTeamMemberParams) WithContext(ctx context.Context) *UpdateTeamMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update team member params
func (o *UpdateTeamMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update team member params
func (o *UpdateTeamMemberParams) WithHTTPClient(client *http.Client) *UpdateTeamMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update team member params
func (o *UpdateTeamMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update team member params
func (o *UpdateTeamMemberParams) WithBody(body *models.UpdateTeamMemberCommand) *UpdateTeamMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update team member params
func (o *UpdateTeamMemberParams) SetBody(body *models.UpdateTeamMemberCommand) {
	o.Body = body
}

// WithTeamID adds the teamID to the update team member params
func (o *UpdateTeamMemberParams) WithTeamID(teamID string) *UpdateTeamMemberParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the update team member params
func (o *UpdateTeamMemberParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithUserID adds the userID to the update team member params
func (o *UpdateTeamMemberParams) WithUserID(userID int64) *UpdateTeamMemberParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update team member params
func (o *UpdateTeamMemberParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTeamMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
