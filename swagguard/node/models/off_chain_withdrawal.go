// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v7/utils"
)

// OffChainWithdrawal off chain withdrawal
// swagger:model OffChainWithdrawal
type OffChainWithdrawal struct {

	// amount
	// Required: true
	Amount utils.BigInt `json:"amount"`

	// Withdrawer of tokens
	// Required: true
	To *string `json:"to"`
}

// Op gets the op of this subtype
func (m *OffChainWithdrawal) Op() string {
	return "OffChainWithdrawal"
}

// SetOp sets the op of this subtype
func (m *OffChainWithdrawal) SetOp(val string) {

}

// Amount gets the amount of this subtype

// To gets the to of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OffChainWithdrawal) UnmarshalJSON(raw []byte) error {
	var data struct {

		// amount
		// Required: true
		Amount utils.BigInt `json:"amount"`

		// Withdrawer of tokens
		// Required: true
		To *string `json:"to"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Op string `json:"op"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result OffChainWithdrawal

	if base.Op != result.Op() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid op value: %q", base.Op)
	}

	result.Amount = data.Amount

	result.To = data.To

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OffChainWithdrawal) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// amount
		// Required: true
		Amount utils.BigInt `json:"amount"`

		// Withdrawer of tokens
		// Required: true
		To *string `json:"to"`
	}{

		Amount: m.Amount,

		To: m.To,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Op string `json:"op"`
	}{

		Op: m.Op(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this off chain withdrawal
func (m *OffChainWithdrawal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OffChainWithdrawal) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *OffChainWithdrawal) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OffChainWithdrawal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffChainWithdrawal) UnmarshalBinary(b []byte) error {
	var res OffChainWithdrawal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
