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

// SSOSettings s s o settings
//
// swagger:model SSOSettings
type SSOSettings struct {

	// id
	ID string `json:"id,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// settings
	Settings interface{} `json:"settings,omitempty"`

	// source
	Source SettingsSource `json:"source,omitempty"`
}

// Validate validates this s s o settings
func (m *SSOSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSOSettings) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if err := m.Source.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source")
		}
		return err
	}

	return nil
}

// ContextValidate validate this s s o settings based on the context it is used
func (m *SSOSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSOSettings) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if err := m.Source.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSOSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSOSettings) UnmarshalBinary(b []byte) error {
	var res SSOSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
