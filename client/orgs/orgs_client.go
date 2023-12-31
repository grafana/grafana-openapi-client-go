// Code generated by go-swagger; DO NOT EDIT.

package orgs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new orgs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for orgs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddOrgUser(orgID int64, body *models.AddOrgUserCommand, opts ...ClientOption) (*AddOrgUserOK, error)
	AddOrgUserWithParams(params *AddOrgUserParams, opts ...ClientOption) (*AddOrgUserOK, error)

	CreateOrg(body *models.CreateOrgCommand, opts ...ClientOption) (*CreateOrgOK, error)
	CreateOrgWithParams(params *CreateOrgParams, opts ...ClientOption) (*CreateOrgOK, error)

	DeleteOrgByID(orgID int64, opts ...ClientOption) (*DeleteOrgByIDOK, error)
	DeleteOrgByIDWithParams(params *DeleteOrgByIDParams, opts ...ClientOption) (*DeleteOrgByIDOK, error)

	GetOrgByID(orgID int64, opts ...ClientOption) (*GetOrgByIDOK, error)
	GetOrgByIDWithParams(params *GetOrgByIDParams, opts ...ClientOption) (*GetOrgByIDOK, error)

	GetOrgByName(orgName string, opts ...ClientOption) (*GetOrgByNameOK, error)
	GetOrgByNameWithParams(params *GetOrgByNameParams, opts ...ClientOption) (*GetOrgByNameOK, error)

	GetOrgQuota(orgID int64, opts ...ClientOption) (*GetOrgQuotaOK, error)
	GetOrgQuotaWithParams(params *GetOrgQuotaParams, opts ...ClientOption) (*GetOrgQuotaOK, error)

	GetOrgUsers(orgID int64, opts ...ClientOption) (*GetOrgUsersOK, error)
	GetOrgUsersWithParams(params *GetOrgUsersParams, opts ...ClientOption) (*GetOrgUsersOK, error)

	RemoveOrgUser(userID int64, orgID int64, opts ...ClientOption) (*RemoveOrgUserOK, error)
	RemoveOrgUserWithParams(params *RemoveOrgUserParams, opts ...ClientOption) (*RemoveOrgUserOK, error)

	SearchOrgUsers(orgID int64, opts ...ClientOption) (*SearchOrgUsersOK, error)
	SearchOrgUsersWithParams(params *SearchOrgUsersParams, opts ...ClientOption) (*SearchOrgUsersOK, error)

	SearchOrgs(params *SearchOrgsParams, opts ...ClientOption) (*SearchOrgsOK, error)

	UpdateOrg(orgID int64, body *models.UpdateOrgForm, opts ...ClientOption) (*UpdateOrgOK, error)
	UpdateOrgWithParams(params *UpdateOrgParams, opts ...ClientOption) (*UpdateOrgOK, error)

	UpdateOrgAddress(orgID int64, body *models.UpdateOrgAddressForm, opts ...ClientOption) (*UpdateOrgAddressOK, error)
	UpdateOrgAddressWithParams(params *UpdateOrgAddressParams, opts ...ClientOption) (*UpdateOrgAddressOK, error)

	UpdateOrgQuota(params *UpdateOrgQuotaParams, opts ...ClientOption) (*UpdateOrgQuotaOK, error)

	UpdateOrgUser(params *UpdateOrgUserParams, opts ...ClientOption) (*UpdateOrgUserOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddOrgUser adds a new user to the current organization

Adds a global user to the current organization.

If you are running Grafana Enterprise and have Fine-grained access control enabled
you need to have a permission with action: `org.users:add` with scope `users:*`.
*/
func (a *Client) AddOrgUser(orgID int64, body *models.AddOrgUserCommand, opts ...ClientOption) (*AddOrgUserOK, error) {
	params := NewAddOrgUserParams().WithBody(body).WithOrgID(orgID)
	return a.AddOrgUserWithParams(params, opts...)
}

func (a *Client) AddOrgUserWithParams(params *AddOrgUserParams, opts ...ClientOption) (*AddOrgUserOK, error) {
	if params == nil {
		params = NewAddOrgUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addOrgUser",
		Method:             "POST",
		PathPattern:        "/orgs/{org_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddOrgUserReader{formats: a.formats},
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
	success, ok := result.(*AddOrgUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOrgUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateOrg creates organization

Only works if [users.allow_org_create](https://grafana.com/docs/grafana/latest/administration/configuration/#allow_org_create) is set.
*/
func (a *Client) CreateOrg(body *models.CreateOrgCommand, opts ...ClientOption) (*CreateOrgOK, error) {
	params := NewCreateOrgParams().WithBody(body)
	return a.CreateOrgWithParams(params, opts...)
}

func (a *Client) CreateOrgWithParams(params *CreateOrgParams, opts ...ClientOption) (*CreateOrgOK, error) {
	if params == nil {
		params = NewCreateOrgParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createOrg",
		Method:             "POST",
		PathPattern:        "/orgs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateOrgReader{formats: a.formats},
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
	success, ok := result.(*CreateOrgOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createOrg: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteOrgByID deletes organization
*/
func (a *Client) DeleteOrgByID(orgID int64, opts ...ClientOption) (*DeleteOrgByIDOK, error) {
	params := NewDeleteOrgByIDParams().WithOrgID(orgID)
	return a.DeleteOrgByIDWithParams(params, opts...)
}

func (a *Client) DeleteOrgByIDWithParams(params *DeleteOrgByIDParams, opts ...ClientOption) (*DeleteOrgByIDOK, error) {
	if params == nil {
		params = NewDeleteOrgByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteOrgByID",
		Method:             "DELETE",
		PathPattern:        "/orgs/{org_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrgByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteOrgByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOrgByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrgByID gets organization by ID
*/
func (a *Client) GetOrgByID(orgID int64, opts ...ClientOption) (*GetOrgByIDOK, error) {
	params := NewGetOrgByIDParams().WithOrgID(orgID)
	return a.GetOrgByIDWithParams(params, opts...)
}

func (a *Client) GetOrgByIDWithParams(params *GetOrgByIDParams, opts ...ClientOption) (*GetOrgByIDOK, error) {
	if params == nil {
		params = NewGetOrgByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgByID",
		Method:             "GET",
		PathPattern:        "/orgs/{org_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrgByIDReader{formats: a.formats},
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
	success, ok := result.(*GetOrgByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrgByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrgByName gets organization by ID
*/
func (a *Client) GetOrgByName(orgName string, opts ...ClientOption) (*GetOrgByNameOK, error) {
	params := NewGetOrgByNameParams().WithOrgName(orgName)
	return a.GetOrgByNameWithParams(params, opts...)
}

func (a *Client) GetOrgByNameWithParams(params *GetOrgByNameParams, opts ...ClientOption) (*GetOrgByNameOK, error) {
	if params == nil {
		params = NewGetOrgByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgByName",
		Method:             "GET",
		PathPattern:        "/orgs/name/{org_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrgByNameReader{formats: a.formats},
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
	success, ok := result.(*GetOrgByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrgByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrgQuota fetches organization quota

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `orgs.quotas:read` and scope `org:id:1` (orgIDScope).
*/
func (a *Client) GetOrgQuota(orgID int64, opts ...ClientOption) (*GetOrgQuotaOK, error) {
	params := NewGetOrgQuotaParams().WithOrgID(orgID)
	return a.GetOrgQuotaWithParams(params, opts...)
}

func (a *Client) GetOrgQuotaWithParams(params *GetOrgQuotaParams, opts ...ClientOption) (*GetOrgQuotaOK, error) {
	if params == nil {
		params = NewGetOrgQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgQuota",
		Method:             "GET",
		PathPattern:        "/orgs/{org_id}/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrgQuotaReader{formats: a.formats},
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
	success, ok := result.(*GetOrgQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrgQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrgUsers gets users in organization

If you are running Grafana Enterprise and have Fine-grained access control enabled
you need to have a permission with action: `org.users:read` with scope `users:*`.
*/
func (a *Client) GetOrgUsers(orgID int64, opts ...ClientOption) (*GetOrgUsersOK, error) {
	params := NewGetOrgUsersParams().WithOrgID(orgID)
	return a.GetOrgUsersWithParams(params, opts...)
}

func (a *Client) GetOrgUsersWithParams(params *GetOrgUsersParams, opts ...ClientOption) (*GetOrgUsersOK, error) {
	if params == nil {
		params = NewGetOrgUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgUsers",
		Method:             "GET",
		PathPattern:        "/orgs/{org_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrgUsersReader{formats: a.formats},
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
	success, ok := result.(*GetOrgUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrgUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveOrgUser deletes user in current organization

If you are running Grafana Enterprise and have Fine-grained access control enabled
you need to have a permission with action: `org.users:remove` with scope `users:*`.
*/
func (a *Client) RemoveOrgUser(userID int64, orgID int64, opts ...ClientOption) (*RemoveOrgUserOK, error) {
	params := NewRemoveOrgUserParams().WithOrgID(orgID).WithUserID(userID)
	return a.RemoveOrgUserWithParams(params, opts...)
}

func (a *Client) RemoveOrgUserWithParams(params *RemoveOrgUserParams, opts ...ClientOption) (*RemoveOrgUserOK, error) {
	if params == nil {
		params = NewRemoveOrgUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeOrgUser",
		Method:             "DELETE",
		PathPattern:        "/orgs/{org_id}/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveOrgUserReader{formats: a.formats},
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
	success, ok := result.(*RemoveOrgUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeOrgUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchOrgUsers searches users in organization

If you are running Grafana Enterprise and have Fine-grained access control enabled
you need to have a permission with action: `org.users:read` with scope `users:*`.
*/
func (a *Client) SearchOrgUsers(orgID int64, opts ...ClientOption) (*SearchOrgUsersOK, error) {
	params := NewSearchOrgUsersParams().WithOrgID(orgID)
	return a.SearchOrgUsersWithParams(params, opts...)
}

func (a *Client) SearchOrgUsersWithParams(params *SearchOrgUsersParams, opts ...ClientOption) (*SearchOrgUsersOK, error) {
	if params == nil {
		params = NewSearchOrgUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchOrgUsers",
		Method:             "GET",
		PathPattern:        "/orgs/{org_id}/users/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchOrgUsersReader{formats: a.formats},
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
	success, ok := result.(*SearchOrgUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchOrgUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchOrgs searches all organizations
*/

func (a *Client) SearchOrgs(params *SearchOrgsParams, opts ...ClientOption) (*SearchOrgsOK, error) {
	if params == nil {
		params = NewSearchOrgsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchOrgs",
		Method:             "GET",
		PathPattern:        "/orgs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchOrgsReader{formats: a.formats},
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
	success, ok := result.(*SearchOrgsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchOrgs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrg updates organization
*/
func (a *Client) UpdateOrg(orgID int64, body *models.UpdateOrgForm, opts ...ClientOption) (*UpdateOrgOK, error) {
	params := NewUpdateOrgParams().WithBody(body).WithOrgID(orgID)
	return a.UpdateOrgWithParams(params, opts...)
}

func (a *Client) UpdateOrgWithParams(params *UpdateOrgParams, opts ...ClientOption) (*UpdateOrgOK, error) {
	if params == nil {
		params = NewUpdateOrgParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrg",
		Method:             "PUT",
		PathPattern:        "/orgs/{org_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrgReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrg: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrgAddress updates organization s address
*/
func (a *Client) UpdateOrgAddress(orgID int64, body *models.UpdateOrgAddressForm, opts ...ClientOption) (*UpdateOrgAddressOK, error) {
	params := NewUpdateOrgAddressParams().WithBody(body).WithOrgID(orgID)
	return a.UpdateOrgAddressWithParams(params, opts...)
}

func (a *Client) UpdateOrgAddressWithParams(params *UpdateOrgAddressParams, opts ...ClientOption) (*UpdateOrgAddressOK, error) {
	if params == nil {
		params = NewUpdateOrgAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrgAddress",
		Method:             "PUT",
		PathPattern:        "/orgs/{org_id}/address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrgAddressReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrgAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrgQuota updates user quota

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `orgs.quotas:write` and scope `org:id:1` (orgIDScope).
*/

func (a *Client) UpdateOrgQuota(params *UpdateOrgQuotaParams, opts ...ClientOption) (*UpdateOrgQuotaOK, error) {
	if params == nil {
		params = NewUpdateOrgQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrgQuota",
		Method:             "PUT",
		PathPattern:        "/orgs/{org_id}/quotas/{quota_target}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrgQuotaReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrgQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrgUser updates users in organization

If you are running Grafana Enterprise and have Fine-grained access control enabled
you need to have a permission with action: `org.users.role:update` with scope `users:*`.
*/

func (a *Client) UpdateOrgUser(params *UpdateOrgUserParams, opts ...ClientOption) (*UpdateOrgUserOK, error) {
	if params == nil {
		params = NewUpdateOrgUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrgUser",
		Method:             "PATCH",
		PathPattern:        "/orgs/{org_id}/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrgUserReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrgUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
