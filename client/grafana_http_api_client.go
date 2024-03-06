// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command
// To edit this file, edit the custom template in templates/clientFacade.gotmpl
// More info on custom templates can be found in the README or here: https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md
// The template itself can be found here: https://github.com/go-swagger/go-swagger/blob/master/generator/templates/client/facade.gotmpl

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/pkg/transport"

	"github.com/grafana/grafana-openapi-client-go/client/access_control"
	"github.com/grafana/grafana-openapi-client-go/client/access_control_provisioning"
	"github.com/grafana/grafana-openapi-client-go/client/admin"
	"github.com/grafana/grafana-openapi-client-go/client/admin_ldap"
	"github.com/grafana/grafana-openapi-client-go/client/admin_provisioning"
	"github.com/grafana/grafana-openapi-client-go/client/admin_users"
	"github.com/grafana/grafana-openapi-client-go/client/annotations"
	"github.com/grafana/grafana-openapi-client-go/client/api_keys"
	"github.com/grafana/grafana-openapi-client-go/client/correlations"
	"github.com/grafana/grafana-openapi-client-go/client/dashboard_permissions"
	"github.com/grafana/grafana-openapi-client-go/client/dashboard_public"
	"github.com/grafana/grafana-openapi-client-go/client/dashboard_versions"
	"github.com/grafana/grafana-openapi-client-go/client/dashboards"
	"github.com/grafana/grafana-openapi-client-go/client/datasources"
	"github.com/grafana/grafana-openapi-client-go/client/devices"
	"github.com/grafana/grafana-openapi-client-go/client/ds"
	"github.com/grafana/grafana-openapi-client-go/client/enterprise"
	"github.com/grafana/grafana-openapi-client-go/client/folder_permissions"
	"github.com/grafana/grafana-openapi-client-go/client/folders"
	"github.com/grafana/grafana-openapi-client-go/client/get_current_org"
	"github.com/grafana/grafana-openapi-client-go/client/ldap_debug"
	"github.com/grafana/grafana-openapi-client-go/client/legacy_alerts"
	"github.com/grafana/grafana-openapi-client-go/client/legacy_alerts_notification_channels"
	"github.com/grafana/grafana-openapi-client-go/client/library_elements"
	"github.com/grafana/grafana-openapi-client-go/client/licensing"
	"github.com/grafana/grafana-openapi-client-go/client/org"
	"github.com/grafana/grafana-openapi-client-go/client/org_invites"
	"github.com/grafana/grafana-openapi-client-go/client/org_preferences"
	"github.com/grafana/grafana-openapi-client-go/client/orgs"
	"github.com/grafana/grafana-openapi-client-go/client/playlists"
	"github.com/grafana/grafana-openapi-client-go/client/provisioning"
	"github.com/grafana/grafana-openapi-client-go/client/query_history"
	"github.com/grafana/grafana-openapi-client-go/client/recording_rules"
	"github.com/grafana/grafana-openapi-client-go/client/reports"
	"github.com/grafana/grafana-openapi-client-go/client/saml"
	"github.com/grafana/grafana-openapi-client-go/client/search"
	"github.com/grafana/grafana-openapi-client-go/client/service_accounts"
	"github.com/grafana/grafana-openapi-client-go/client/signed_in_user"
	"github.com/grafana/grafana-openapi-client-go/client/signing_keys"
	"github.com/grafana/grafana-openapi-client-go/client/snapshots"
	"github.com/grafana/grafana-openapi-client-go/client/sso_settings"
	"github.com/grafana/grafana-openapi-client-go/client/sync_team_groups"
	"github.com/grafana/grafana-openapi-client-go/client/teams"
	"github.com/grafana/grafana-openapi-client-go/client/user"
	"github.com/grafana/grafana-openapi-client-go/client/user_preferences"
	"github.com/grafana/grafana-openapi-client-go/client/users"
)

// Default grafana HTTP API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
	// Optional property that specifies the org.
	// For more info, see: https://grafana.com/docs/grafana/latest/developers/http_api/auth/
	OrgIDHeader = "X-Grafana-Org-Id"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new grafana HTTP API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *GrafanaHTTPAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new grafana HTTP API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *GrafanaHTTPAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := newTransportWithConfig(cfg)
	return New(transport, cfg, formats)
}

// New creates a new grafana HTTP API client
func New(transport runtime.ClientTransport, cfg *TransportConfig, formats strfmt.Registry) *GrafanaHTTPAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(GrafanaHTTPAPI)
	cli.cfg = cfg
	cli.Transport = transport
	cli.formats = formats
	cli.AccessControl = access_control.New(transport, formats)
	cli.AccessControlProvisioning = access_control_provisioning.New(transport, formats)
	cli.Admin = admin.New(transport, formats)
	cli.AdminLDAP = admin_ldap.New(transport, formats)
	cli.AdminProvisioning = admin_provisioning.New(transport, formats)
	cli.AdminUsers = admin_users.New(transport, formats)
	cli.Annotations = annotations.New(transport, formats)
	cli.APIKeys = api_keys.New(transport, formats)
	cli.Correlations = correlations.New(transport, formats)
	cli.DashboardPermissions = dashboard_permissions.New(transport, formats)
	cli.DashboardPublic = dashboard_public.New(transport, formats)
	cli.DashboardVersions = dashboard_versions.New(transport, formats)
	cli.Dashboards = dashboards.New(transport, formats)
	cli.Datasources = datasources.New(transport, formats)
	cli.Devices = devices.New(transport, formats)
	cli.Ds = ds.New(transport, formats)
	cli.Enterprise = enterprise.New(transport, formats)
	cli.FolderPermissions = folder_permissions.New(transport, formats)
	cli.Folders = folders.New(transport, formats)
	cli.GetCurrentOrg = get_current_org.New(transport, formats)
	cli.LDAPDebug = ldap_debug.New(transport, formats)
	cli.LegacyAlerts = legacy_alerts.New(transport, formats)
	cli.LegacyAlertsNotificationChannels = legacy_alerts_notification_channels.New(transport, formats)
	cli.LibraryElements = library_elements.New(transport, formats)
	cli.Licensing = licensing.New(transport, formats)
	cli.Org = org.New(transport, formats)
	cli.OrgInvites = org_invites.New(transport, formats)
	cli.OrgPreferences = org_preferences.New(transport, formats)
	cli.Orgs = orgs.New(transport, formats)
	cli.Playlists = playlists.New(transport, formats)
	cli.Provisioning = provisioning.New(transport, formats)
	cli.QueryHistory = query_history.New(transport, formats)
	cli.RecordingRules = recording_rules.New(transport, formats)
	cli.Reports = reports.New(transport, formats)
	cli.Saml = saml.New(transport, formats)
	cli.Search = search.New(transport, formats)
	cli.ServiceAccounts = service_accounts.New(transport, formats)
	cli.SignedInUser = signed_in_user.New(transport, formats)
	cli.SigningKeys = signing_keys.New(transport, formats)
	cli.Snapshots = snapshots.New(transport, formats)
	cli.SsoSettings = sso_settings.New(transport, formats)
	cli.SyncTeamGroups = sync_team_groups.New(transport, formats)
	cli.Teams = teams.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.UserPreferences = user_preferences.New(transport, formats)
	cli.Users = users.New(transport, formats)
	return cli
}

// Clone creates a clone of the grafana HTTP API client.
// This new client can then be modified independently of the original client with the With* methods.
func (c *GrafanaHTTPAPI) Clone() *GrafanaHTTPAPI {
	cfg := *c.cfg
	return New(c.Transport, &cfg, c.formats)
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	// Host is the doman name or IP address of the host that serves the API.
	Host string
	// BasePath is the URL prefix for all API paths, relative to the host root.
	BasePath string
	// Schemes are the transfer protocols used by the API (http or https).
	Schemes []string
	// APIKey is an optional API key or service account token.
	APIKey string
	// BasicAuth is optional basic auth credentials.
	BasicAuth *url.Userinfo
	// OrgID provides an optional organization ID.
	// OrgID is only supported with BasicAuth since API keys are already org-scoped.
	OrgID int64
	// TLSConfig provides an optional configuration for a TLS client
	TLSConfig *tls.Config
	// NumRetries contains the optional number of attempted retries
	NumRetries int
	// RetryTimeout sets an optional time to wait before retrying a request
	RetryTimeout time.Duration
	// RetryStatusCodes contains the optional list of status codes to retry
	// Use "x" as a wildcard for a single digit (default: [429, 5xx])
	RetryStatusCodes []string
	// HTTPHeaders contains an optional map of HTTP headers to add to each request
	HTTPHeaders map[string]string
	// Debug sets the optional debug mode for the transport
	Debug bool
	// Client sets the net/http client used for the transport
	Client *http.Client
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// GrafanaHTTPAPI is a client for grafana HTTP API
type GrafanaHTTPAPI struct {
	AccessControl access_control.ClientService

	AccessControlProvisioning access_control_provisioning.ClientService

	Admin admin.ClientService

	AdminLDAP admin_ldap.ClientService

	AdminProvisioning admin_provisioning.ClientService

	AdminUsers admin_users.ClientService

	Annotations annotations.ClientService

	APIKeys api_keys.ClientService

	Correlations correlations.ClientService

	DashboardPermissions dashboard_permissions.ClientService

	DashboardPublic dashboard_public.ClientService

	DashboardVersions dashboard_versions.ClientService

	Dashboards dashboards.ClientService

	Datasources datasources.ClientService

	Devices devices.ClientService

	Ds ds.ClientService

	Enterprise enterprise.ClientService

	FolderPermissions folder_permissions.ClientService

	Folders folders.ClientService

	GetCurrentOrg get_current_org.ClientService

	LDAPDebug ldap_debug.ClientService

	LegacyAlerts legacy_alerts.ClientService

	LegacyAlertsNotificationChannels legacy_alerts_notification_channels.ClientService

	LibraryElements library_elements.ClientService

	Licensing licensing.ClientService

	Org org.ClientService

	OrgInvites org_invites.ClientService

	OrgPreferences org_preferences.ClientService

	Orgs orgs.ClientService

	Playlists playlists.ClientService

	Provisioning provisioning.ClientService

	QueryHistory query_history.ClientService

	RecordingRules recording_rules.ClientService

	Reports reports.ClientService

	Saml saml.ClientService

	Search search.ClientService

	ServiceAccounts service_accounts.ClientService

	SignedInUser signed_in_user.ClientService

	SigningKeys signing_keys.ClientService

	Snapshots snapshots.ClientService

	SsoSettings sso_settings.ClientService

	SyncTeamGroups sync_team_groups.ClientService

	Teams teams.ClientService

	User user.ClientService

	UserPreferences user_preferences.ClientService

	Users users.ClientService

	Transport runtime.ClientTransport
	// cfg is private because it should only be read (for example, to get the OrgID) or set (and then the transport must be created again)
	cfg     *TransportConfig
	formats strfmt.Registry
}

// SetTransport changes the transport on the client and all its subresources
func (c *GrafanaHTTPAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccessControl.SetTransport(transport)
	c.AccessControlProvisioning.SetTransport(transport)
	c.Admin.SetTransport(transport)
	c.AdminLDAP.SetTransport(transport)
	c.AdminProvisioning.SetTransport(transport)
	c.AdminUsers.SetTransport(transport)
	c.Annotations.SetTransport(transport)
	c.APIKeys.SetTransport(transport)
	c.Correlations.SetTransport(transport)
	c.DashboardPermissions.SetTransport(transport)
	c.DashboardPublic.SetTransport(transport)
	c.DashboardVersions.SetTransport(transport)
	c.Dashboards.SetTransport(transport)
	c.Datasources.SetTransport(transport)
	c.Devices.SetTransport(transport)
	c.Ds.SetTransport(transport)
	c.Enterprise.SetTransport(transport)
	c.FolderPermissions.SetTransport(transport)
	c.Folders.SetTransport(transport)
	c.GetCurrentOrg.SetTransport(transport)
	c.LDAPDebug.SetTransport(transport)
	c.LegacyAlerts.SetTransport(transport)
	c.LegacyAlertsNotificationChannels.SetTransport(transport)
	c.LibraryElements.SetTransport(transport)
	c.Licensing.SetTransport(transport)
	c.Org.SetTransport(transport)
	c.OrgInvites.SetTransport(transport)
	c.OrgPreferences.SetTransport(transport)
	c.Orgs.SetTransport(transport)
	c.Playlists.SetTransport(transport)
	c.Provisioning.SetTransport(transport)
	c.QueryHistory.SetTransport(transport)
	c.RecordingRules.SetTransport(transport)
	c.Reports.SetTransport(transport)
	c.Saml.SetTransport(transport)
	c.Search.SetTransport(transport)
	c.ServiceAccounts.SetTransport(transport)
	c.SignedInUser.SetTransport(transport)
	c.SigningKeys.SetTransport(transport)
	c.Snapshots.SetTransport(transport)
	c.SsoSettings.SetTransport(transport)
	c.SyncTeamGroups.SetTransport(transport)
	c.Teams.SetTransport(transport)
	c.User.SetTransport(transport)
	c.UserPreferences.SetTransport(transport)
	c.Users.SetTransport(transport)
}

// OrgID returns the organization ID that was set in the transport config
func (c *GrafanaHTTPAPI) OrgID() int64 {
	return c.cfg.OrgID
}

// WithOrgID sets the organization ID and returns the client
func (c *GrafanaHTTPAPI) WithOrgID(orgID int64) *GrafanaHTTPAPI {
	c.cfg.OrgID = orgID
	c.SetTransport(newTransportWithConfig(c.cfg))
	return c
}

// WithRetries sets retry parameters and returns the client
func (c *GrafanaHTTPAPI) WithRetries(numRetries int, retryTimeout time.Duration, retryStatusCodes ...string) *GrafanaHTTPAPI {
	c.cfg.NumRetries = numRetries
	c.cfg.RetryTimeout = retryTimeout
	c.cfg.RetryStatusCodes = retryStatusCodes
	c.SetTransport(newTransportWithConfig(c.cfg))
	return c
}

// WithHTTPClient sets the used net/http client and returns the API client
func (c *GrafanaHTTPAPI) WithHTTPClient(client *http.Client) *GrafanaHTTPAPI {
	c.cfg.Client = client
	c.SetTransport(newTransportWithConfig(c.cfg))
	return c
}

func newTransportWithConfig(cfg *TransportConfig) *httptransport.Runtime {
	httpTransport := http.DefaultTransport.(*http.Transport)
	httpTransport.TLSClientConfig = cfg.TLSConfig

	retryableTransport := &transport.RetryableTransport{
		Transport:        httpTransport,
		NumRetries:       cfg.NumRetries,
		RetryTimeout:     cfg.RetryTimeout,
		RetryStatusCodes: cfg.RetryStatusCodes,
		HTTPHeaders:      cfg.HTTPHeaders,
	}

	tr := httptransport.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, cfg.Client)
	tr.Transport = retryableTransport

	var auth []runtime.ClientAuthInfoWriter
	if cfg.BasicAuth != nil {
		pwd, _ := cfg.BasicAuth.Password()
		basicAuth := httptransport.BasicAuth(cfg.BasicAuth.Username(), pwd)
		auth = append(auth, basicAuth)
	}
	if cfg.OrgID != 0 {
		orgIDHeader := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
			return r.SetHeaderParam(OrgIDHeader, strconv.FormatInt(cfg.OrgID, 10))
		})
		auth = append(auth, orgIDHeader)
	}
	if cfg.APIKey != "" {
		APIKey := httptransport.BearerToken(cfg.APIKey)
		auth = append(auth, APIKey)
	}

	tr.DefaultAuthentication = httptransport.Compose(auth...)

	// The default runtime.JSONConsumer uses `json.Number` for numbers which is unwieldy to use.
	tr.Consumers[runtime.JSONMime] = runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		return json.NewDecoder(reader).Decode(data)
	})

	tr.Debug = cfg.Debug

	return tr
}
