// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorTooManyRequests error too many requests
//
// swagger:model error_too_many_requests
type ErrorTooManyRequests struct {

	// code
	// Example: 429
	Code int64 `json:"code,omitempty"`

	// message
	// Example: too many requests
	Message string `json:"message,omitempty"`
}

// Validate validates this error too many requests
func (m *ErrorTooManyRequests) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error too many requests based on context it is used
func (m *ErrorTooManyRequests) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorTooManyRequests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorTooManyRequests) UnmarshalBinary(b []byte) error {
	var res ErrorTooManyRequests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
