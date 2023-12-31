// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetPermissionCommand set permission command
//
// swagger:model setPermissionCommand
type SetPermissionCommand struct {

	// permission
	Permission string `json:"permission,omitempty"`
}

// Validate validates this set permission command
func (m *SetPermissionCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set permission command based on context it is used
func (m *SetPermissionCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetPermissionCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetPermissionCommand) UnmarshalBinary(b []byte) error {
	var res SetPermissionCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
