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
	"github.com/grafana/grafana-openapi-client-go/pkg/custom_models"
)

// EmbeddedContactPoint EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
//
// swagger:model EmbeddedContactPoint
type EmbeddedContactPoint struct {

	// disable resolve message
	// Example: false
	DisableResolveMessage bool `json:"disableResolveMessage,omitempty"`

	// Name is used as grouping key in the UI. Contact points with the
	// same name will be grouped in the UI.
	// Example: webhook_1
	Name string `json:"name,omitempty"`

	// provenance
	// Read Only: true
	Provenance string `json:"provenance,omitempty"`

	// settings
	// Required: true
	Settings *custom_models.JSON `json:"settings"`

	// type
	// Example: webhook
	// Required: true
	// Enum: [alertmanager  dingding  discord  email  googlechat  kafka  line  opsgenie  pagerduty  pushover  sensugo  slack  teams  telegram  threema  victorops  webhook  wecom]
	Type *string `json:"type"`

	// UID is the unique identifier of the contact point. The UID can be
	// set by the user.
	// Example: my_external_reference
	// Max Length: 40
	// Min Length: 1
	// Pattern: ^[a-zA-Z0-9\-\_]+$
	UID string `json:"uid,omitempty"`
}

// Validate validates this embedded contact point
func (m *EmbeddedContactPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmbeddedContactPoint) validateSettings(formats strfmt.Registry) error {

	if err := validate.Required("settings", "body", m.Settings); err != nil {
		return err
	}

	return nil
}

var embeddedContactPointTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alertmanager"," dingding"," discord"," email"," googlechat"," kafka"," line"," opsgenie"," pagerduty"," pushover"," sensugo"," slack"," teams"," telegram"," threema"," victorops"," webhook"," wecom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		embeddedContactPointTypeTypePropEnum = append(embeddedContactPointTypeTypePropEnum, v)
	}
}

const (

	// EmbeddedContactPointTypeAlertmanager captures enum value "alertmanager"
	EmbeddedContactPointTypeAlertmanager string = "alertmanager"

	// EmbeddedContactPointTypeDingding captures enum value " dingding"
	EmbeddedContactPointTypeDingding string = " dingding"

	// EmbeddedContactPointTypeDiscord captures enum value " discord"
	EmbeddedContactPointTypeDiscord string = " discord"

	// EmbeddedContactPointTypeEmail captures enum value " email"
	EmbeddedContactPointTypeEmail string = " email"

	// EmbeddedContactPointTypeGooglechat captures enum value " googlechat"
	EmbeddedContactPointTypeGooglechat string = " googlechat"

	// EmbeddedContactPointTypeKafka captures enum value " kafka"
	EmbeddedContactPointTypeKafka string = " kafka"

	// EmbeddedContactPointTypeLine captures enum value " line"
	EmbeddedContactPointTypeLine string = " line"

	// EmbeddedContactPointTypeOpsgenie captures enum value " opsgenie"
	EmbeddedContactPointTypeOpsgenie string = " opsgenie"

	// EmbeddedContactPointTypePagerduty captures enum value " pagerduty"
	EmbeddedContactPointTypePagerduty string = " pagerduty"

	// EmbeddedContactPointTypePushover captures enum value " pushover"
	EmbeddedContactPointTypePushover string = " pushover"

	// EmbeddedContactPointTypeSensugo captures enum value " sensugo"
	EmbeddedContactPointTypeSensugo string = " sensugo"

	// EmbeddedContactPointTypeSlack captures enum value " slack"
	EmbeddedContactPointTypeSlack string = " slack"

	// EmbeddedContactPointTypeTeams captures enum value " teams"
	EmbeddedContactPointTypeTeams string = " teams"

	// EmbeddedContactPointTypeTelegram captures enum value " telegram"
	EmbeddedContactPointTypeTelegram string = " telegram"

	// EmbeddedContactPointTypeThreema captures enum value " threema"
	EmbeddedContactPointTypeThreema string = " threema"

	// EmbeddedContactPointTypeVictorops captures enum value " victorops"
	EmbeddedContactPointTypeVictorops string = " victorops"

	// EmbeddedContactPointTypeWebhook captures enum value " webhook"
	EmbeddedContactPointTypeWebhook string = " webhook"

	// EmbeddedContactPointTypeWecom captures enum value " wecom"
	EmbeddedContactPointTypeWecom string = " wecom"
)

// prop value enum
func (m *EmbeddedContactPoint) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, embeddedContactPointTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmbeddedContactPoint) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *EmbeddedContactPoint) validateUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UID) { // not required
		return nil
	}

	if err := validate.MinLength("uid", "body", m.UID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("uid", "body", m.UID, 40); err != nil {
		return err
	}

	if err := validate.Pattern("uid", "body", m.UID, `^[a-zA-Z0-9\-\_]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this embedded contact point based on the context it is used
func (m *EmbeddedContactPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProvenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmbeddedContactPoint) contextValidateProvenance(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "provenance", "body", string(m.Provenance)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmbeddedContactPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmbeddedContactPoint) UnmarshalBinary(b []byte) error {
	var res EmbeddedContactPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
