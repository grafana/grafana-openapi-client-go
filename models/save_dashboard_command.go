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
	"github.com/grafana/grafana-openapi-client-go/pkg/custom_models"
)

// SaveDashboardCommand save dashboard command
//
// swagger:model SaveDashboardCommand
type SaveDashboardCommand struct {

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"UpdatedAt,omitempty"`

	// dashboard
	Dashboard custom_models.JSON `json:"dashboard,omitempty"`

	// folder Id
	FolderID int64 `json:"folderId,omitempty"`

	// folder Uid
	FolderUID string `json:"folderUid,omitempty"`

	// is folder
	IsFolder bool `json:"isFolder,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// overwrite
	Overwrite bool `json:"overwrite,omitempty"`

	// user Id
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this save dashboard command
func (m *SaveDashboardCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SaveDashboardCommand) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SaveDashboardCommand) validateDashboard(formats strfmt.Registry) error {
	if swag.IsZero(m.Dashboard) { // not required
		return nil
	}

	return nil
}

// ContextValidate validates this save dashboard command based on context it is used
func (m *SaveDashboardCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SaveDashboardCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SaveDashboardCommand) UnmarshalBinary(b []byte) error {
	var res SaveDashboardCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
