// Code generated by go-swagger; DO NOT EDIT.

package clientmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SubscriptionData subscription data
// swagger:model SubscriptionData
type SubscriptionData struct {

	// endpoint
	Endpoint []string `json:"Endpoint"`

	// meid
	Meid string `json:"Meid,omitempty"`

	// subscription Id
	SubscriptionID int64 `json:"SubscriptionId,omitempty"`
}

// Validate validates this subscription data
func (m *SubscriptionData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionData) UnmarshalBinary(b []byte) error {
	var res SubscriptionData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
