// Code generated by go-swagger; DO NOT EDIT.

package licensing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new licensing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for licensing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteLicenseToken(params *DeleteLicenseTokenParams, opts ...ClientOption) (*DeleteLicenseTokenAccepted, error)

	GetCustomPermissionsCSV(params *GetCustomPermissionsCSVParams, opts ...ClientOption) (*GetCustomPermissionsCSVOK, error)

	GetCustomPermissionsReport(params *GetCustomPermissionsReportParams, opts ...ClientOption) (*GetCustomPermissionsReportOK, error)

	GetLicenseToken(params *GetLicenseTokenParams, opts ...ClientOption) (*GetLicenseTokenOK, error)

	GetStatus(params *GetStatusParams, opts ...ClientOption) (*GetStatusOK, error)

	PostLicenseToken(params *PostLicenseTokenParams, opts ...ClientOption) (*PostLicenseTokenOK, error)

	PostRenewLicenseToken(params *PostRenewLicenseTokenParams, opts ...ClientOption) (*PostRenewLicenseTokenOK, error)

	RefreshLicenseStats(params *RefreshLicenseStatsParams, opts ...ClientOption) (*RefreshLicenseStatsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	DeleteLicenseToken removes license from database

	Removes the license stored in the Grafana database. Available in Grafana Enterprise v7.4+.

You need to have a permission with action `licensing:delete`.
*/
func (a *Client) DeleteLicenseToken(params *DeleteLicenseTokenParams, opts ...ClientOption) (*DeleteLicenseTokenAccepted, error) {
	if params == nil {
		params = NewDeleteLicenseTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteLicenseToken",
		Method:             "DELETE",
		PathPattern:        "/licensing/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteLicenseTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLicenseTokenAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteLicenseToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCustomPermissionsCSV gets custom permissions report in CSV format

You need to have a permission with action `licensing.reports:read`.
*/
func (a *Client) GetCustomPermissionsCSV(params *GetCustomPermissionsCSVParams, opts ...ClientOption) (*GetCustomPermissionsCSVOK, error) {
	if params == nil {
		params = NewGetCustomPermissionsCSVParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCustomPermissionsCSV",
		Method:             "GET",
		PathPattern:        "/licensing/custom-permissions-csv",
		ProducesMediaTypes: []string{"text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCustomPermissionsCSVReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomPermissionsCSVOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCustomPermissionsCSV: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCustomPermissionsReport gets custom permissions report

You need to have a permission with action `licensing.reports:read`.
*/
func (a *Client) GetCustomPermissionsReport(params *GetCustomPermissionsReportParams, opts ...ClientOption) (*GetCustomPermissionsReportOK, error) {
	if params == nil {
		params = NewGetCustomPermissionsReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCustomPermissionsReport",
		Method:             "GET",
		PathPattern:        "/licensing/custom-permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCustomPermissionsReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomPermissionsReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCustomPermissionsReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLicenseToken gets license token

You need to have a permission with action `licensing:read`.
*/
func (a *Client) GetLicenseToken(params *GetLicenseTokenParams, opts ...ClientOption) (*GetLicenseTokenOK, error) {
	if params == nil {
		params = NewGetLicenseTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLicenseToken",
		Method:             "GET",
		PathPattern:        "/licensing/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLicenseTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLicenseTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLicenseToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetStatus checks license availability
*/
func (a *Client) GetStatus(params *GetStatusParams, opts ...ClientOption) (*GetStatusOK, error) {
	if params == nil {
		params = NewGetStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStatus",
		Method:             "GET",
		PathPattern:        "/licensing/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostLicenseToken creates license token

You need to have a permission with action `licensing:update`.
*/
func (a *Client) PostLicenseToken(params *PostLicenseTokenParams, opts ...ClientOption) (*PostLicenseTokenOK, error) {
	if params == nil {
		params = NewPostLicenseTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postLicenseToken",
		Method:             "POST",
		PathPattern:        "/licensing/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostLicenseTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLicenseTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postLicenseToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostRenewLicenseToken manuallies force license refresh

	Manually ask license issuer for a new token. Available in Grafana Enterprise v7.4+.

You need to have a permission with action `licensing:update`.
*/
func (a *Client) PostRenewLicenseToken(params *PostRenewLicenseTokenParams, opts ...ClientOption) (*PostRenewLicenseTokenOK, error) {
	if params == nil {
		params = NewPostRenewLicenseTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postRenewLicenseToken",
		Method:             "POST",
		PathPattern:        "/licensing/token/renew",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRenewLicenseTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostRenewLicenseTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postRenewLicenseToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RefreshLicenseStats refreshes license stats

You need to have a permission with action `licensing:read`.
*/
func (a *Client) RefreshLicenseStats(params *RefreshLicenseStatsParams, opts ...ClientOption) (*RefreshLicenseStatsOK, error) {
	if params == nil {
		params = NewRefreshLicenseStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "refreshLicenseStats",
		Method:             "GET",
		PathPattern:        "/licensing/refresh-stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RefreshLicenseStatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RefreshLicenseStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for refreshLicenseStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
