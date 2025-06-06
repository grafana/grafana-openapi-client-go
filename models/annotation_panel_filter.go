// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnnotationPanelFilter annotation panel filter
//
// swagger:model AnnotationPanelFilter
type AnnotationPanelFilter struct {

	// Should the specified panels be included or excluded
	Exclude bool `json:"exclude,omitempty"`

	// Panel IDs that should be included or excluded
	Ids []uint8 `json:"ids"`
}

// Validate validates this annotation panel filter
func (m *AnnotationPanelFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this annotation panel filter based on context it is used
func (m *AnnotationPanelFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnnotationPanelFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnnotationPanelFilter) UnmarshalBinary(b []byte) error {
	var res AnnotationPanelFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
