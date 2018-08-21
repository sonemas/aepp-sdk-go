// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// NewPostContractCreateComputeParams creates a new PostContractCreateComputeParams object
// with the default values initialized.
func NewPostContractCreateComputeParams() *PostContractCreateComputeParams {
	var ()
	return &PostContractCreateComputeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostContractCreateComputeParamsWithTimeout creates a new PostContractCreateComputeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostContractCreateComputeParamsWithTimeout(timeout time.Duration) *PostContractCreateComputeParams {
	var ()
	return &PostContractCreateComputeParams{

		timeout: timeout,
	}
}

// NewPostContractCreateComputeParamsWithContext creates a new PostContractCreateComputeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostContractCreateComputeParamsWithContext(ctx context.Context) *PostContractCreateComputeParams {
	var ()
	return &PostContractCreateComputeParams{

		Context: ctx,
	}
}

// NewPostContractCreateComputeParamsWithHTTPClient creates a new PostContractCreateComputeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostContractCreateComputeParamsWithHTTPClient(client *http.Client) *PostContractCreateComputeParams {
	var ()
	return &PostContractCreateComputeParams{
		HTTPClient: client,
	}
}

/*PostContractCreateComputeParams contains all the parameters to send to the API endpoint
for the post contract create compute operation typically these are written to a http.Request
*/
type PostContractCreateComputeParams struct {

	/*Body*/
	Body *models.ContractCreateCompute

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post contract create compute params
func (o *PostContractCreateComputeParams) WithTimeout(timeout time.Duration) *PostContractCreateComputeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post contract create compute params
func (o *PostContractCreateComputeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post contract create compute params
func (o *PostContractCreateComputeParams) WithContext(ctx context.Context) *PostContractCreateComputeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post contract create compute params
func (o *PostContractCreateComputeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post contract create compute params
func (o *PostContractCreateComputeParams) WithHTTPClient(client *http.Client) *PostContractCreateComputeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post contract create compute params
func (o *PostContractCreateComputeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post contract create compute params
func (o *PostContractCreateComputeParams) WithBody(body *models.ContractCreateCompute) *PostContractCreateComputeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post contract create compute params
func (o *PostContractCreateComputeParams) SetBody(body *models.ContractCreateCompute) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostContractCreateComputeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
