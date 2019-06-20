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

// OracleQuery oracle query
// swagger:model OracleQuery
type OracleQuery struct {

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// id
	// Required: true
	ID *string `json:"id"`

	// oracle id
	// Required: true
	OracleID *string `json:"oracle_id"`

	// query
	// Required: true
	Query *string `json:"query"`

	// response
	// Required: true
	Response *string `json:"response"`

	// response ttl
	// Required: true
	ResponseTTL *TTL `json:"response_ttl"`

	// sender id
	// Required: true
	SenderID *string `json:"sender_id"`

	// sender nonce
	// Required: true
	SenderNonce *uint64 `json:"sender_nonce"`

	// ttl
	// Required: true
	TTL *uint64 `json:"ttl"`
}

// Validate validates this oracle query
func (m *OracleQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleQuery) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *OracleQuery) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *OracleQuery) validateOracleID(formats strfmt.Registry) error {

	if err := validate.Required("oracle_id", "body", m.OracleID); err != nil {
		return err
	}

	return nil
}

func (m *OracleQuery) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

func (m *OracleQuery) validateResponse(formats strfmt.Registry) error {

	if err := validate.Required("response", "body", m.Response); err != nil {
		return err
	}

	return nil
}

func (m *OracleQuery) validateResponseTTL(formats strfmt.Registry) error {

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

func (m *OracleQuery) validateSenderID(formats strfmt.Registry) error {

	if err := validate.Required("sender_id", "body", m.SenderID); err != nil {
		return err
	}

	return nil
}

func (m *OracleQuery) validateSenderNonce(formats strfmt.Registry) error {

	if err := validate.Required("sender_nonce", "body", m.SenderNonce); err != nil {
		return err
	}

	return nil
}

func (m *OracleQuery) validateTTL(formats strfmt.Registry) error {

	if err := validate.Required("ttl", "body", m.TTL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleQuery) UnmarshalBinary(b []byte) error {
	var res OracleQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}