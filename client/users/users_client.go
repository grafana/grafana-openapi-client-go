// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command
// To edit this file, edit the custom template in templates/clientClient.gotmpl
// More info on custom templates can be found in the README or here: https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md
// The template itself can be found here: https://github.com/go-swagger/go-swagger/blob/master/generator/templates/client/client.gotmpl

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetUserByID(params *GetUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserByIDOK, error)

	GetUserByLoginOrEmail(params *GetUserByLoginOrEmailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserByLoginOrEmailOK, error)

	GetUserOrgList(params *GetUserOrgListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserOrgListOK, error)

	GetUserTeams(params *GetUserTeamsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserTeamsOK, error)

	SearchUsers(params *SearchUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchUsersOK, error)

	SearchUsersWithPaging(params *SearchUsersWithPagingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchUsersWithPagingOK, error)

	UpdateUser(params *UpdateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetUserByID gets user by id
*/
func (a *Client) GetUserByID(params *GetUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserByID",
		Method:             "GET",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetUserByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserByLoginOrEmail gets user by login or email
*/
func (a *Client) GetUserByLoginOrEmail(params *GetUserByLoginOrEmailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserByLoginOrEmailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserByLoginOrEmailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserByLoginOrEmail",
		Method:             "GET",
		PathPattern:        "/users/lookup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserByLoginOrEmailReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetUserByLoginOrEmailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserByLoginOrEmail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserOrgList gets organizations for user

Get organizations for user identified by id.
*/
func (a *Client) GetUserOrgList(params *GetUserOrgListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserOrgListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserOrgListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserOrgList",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/orgs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserOrgListReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetUserOrgListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserOrgList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserTeams gets teams for user

Get teams for user identified by id.
*/
func (a *Client) GetUserTeams(params *GetUserTeamsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserTeamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserTeamsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserTeams",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserTeamsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetUserTeamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserTeams: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchUsers gets users

Returns all users that the authenticated user has permission to view, admin permission required.
*/
func (a *Client) SearchUsers(params *SearchUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*SearchUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchUsersWithPaging gets users with paging
*/
func (a *Client) SearchUsersWithPaging(params *SearchUsersWithPagingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchUsersWithPagingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchUsersWithPagingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchUsersWithPaging",
		Method:             "GET",
		PathPattern:        "/users/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchUsersWithPagingReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*SearchUsersWithPagingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchUsersWithPaging: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUser updates user

Update the user identified by id.
*/
func (a *Client) UpdateUser(params *UpdateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUser",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}