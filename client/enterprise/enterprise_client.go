// Code generated by go-swagger; DO NOT EDIT.

package enterprise

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command
// To edit this file, edit the custom template in templates/clientClient.gotmpl
// More info on custom templates can be found in the README or here: https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md
// The template itself can be found here: https://github.com/go-swagger/go-swagger/blob/master/generator/templates/client/client.gotmpl

import (
	"fmt"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/utils"
)

// New creates a new enterprise API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, cfg *utils.ClientConfig) ClientService {
	return &Client{transport: transport, formats: formats, cfg: cfg}
}

/*
Client for enterprise API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	cfg       *utils.ClientConfig
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SearchResult(params *SearchResultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchResultOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	SearchResult debugs permissions

	Returns the result of the search through access-control role assignments.

You need to have a permission with action `teams.roles:read` on scope `teams:*`
and a permission with action `users.roles:read` on scope `users:*`.
*/
func (a *Client) SearchResult(params *SearchResultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchResultOK, error) {
	var (
		result interface{}
		err    error
	)
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchResult",
		Method:             "POST",
		PathPattern:        "/access-control/assignments/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchResultReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	timeout := utils.GetTimeout(a.cfg.RetryTimeout)
	for n := 0; n <= a.cfg.NumRetries; n++ {
		// Wait a bit if it is not the first request
		if n != 0 {
			time.Sleep(timeout)
		}

		result, err = a.transport.Submit(op)

		// If err is not nil, retry again
		// That's either caused by client policy, or failure to speak HTTP (such as network connectivity problem). A
		// non-2xx status code doesn't cause an error.
		if err != nil {
			continue
		}

		shouldRetry, err := utils.MatchRetryCode(err, a.cfg.RetryStatusCodes)
		if err != nil {
			return nil, err
		}
		if !shouldRetry {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
