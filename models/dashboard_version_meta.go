// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DashboardVersionMeta DashboardVersionMeta extends the DashboardVersionDTO with the names
// associated with the UserIds, overriding the field with the same name from
// the DashboardVersionDTO model.
//
// swagger:model DashboardVersionMeta
type DashboardVersionMeta struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// dashboard Id
	DashboardID int64 `json:"dashboardId,omitempty"`

	// data
	Data JSON `json:"data,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// parent version
	ParentVersion int64 `json:"parentVersion,omitempty"`

	// restored from
	RestoredFrom int64 `json:"restoredFrom,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this dashboard version meta
func (m *DashboardVersionMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardVersionMeta) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dashboard version meta based on context it is used
func (m *DashboardVersionMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DashboardVersionMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardVersionMeta) UnmarshalBinary(b []byte) error {
	var res DashboardVersionMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
