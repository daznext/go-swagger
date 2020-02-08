// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SocialID social id
// swagger:model social_id
type SocialID struct {

	// ssn
	// Required: true
	// Min Length: 11
	Ssn *string `json:"ssn"`
}

// Validate validates this social id
func (m *SocialID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSsn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SocialID) validateSsn(formats strfmt.Registry) error {

	if err := validate.Required("ssn", "body", m.Ssn); err != nil {
		return err
	}

	if err := validate.MinLength("ssn", "body", string(*m.Ssn), 11); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SocialID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SocialID) UnmarshalBinary(b []byte) error {
	var res SocialID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
