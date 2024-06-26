// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MigrateDataResponseListDTO migrate data response list DTO
//
// swagger:model MigrateDataResponseListDTO
type MigrateDataResponseListDTO struct {

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this migrate data response list DTO
func (m *MigrateDataResponseListDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this migrate data response list DTO based on context it is used
func (m *MigrateDataResponseListDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MigrateDataResponseListDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateDataResponseListDTO) UnmarshalBinary(b []byte) error {
	var res MigrateDataResponseListDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
