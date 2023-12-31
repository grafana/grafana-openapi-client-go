// Code generated by go-swagger; DO NOT EDIT.

package library_elements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new library elements API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for library elements API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLibraryElement(body *models.CreateLibraryElementCommand, opts ...ClientOption) (*CreateLibraryElementOK, error)
	CreateLibraryElementWithParams(params *CreateLibraryElementParams, opts ...ClientOption) (*CreateLibraryElementOK, error)

	DeleteLibraryElementByUID(libraryElementUID string, opts ...ClientOption) (*DeleteLibraryElementByUIDOK, error)
	DeleteLibraryElementByUIDWithParams(params *DeleteLibraryElementByUIDParams, opts ...ClientOption) (*DeleteLibraryElementByUIDOK, error)

	GetLibraryElementByName(libraryElementName string, opts ...ClientOption) (*GetLibraryElementByNameOK, error)
	GetLibraryElementByNameWithParams(params *GetLibraryElementByNameParams, opts ...ClientOption) (*GetLibraryElementByNameOK, error)

	GetLibraryElementByUID(libraryElementUID string, opts ...ClientOption) (*GetLibraryElementByUIDOK, error)
	GetLibraryElementByUIDWithParams(params *GetLibraryElementByUIDParams, opts ...ClientOption) (*GetLibraryElementByUIDOK, error)

	GetLibraryElementConnections(libraryElementUID string, opts ...ClientOption) (*GetLibraryElementConnectionsOK, error)
	GetLibraryElementConnectionsWithParams(params *GetLibraryElementConnectionsParams, opts ...ClientOption) (*GetLibraryElementConnectionsOK, error)

	GetLibraryElements(params *GetLibraryElementsParams, opts ...ClientOption) (*GetLibraryElementsOK, error)

	UpdateLibraryElement(libraryElementUID string, body *models.PatchLibraryElementCommand, opts ...ClientOption) (*UpdateLibraryElementOK, error)
	UpdateLibraryElementWithParams(params *UpdateLibraryElementParams, opts ...ClientOption) (*UpdateLibraryElementOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateLibraryElement creates library element

Creates a new library element.
*/
func (a *Client) CreateLibraryElement(body *models.CreateLibraryElementCommand, opts ...ClientOption) (*CreateLibraryElementOK, error) {
	params := NewCreateLibraryElementParams().WithBody(body)
	return a.CreateLibraryElementWithParams(params, opts...)
}

func (a *Client) CreateLibraryElementWithParams(params *CreateLibraryElementParams, opts ...ClientOption) (*CreateLibraryElementOK, error) {
	if params == nil {
		params = NewCreateLibraryElementParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createLibraryElement",
		Method:             "POST",
		PathPattern:        "/library-elements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateLibraryElementReader{formats: a.formats},
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
	success, ok := result.(*CreateLibraryElementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createLibraryElement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteLibraryElementByUID deletes library element

Deletes an existing library element as specified by the UID. This operation cannot be reverted.
You cannot delete a library element that is connected. This operation cannot be reverted.
*/
func (a *Client) DeleteLibraryElementByUID(libraryElementUID string, opts ...ClientOption) (*DeleteLibraryElementByUIDOK, error) {
	params := NewDeleteLibraryElementByUIDParams().WithLibraryElementUID(libraryElementUID)
	return a.DeleteLibraryElementByUIDWithParams(params, opts...)
}

func (a *Client) DeleteLibraryElementByUIDWithParams(params *DeleteLibraryElementByUIDParams, opts ...ClientOption) (*DeleteLibraryElementByUIDOK, error) {
	if params == nil {
		params = NewDeleteLibraryElementByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteLibraryElementByUID",
		Method:             "DELETE",
		PathPattern:        "/library-elements/{library_element_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteLibraryElementByUIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteLibraryElementByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteLibraryElementByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLibraryElementByName gets library element by name

Returns a library element with the given name.
*/
func (a *Client) GetLibraryElementByName(libraryElementName string, opts ...ClientOption) (*GetLibraryElementByNameOK, error) {
	params := NewGetLibraryElementByNameParams().WithLibraryElementName(libraryElementName)
	return a.GetLibraryElementByNameWithParams(params, opts...)
}

func (a *Client) GetLibraryElementByNameWithParams(params *GetLibraryElementByNameParams, opts ...ClientOption) (*GetLibraryElementByNameOK, error) {
	if params == nil {
		params = NewGetLibraryElementByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLibraryElementByName",
		Method:             "GET",
		PathPattern:        "/library-elements/name/{library_element_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLibraryElementByNameReader{formats: a.formats},
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
	success, ok := result.(*GetLibraryElementByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLibraryElementByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLibraryElementByUID gets library element by UID

Returns a library element with the given UID.
*/
func (a *Client) GetLibraryElementByUID(libraryElementUID string, opts ...ClientOption) (*GetLibraryElementByUIDOK, error) {
	params := NewGetLibraryElementByUIDParams().WithLibraryElementUID(libraryElementUID)
	return a.GetLibraryElementByUIDWithParams(params, opts...)
}

func (a *Client) GetLibraryElementByUIDWithParams(params *GetLibraryElementByUIDParams, opts ...ClientOption) (*GetLibraryElementByUIDOK, error) {
	if params == nil {
		params = NewGetLibraryElementByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLibraryElementByUID",
		Method:             "GET",
		PathPattern:        "/library-elements/{library_element_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLibraryElementByUIDReader{formats: a.formats},
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
	success, ok := result.(*GetLibraryElementByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLibraryElementByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLibraryElementConnections gets library element connections

Returns a list of connections for a library element based on the UID specified.
*/
func (a *Client) GetLibraryElementConnections(libraryElementUID string, opts ...ClientOption) (*GetLibraryElementConnectionsOK, error) {
	params := NewGetLibraryElementConnectionsParams().WithLibraryElementUID(libraryElementUID)
	return a.GetLibraryElementConnectionsWithParams(params, opts...)
}

func (a *Client) GetLibraryElementConnectionsWithParams(params *GetLibraryElementConnectionsParams, opts ...ClientOption) (*GetLibraryElementConnectionsOK, error) {
	if params == nil {
		params = NewGetLibraryElementConnectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLibraryElementConnections",
		Method:             "GET",
		PathPattern:        "/library-elements/{library_element_uid}/connections/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLibraryElementConnectionsReader{formats: a.formats},
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
	success, ok := result.(*GetLibraryElementConnectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLibraryElementConnections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLibraryElements gets all library elements

Returns a list of all library elements the authenticated user has permission to view.
Use the `perPage` query parameter to control the maximum number of library elements returned; the default limit is `100`.
You can also use the `page` query parameter to fetch library elements from any page other than the first one.
*/

func (a *Client) GetLibraryElements(params *GetLibraryElementsParams, opts ...ClientOption) (*GetLibraryElementsOK, error) {
	if params == nil {
		params = NewGetLibraryElementsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLibraryElements",
		Method:             "GET",
		PathPattern:        "/library-elements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLibraryElementsReader{formats: a.formats},
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
	success, ok := result.(*GetLibraryElementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLibraryElements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateLibraryElement updates library element

Updates an existing library element identified by uid.
*/
func (a *Client) UpdateLibraryElement(libraryElementUID string, body *models.PatchLibraryElementCommand, opts ...ClientOption) (*UpdateLibraryElementOK, error) {
	params := NewUpdateLibraryElementParams().WithBody(body).WithLibraryElementUID(libraryElementUID)
	return a.UpdateLibraryElementWithParams(params, opts...)
}

func (a *Client) UpdateLibraryElementWithParams(params *UpdateLibraryElementParams, opts ...ClientOption) (*UpdateLibraryElementOK, error) {
	if params == nil {
		params = NewUpdateLibraryElementParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateLibraryElement",
		Method:             "PATCH",
		PathPattern:        "/library-elements/{library_element_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateLibraryElementReader{formats: a.formats},
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
	success, ok := result.(*UpdateLibraryElementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateLibraryElement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
