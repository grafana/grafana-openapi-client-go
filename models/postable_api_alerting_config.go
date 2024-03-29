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

// PostableAPIAlertingConfig nolint:revive
//
// swagger:model PostableApiAlertingConfig
type PostableAPIAlertingConfig struct {

	// global
	Global *GlobalConfig `json:"global,omitempty"`

	// inhibit rules
	InhibitRules []*InhibitRule `json:"inhibit_rules"`

	// MuteTimeIntervals is deprecated and will be removed before Alertmanager 1.0.
	MuteTimeIntervals []*MuteTimeInterval `json:"mute_time_intervals"`

	// Override with our superset receiver type
	Receivers []*PostableAPIReceiver `json:"receivers"`

	// route
	Route *Route `json:"route,omitempty"`

	// templates
	Templates []string `json:"templates"`

	// time intervals
	TimeIntervals []*TimeInterval `json:"time_intervals"`
}

// Validate validates this postable Api alerting config
func (m *PostableAPIAlertingConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInhibitRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMuteTimeIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeIntervals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableAPIAlertingConfig) validateGlobal(formats strfmt.Registry) error {
	if swag.IsZero(m.Global) { // not required
		return nil
	}

	if m.Global != nil {
		if err := m.Global.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global")
			}
			return err
		}
	}

	return nil
}

func (m *PostableAPIAlertingConfig) validateInhibitRules(formats strfmt.Registry) error {
	if swag.IsZero(m.InhibitRules) { // not required
		return nil
	}

	for i := 0; i < len(m.InhibitRules); i++ {
		if swag.IsZero(m.InhibitRules[i]) { // not required
			continue
		}

		if m.InhibitRules[i] != nil {
			if err := m.InhibitRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIAlertingConfig) validateMuteTimeIntervals(formats strfmt.Registry) error {
	if swag.IsZero(m.MuteTimeIntervals) { // not required
		return nil
	}

	for i := 0; i < len(m.MuteTimeIntervals); i++ {
		if swag.IsZero(m.MuteTimeIntervals[i]) { // not required
			continue
		}

		if m.MuteTimeIntervals[i] != nil {
			if err := m.MuteTimeIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIAlertingConfig) validateReceivers(formats strfmt.Registry) error {
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

func (m *PostableAPIAlertingConfig) validateRoute(formats strfmt.Registry) error {
	if swag.IsZero(m.Route) { // not required
		return nil
	}

	if m.Route != nil {
		if err := m.Route.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

func (m *PostableAPIAlertingConfig) validateTimeIntervals(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeIntervals) { // not required
		return nil
	}

	for i := 0; i < len(m.TimeIntervals); i++ {
		if swag.IsZero(m.TimeIntervals[i]) { // not required
			continue
		}

		if m.TimeIntervals[i] != nil {
			if err := m.TimeIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this postable Api alerting config based on the context it is used
func (m *PostableAPIAlertingConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGlobal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInhibitRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMuteTimeIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReceivers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableAPIAlertingConfig) contextValidateGlobal(ctx context.Context, formats strfmt.Registry) error {

	if m.Global != nil {

		if swag.IsZero(m.Global) { // not required
			return nil
		}

		if err := m.Global.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global")
			}
			return err
		}
	}

	return nil
}

func (m *PostableAPIAlertingConfig) contextValidateInhibitRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InhibitRules); i++ {

		if m.InhibitRules[i] != nil {

			if swag.IsZero(m.InhibitRules[i]) { // not required
				return nil
			}

			if err := m.InhibitRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIAlertingConfig) contextValidateMuteTimeIntervals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MuteTimeIntervals); i++ {

		if m.MuteTimeIntervals[i] != nil {

			if swag.IsZero(m.MuteTimeIntervals[i]) { // not required
				return nil
			}

			if err := m.MuteTimeIntervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIAlertingConfig) contextValidateReceivers(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PostableAPIAlertingConfig) contextValidateRoute(ctx context.Context, formats strfmt.Registry) error {

	if m.Route != nil {

		if swag.IsZero(m.Route) { // not required
			return nil
		}

		if err := m.Route.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

func (m *PostableAPIAlertingConfig) contextValidateTimeIntervals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TimeIntervals); i++ {

		if m.TimeIntervals[i] != nil {

			if swag.IsZero(m.TimeIntervals[i]) { // not required
				return nil
			}

			if err := m.TimeIntervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostableAPIAlertingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostableAPIAlertingConfig) UnmarshalBinary(b []byte) error {
	var res PostableAPIAlertingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
