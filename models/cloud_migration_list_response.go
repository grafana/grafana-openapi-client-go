// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudMigrationListResponse cloud migration list response
//
// swagger:model CloudMigrationListResponse
type CloudMigrationListResponse struct {

	// migrations
	Migrations []*CloudMigrationResponse `json:"migrations"`
}

// Validate validates this cloud migration list response
func (m *CloudMigrationListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMigrations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudMigrationListResponse) validateMigrations(formats strfmt.Registry) error {
	if swag.IsZero(m.Migrations) { // not required
		return nil
	}

	for i := 0; i < len(m.Migrations); i++ {
		if swag.IsZero(m.Migrations[i]) { // not required
			continue
		}

		if m.Migrations[i] != nil {
			if err := m.Migrations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("migrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("migrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cloud migration list response based on the context it is used
func (m *CloudMigrationListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMigrations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudMigrationListResponse) contextValidateMigrations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Migrations); i++ {

		if m.Migrations[i] != nil {

			if swag.IsZero(m.Migrations[i]) { // not required
				return nil
			}

			if err := m.Migrations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("migrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("migrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudMigrationListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudMigrationListResponse) UnmarshalBinary(b []byte) error {
	var res CloudMigrationListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}