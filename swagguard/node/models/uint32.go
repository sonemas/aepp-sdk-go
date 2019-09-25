// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Uint32 uint32
// swagger:model UInt32
type Uint32 uint64

// Validate validates this uint32
func (m Uint32) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinimumInt("", "body", int64(m), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("", "body", int64(m), 4.294967295e+09, false); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
