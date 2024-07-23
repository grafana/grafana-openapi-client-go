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

// MigrateDataResponseItemDTO migrate data response item DTO
//
// swagger:model MigrateDataResponseItemDTO
type MigrateDataResponseItemDTO struct {

	// error
	Error string `json:"error,omitempty"`

	// ref Id
	// Required: true
	RefID *string `json:"refId"`

	// status
	// Required: true
	// Enum: [OK ERROR PENDING UNKNOWN]
	Status *string `json:"status"`

	// type
	// Required: true
	// Enum: [DASHBOARD DATASOURCE FOLDER]
	Type *string `json:"type"`
}

// Validate validates this migrate data response item DTO
func (m *MigrateDataResponseItemDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRefID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateDataResponseItemDTO) validateRefID(formats strfmt.Registry) error {

	if err := validate.Required("refId", "body", m.RefID); err != nil {
		return err
	}

	return nil
}

var migrateDataResponseItemDtoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","ERROR","PENDING","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		migrateDataResponseItemDtoTypeStatusPropEnum = append(migrateDataResponseItemDtoTypeStatusPropEnum, v)
	}
}

const (

	// MigrateDataResponseItemDTOStatusOK captures enum value "OK"
	MigrateDataResponseItemDTOStatusOK string = "OK"

	// MigrateDataResponseItemDTOStatusERROR captures enum value "ERROR"
	MigrateDataResponseItemDTOStatusERROR string = "ERROR"

	// MigrateDataResponseItemDTOStatusPENDING captures enum value "PENDING"
	MigrateDataResponseItemDTOStatusPENDING string = "PENDING"

	// MigrateDataResponseItemDTOStatusUNKNOWN captures enum value "UNKNOWN"
	MigrateDataResponseItemDTOStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *MigrateDataResponseItemDTO) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, migrateDataResponseItemDtoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MigrateDataResponseItemDTO) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var migrateDataResponseItemDtoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DASHBOARD","DATASOURCE","FOLDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		migrateDataResponseItemDtoTypeTypePropEnum = append(migrateDataResponseItemDtoTypeTypePropEnum, v)
	}
}

const (

	// MigrateDataResponseItemDTOTypeDASHBOARD captures enum value "DASHBOARD"
	MigrateDataResponseItemDTOTypeDASHBOARD string = "DASHBOARD"

	// MigrateDataResponseItemDTOTypeDATASOURCE captures enum value "DATASOURCE"
	MigrateDataResponseItemDTOTypeDATASOURCE string = "DATASOURCE"

	// MigrateDataResponseItemDTOTypeFOLDER captures enum value "FOLDER"
	MigrateDataResponseItemDTOTypeFOLDER string = "FOLDER"
)

// prop value enum
func (m *MigrateDataResponseItemDTO) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, migrateDataResponseItemDtoTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MigrateDataResponseItemDTO) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this migrate data response item DTO based on context it is used
func (m *MigrateDataResponseItemDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MigrateDataResponseItemDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateDataResponseItemDTO) UnmarshalBinary(b []byte) error {
	var res MigrateDataResponseItemDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
