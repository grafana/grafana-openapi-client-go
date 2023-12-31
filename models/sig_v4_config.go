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

// SigV4Config SigV4Config is the configuration for signing remote write requests with
// AWS's SigV4 verification process. Empty values will be retrieved using the
// AWS default credentials chain.
//
// swagger:model SigV4Config
type SigV4Config struct {

	// access key
	AccessKey string `json:"AccessKey,omitempty"`

	// profile
	Profile string `json:"Profile,omitempty"`

	// region
	Region string `json:"Region,omitempty"`

	// role a r n
	RoleARN string `json:"RoleARN,omitempty"`

	// secret key
	SecretKey Secret `json:"SecretKey,omitempty"`
}

// Validate validates this sig v4 config
func (m *SigV4Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigV4Config) validateSecretKey(formats strfmt.Registry) error {
	if swag.IsZero(m.SecretKey) { // not required
		return nil
	}

	if err := m.SecretKey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SecretKey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SecretKey")
		}
		return err
	}

	return nil
}

// ContextValidate validate this sig v4 config based on the context it is used
func (m *SigV4Config) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecretKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigV4Config) contextValidateSecretKey(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SecretKey) { // not required
		return nil
	}

	if err := m.SecretKey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SecretKey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SecretKey")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SigV4Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigV4Config) UnmarshalBinary(b []byte) error {
	var res SigV4Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
