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

// Users users
//
// swagger:model Users
type Users struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamp;autoCreateTime"`

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty" gorm:"type:varchar(255)"`

	// full name
	FullName string `json:"full_name,omitempty" gorm:"type:varchar(255)"`

	// id
	ID int64 `json:"id,omitempty" gorm:"type:int primary key auto_increment"`

	// is login
	IsLogin bool `json:"is_login,omitempty" gorm:"type:tinyint(1);default:0"`

	// password
	Password string `json:"password,omitempty" gorm:"type:varchar"`

	// unique id
	UniqueID string `json:"unique_id,omitempty" gorm:"type:varchar(255)"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamp;autoUpdateTime"`

	// user photo
	UserPhoto string `json:"user_photo,omitempty" gorm:"type:varchar(255);default:up0.jpg"`

	// username
	Username string `json:"username,omitempty" gorm:"type:varchar(255)"`
}

// Validate validates this users
func (m *Users) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Users) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Users) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Users) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this users based on context it is used
func (m *Users) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Users) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Users) UnmarshalBinary(b []byte) error {
	var res Users
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
