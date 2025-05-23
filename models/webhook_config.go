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

// WebhookConfig WebhookConfig configures notifications via a generic webhook.
//
// swagger:model WebhookConfig
type WebhookConfig struct {

	// http config
	HTTPConfig *HTTPClientConfig `json:"http_config,omitempty"`

	// MaxAlerts is the maximum number of alerts to be sent per webhook message.
	// Alerts exceeding this threshold will be truncated. Setting this to 0
	// allows an unlimited number of alerts.
	MaxAlerts uint64 `json:"max_alerts,omitempty"`

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// timeout
	Timeout Duration `json:"timeout,omitempty"`

	// url
	URL *SecretURL `json:"url,omitempty"`

	// url file
	URLFile string `json:"url_file,omitempty"`
}

// Validate validates this webhook config
func (m *WebhookConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConfig) { // not required
		return nil
	}

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

func (m *WebhookConfig) validateTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if err := m.Timeout.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timeout")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("timeout")
		}
		return err
	}

	return nil
}

func (m *WebhookConfig) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if m.URL != nil {
		if err := m.URL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("url")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this webhook config based on the context it is used
func (m *WebhookConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeout(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPConfig != nil {

		if swag.IsZero(m.HTTPConfig) { // not required
			return nil
		}

		if err := m.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

func (m *WebhookConfig) contextValidateTimeout(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if err := m.Timeout.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timeout")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("timeout")
		}
		return err
	}

	return nil
}

func (m *WebhookConfig) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if m.URL != nil {

		if swag.IsZero(m.URL) { // not required
			return nil
		}

		if err := m.URL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("url")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebhookConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookConfig) UnmarshalBinary(b []byte) error {
	var res WebhookConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
