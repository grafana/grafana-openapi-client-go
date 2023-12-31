// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnnotationTarget TODO: this should be a regular DataQuery that depends on the selected dashboard
// these match the properties of the "grafana" datasouce that is default in most dashboards
//
// swagger:model AnnotationTarget
type AnnotationTarget struct {

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Limit int64 `json:"limit,omitempty"`

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	MatchAny bool `json:"matchAny,omitempty"`

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Tags []string `json:"tags"`

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Type string `json:"type,omitempty"`
}

// Validate validates this annotation target
func (m *AnnotationTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this annotation target based on context it is used
func (m *AnnotationTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnnotationTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnnotationTarget) UnmarshalBinary(b []byte) error {
	var res AnnotationTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
