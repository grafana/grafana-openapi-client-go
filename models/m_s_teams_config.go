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

// MSTeamsConfig m s teams config
//
// swagger:model MSTeamsConfig
type MSTeamsConfig struct {

	// http config
	HTTPConfig *HTTPClientConfig `json:"http_config,omitempty"`

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// webhook url
	WebhookURL *SecretURL `json:"webhook_url,omitempty"`

	// webhook url file
	WebhookURLFile string `json:"webhook_url_file,omitempty"`
}

// Validate validates this m s teams config
func (m *MSTeamsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MSTeamsConfig) validateHTTPConfig(formats strfmt.Registry) error {
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

func (m *MSTeamsConfig) validateWebhookURL(formats strfmt.Registry) error {
	if swag.IsZero(m.WebhookURL) { // not required
		return nil
	}

	if m.WebhookURL != nil {
		if err := m.WebhookURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_url")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this m s teams config based on the context it is used
func (m *MSTeamsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhookURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MSTeamsConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MSTeamsConfig) contextValidateWebhookURL(ctx context.Context, formats strfmt.Registry) error {

	if m.WebhookURL != nil {

		if swag.IsZero(m.WebhookURL) { // not required
			return nil
		}

		if err := m.WebhookURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_url")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MSTeamsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MSTeamsConfig) UnmarshalBinary(b []byte) error {
	var res MSTeamsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
