// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrometheusRule prometheus rule
//
// swagger:model PrometheusRule
type PrometheusRule struct {

	// alert
	Alert string `json:"Alert,omitempty"`

	// annotations
	Annotations map[string]string `json:"Annotations,omitempty"`

	// expr
	Expr string `json:"Expr,omitempty"`

	// for
	For string `json:"For,omitempty"`

	// keep firing for
	KeepFiringFor string `json:"KeepFiringFor,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// record
	Record string `json:"Record,omitempty"`
}

// Validate validates this prometheus rule
func (m *PrometheusRule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prometheus rule based on context it is used
func (m *PrometheusRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRule) UnmarshalBinary(b []byte) error {
	var res PrometheusRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
