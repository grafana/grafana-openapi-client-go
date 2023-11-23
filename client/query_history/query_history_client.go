// Code generated by go-swagger; DO NOT EDIT.

package query_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new query history API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for query history API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateQuery(params *CreateQueryParams, opts ...ClientOption) (*CreateQueryOK, error)

	DeleteQuery(params *DeleteQueryParams, opts ...ClientOption) (*DeleteQueryOK, error)

	PatchQueryComment(params *PatchQueryCommentParams, opts ...ClientOption) (*PatchQueryCommentOK, error)

	SearchQueries(params *SearchQueriesParams, opts ...ClientOption) (*SearchQueriesOK, error)

	StarQuery(params *StarQueryParams, opts ...ClientOption) (*StarQueryOK, error)

	UnstarQuery(params *UnstarQueryParams, opts ...ClientOption) (*UnstarQueryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateQuery adds query to query history

Adds new query to query history.
*/
func (a *Client) CreateQuery(params *CreateQueryParams, opts ...ClientOption) (*CreateQueryOK, error) {
	if params == nil {
		params = NewCreateQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createQuery",
		Method:             "POST",
		PathPattern:        "/query-history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateQueryReader{formats: a.formats},
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
	success, ok := result.(*CreateQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteQuery deletes query in query history

Deletes an existing query in query history as specified by the UID. This operation cannot be reverted.
*/
func (a *Client) DeleteQuery(params *DeleteQueryParams, opts ...ClientOption) (*DeleteQueryOK, error) {
	if params == nil {
		params = NewDeleteQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteQuery",
		Method:             "DELETE",
		PathPattern:        "/query-history/{query_history_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteQueryReader{formats: a.formats},
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
	success, ok := result.(*DeleteQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchQueryComment updates comment for query in query history

Updates comment for query in query history as specified by the UID.
*/
func (a *Client) PatchQueryComment(params *PatchQueryCommentParams, opts ...ClientOption) (*PatchQueryCommentOK, error) {
	if params == nil {
		params = NewPatchQueryCommentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchQueryComment",
		Method:             "PATCH",
		PathPattern:        "/query-history/{query_history_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchQueryCommentReader{formats: a.formats},
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
	success, ok := result.(*PatchQueryCommentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchQueryComment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SearchQueries queries history search

	Returns a list of queries in the query history that matches the search criteria.

Query history search supports pagination. Use the `limit` parameter to control the maximum number of queries returned; the default limit is 100.
You can also use the `page` query parameter to fetch queries from any page other than the first one.
*/
func (a *Client) SearchQueries(params *SearchQueriesParams, opts ...ClientOption) (*SearchQueriesOK, error) {
	if params == nil {
		params = NewSearchQueriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchQueries",
		Method:             "GET",
		PathPattern:        "/query-history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchQueriesReader{formats: a.formats},
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
	success, ok := result.(*SearchQueriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchQueries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StarQuery adds star to query in query history

Adds star to query in query history as specified by the UID.
*/
func (a *Client) StarQuery(params *StarQueryParams, opts ...ClientOption) (*StarQueryOK, error) {
	if params == nil {
		params = NewStarQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "starQuery",
		Method:             "POST",
		PathPattern:        "/query-history/star/{query_history_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StarQueryReader{formats: a.formats},
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
	success, ok := result.(*StarQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for starQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UnstarQuery removes star to query in query history

Removes star from query in query history as specified by the UID.
*/
func (a *Client) UnstarQuery(params *UnstarQueryParams, opts ...ClientOption) (*UnstarQueryOK, error) {
	if params == nil {
		params = NewUnstarQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unstarQuery",
		Method:             "DELETE",
		PathPattern:        "/query-history/star/{query_history_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UnstarQueryReader{formats: a.formats},
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
	success, ok := result.(*UnstarQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unstarQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
