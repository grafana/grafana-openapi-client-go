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

// ContactPointExport ContactPointExport is the provisioned file export of alerting.ContactPointV1.
//
// swagger:model ContactPointExport
type ContactPointExport struct {

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID int64 `json:"orgId,omitempty"`

	// receivers
	Receivers []*ReceiverExport `json:"receivers"`
}

// Validate validates this contact point export
func (m *ContactPointExport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReceivers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactPointExport) validateReceivers(formats strfmt.Registry) error {
	if swag.IsZero(m.Receivers) { // not required
		return nil
	}

	for i := 0; i < len(m.Receivers); i++ {
		if swag.IsZero(m.Receivers[i]) { // not required
			continue
		}

		if m.Receivers[i] != nil {
			if err := m.Receivers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contact point export based on the context it is used
func (m *ContactPointExport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReceivers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactPointExport) contextValidateReceivers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Receivers); i++ {

		if m.Receivers[i] != nil {

			if swag.IsZero(m.Receivers[i]) { // not required
				return nil
			}

			if err := m.Receivers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactPointExport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactPointExport) UnmarshalBinary(b []byte) error {
	var res ContactPointExport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
