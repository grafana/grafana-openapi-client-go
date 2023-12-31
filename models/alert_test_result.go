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

// AlertTestResult alert test result
//
// swagger:model AlertTestResult
type AlertTestResult struct {

	// condition evals
	ConditionEvals string `json:"conditionEvals,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// firing
	Firing bool `json:"firing,omitempty"`

	// logs
	Logs []*AlertTestResultLog `json:"logs"`

	// matches
	Matches []*EvalMatch `json:"matches"`

	// state
	State AlertStateType `json:"state,omitempty"`

	// time ms
	TimeMs string `json:"timeMs,omitempty"`
}

// Validate validates this alert test result
func (m *AlertTestResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertTestResult) validateLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertTestResult) validateMatches(formats strfmt.Registry) error {
	if swag.IsZero(m.Matches) { // not required
		return nil
	}

	for i := 0; i < len(m.Matches); i++ {
		if swag.IsZero(m.Matches[i]) { // not required
			continue
		}

		if m.Matches[i] != nil {
			if err := m.Matches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matches" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertTestResult) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

// ContextValidate validate this alert test result based on the context it is used
func (m *AlertTestResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMatches(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertTestResult) contextValidateLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Logs); i++ {

		if m.Logs[i] != nil {

			if swag.IsZero(m.Logs[i]) { // not required
				return nil
			}

			if err := m.Logs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertTestResult) contextValidateMatches(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Matches); i++ {

		if m.Matches[i] != nil {

			if swag.IsZero(m.Matches[i]) { // not required
				return nil
			}

			if err := m.Matches[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matches" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertTestResult) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertTestResult) UnmarshalBinary(b []byte) error {
	var res AlertTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
