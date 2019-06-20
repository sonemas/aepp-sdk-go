// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/utils"
)

// OracleQueryTx oracle query tx
// swagger:model OracleQueryTx
type OracleQueryTx struct {

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// Sender nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// oracle id
	// Required: true
	OracleID *string `json:"oracle_id"`

	// query
	// Required: true
	Query *string `json:"query"`

	// query fee
	// Required: true
	QueryFee utils.BigInt `json:"query_fee"`

	// query ttl
	// Required: true
	QueryTTL *TTL `json:"query_ttl"`

	// response ttl
	// Required: true
	ResponseTTL *RelativeTTL `json:"response_ttl"`

	// sender id
	// Required: true
	SenderID *string `json:"sender_id"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this oracle query tx
func (m *OracleQueryTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleQueryTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *OracleQueryTx) validateOracleID(formats strfmt.Registry) error {

	if err := validate.Required("oracle_id", "body", m.OracleID); err != nil {
		return err
	}

	return nil
}

func (m *OracleQueryTx) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

func (m *OracleQueryTx) validateQueryFee(formats strfmt.Registry) error {

	if err := m.QueryFee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("query_fee")
		}
		return err
	}

	return nil
}

func (m *OracleQueryTx) validateQueryTTL(formats strfmt.Registry) error {

	if err := validate.Required("query_ttl", "body", m.QueryTTL); err != nil {
		return err
	}

	if m.QueryTTL != nil {
		if err := m.QueryTTL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_ttl")
			}
			return err
		}
	}

	return nil
}

func (m *OracleQueryTx) validateResponseTTL(formats strfmt.Registry) error {

	if err := validate.Required("response_ttl", "body", m.ResponseTTL); err != nil {
		return err
	}

	if m.ResponseTTL != nil {
		if err := m.ResponseTTL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response_ttl")
			}
			return err
		}
	}

	return nil
}

func (m *OracleQueryTx) validateSenderID(formats strfmt.Registry) error {

	if err := validate.Required("sender_id", "body", m.SenderID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleQueryTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleQueryTx) UnmarshalBinary(b []byte) error {
	var res OracleQueryTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}