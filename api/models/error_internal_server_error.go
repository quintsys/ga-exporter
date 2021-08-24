// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorInternalServerError error internal server error
//
// swagger:model error_internal_server_error
type ErrorInternalServerError struct {

	// code
	// Example: 500
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Internal Server Error
	Message string `json:"message,omitempty"`
}

// Validate validates this error internal server error
func (m *ErrorInternalServerError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error internal server error based on context it is used
func (m *ErrorInternalServerError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorInternalServerError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorInternalServerError) UnmarshalBinary(b []byte) error {
	var res ErrorInternalServerError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}