// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/grafana/grafana-openapi-client-go/pkg/custom_models"
)

// UpdateDataSourceCommand Also acts as api DTO
//
// swagger:model UpdateDataSourceCommand
type UpdateDataSourceCommand struct {

	// access
	Access DsAccess `json:"access,omitempty"`

	// basic auth
	BasicAuth bool `json:"basicAuth,omitempty"`

	// basic auth user
	BasicAuthUser string `json:"basicAuthUser,omitempty"`

	// database
	Database string `json:"database,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// json data
	JSONData custom_models.JSON `json:"jsonData,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// secure Json data
	SecureJSONData map[string]string `json:"secureJsonData,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`

	// with credentials
	WithCredentials bool `json:"withCredentials,omitempty"`
}

// Validate validates this update data source command
func (m *UpdateDataSourceCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJSONData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDataSourceCommand) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if err := m.Access.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("access")
		}
		return err
	}

	return nil
}

func (m *UpdateDataSourceCommand) validateJSONData(formats strfmt.Registry) error {
	if swag.IsZero(m.JSONData) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this update data source command based on the context it is used
func (m *UpdateDataSourceCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDataSourceCommand) contextValidateAccess(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if err := m.Access.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("access")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDataSourceCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDataSourceCommand) UnmarshalBinary(b []byte) error {
	var res UpdateDataSourceCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
