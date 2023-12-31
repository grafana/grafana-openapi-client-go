// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestTemplatesErrorResult test templates error result
//
// swagger:model TestTemplatesErrorResult
type TestTemplatesErrorResult struct {

	// Kind of template error that occurred.
	// Enum: [invalid_template execution_error]
	Kind string `json:"kind,omitempty"`

	// Error message.
	Message string `json:"message,omitempty"`

	// Name of the associated template for this error. Will be empty if the Kind is "invalid_template".
	Name string `json:"name,omitempty"`
}

// Validate validates this test templates error result
func (m *TestTemplatesErrorResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var testTemplatesErrorResultTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["invalid_template","execution_error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		testTemplatesErrorResultTypeKindPropEnum = append(testTemplatesErrorResultTypeKindPropEnum, v)
	}
}

const (

	// TestTemplatesErrorResultKindInvalidTemplate captures enum value "invalid_template"
	TestTemplatesErrorResultKindInvalidTemplate string = "invalid_template"

	// TestTemplatesErrorResultKindExecutionError captures enum value "execution_error"
	TestTemplatesErrorResultKindExecutionError string = "execution_error"
)

// prop value enum
func (m *TestTemplatesErrorResult) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, testTemplatesErrorResultTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TestTemplatesErrorResult) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this test templates error result based on context it is used
func (m *TestTemplatesErrorResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestTemplatesErrorResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestTemplatesErrorResult) UnmarshalBinary(b []byte) error {
	var res TestTemplatesErrorResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
