// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Format1ActionDefinition format1 action definition
// swagger:model Format1ActionDefinition
type Format1ActionDefinition struct {

	// action parameters
	// Required: true
	ActionParameters []*ActionParameters `json:"ActionParameters"`

	// style ID
	// Required: true
	StyleID *int64 `json:"StyleID"`
}

// Validate validates this format1 action definition
func (m *Format1ActionDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStyleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Format1ActionDefinition) validateActionParameters(formats strfmt.Registry) error {

	if err := validate.Required("ActionParameters", "body", m.ActionParameters); err != nil {
		return err
	}

	for i := 0; i < len(m.ActionParameters); i++ {
		if swag.IsZero(m.ActionParameters[i]) { // not required
			continue
		}

		if m.ActionParameters[i] != nil {
			if err := m.ActionParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ActionParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Format1ActionDefinition) validateStyleID(formats strfmt.Registry) error {

	if err := validate.Required("StyleID", "body", m.StyleID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Format1ActionDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Format1ActionDefinition) UnmarshalBinary(b []byte) error {
	var res Format1ActionDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
