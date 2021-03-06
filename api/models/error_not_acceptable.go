// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorNotAcceptable error not acceptable
//
// swagger:model error_not_acceptable
type ErrorNotAcceptable struct {

	// code
	// Example: 406
	Code int64 `json:"code,omitempty"`

	// message
	// Example: unsupported media type requested, only [application/json] are available
	Message string `json:"message,omitempty"`
}

// Validate validates this error not acceptable
func (m *ErrorNotAcceptable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error not acceptable based on context it is used
func (m *ErrorNotAcceptable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorNotAcceptable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorNotAcceptable) UnmarshalBinary(b []byte) error {
	var res ErrorNotAcceptable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
