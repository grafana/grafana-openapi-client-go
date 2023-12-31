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

// DeleteDataSourceByNameOKBody delete data source by name Ok body
//
// swagger:model deleteDataSourceByNameOkBody
type DeleteDataSourceByNameOKBody struct {

	// ID Identifier of the deleted data source.
	// Example: 65
	// Required: true
	ID *int64 `json:"id"`

	// Message Message of the deleted dashboard.
	// Example: Dashboard My Dashboard deleted
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete data source by name Ok body
func (m *DeleteDataSourceByNameOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteDataSourceByNameOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DeleteDataSourceByNameOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete data source by name Ok body based on context it is used
func (m *DeleteDataSourceByNameOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteDataSourceByNameOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteDataSourceByNameOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteDataSourceByNameOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
