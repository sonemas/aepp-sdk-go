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

// OracleRegisterTx oracle register tx
// swagger:model OracleRegisterTx
type OracleRegisterTx struct {

	// abi version
	AbiVersion uint16 `json:"abi_version,omitempty"`

	// account id
	// Required: true
	AccountID *string `json:"account_id"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// oracle ttl
	// Required: true
	OracleTTL *TTL `json:"oracle_ttl"`

	// query fee
	// Required: true
	QueryFee utils.BigInt `json:"query_fee"`

	// query format
	// Required: true
	QueryFormat *string `json:"query_format"`

	// response format
	// Required: true
	ResponseFormat *string `json:"response_format"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this oracle register tx
func (m *OracleRegisterTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleRegisterTx) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *OracleRegisterTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *OracleRegisterTx) validateOracleTTL(formats strfmt.Registry) error {

	if err := validate.Required("oracle_ttl", "body", m.OracleTTL); err != nil {
		return err
	}

	if m.OracleTTL != nil {
		if err := m.OracleTTL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracle_ttl")
			}
			return err
		}
	}

	return nil
}

func (m *OracleRegisterTx) validateQueryFee(formats strfmt.Registry) error {

	if err := m.QueryFee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("query_fee")
		}
		return err
	}

	return nil
}

func (m *OracleRegisterTx) validateQueryFormat(formats strfmt.Registry) error {

	if err := validate.Required("query_format", "body", m.QueryFormat); err != nil {
		return err
	}

	return nil
}

func (m *OracleRegisterTx) validateResponseFormat(formats strfmt.Registry) error {

	if err := validate.Required("response_format", "body", m.ResponseFormat); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleRegisterTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleRegisterTx) UnmarshalBinary(b []byte) error {
	var res OracleRegisterTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
