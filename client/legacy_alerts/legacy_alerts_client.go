// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new legacy alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for legacy alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAlertByID(alertID string, opts ...ClientOption) (*GetAlertByIDOK, error)
	GetAlertByIDWithParams(params *GetAlertByIDParams, opts ...ClientOption) (*GetAlertByIDOK, error)

	GetAlerts(params *GetAlertsParams, opts ...ClientOption) (*GetAlertsOK, error)

	GetDashboardStates(dashboardID int64, opts ...ClientOption) (*GetDashboardStatesOK, error)
	GetDashboardStatesWithParams(params *GetDashboardStatesParams, opts ...ClientOption) (*GetDashboardStatesOK, error)

	PauseAlert(alertID string, body *models.PauseAlertCommand, opts ...ClientOption) (*PauseAlertOK, error)
	PauseAlertWithParams(params *PauseAlertParams, opts ...ClientOption) (*PauseAlertOK, error)

	TestAlert(params *TestAlertParams, opts ...ClientOption) (*TestAlertOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAlertByID gets alert by ID

“evalMatches” data in the response is cached in the db when and only when the state of the alert changes (e.g. transitioning from “ok” to “alerting” state).
If data from one server triggers the alert first and, before that server is seen leaving alerting state, a second server also enters a state that would trigger the alert, the second server will not be visible in “evalMatches” data.
*/
func (a *Client) GetAlertByID(alertID string, opts ...ClientOption) (*GetAlertByIDOK, error) {
	params := NewGetAlertByIDParams().WithAlertID(alertID)
	return a.GetAlertByIDWithParams(params, opts...)
}

func (a *Client) GetAlertByIDWithParams(params *GetAlertByIDParams, opts ...ClientOption) (*GetAlertByIDOK, error) {
	if params == nil {
		params = NewGetAlertByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertByID",
		Method:             "GET",
		PathPattern:        "/alerts/{alert_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertByIDReader{formats: a.formats},
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
	success, ok := result.(*GetAlertByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAlerts gets legacy alerts
*/

func (a *Client) GetAlerts(params *GetAlertsParams, opts ...ClientOption) (*GetAlertsOK, error) {
	if params == nil {
		params = NewGetAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlerts",
		Method:             "GET",
		PathPattern:        "/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertsReader{formats: a.formats},
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
	success, ok := result.(*GetAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardStates gets alert states for a dashboard
*/
func (a *Client) GetDashboardStates(dashboardID int64, opts ...ClientOption) (*GetDashboardStatesOK, error) {
	params := NewGetDashboardStatesParams().WithDashboardID(dashboardID)
	return a.GetDashboardStatesWithParams(params, opts...)
}

func (a *Client) GetDashboardStatesWithParams(params *GetDashboardStatesParams, opts ...ClientOption) (*GetDashboardStatesOK, error) {
	if params == nil {
		params = NewGetDashboardStatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardStates",
		Method:             "GET",
		PathPattern:        "/alerts/states-for-dashboard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardStatesReader{formats: a.formats},
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
	success, ok := result.(*GetDashboardStatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardStates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PauseAlert pauses unpause alert by id
*/
func (a *Client) PauseAlert(alertID string, body *models.PauseAlertCommand, opts ...ClientOption) (*PauseAlertOK, error) {
	params := NewPauseAlertParams().WithAlertID(alertID).WithBody(body)
	return a.PauseAlertWithParams(params, opts...)
}

func (a *Client) PauseAlertWithParams(params *PauseAlertParams, opts ...ClientOption) (*PauseAlertOK, error) {
	if params == nil {
		params = NewPauseAlertParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pauseAlert",
		Method:             "POST",
		PathPattern:        "/alerts/{alert_id}/pause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PauseAlertReader{formats: a.formats},
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
	success, ok := result.(*PauseAlertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pauseAlert: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TestAlert tests alert
*/

func (a *Client) TestAlert(params *TestAlertParams, opts ...ClientOption) (*TestAlertOK, error) {
	if params == nil {
		params = NewTestAlertParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "testAlert",
		Method:             "POST",
		PathPattern:        "/alerts/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestAlertReader{formats: a.formats},
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
	success, ok := result.(*TestAlertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for testAlert: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
