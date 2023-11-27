// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts_notification_channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new legacy alerts notification channels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for legacy alerts notification channels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAlertNotificationChannel(body *models.CreateAlertNotificationCommand, opts ...ClientOption) (*CreateAlertNotificationChannelOK, error)
	CreateAlertNotificationChannelWithParams(params *CreateAlertNotificationChannelParams, opts ...ClientOption) (*CreateAlertNotificationChannelOK, error)

	DeleteAlertNotificationChannel(notificationChannelID int64, opts ...ClientOption) (*DeleteAlertNotificationChannelOK, error)
	DeleteAlertNotificationChannelWithParams(params *DeleteAlertNotificationChannelParams, opts ...ClientOption) (*DeleteAlertNotificationChannelOK, error)

	DeleteAlertNotificationChannelByUID(notificationChannelUID string, opts ...ClientOption) (*DeleteAlertNotificationChannelByUIDOK, error)
	DeleteAlertNotificationChannelByUIDWithParams(params *DeleteAlertNotificationChannelByUIDParams, opts ...ClientOption) (*DeleteAlertNotificationChannelByUIDOK, error)

	GetAlertNotificationChannelByID(notificationChannelID int64, opts ...ClientOption) (*GetAlertNotificationChannelByIDOK, error)
	GetAlertNotificationChannelByIDWithParams(params *GetAlertNotificationChannelByIDParams, opts ...ClientOption) (*GetAlertNotificationChannelByIDOK, error)

	GetAlertNotificationChannelByUID(notificationChannelUID string, opts ...ClientOption) (*GetAlertNotificationChannelByUIDOK, error)
	GetAlertNotificationChannelByUIDWithParams(params *GetAlertNotificationChannelByUIDParams, opts ...ClientOption) (*GetAlertNotificationChannelByUIDOK, error)

	GetAlertNotificationChannels(opts ...ClientOption) (*GetAlertNotificationChannelsOK, error)
	GetAlertNotificationChannelsWithParams(params *GetAlertNotificationChannelsParams, opts ...ClientOption) (*GetAlertNotificationChannelsOK, error)

	GetAlertNotificationLookup(opts ...ClientOption) (*GetAlertNotificationLookupOK, error)
	GetAlertNotificationLookupWithParams(params *GetAlertNotificationLookupParams, opts ...ClientOption) (*GetAlertNotificationLookupOK, error)

	NotificationChannelTest(body *models.NotificationTestCommand, opts ...ClientOption) (*NotificationChannelTestOK, error)
	NotificationChannelTestWithParams(params *NotificationChannelTestParams, opts ...ClientOption) (*NotificationChannelTestOK, error)

	UpdateAlertNotificationChannel(notificationChannelID int64, body *models.UpdateAlertNotificationCommand, opts ...ClientOption) (*UpdateAlertNotificationChannelOK, error)
	UpdateAlertNotificationChannelWithParams(params *UpdateAlertNotificationChannelParams, opts ...ClientOption) (*UpdateAlertNotificationChannelOK, error)

	UpdateAlertNotificationChannelByUID(notificationChannelUID string, body *models.UpdateAlertNotificationWithUIDCommand, opts ...ClientOption) (*UpdateAlertNotificationChannelByUIDOK, error)
	UpdateAlertNotificationChannelByUIDWithParams(params *UpdateAlertNotificationChannelByUIDParams, opts ...ClientOption) (*UpdateAlertNotificationChannelByUIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAlertNotificationChannel creates notification channel

You can find the full list of [supported notifiers](https://grafana.com/docs/grafana/latest/alerting/old-alerting/notifications/#list-of-supported-notifiers) on the alert notifiers page.
*/
func (a *Client) CreateAlertNotificationChannel(body *models.CreateAlertNotificationCommand, opts ...ClientOption) (*CreateAlertNotificationChannelOK, error) {
	params := NewCreateAlertNotificationChannelParams().WithBody(body)
	return a.CreateAlertNotificationChannelWithParams(params, opts...)
}

func (a *Client) CreateAlertNotificationChannelWithParams(params *CreateAlertNotificationChannelParams, opts ...ClientOption) (*CreateAlertNotificationChannelOK, error) {
	if params == nil {
		params = NewCreateAlertNotificationChannelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAlertNotificationChannel",
		Method:             "POST",
		PathPattern:        "/alert-notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAlertNotificationChannelReader{formats: a.formats},
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
	success, ok := result.(*CreateAlertNotificationChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAlertNotificationChannel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAlertNotificationChannel deletes alert notification by ID

Deletes an existing notification channel identified by ID.
*/
func (a *Client) DeleteAlertNotificationChannel(notificationChannelID int64, opts ...ClientOption) (*DeleteAlertNotificationChannelOK, error) {
	params := NewDeleteAlertNotificationChannelParams().WithNotificationChannelID(notificationChannelID)
	return a.DeleteAlertNotificationChannelWithParams(params, opts...)
}

func (a *Client) DeleteAlertNotificationChannelWithParams(params *DeleteAlertNotificationChannelParams, opts ...ClientOption) (*DeleteAlertNotificationChannelOK, error) {
	if params == nil {
		params = NewDeleteAlertNotificationChannelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAlertNotificationChannel",
		Method:             "DELETE",
		PathPattern:        "/alert-notifications/{notification_channel_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAlertNotificationChannelReader{formats: a.formats},
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
	success, ok := result.(*DeleteAlertNotificationChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAlertNotificationChannel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAlertNotificationChannelByUID deletes alert notification by UID

Deletes an existing notification channel identified by UID.
*/
func (a *Client) DeleteAlertNotificationChannelByUID(notificationChannelUID string, opts ...ClientOption) (*DeleteAlertNotificationChannelByUIDOK, error) {
	params := NewDeleteAlertNotificationChannelByUIDParams().WithNotificationChannelUID(notificationChannelUID)
	return a.DeleteAlertNotificationChannelByUIDWithParams(params, opts...)
}

func (a *Client) DeleteAlertNotificationChannelByUIDWithParams(params *DeleteAlertNotificationChannelByUIDParams, opts ...ClientOption) (*DeleteAlertNotificationChannelByUIDOK, error) {
	if params == nil {
		params = NewDeleteAlertNotificationChannelByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAlertNotificationChannelByUID",
		Method:             "DELETE",
		PathPattern:        "/alert-notifications/uid/{notification_channel_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAlertNotificationChannelByUIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteAlertNotificationChannelByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAlertNotificationChannelByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAlertNotificationChannelByID gets notification channel by ID

Returns the notification channel given the notification channel ID.
*/
func (a *Client) GetAlertNotificationChannelByID(notificationChannelID int64, opts ...ClientOption) (*GetAlertNotificationChannelByIDOK, error) {
	params := NewGetAlertNotificationChannelByIDParams().WithNotificationChannelID(notificationChannelID)
	return a.GetAlertNotificationChannelByIDWithParams(params, opts...)
}

func (a *Client) GetAlertNotificationChannelByIDWithParams(params *GetAlertNotificationChannelByIDParams, opts ...ClientOption) (*GetAlertNotificationChannelByIDOK, error) {
	if params == nil {
		params = NewGetAlertNotificationChannelByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertNotificationChannelByID",
		Method:             "GET",
		PathPattern:        "/alert-notifications/{notification_channel_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertNotificationChannelByIDReader{formats: a.formats},
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
	success, ok := result.(*GetAlertNotificationChannelByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertNotificationChannelByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAlertNotificationChannelByUID gets notification channel by UID

Returns the notification channel given the notification channel UID.
*/
func (a *Client) GetAlertNotificationChannelByUID(notificationChannelUID string, opts ...ClientOption) (*GetAlertNotificationChannelByUIDOK, error) {
	params := NewGetAlertNotificationChannelByUIDParams().WithNotificationChannelUID(notificationChannelUID)
	return a.GetAlertNotificationChannelByUIDWithParams(params, opts...)
}

func (a *Client) GetAlertNotificationChannelByUIDWithParams(params *GetAlertNotificationChannelByUIDParams, opts ...ClientOption) (*GetAlertNotificationChannelByUIDOK, error) {
	if params == nil {
		params = NewGetAlertNotificationChannelByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertNotificationChannelByUID",
		Method:             "GET",
		PathPattern:        "/alert-notifications/uid/{notification_channel_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertNotificationChannelByUIDReader{formats: a.formats},
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
	success, ok := result.(*GetAlertNotificationChannelByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertNotificationChannelByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAlertNotificationChannels gets all notification channels

Returns all notification channels that the authenticated user has permission to view.
*/
func (a *Client) GetAlertNotificationChannels(opts ...ClientOption) (*GetAlertNotificationChannelsOK, error) {
	params := NewGetAlertNotificationChannelsParams()
	return a.GetAlertNotificationChannelsWithParams(params, opts...)
}

func (a *Client) GetAlertNotificationChannelsWithParams(params *GetAlertNotificationChannelsParams, opts ...ClientOption) (*GetAlertNotificationChannelsOK, error) {
	if params == nil {
		params = NewGetAlertNotificationChannelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertNotificationChannels",
		Method:             "GET",
		PathPattern:        "/alert-notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertNotificationChannelsReader{formats: a.formats},
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
	success, ok := result.(*GetAlertNotificationChannelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertNotificationChannels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAlertNotificationLookup gets all notification channels lookup

Returns all notification channels, but with less detailed information. Accessible by any authenticated user and is mainly used by providing alert notification channels in Grafana UI when configuring alert rule.
*/
func (a *Client) GetAlertNotificationLookup(opts ...ClientOption) (*GetAlertNotificationLookupOK, error) {
	params := NewGetAlertNotificationLookupParams()
	return a.GetAlertNotificationLookupWithParams(params, opts...)
}

func (a *Client) GetAlertNotificationLookupWithParams(params *GetAlertNotificationLookupParams, opts ...ClientOption) (*GetAlertNotificationLookupOK, error) {
	if params == nil {
		params = NewGetAlertNotificationLookupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAlertNotificationLookup",
		Method:             "GET",
		PathPattern:        "/alert-notifications/lookup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAlertNotificationLookupReader{formats: a.formats},
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
	success, ok := result.(*GetAlertNotificationLookupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlertNotificationLookup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
NotificationChannelTest tests notification channel

Sends a test notification to the channel.
*/
func (a *Client) NotificationChannelTest(body *models.NotificationTestCommand, opts ...ClientOption) (*NotificationChannelTestOK, error) {
	params := NewNotificationChannelTestParams().WithBody(body)
	return a.NotificationChannelTestWithParams(params, opts...)
}

func (a *Client) NotificationChannelTestWithParams(params *NotificationChannelTestParams, opts ...ClientOption) (*NotificationChannelTestOK, error) {
	if params == nil {
		params = NewNotificationChannelTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationChannelTest",
		Method:             "POST",
		PathPattern:        "/alert-notifications/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NotificationChannelTestReader{formats: a.formats},
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
	success, ok := result.(*NotificationChannelTestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationChannelTest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAlertNotificationChannel updates notification channel by ID

Updates an existing notification channel identified by ID.
*/
func (a *Client) UpdateAlertNotificationChannel(notificationChannelID int64, body *models.UpdateAlertNotificationCommand, opts ...ClientOption) (*UpdateAlertNotificationChannelOK, error) {
	params := NewUpdateAlertNotificationChannelParams().WithBody(body).WithNotificationChannelID(notificationChannelID)
	return a.UpdateAlertNotificationChannelWithParams(params, opts...)
}

func (a *Client) UpdateAlertNotificationChannelWithParams(params *UpdateAlertNotificationChannelParams, opts ...ClientOption) (*UpdateAlertNotificationChannelOK, error) {
	if params == nil {
		params = NewUpdateAlertNotificationChannelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAlertNotificationChannel",
		Method:             "PUT",
		PathPattern:        "/alert-notifications/{notification_channel_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAlertNotificationChannelReader{formats: a.formats},
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
	success, ok := result.(*UpdateAlertNotificationChannelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAlertNotificationChannel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAlertNotificationChannelByUID updates notification channel by UID

Updates an existing notification channel identified by uid.
*/
func (a *Client) UpdateAlertNotificationChannelByUID(notificationChannelUID string, body *models.UpdateAlertNotificationWithUIDCommand, opts ...ClientOption) (*UpdateAlertNotificationChannelByUIDOK, error) {
	params := NewUpdateAlertNotificationChannelByUIDParams().WithBody(body).WithNotificationChannelUID(notificationChannelUID)
	return a.UpdateAlertNotificationChannelByUIDWithParams(params, opts...)
}

func (a *Client) UpdateAlertNotificationChannelByUIDWithParams(params *UpdateAlertNotificationChannelByUIDParams, opts ...ClientOption) (*UpdateAlertNotificationChannelByUIDOK, error) {
	if params == nil {
		params = NewUpdateAlertNotificationChannelByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateAlertNotificationChannelByUID",
		Method:             "PUT",
		PathPattern:        "/alert-notifications/uid/{notification_channel_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAlertNotificationChannelByUIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateAlertNotificationChannelByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAlertNotificationChannelByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
