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

// ContractCallTx contract call tx
// swagger:model ContractCallTx
type ContractCallTx struct {

  // Amount
  // Required: true
  // Minimum: 0
  Amount *int64 `json:"amount"`

  // Contract call data
  // Required: true
  CallData EncodedByteArray `json:"call_data"`

  // Contract caller pub_key
  // Required: true
  CallerID EncodedHash `json:"caller_id"`

  // Contract's pub_key
  // Required: true
  ContractID EncodedHash `json:"contract_id"`

  // Transaction fee
  // Required: true
  // Minimum: 0
  Fee *int64 `json:"fee"`

  // Contract gas
  // Required: true
  // Minimum: 0
  Gas *int64 `json:"gas"`

  // Gas price
  // Required: true
  // Minimum: 0
  GasPrice *int64 `json:"gas_price"`

  // Caller's nonce
  Nonce int64 `json:"nonce,omitempty"`

  // Transaction TTL
  // Minimum: 0
  TTL *int64 `json:"ttl,omitempty"`

  // Virtual machine's version
  // Required: true
  // Maximum: 255
  // Minimum: 0
  VMVersion *string `json:"vm_version"`
}

// Validate validates this contract call tx
func (m *ContractCallTx) Validate(formats strfmt.Registry) error {
  var res []error

  if err := m.validateAmount(formats); err != nil {
    res = append(res, err)
  }

  if err := m.validateCallData(formats); err != nil {
    res = append(res, err)
  }

  if err := m.validateCallerID(formats); err != nil {
    res = append(res, err)
  }

  if err := m.validateContractID(formats); err != nil {
    res = append(res, err)
  }

  if err := m.validateFee(formats); err != nil {
    res = append(res, err)
  }

  if err := m.validateGas(formats); err != nil {
    res = append(res, err)
  }

  if err := m.validateGasPrice(formats); err != nil {
    res = append(res, err)
  }

  if err := m.validateTTL(formats); err != nil {
    res = append(res, err)
  }

  if err := m.validateVMVersion(formats); err != nil {
    res = append(res, err)
  }

  if len(res) > 0 {
    return errors.CompositeValidationError(res...)
  }
  return nil
}

func (m *ContractCallTx) validateAmount(formats strfmt.Registry) error {

  if err := validate.Required("amount", "body", m.Amount); err != nil {
    return err
  }

  if err := validate.MinimumInt("amount", "body", int64(*m.Amount), 0, false); err != nil {
    return err
  }

  return nil
}

func (m *ContractCallTx) validateCallData(formats strfmt.Registry) error {

  if err := m.CallData.Validate(formats); err != nil {
    if ve, ok := err.(*errors.Validation); ok {
      return ve.ValidateName("call_data")
    }
    return err
  }

  return nil
}

func (m *ContractCallTx) validateCallerID(formats strfmt.Registry) error {

  if err := m.CallerID.Validate(formats); err != nil {
    if ve, ok := err.(*errors.Validation); ok {
      return ve.ValidateName("caller_id")
    }
    return err
  }

  return nil
}

func (m *ContractCallTx) validateContractID(formats strfmt.Registry) error {

  if err := m.ContractID.Validate(formats); err != nil {
    if ve, ok := err.(*errors.Validation); ok {
      return ve.ValidateName("contract_id")
    }
    return err
  }

  return nil
}

func (m *ContractCallTx) validateFee(formats strfmt.Registry) error {

  if err := validate.Required("fee", "body", m.Fee); err != nil {
    return err
  }

  if err := validate.MinimumInt("fee", "body", int64(*m.Fee), 0, false); err != nil {
    return err
  }

  return nil
}

func (m *ContractCallTx) validateGas(formats strfmt.Registry) error {

  if err := validate.Required("gas", "body", m.Gas); err != nil {
    return err
  }

  if err := validate.MinimumInt("gas", "body", int64(*m.Gas), 0, false); err != nil {
    return err
  }

  return nil
}

func (m *ContractCallTx) validateGasPrice(formats strfmt.Registry) error {

  if err := validate.Required("gas_price", "body", m.GasPrice); err != nil {
    return err
  }

  if err := validate.MinimumInt("gas_price", "body", int64(*m.GasPrice), 0, false); err != nil {
    return err
  }

  return nil
}

func (m *ContractCallTx) validateTTL(formats strfmt.Registry) error {

  if swag.IsZero(m.TTL) { // not required
    return nil
  }

  if err := validate.MinimumInt("ttl", "body", int64(*m.TTL), 0, false); err != nil {
    return err
  }

  return nil
}

func (m *ContractCallTx) validateVMVersion(formats strfmt.Registry) error {

  if err := validate.Required("vm_version", "body", m.VMVersion); err != nil {
    return err
  }

  // if err := validate.MinimumInt("vm_version", "body", int64(*m.VMVersion), 0, false); err != nil {
  //   return err
  // }

  // if err := validate.MaximumInt("vm_version", "body", int64(*m.VMVersion), 255, false); err != nil {
  //   return err
  // }

  return nil
}

// MarshalBinary interface implementation
func (m *ContractCallTx) MarshalBinary() ([]byte, error) {
  if m == nil {
    return nil, nil
  }
  return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractCallTx) UnmarshalBinary(b []byte) error {
  var res ContractCallTx
  if err := swag.ReadJSON(b, &res); err != nil {
    return err
  }
  *m = res
  return nil
}
