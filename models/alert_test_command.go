// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/grafana/grafana-openapi-client-go/pkg/custom_models"
)

// AlertTestCommand alert test command
//
// swagger:model AlertTestCommand
type AlertTestCommand struct {

	// dashboard
	Dashboard custom_models.JSON `json:"dashboard,omitempty"`

	// panel Id
	PanelID int64 `json:"panelId,omitempty"`
}

// Validate validates this alert test command
func (m *AlertTestCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertTestCommand) validateDashboard(formats strfmt.Registry) error {
	if swag.IsZero(m.Dashboard) { // not required
		return nil
	}

	return nil
}

// ContextValidate validates this alert test command based on context it is used
func (m *AlertTestCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertTestCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertTestCommand) UnmarshalBinary(b []byte) error {
	var res AlertTestCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
