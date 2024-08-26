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

// Authorization Authorization contains HTTP authorization credentials.
//
// swagger:model Authorization
type Authorization struct {

	// credentials
	Credentials Secret `json:"credentials,omitempty"`

	// credentials file
	CredentialsFile string `json:"credentials_file,omitempty"`

	// CredentialsRef is the name of the secret within the secret manager to use as credentials.
	CredentialsRef string `json:"credentials_ref,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this authorization
func (m *Authorization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Authorization) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if err := m.Credentials.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentials")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("credentials")
		}
		return err
	}

	return nil
}

// ContextValidate validate this authorization based on the context it is used
func (m *Authorization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Authorization) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentials")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("credentials")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Authorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Authorization) UnmarshalBinary(b []byte) error {
	var res Authorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
