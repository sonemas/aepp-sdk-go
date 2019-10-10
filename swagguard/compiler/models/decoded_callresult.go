// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DecodedCallresult decoded callresult
// swagger:model DecodedCallresult
type DecodedCallresult struct {

	// function
	// Required: true
	Function *string `json:"function"`

	// result
	// Required: true
	Result interface{} `json:"result"`
}

// Validate validates this decoded callresult
func (m *DecodedCallresult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DecodedCallresult) validateFunction(formats strfmt.Registry) error {

	if err := validate.Required("function", "body", m.Function); err != nil {
		return err
	}

	return nil
}

func (m *DecodedCallresult) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DecodedCallresult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DecodedCallresult) UnmarshalBinary(b []byte) error {
	var res DecodedCallresult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
