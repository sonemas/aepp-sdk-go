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

// OracleRespondTx oracle respond tx
// swagger:model OracleRespondTx
type OracleRespondTx struct {

	// fee
	// Required: true
	Fee *int64 `json:"fee"`

	// Oracle nonce
	Nonce int64 `json:"nonce,omitempty"`

	// oracle id
	// Required: true
	OracleID EncodedHash `json:"oracle_id"`

	// query id
	// Required: true
	QueryID EncodedHash `json:"query_id"`

	// response
	// Required: true
	Response *string `json:"response"`

	// ttl
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this oracle respond tx
func (m *OracleRespondTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleRespondTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	return nil
}

func (m *OracleRespondTx) validateOracleID(formats strfmt.Registry) error {

	if err := m.OracleID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("oracle_id")
		}
		return err
	}

	return nil
}

func (m *OracleRespondTx) validateQueryID(formats strfmt.Registry) error {

	if err := m.QueryID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("query_id")
		}
		return err
	}

	return nil
}

func (m *OracleRespondTx) validateResponse(formats strfmt.Registry) error {

	if err := validate.Required("response", "body", m.Response); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleRespondTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleRespondTx) UnmarshalBinary(b []byte) error {
	var res OracleRespondTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
