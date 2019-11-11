// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v7/utils"
)

// ChannelSnapshotSoloTx channel snapshot solo tx
// swagger:model ChannelSnapshotSoloTx
type ChannelSnapshotSoloTx struct {

	// channel id
	// Required: true
	ChannelID *string `json:"channel_id"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// from id
	// Required: true
	FromID *string `json:"from_id"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// payload
	// Required: true
	Payload *string `json:"payload"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this channel snapshot solo tx
func (m *ChannelSnapshotSoloTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelSnapshotSoloTx) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("channel_id", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSnapshotSoloTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *ChannelSnapshotSoloTx) validateFromID(formats strfmt.Registry) error {

	if err := validate.Required("from_id", "body", m.FromID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSnapshotSoloTx) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelSnapshotSoloTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelSnapshotSoloTx) UnmarshalBinary(b []byte) error {
	var res ChannelSnapshotSoloTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
