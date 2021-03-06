// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAccountByPubkeyAndHeightParams creates a new GetAccountByPubkeyAndHeightParams object
// with the default values initialized.
func NewGetAccountByPubkeyAndHeightParams() *GetAccountByPubkeyAndHeightParams {
	var ()
	return &GetAccountByPubkeyAndHeightParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountByPubkeyAndHeightParamsWithTimeout creates a new GetAccountByPubkeyAndHeightParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountByPubkeyAndHeightParamsWithTimeout(timeout time.Duration) *GetAccountByPubkeyAndHeightParams {
	var ()
	return &GetAccountByPubkeyAndHeightParams{

		timeout: timeout,
	}
}

// NewGetAccountByPubkeyAndHeightParamsWithContext creates a new GetAccountByPubkeyAndHeightParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountByPubkeyAndHeightParamsWithContext(ctx context.Context) *GetAccountByPubkeyAndHeightParams {
	var ()
	return &GetAccountByPubkeyAndHeightParams{

		Context: ctx,
	}
}

// NewGetAccountByPubkeyAndHeightParamsWithHTTPClient creates a new GetAccountByPubkeyAndHeightParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountByPubkeyAndHeightParamsWithHTTPClient(client *http.Client) *GetAccountByPubkeyAndHeightParams {
	var ()
	return &GetAccountByPubkeyAndHeightParams{
		HTTPClient: client,
	}
}

/*GetAccountByPubkeyAndHeightParams contains all the parameters to send to the API endpoint
for the get account by pubkey and height operation typically these are written to a http.Request
*/
type GetAccountByPubkeyAndHeightParams struct {

	/*Height
	  The height

	*/
	Height uint64
	/*Pubkey
	  The public key of the account

	*/
	Pubkey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) WithTimeout(timeout time.Duration) *GetAccountByPubkeyAndHeightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) WithContext(ctx context.Context) *GetAccountByPubkeyAndHeightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) WithHTTPClient(client *http.Client) *GetAccountByPubkeyAndHeightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) WithHeight(height uint64) *GetAccountByPubkeyAndHeightParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) SetHeight(height uint64) {
	o.Height = height
}

// WithPubkey adds the pubkey to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) WithPubkey(pubkey string) *GetAccountByPubkeyAndHeightParams {
	o.SetPubkey(pubkey)
	return o
}

// SetPubkey adds the pubkey to the get account by pubkey and height params
func (o *GetAccountByPubkeyAndHeightParams) SetPubkey(pubkey string) {
	o.Pubkey = pubkey
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountByPubkeyAndHeightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param height
	if err := r.SetPathParam("height", swag.FormatUint64(o.Height)); err != nil {
		return err
	}

	// path param pubkey
	if err := r.SetPathParam("pubkey", o.Pubkey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
