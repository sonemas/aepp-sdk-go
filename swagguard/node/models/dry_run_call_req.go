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

// DryRunCallReq dry run call req
// swagger:model DryRunCallReq
type DryRunCallReq struct {

	// abi version
	AbiVersion uint16 `json:"abi_version,omitempty"`

	// amount
	Amount utils.BigInt `json:"amount,omitempty"`

	// calldata
	// Required: true
	Calldata *string `json:"calldata"`

	// caller
	Caller string `json:"caller,omitempty"`

	// context
	Context *DryRunCallContext `json:"context,omitempty"`

	// contract
	// Required: true
	Contract *string `json:"contract"`

	// gas
	Gas utils.BigInt `json:"gas,omitempty"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`
}

// Validate validates this dry run call req
func (m *DryRunCallReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalldata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DryRunCallReq) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *DryRunCallReq) validateCalldata(formats strfmt.Registry) error {

	if err := validate.Required("calldata", "body", m.Calldata); err != nil {
		return err
	}

	return nil
}

func (m *DryRunCallReq) validateContext(formats strfmt.Registry) error {

	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

func (m *DryRunCallReq) validateContract(formats strfmt.Registry) error {

	if err := validate.Required("contract", "body", m.Contract); err != nil {
		return err
	}

	return nil
}

func (m *DryRunCallReq) validateGas(formats strfmt.Registry) error {

	if swag.IsZero(m.Gas) { // not required
		return nil
	}

	if err := m.Gas.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gas")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DryRunCallReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DryRunCallReq) UnmarshalBinary(b []byte) error {
	var res DryRunCallReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
