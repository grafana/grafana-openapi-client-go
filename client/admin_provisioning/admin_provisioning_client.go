// Code generated by go-swagger; DO NOT EDIT.

package admin_provisioning

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

// New creates a new admin provisioning API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, cfg *utils.ClientConfig) ClientService {
	return &Client{transport: transport, formats: formats, cfg: cfg}
}

/*
Client for admin provisioning API
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
	AdminProvisioningReloadDashboards(params *AdminProvisioningReloadDashboardsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadDashboardsOK, error)

	AdminProvisioningReloadDatasources(params *AdminProvisioningReloadDatasourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadDatasourcesOK, error)

	AdminProvisioningReloadNotifications(params *AdminProvisioningReloadNotificationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadNotificationsOK, error)

	AdminProvisioningReloadPlugins(params *AdminProvisioningReloadPluginsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadPluginsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AdminProvisioningReloadDashboards reloads dashboard provisioning configurations

	Reloads the provisioning config files for dashboards again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:dashboards`.
*/
func (a *Client) AdminProvisioningReloadDashboards(params *AdminProvisioningReloadDashboardsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadDashboardsOK, error) {
	var (
		result interface{}
		err    error
	)
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminProvisioningReloadDashboardsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminProvisioningReloadDashboards",
		Method:             "POST",
		PathPattern:        "/admin/provisioning/dashboards/reload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminProvisioningReloadDashboardsReader{formats: a.formats},
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
	success, ok := result.(*AdminProvisioningReloadDashboardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminProvisioningReloadDashboards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AdminProvisioningReloadDatasources reloads datasource provisioning configurations

	Reloads the provisioning config files for datasources again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:datasources`.
*/
func (a *Client) AdminProvisioningReloadDatasources(params *AdminProvisioningReloadDatasourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadDatasourcesOK, error) {
	var (
		result interface{}
		err    error
	)
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminProvisioningReloadDatasourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminProvisioningReloadDatasources",
		Method:             "POST",
		PathPattern:        "/admin/provisioning/datasources/reload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminProvisioningReloadDatasourcesReader{formats: a.formats},
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
	success, ok := result.(*AdminProvisioningReloadDatasourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminProvisioningReloadDatasources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AdminProvisioningReloadNotifications reloads legacy alert notifier provisioning configurations

	Reloads the provisioning config files for legacy alert notifiers again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:notifications`.
*/
func (a *Client) AdminProvisioningReloadNotifications(params *AdminProvisioningReloadNotificationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadNotificationsOK, error) {
	var (
		result interface{}
		err    error
	)
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminProvisioningReloadNotificationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminProvisioningReloadNotifications",
		Method:             "POST",
		PathPattern:        "/admin/provisioning/notifications/reload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminProvisioningReloadNotificationsReader{formats: a.formats},
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
	success, ok := result.(*AdminProvisioningReloadNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminProvisioningReloadNotifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AdminProvisioningReloadPlugins reloads plugin provisioning configurations

	Reloads the provisioning config files for plugins again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:plugin`.
*/
func (a *Client) AdminProvisioningReloadPlugins(params *AdminProvisioningReloadPluginsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadPluginsOK, error) {
	var (
		result interface{}
		err    error
	)
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminProvisioningReloadPluginsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminProvisioningReloadPlugins",
		Method:             "POST",
		PathPattern:        "/admin/provisioning/plugins/reload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminProvisioningReloadPluginsReader{formats: a.formats},
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
	success, ok := result.(*AdminProvisioningReloadPluginsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminProvisioningReloadPlugins: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
