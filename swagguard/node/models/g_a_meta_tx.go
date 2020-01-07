// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v8/utils"
)

// GAMetaTx g a meta tx
// swagger:model GAMetaTx
type GAMetaTx struct {

	// ABI version
	// Required: true
	AbiVersion *uint16 `json:"abi_version"`

	// Contract authorization function call data
	// Required: true
	AuthData *string `json:"auth_data"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// Account owner pub_key
	// Required: true
	GaID *string `json:"ga_id"`

	// gas
	// Required: true
	Gas *uint64 `json:"gas"`

	// gas price
	// Required: true
	GasPrice utils.BigInt `json:"gas_price"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`

	// Enclosed signed transaction
	// Required: true
	Tx *GenericSignedTx `json:"tx"`
}

// Validate validates this g a meta tx
func (m *GAMetaTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbiVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGaID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GAMetaTx) validateAbiVersion(formats strfmt.Registry) error {

	if err := validate.Required("abi_version", "body", m.AbiVersion); err != nil {
		return err
	}

	return nil
}

func (m *GAMetaTx) validateAuthData(formats strfmt.Registry) error {

	if err := validate.Required("auth_data", "body", m.AuthData); err != nil {
		return err
	}

	return nil
}

func (m *GAMetaTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *GAMetaTx) validateGaID(formats strfmt.Registry) error {

	if err := validate.Required("ga_id", "body", m.GaID); err != nil {
		return err
	}

	return nil
}

func (m *GAMetaTx) validateGas(formats strfmt.Registry) error {

	if err := validate.Required("gas", "body", m.Gas); err != nil {
		return err
	}

	return nil
}

func (m *GAMetaTx) validateGasPrice(formats strfmt.Registry) error {

	if err := m.GasPrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gas_price")
		}
		return err
	}

	return nil
}

func (m *GAMetaTx) validateTx(formats strfmt.Registry) error {

	if err := validate.Required("tx", "body", m.Tx); err != nil {
		return err
	}

	if m.Tx != nil {
		if err := m.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GAMetaTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GAMetaTx) UnmarshalBinary(b []byte) error {
	var res GAMetaTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
