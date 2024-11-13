// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceAccountProfileDTO service account profile DTO
//
// swagger:model ServiceAccountProfileDTO
type ServiceAccountProfileDTO struct {

	// access control
	AccessControl map[string]bool `json:"accessControl,omitempty"`

	// avatar Url
	// Example: /avatar/8ea890a677d6a223c591a1beea6ea9d2
	AvatarURL string `json:"avatarUrl,omitempty"`

	// created at
	// Example: 2022-03-21T14:35:33Z
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// id
	// Example: 2
	ID int64 `json:"id,omitempty"`

	// is disabled
	// Example: false
	IsDisabled bool `json:"isDisabled,omitempty"`

	// is external
	// Example: false
	IsExternal bool `json:"isExternal,omitempty"`

	// login
	// Example: sa-grafana
	Login string `json:"login,omitempty"`

	// name
	// Example: test
	Name string `json:"name,omitempty"`

	// org Id
	// Example: 1
	OrgID int64 `json:"orgId,omitempty"`

	// required by
	// Example: grafana-app
	RequiredBy string `json:"requiredBy,omitempty"`

	// role
	// Example: Editor
	Role string `json:"role,omitempty"`

	// teams
	// Example: []
	Teams []string `json:"teams"`

	// tokens
	Tokens int64 `json:"tokens,omitempty"`

	// uid
	// Example: fe1xejlha91xce
	UID string `json:"uid,omitempty"`

	// updated at
	// Example: 2022-03-21T14:35:33Z
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this service account profile DTO
func (m *ServiceAccountProfileDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceAccountProfileDTO) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceAccountProfileDTO) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service account profile DTO based on context it is used
func (m *ServiceAccountProfileDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceAccountProfileDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceAccountProfileDTO) UnmarshalBinary(b []byte) error {
	var res ServiceAccountProfileDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
