// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new admin API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admin API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AdminGetSettings(opts ...ClientOption) (*AdminGetSettingsOK, error)
	AdminGetSettingsWithParams(params *AdminGetSettingsParams, opts ...ClientOption) (*AdminGetSettingsOK, error)

	AdminGetStats(opts ...ClientOption) (*AdminGetStatsOK, error)
	AdminGetStatsWithParams(params *AdminGetStatsParams, opts ...ClientOption) (*AdminGetStatsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AdminGetSettings fetches settings

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `settings:read` and scopes: `settings:*`, `settings:auth.saml:` and `settings:auth.saml:enabled` (property level).
*/
func (a *Client) AdminGetSettings(opts ...ClientOption) (*AdminGetSettingsOK, error) {
	params := NewAdminGetSettingsParams()
	return a.AdminGetSettingsWithParams(params, opts...)
}

func (a *Client) AdminGetSettingsWithParams(params *AdminGetSettingsParams, opts ...ClientOption) (*AdminGetSettingsOK, error) {
	if params == nil {
		params = NewAdminGetSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminGetSettings",
		Method:             "GET",
		PathPattern:        "/admin/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminGetSettingsReader{formats: a.formats},
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
	success, ok := result.(*AdminGetSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminGetSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AdminGetStats fetches grafana stats

Only works with Basic Authentication (username and password). See introduction for an explanation.
If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `server:stats:read`.
*/
func (a *Client) AdminGetStats(opts ...ClientOption) (*AdminGetStatsOK, error) {
	params := NewAdminGetStatsParams()
	return a.AdminGetStatsWithParams(params, opts...)
}

func (a *Client) AdminGetStatsWithParams(params *AdminGetStatsParams, opts ...ClientOption) (*AdminGetStatsOK, error) {
	if params == nil {
		params = NewAdminGetStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminGetStats",
		Method:             "GET",
		PathPattern:        "/admin/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminGetStatsReader{formats: a.formats},
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
	success, ok := result.(*AdminGetStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminGetStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
