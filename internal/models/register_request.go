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

// RegisterRequest register request
//
// swagger:model RegisterRequest
type RegisterRequest struct {

	// email
	// Example: hulhay@gmail.com
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// first name
	// Required: true
	// Min Length: 3
	FirstName *string `json:"first_name"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// last name
	// Min Length: 3
	LastName string `json:"last_name,omitempty"`

	// password
	// Required: true
	// Min Length: 6
	Password *string `json:"password"`

	// role
	Role string `json:"role,omitempty"`

	// unique id
	UniqueID string `json:"unique_id,omitempty"`

	// username
	// Example: Hulhay169
	// Required: true
	// Min Length: 5
	Username *string `json:"username"`
}

// Validate validates this register request
func (m *RegisterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterRequest) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RegisterRequest) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	if err := validate.MinLength("first_name", "body", *m.FirstName, 3); err != nil {
		return err
	}

	return nil
}

func (m *RegisterRequest) validateLastName(formats strfmt.Registry) error {
	if swag.IsZero(m.LastName) { // not required
		return nil
	}

	if err := validate.MinLength("last_name", "body", m.LastName, 3); err != nil {
		return err
	}

	return nil
}

func (m *RegisterRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", *m.Password, 6); err != nil {
		return err
	}

	return nil
}

func (m *RegisterRequest) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 5); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this register request based on context it is used
func (m *RegisterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterRequest) UnmarshalBinary(b []byte) error {
	var res RegisterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
