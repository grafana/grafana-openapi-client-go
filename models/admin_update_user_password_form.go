// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminUpdateUserPasswordForm admin update user password form
//
// swagger:model AdminUpdateUserPasswordForm
type AdminUpdateUserPasswordForm struct {

	// password
	Password Password `json:"password,omitempty"`
}

// Validate validates this admin update user password form
func (m *AdminUpdateUserPasswordForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminUpdateUserPasswordForm) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := m.Password.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("password")
		}
		return err
	}

	return nil
}

// ContextValidate validate this admin update user password form based on the context it is used
func (m *AdminUpdateUserPasswordForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePassword(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminUpdateUserPasswordForm) contextValidatePassword(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := m.Password.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("password")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdminUpdateUserPasswordForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminUpdateUserPasswordForm) UnmarshalBinary(b []byte) error {
	var res AdminUpdateUserPasswordForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
