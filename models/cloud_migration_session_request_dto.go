// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudMigrationSessionRequestDTO cloud migration session request DTO
//
// swagger:model CloudMigrationSessionRequestDTO
type CloudMigrationSessionRequestDTO struct {

	// auth token
	AuthToken string `json:"authToken,omitempty"`
}

// Validate validates this cloud migration session request DTO
func (m *CloudMigrationSessionRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud migration session request DTO based on context it is used
func (m *CloudMigrationSessionRequestDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudMigrationSessionRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudMigrationSessionRequestDTO) UnmarshalBinary(b []byte) error {
	var res CloudMigrationSessionRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
