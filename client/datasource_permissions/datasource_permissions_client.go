// Code generated by go-swagger; DO NOT EDIT.

package datasource_permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new datasource permissions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for datasource permissions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddPermission(params *AddPermissionParams, opts ...ClientOption) (*AddPermissionOK, error)

	DeletePermissions(params *DeletePermissionsParams, opts ...ClientOption) (*DeletePermissionsOK, error)

	DisablePermissions(params *DisablePermissionsParams, opts ...ClientOption) (*DisablePermissionsOK, error)

	EnablePermissions(params *EnablePermissionsParams, opts ...ClientOption) (*EnablePermissionsOK, error)

	GetAllPermissions(params *GetAllPermissionsParams, opts ...ClientOption) (*GetAllPermissionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AddPermission adds permissions for a data source

	You need to have a permission with action `datasources.permissions:read` and scopes `datasources:*`, `datasources:id:*`, `datasources:id:1` (single data source).

Deprecated: true.
Deprecated. Please use POST /api/access-control/datasources/:uid/users/:id, /api/access-control/datasources/:uid/teams/:id or /api/access-control/datasources/:uid/buildInRoles/:id
*/
func (a *Client) AddPermission(params *AddPermissionParams, opts ...ClientOption) (*AddPermissionOK, error) {
	if params == nil {
		params = NewAddPermissionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addPermission",
		Method:             "POST",
		PathPattern:        "/datasources/{datasourceId}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPermissionReader{formats: a.formats},
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
	success, ok := result.(*AddPermissionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addPermission: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DeletePermissions removes permission for a data source

	Removes the permission with the given permissionId for the data source with the given id.

You need to have a permission with action `datasources.permissions:delete` and scopes `datasources:*`, `datasources:id:*`, `datasources:id:1` (single data source).

Deprecated: true.
Deprecated. Please use POST /api/access-control/datasources/:uid/users/:id, /api/access-control/datasources/:uid/teams/:id or /api/access-control/datasources/:uid/buildInRoles/:id
*/
func (a *Client) DeletePermissions(params *DeletePermissionsParams, opts ...ClientOption) (*DeletePermissionsOK, error) {
	if params == nil {
		params = NewDeletePermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePermissions",
		Method:             "DELETE",
		PathPattern:        "/datasources/{datasourceId}/permissions/{permissionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePermissionsReader{formats: a.formats},
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
	success, ok := result.(*DeletePermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DisablePermissions disables permissions for a data source

	Disables permissions for the data source with the given id. All existing permissions will be removed and anyone will be able to query the data source.

You need to have a permission with action `datasources.permissions:toggle` and scopes `datasources:*`, `datasources:id:*`, `datasources:id:1` (single data source).

Deprecated: true.
*/
func (a *Client) DisablePermissions(params *DisablePermissionsParams, opts ...ClientOption) (*DisablePermissionsOK, error) {
	if params == nil {
		params = NewDisablePermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disablePermissions",
		Method:             "POST",
		PathPattern:        "/datasources/{datasourceId}/disable-permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisablePermissionsReader{formats: a.formats},
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
	success, ok := result.(*DisablePermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disablePermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	EnablePermissions enables permissions for a data source

	Enables permissions for the data source with the given id.

No one except Org Admins will be able to query the data source until permissions have been added
which permit certain users or teams to query the data source.

You need to have a permission with action `datasources.permissions:toggle` and scopes `datasources:*`, `datasources:id:*`, `datasources:id:1` (single data source).

Deprecated: true.
*/
func (a *Client) EnablePermissions(params *EnablePermissionsParams, opts ...ClientOption) (*EnablePermissionsOK, error) {
	if params == nil {
		params = NewEnablePermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enablePermissions",
		Method:             "POST",
		PathPattern:        "/datasources/{datasourceId}/enable-permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EnablePermissionsReader{formats: a.formats},
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
	success, ok := result.(*EnablePermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enablePermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetAllPermissions gets permissions for a data source

	Gets all existing permissions for the data source with the given id.

You need to have a permission with action `datasources.permissions:read` and scopes `datasources:*`, `datasources:id:*`, `datasources:id:1` (single data source).

Deprecated: true.
Deprecated. Please use GET /api/access-control/datasources/:uid
*/
func (a *Client) GetAllPermissions(params *GetAllPermissionsParams, opts ...ClientOption) (*GetAllPermissionsOK, error) {
	if params == nil {
		params = NewGetAllPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllPermissions",
		Method:             "GET",
		PathPattern:        "/datasources/{datasourceId}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllPermissionsReader{formats: a.formats},
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
	success, ok := result.(*GetAllPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllPermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
