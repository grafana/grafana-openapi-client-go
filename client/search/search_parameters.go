// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSearchParams creates a new SearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchParams() *SearchParams {
	return &SearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchParamsWithTimeout creates a new SearchParams object
// with the ability to set a timeout on a request.
func NewSearchParamsWithTimeout(timeout time.Duration) *SearchParams {
	return &SearchParams{
		timeout: timeout,
	}
}

// NewSearchParamsWithContext creates a new SearchParams object
// with the ability to set a context for a request.
func NewSearchParamsWithContext(ctx context.Context) *SearchParams {
	return &SearchParams{
		Context: ctx,
	}
}

// NewSearchParamsWithHTTPClient creates a new SearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchParamsWithHTTPClient(client *http.Client) *SearchParams {
	return &SearchParams{
		HTTPClient: client,
	}
}

/*
SearchParams contains all the parameters to send to the API endpoint

	for the search operation.

	Typically these are written to a http.Request.
*/
type SearchParams struct {

	/* DashboardIds.

	     List of dashboard id’s to search for
	This is deprecated: users should use the `dashboardUIDs` query parameter instead
	*/
	DashboardIds []int64

	/* DashboardUIDs.

	   List of dashboard uid’s to search for
	*/
	DashboardUIDs []string

	/* Deleted.

	   Flag indicating if only soft deleted Dashboards should be returned
	*/
	Deleted *bool

	/* FolderIds.

	     List of folder id’s to search in for dashboards
	If it's `0` then it will query for the top level folders
	This is deprecated: users should use the `folderUIDs` query parameter instead
	*/
	FolderIds []int64

	/* FolderUIDs.

	     List of folder UID’s to search in for dashboards
	If it's an empty string then it will query for the top level folders
	*/
	FolderUIDs []string

	/* Limit.

	   Limit the number of returned results (max 5000)

	   Format: int64
	*/
	Limit *int64

	/* Page.

	   Use this parameter to access hits beyond limit. Numbering starts at 1. limit param acts as page size. Only available in Grafana v6.2+.

	   Format: int64
	*/
	Page *int64

	/* Permission.

	   Set to `Edit` to return dashboards/folders that the user can edit

	   Default: "View"
	*/
	Permission *string

	/* Query.

	   Search Query
	*/
	Query *string

	/* Sort.

	   Sort method; for listing all the possible sort methods use the search sorting endpoint.

	   Default: "alpha-asc"
	*/
	Sort *string

	/* Starred.

	   Flag indicating if only starred Dashboards should be returned
	*/
	Starred *bool

	/* Tag.

	   List of tags to search for
	*/
	Tag []string

	/* Type.

	   Type to search for, dash-folder or dash-db
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchParams) WithDefaults() *SearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchParams) SetDefaults() {
	var (
		permissionDefault = string("View")

		sortDefault = string("alpha-asc")
	)

	val := SearchParams{
		Permission: &permissionDefault,
		Sort:       &sortDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search params
func (o *SearchParams) WithTimeout(timeout time.Duration) *SearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search params
func (o *SearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search params
func (o *SearchParams) WithContext(ctx context.Context) *SearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search params
func (o *SearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) WithHTTPClient(client *http.Client) *SearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardIds adds the dashboardIds to the search params
func (o *SearchParams) WithDashboardIds(dashboardIds []int64) *SearchParams {
	o.SetDashboardIds(dashboardIds)
	return o
}

// SetDashboardIds adds the dashboardIds to the search params
func (o *SearchParams) SetDashboardIds(dashboardIds []int64) {
	o.DashboardIds = dashboardIds
}

// WithDashboardUIDs adds the dashboardUIDs to the search params
func (o *SearchParams) WithDashboardUIDs(dashboardUIDs []string) *SearchParams {
	o.SetDashboardUIDs(dashboardUIDs)
	return o
}

// SetDashboardUIDs adds the dashboardUiDs to the search params
func (o *SearchParams) SetDashboardUIDs(dashboardUIDs []string) {
	o.DashboardUIDs = dashboardUIDs
}

// WithDeleted adds the deleted to the search params
func (o *SearchParams) WithDeleted(deleted *bool) *SearchParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the search params
func (o *SearchParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithFolderIds adds the folderIds to the search params
func (o *SearchParams) WithFolderIds(folderIds []int64) *SearchParams {
	o.SetFolderIds(folderIds)
	return o
}

// SetFolderIds adds the folderIds to the search params
func (o *SearchParams) SetFolderIds(folderIds []int64) {
	o.FolderIds = folderIds
}

// WithFolderUIDs adds the folderUIDs to the search params
func (o *SearchParams) WithFolderUIDs(folderUIDs []string) *SearchParams {
	o.SetFolderUIDs(folderUIDs)
	return o
}

// SetFolderUIDs adds the folderUiDs to the search params
func (o *SearchParams) SetFolderUIDs(folderUIDs []string) {
	o.FolderUIDs = folderUIDs
}

// WithLimit adds the limit to the search params
func (o *SearchParams) WithLimit(limit *int64) *SearchParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search params
func (o *SearchParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the search params
func (o *SearchParams) WithPage(page *int64) *SearchParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search params
func (o *SearchParams) SetPage(page *int64) {
	o.Page = page
}

// WithPermission adds the permission to the search params
func (o *SearchParams) WithPermission(permission *string) *SearchParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the search params
func (o *SearchParams) SetPermission(permission *string) {
	o.Permission = permission
}

// WithQuery adds the query to the search params
func (o *SearchParams) WithQuery(query *string) *SearchParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the search params
func (o *SearchParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the search params
func (o *SearchParams) WithSort(sort *string) *SearchParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search params
func (o *SearchParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStarred adds the starred to the search params
func (o *SearchParams) WithStarred(starred *bool) *SearchParams {
	o.SetStarred(starred)
	return o
}

// SetStarred adds the starred to the search params
func (o *SearchParams) SetStarred(starred *bool) {
	o.Starred = starred
}

// WithTag adds the tag to the search params
func (o *SearchParams) WithTag(tag []string) *SearchParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the search params
func (o *SearchParams) SetTag(tag []string) {
	o.Tag = tag
}

// WithType adds the typeVar to the search params
func (o *SearchParams) WithType(typeVar *string) *SearchParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the search params
func (o *SearchParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DashboardIds != nil {

		// binding items for dashboardIds
		joinedDashboardIds := o.bindParamDashboardIds(reg)

		// query array param dashboardIds
		if err := r.SetQueryParam("dashboardIds", joinedDashboardIds...); err != nil {
			return err
		}
	}

	if o.DashboardUIDs != nil {

		// binding items for dashboardUIDs
		joinedDashboardUIDs := o.bindParamDashboardUIDs(reg)

		// query array param dashboardUIDs
		if err := r.SetQueryParam("dashboardUIDs", joinedDashboardUIDs...); err != nil {
			return err
		}
	}

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool

		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {

			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}
	}

	if o.FolderIds != nil {

		// binding items for folderIds
		joinedFolderIds := o.bindParamFolderIds(reg)

		// query array param folderIds
		if err := r.SetQueryParam("folderIds", joinedFolderIds...); err != nil {
			return err
		}
	}

	if o.FolderUIDs != nil {

		// binding items for folderUIDs
		joinedFolderUIDs := o.bindParamFolderUIDs(reg)

		// query array param folderUIDs
		if err := r.SetQueryParam("folderUIDs", joinedFolderUIDs...); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Permission != nil {

		// query param permission
		var qrPermission string

		if o.Permission != nil {
			qrPermission = *o.Permission
		}
		qPermission := qrPermission
		if qPermission != "" {

			if err := r.SetQueryParam("permission", qPermission); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Starred != nil {

		// query param starred
		var qrStarred bool

		if o.Starred != nil {
			qrStarred = *o.Starred
		}
		qStarred := swag.FormatBool(qrStarred)
		if qStarred != "" {

			if err := r.SetQueryParam("starred", qStarred); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// binding items for tag
		joinedTag := o.bindParamTag(reg)

		// query array param tag
		if err := r.SetQueryParam("tag", joinedTag...); err != nil {
			return err
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearch binds the parameter dashboardIds
func (o *SearchParams) bindParamDashboardIds(formats strfmt.Registry) []string {
	dashboardIdsIR := o.DashboardIds

	var dashboardIdsIC []string
	for _, dashboardIdsIIR := range dashboardIdsIR { // explode []int64

		dashboardIdsIIV := swag.FormatInt64(dashboardIdsIIR) // int64 as string
		dashboardIdsIC = append(dashboardIdsIC, dashboardIdsIIV)
	}

	// items.CollectionFormat: ""
	dashboardIdsIS := swag.JoinByFormat(dashboardIdsIC, "")

	return dashboardIdsIS
}

// bindParamSearch binds the parameter dashboardUIDs
func (o *SearchParams) bindParamDashboardUIDs(formats strfmt.Registry) []string {
	dashboardUIDsIR := o.DashboardUIDs

	var dashboardUIDsIC []string
	for _, dashboardUIDsIIR := range dashboardUIDsIR { // explode []string

		dashboardUIDsIIV := dashboardUIDsIIR // string as string
		dashboardUIDsIC = append(dashboardUIDsIC, dashboardUIDsIIV)
	}

	// items.CollectionFormat: ""
	dashboardUIDsIS := swag.JoinByFormat(dashboardUIDsIC, "")

	return dashboardUIDsIS
}

// bindParamSearch binds the parameter folderIds
func (o *SearchParams) bindParamFolderIds(formats strfmt.Registry) []string {
	folderIdsIR := o.FolderIds

	var folderIdsIC []string
	for _, folderIdsIIR := range folderIdsIR { // explode []int64

		folderIdsIIV := swag.FormatInt64(folderIdsIIR) // int64 as string
		folderIdsIC = append(folderIdsIC, folderIdsIIV)
	}

	// items.CollectionFormat: ""
	folderIdsIS := swag.JoinByFormat(folderIdsIC, "")

	return folderIdsIS
}

// bindParamSearch binds the parameter folderUIDs
func (o *SearchParams) bindParamFolderUIDs(formats strfmt.Registry) []string {
	folderUIDsIR := o.FolderUIDs

	var folderUIDsIC []string
	for _, folderUIDsIIR := range folderUIDsIR { // explode []string

		folderUIDsIIV := folderUIDsIIR // string as string
		folderUIDsIC = append(folderUIDsIC, folderUIDsIIV)
	}

	// items.CollectionFormat: ""
	folderUIDsIS := swag.JoinByFormat(folderUIDsIC, "")

	return folderUIDsIS
}

// bindParamSearch binds the parameter tag
func (o *SearchParams) bindParamTag(formats strfmt.Registry) []string {
	tagIR := o.Tag

	var tagIC []string
	for _, tagIIR := range tagIR { // explode []string

		tagIIV := tagIIR // string as string
		tagIC = append(tagIC, tagIIV)
	}

	// items.CollectionFormat: "multi"
	tagIS := swag.JoinByFormat(tagIC, "multi")

	return tagIS
}
