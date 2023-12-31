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

// PostableExtendedRuleNode postable extended rule node
//
// swagger:model PostableExtendedRuleNode
type PostableExtendedRuleNode struct {

	// alert
	Alert string `json:"alert,omitempty"`

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// expr
	Expr string `json:"expr,omitempty"`

	// for
	For string `json:"for,omitempty"`

	// grafana alert
	GrafanaAlert *PostableGrafanaRule `json:"grafana_alert,omitempty"`

	// keep firing for
	KeepFiringFor string `json:"keep_firing_for,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// record
	Record string `json:"record,omitempty"`
}

// Validate validates this postable extended rule node
func (m *PostableExtendedRuleNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrafanaAlert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableExtendedRuleNode) validateGrafanaAlert(formats strfmt.Registry) error {
	if swag.IsZero(m.GrafanaAlert) { // not required
		return nil
	}

	if m.GrafanaAlert != nil {
		if err := m.GrafanaAlert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grafana_alert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grafana_alert")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this postable extended rule node based on the context it is used
func (m *PostableExtendedRuleNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGrafanaAlert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableExtendedRuleNode) contextValidateGrafanaAlert(ctx context.Context, formats strfmt.Registry) error {

	if m.GrafanaAlert != nil {

		if swag.IsZero(m.GrafanaAlert) { // not required
			return nil
		}

		if err := m.GrafanaAlert.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grafana_alert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grafana_alert")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostableExtendedRuleNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostableExtendedRuleNode) UnmarshalBinary(b []byte) error {
	var res PostableExtendedRuleNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
