// Code generated by go-swagger; DO NOT EDIT.

package recording_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new recording rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for recording rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRecordingRule(body *models.RecordingRuleJSON, opts ...ClientOption) (*CreateRecordingRuleOK, error)
	CreateRecordingRuleWithParams(params *CreateRecordingRuleParams, opts ...ClientOption) (*CreateRecordingRuleOK, error)

	CreateRecordingRuleWriteTarget(body *models.PrometheusRemoteWriteTargetJSON, opts ...ClientOption) (*CreateRecordingRuleWriteTargetOK, error)
	CreateRecordingRuleWriteTargetWithParams(params *CreateRecordingRuleWriteTargetParams, opts ...ClientOption) (*CreateRecordingRuleWriteTargetOK, error)

	DeleteRecordingRule(recordingRuleID int64, opts ...ClientOption) (*DeleteRecordingRuleOK, error)
	DeleteRecordingRuleWithParams(params *DeleteRecordingRuleParams, opts ...ClientOption) (*DeleteRecordingRuleOK, error)

	DeleteRecordingRuleWriteTarget(opts ...ClientOption) (*DeleteRecordingRuleWriteTargetOK, error)
	DeleteRecordingRuleWriteTargetWithParams(params *DeleteRecordingRuleWriteTargetParams, opts ...ClientOption) (*DeleteRecordingRuleWriteTargetOK, error)

	GetRecordingRuleWriteTarget(opts ...ClientOption) (*GetRecordingRuleWriteTargetOK, error)
	GetRecordingRuleWriteTargetWithParams(params *GetRecordingRuleWriteTargetParams, opts ...ClientOption) (*GetRecordingRuleWriteTargetOK, error)

	ListRecordingRules(opts ...ClientOption) (*ListRecordingRulesOK, error)
	ListRecordingRulesWithParams(params *ListRecordingRulesParams, opts ...ClientOption) (*ListRecordingRulesOK, error)

	TestCreateRecordingRule(body *models.RecordingRuleJSON, opts ...ClientOption) (*TestCreateRecordingRuleOK, error)
	TestCreateRecordingRuleWithParams(params *TestCreateRecordingRuleParams, opts ...ClientOption) (*TestCreateRecordingRuleOK, error)

	UpdateRecordingRule(body *models.RecordingRuleJSON, opts ...ClientOption) (*UpdateRecordingRuleOK, error)
	UpdateRecordingRuleWithParams(params *UpdateRecordingRuleParams, opts ...ClientOption) (*UpdateRecordingRuleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRecordingRule creates a recording rule that is then registered and started
*/
func (a *Client) CreateRecordingRule(body *models.RecordingRuleJSON, opts ...ClientOption) (*CreateRecordingRuleOK, error) {
	params := NewCreateRecordingRuleParams().WithBody(body)
	return a.CreateRecordingRuleWithParams(params, opts...)
}

func (a *Client) CreateRecordingRuleWithParams(params *CreateRecordingRuleParams, opts ...ClientOption) (*CreateRecordingRuleOK, error) {
	if params == nil {
		params = NewCreateRecordingRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRecordingRule",
		Method:             "POST",
		PathPattern:        "/recording-rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRecordingRuleReader{formats: a.formats},
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
	success, ok := result.(*CreateRecordingRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRecordingRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateRecordingRuleWriteTarget creates a remote write target

It returns a 422 if there is not an existing prometheus data source configured.
*/
func (a *Client) CreateRecordingRuleWriteTarget(body *models.PrometheusRemoteWriteTargetJSON, opts ...ClientOption) (*CreateRecordingRuleWriteTargetOK, error) {
	params := NewCreateRecordingRuleWriteTargetParams().WithBody(body)
	return a.CreateRecordingRuleWriteTargetWithParams(params, opts...)
}

func (a *Client) CreateRecordingRuleWriteTargetWithParams(params *CreateRecordingRuleWriteTargetParams, opts ...ClientOption) (*CreateRecordingRuleWriteTargetOK, error) {
	if params == nil {
		params = NewCreateRecordingRuleWriteTargetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRecordingRuleWriteTarget",
		Method:             "POST",
		PathPattern:        "/recording-rules/writer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRecordingRuleWriteTargetReader{formats: a.formats},
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
	success, ok := result.(*CreateRecordingRuleWriteTargetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRecordingRuleWriteTarget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteRecordingRule deletes removes the rule from the registry and stops it
*/
func (a *Client) DeleteRecordingRule(recordingRuleID int64, opts ...ClientOption) (*DeleteRecordingRuleOK, error) {
	params := NewDeleteRecordingRuleParams().WithRecordingRuleID(recordingRuleID)
	return a.DeleteRecordingRuleWithParams(params, opts...)
}

func (a *Client) DeleteRecordingRuleWithParams(params *DeleteRecordingRuleParams, opts ...ClientOption) (*DeleteRecordingRuleOK, error) {
	if params == nil {
		params = NewDeleteRecordingRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRecordingRule",
		Method:             "DELETE",
		PathPattern:        "/recording-rules/{recordingRuleID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecordingRuleReader{formats: a.formats},
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
	success, ok := result.(*DeleteRecordingRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRecordingRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteRecordingRuleWriteTarget deletes the remote write target
*/
func (a *Client) DeleteRecordingRuleWriteTarget(opts ...ClientOption) (*DeleteRecordingRuleWriteTargetOK, error) {
	params := NewDeleteRecordingRuleWriteTargetParams()
	return a.DeleteRecordingRuleWriteTargetWithParams(params, opts...)
}

func (a *Client) DeleteRecordingRuleWriteTargetWithParams(params *DeleteRecordingRuleWriteTargetParams, opts ...ClientOption) (*DeleteRecordingRuleWriteTargetOK, error) {
	if params == nil {
		params = NewDeleteRecordingRuleWriteTargetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRecordingRuleWriteTarget",
		Method:             "DELETE",
		PathPattern:        "/recording-rules/writer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecordingRuleWriteTargetReader{formats: a.formats},
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
	success, ok := result.(*DeleteRecordingRuleWriteTargetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRecordingRuleWriteTarget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRecordingRuleWriteTarget returns the prometheus remote write target
*/
func (a *Client) GetRecordingRuleWriteTarget(opts ...ClientOption) (*GetRecordingRuleWriteTargetOK, error) {
	params := NewGetRecordingRuleWriteTargetParams()
	return a.GetRecordingRuleWriteTargetWithParams(params, opts...)
}

func (a *Client) GetRecordingRuleWriteTargetWithParams(params *GetRecordingRuleWriteTargetParams, opts ...ClientOption) (*GetRecordingRuleWriteTargetOK, error) {
	if params == nil {
		params = NewGetRecordingRuleWriteTargetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRecordingRuleWriteTarget",
		Method:             "GET",
		PathPattern:        "/recording-rules/writer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecordingRuleWriteTargetReader{formats: a.formats},
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
	success, ok := result.(*GetRecordingRuleWriteTargetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRecordingRuleWriteTarget: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListRecordingRules lists all rules in the database active or deleted
*/
func (a *Client) ListRecordingRules(opts ...ClientOption) (*ListRecordingRulesOK, error) {
	params := NewListRecordingRulesParams()
	return a.ListRecordingRulesWithParams(params, opts...)
}

func (a *Client) ListRecordingRulesWithParams(params *ListRecordingRulesParams, opts ...ClientOption) (*ListRecordingRulesOK, error) {
	if params == nil {
		params = NewListRecordingRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listRecordingRules",
		Method:             "GET",
		PathPattern:        "/recording-rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRecordingRulesReader{formats: a.formats},
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
	success, ok := result.(*ListRecordingRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRecordingRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TestCreateRecordingRule tests a recording rule
*/
func (a *Client) TestCreateRecordingRule(body *models.RecordingRuleJSON, opts ...ClientOption) (*TestCreateRecordingRuleOK, error) {
	params := NewTestCreateRecordingRuleParams().WithBody(body)
	return a.TestCreateRecordingRuleWithParams(params, opts...)
}

func (a *Client) TestCreateRecordingRuleWithParams(params *TestCreateRecordingRuleParams, opts ...ClientOption) (*TestCreateRecordingRuleOK, error) {
	if params == nil {
		params = NewTestCreateRecordingRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "testCreateRecordingRule",
		Method:             "POST",
		PathPattern:        "/recording-rules/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestCreateRecordingRuleReader{formats: a.formats},
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
	success, ok := result.(*TestCreateRecordingRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for testCreateRecordingRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateRecordingRule updates the active status of a rule
*/
func (a *Client) UpdateRecordingRule(body *models.RecordingRuleJSON, opts ...ClientOption) (*UpdateRecordingRuleOK, error) {
	params := NewUpdateRecordingRuleParams().WithBody(body)
	return a.UpdateRecordingRuleWithParams(params, opts...)
}

func (a *Client) UpdateRecordingRuleWithParams(params *UpdateRecordingRuleParams, opts ...ClientOption) (*UpdateRecordingRuleOK, error) {
	if params == nil {
		params = NewUpdateRecordingRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRecordingRule",
		Method:             "PUT",
		PathPattern:        "/recording-rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRecordingRuleReader{formats: a.formats},
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
	success, ok := result.(*UpdateRecordingRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRecordingRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
