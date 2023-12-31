// Code generated by go-swagger; DO NOT EDIT.

package sync_team_groups

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

// NewRemoveTeamGroupAPIQueryParams creates a new RemoveTeamGroupAPIQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveTeamGroupAPIQueryParams() *RemoveTeamGroupAPIQueryParams {
	return &RemoveTeamGroupAPIQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTeamGroupAPIQueryParamsWithTimeout creates a new RemoveTeamGroupAPIQueryParams object
// with the ability to set a timeout on a request.
func NewRemoveTeamGroupAPIQueryParamsWithTimeout(timeout time.Duration) *RemoveTeamGroupAPIQueryParams {
	return &RemoveTeamGroupAPIQueryParams{
		timeout: timeout,
	}
}

// NewRemoveTeamGroupAPIQueryParamsWithContext creates a new RemoveTeamGroupAPIQueryParams object
// with the ability to set a context for a request.
func NewRemoveTeamGroupAPIQueryParamsWithContext(ctx context.Context) *RemoveTeamGroupAPIQueryParams {
	return &RemoveTeamGroupAPIQueryParams{
		Context: ctx,
	}
}

// NewRemoveTeamGroupAPIQueryParamsWithHTTPClient creates a new RemoveTeamGroupAPIQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveTeamGroupAPIQueryParamsWithHTTPClient(client *http.Client) *RemoveTeamGroupAPIQueryParams {
	return &RemoveTeamGroupAPIQueryParams{
		HTTPClient: client,
	}
}

/*
RemoveTeamGroupAPIQueryParams contains all the parameters to send to the API endpoint

	for the remove team group Api query operation.

	Typically these are written to a http.Request.
*/
type RemoveTeamGroupAPIQueryParams struct {

	// GroupID.
	GroupID *string

	// TeamID.
	//
	// Format: int64
	TeamID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove team group Api query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTeamGroupAPIQueryParams) WithDefaults() *RemoveTeamGroupAPIQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove team group Api query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTeamGroupAPIQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) WithTimeout(timeout time.Duration) *RemoveTeamGroupAPIQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) WithContext(ctx context.Context) *RemoveTeamGroupAPIQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) WithHTTPClient(client *http.Client) *RemoveTeamGroupAPIQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) WithGroupID(groupID *string) *RemoveTeamGroupAPIQueryParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithTeamID adds the teamID to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) WithTeamID(teamID int64) *RemoveTeamGroupAPIQueryParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the remove team group Api query params
func (o *RemoveTeamGroupAPIQueryParams) SetTeamID(teamID int64) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTeamGroupAPIQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GroupID != nil {

		// query param groupId
		var qrGroupID string

		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {

			if err := r.SetQueryParam("groupId", qGroupID); err != nil {
				return err
			}
		}
	}

	// path param teamId
	if err := r.SetPathParam("teamId", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
