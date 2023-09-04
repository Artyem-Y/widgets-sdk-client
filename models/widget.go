// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Widget widget
//
// swagger:model widget
type Widget struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this widget
func (m *Widget) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this widget based on context it is used
func (m *Widget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Widget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Widget) UnmarshalBinary(b []byte) error {
	var res Widget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}