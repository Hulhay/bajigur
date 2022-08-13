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

// Stores stores
//
// swagger:model Stores
type Stores struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamp;autoCreateTime"`

	// id
	ID int64 `json:"id,omitempty" gorm:"type:int primary key auto_increment"`

	// is deleted
	IsDeleted bool `json:"is_deleted,omitempty" gorm:"type:tinyint(1);default:0"`

	// owner
	Owner string `json:"owner,omitempty" gorm:"type:varchar(255)"`

	// store address
	StoreAddress string `json:"store_address,omitempty" gorm:"type:varchar(255)"`

	// store name
	StoreName string `json:"store_name,omitempty" gorm:"type:varchar(255)"`

	// store phone
	StorePhone string `json:"store_phone,omitempty" gorm:"type:varchar(15)"`

	// store photo
	StorePhoto string `json:"store_photo,omitempty" gorm:"type:varchar(255);default:sp0.jpg"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamp;autoUpdateTime"`
}

// Validate validates this stores
func (m *Stores) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
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

func (m *Stores) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Stores) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stores based on context it is used
func (m *Stores) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Stores) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Stores) UnmarshalBinary(b []byte) error {
	var res Stores
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
