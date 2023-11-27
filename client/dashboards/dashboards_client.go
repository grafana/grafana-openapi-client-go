// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new dashboards API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboards API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CalculateDashboardDiff(body *models.CalculateDashboardDiffParamsBody, opts ...ClientOption) (*CalculateDashboardDiffOK, error)
	CalculateDashboardDiffWithParams(params *CalculateDashboardDiffParams, opts ...ClientOption) (*CalculateDashboardDiffOK, error)

	DeleteDashboardByUID(uid string, opts ...ClientOption) (*DeleteDashboardByUIDOK, error)
	DeleteDashboardByUIDWithParams(params *DeleteDashboardByUIDParams, opts ...ClientOption) (*DeleteDashboardByUIDOK, error)

	GetDashboardByUID(uid string, opts ...ClientOption) (*GetDashboardByUIDOK, error)
	GetDashboardByUIDWithParams(params *GetDashboardByUIDParams, opts ...ClientOption) (*GetDashboardByUIDOK, error)

	GetDashboardTags(opts ...ClientOption) (*GetDashboardTagsOK, error)
	GetDashboardTagsWithParams(params *GetDashboardTagsParams, opts ...ClientOption) (*GetDashboardTagsOK, error)

	GetHomeDashboard(opts ...ClientOption) (*GetHomeDashboardOK, error)
	GetHomeDashboardWithParams(params *GetHomeDashboardParams, opts ...ClientOption) (*GetHomeDashboardOK, error)

	ImportDashboard(body *models.ImportDashboardRequest, opts ...ClientOption) (*ImportDashboardOK, error)
	ImportDashboardWithParams(params *ImportDashboardParams, opts ...ClientOption) (*ImportDashboardOK, error)

	PostDashboard(body *models.SaveDashboardCommand, opts ...ClientOption) (*PostDashboardOK, error)
	PostDashboardWithParams(params *PostDashboardParams, opts ...ClientOption) (*PostDashboardOK, error)

	TrimDashboard(body *models.TrimDashboardCommand, opts ...ClientOption) (*TrimDashboardOK, error)
	TrimDashboardWithParams(params *TrimDashboardParams, opts ...ClientOption) (*TrimDashboardOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CalculateDashboardDiff performs diff on two dashboards
*/
func (a *Client) CalculateDashboardDiff(body *models.CalculateDashboardDiffParamsBody, opts ...ClientOption) (*CalculateDashboardDiffOK, error) {
	params := NewCalculateDashboardDiffParams().WithBody(body)
	return a.CalculateDashboardDiffWithParams(params, opts...)
}

func (a *Client) CalculateDashboardDiffWithParams(params *CalculateDashboardDiffParams, opts ...ClientOption) (*CalculateDashboardDiffOK, error) {
	if params == nil {
		params = NewCalculateDashboardDiffParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "calculateDashboardDiff",
		Method:             "POST",
		PathPattern:        "/dashboards/calculate-diff",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CalculateDashboardDiffReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CalculateDashboardDiffOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for calculateDashboardDiff: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDashboardByUID deletes dashboard by uid

Will delete the dashboard given the specified unique identifier (uid).
*/
func (a *Client) DeleteDashboardByUID(uid string, opts ...ClientOption) (*DeleteDashboardByUIDOK, error) {
	params := NewDeleteDashboardByUIDParams().WithUID(uid)
	return a.DeleteDashboardByUIDWithParams(params, opts...)
}

func (a *Client) DeleteDashboardByUIDWithParams(params *DeleteDashboardByUIDParams, opts ...ClientOption) (*DeleteDashboardByUIDOK, error) {
	if params == nil {
		params = NewDeleteDashboardByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDashboardByUID",
		Method:             "DELETE",
		PathPattern:        "/dashboards/uid/{uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDashboardByUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDashboardByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDashboardByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardByUID gets dashboard by uid

Will return the dashboard given the dashboard unique identifier (uid).
*/
func (a *Client) GetDashboardByUID(uid string, opts ...ClientOption) (*GetDashboardByUIDOK, error) {
	params := NewGetDashboardByUIDParams().WithUID(uid)
	return a.GetDashboardByUIDWithParams(params, opts...)
}

func (a *Client) GetDashboardByUIDWithParams(params *GetDashboardByUIDParams, opts ...ClientOption) (*GetDashboardByUIDOK, error) {
	if params == nil {
		params = NewGetDashboardByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardByUID",
		Method:             "GET",
		PathPattern:        "/dashboards/uid/{uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardByUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardTags gets all dashboards tags of an organisation
*/
func (a *Client) GetDashboardTags(opts ...ClientOption) (*GetDashboardTagsOK, error) {
	params := NewGetDashboardTagsParams()
	return a.GetDashboardTagsWithParams(params, opts...)
}

func (a *Client) GetDashboardTagsWithParams(params *GetDashboardTagsParams, opts ...ClientOption) (*GetDashboardTagsOK, error) {
	if params == nil {
		params = NewGetDashboardTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardTags",
		Method:             "GET",
		PathPattern:        "/dashboards/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardTagsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHomeDashboard gets home dashboard
*/
func (a *Client) GetHomeDashboard(opts ...ClientOption) (*GetHomeDashboardOK, error) {
	params := NewGetHomeDashboardParams()
	return a.GetHomeDashboardWithParams(params, opts...)
}

func (a *Client) GetHomeDashboardWithParams(params *GetHomeDashboardParams, opts ...ClientOption) (*GetHomeDashboardOK, error) {
	if params == nil {
		params = NewGetHomeDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHomeDashboard",
		Method:             "GET",
		PathPattern:        "/dashboards/home",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHomeDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHomeDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHomeDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImportDashboard imports dashboard
*/
func (a *Client) ImportDashboard(body *models.ImportDashboardRequest, opts ...ClientOption) (*ImportDashboardOK, error) {
	params := NewImportDashboardParams().WithBody(body)
	return a.ImportDashboardWithParams(params, opts...)
}

func (a *Client) ImportDashboardWithParams(params *ImportDashboardParams, opts ...ClientOption) (*ImportDashboardOK, error) {
	if params == nil {
		params = NewImportDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "importDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImportDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostDashboard creates update dashboard

Creates a new dashboard or updates an existing dashboard.
*/
func (a *Client) PostDashboard(body *models.SaveDashboardCommand, opts ...ClientOption) (*PostDashboardOK, error) {
	params := NewPostDashboardParams().WithBody(body)
	return a.PostDashboardWithParams(params, opts...)
}

func (a *Client) PostDashboardWithParams(params *PostDashboardParams, opts ...ClientOption) (*PostDashboardOK, error) {
	if params == nil {
		params = NewPostDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/db",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TrimDashboard trims defaults from dashboard
*/
func (a *Client) TrimDashboard(body *models.TrimDashboardCommand, opts ...ClientOption) (*TrimDashboardOK, error) {
	params := NewTrimDashboardParams().WithBody(body)
	return a.TrimDashboardWithParams(params, opts...)
}

func (a *Client) TrimDashboardWithParams(params *TrimDashboardParams, opts ...ClientOption) (*TrimDashboardOK, error) {
	if params == nil {
		params = NewTrimDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "trimDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/trim",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TrimDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TrimDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for trimDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// WithAuthInfo changes the transport on the client
func WithAuthInfo(authInfo runtime.ClientAuthInfoWriter) ClientOption {
	return func(op *runtime.ClientOperation) {
		op.AuthInfo = authInfo
	}
}
