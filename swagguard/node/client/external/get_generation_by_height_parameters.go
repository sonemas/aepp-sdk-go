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

// NewGetGenerationByHeightParams creates a new GetGenerationByHeightParams object
// with the default values initialized.
func NewGetGenerationByHeightParams() *GetGenerationByHeightParams {
	var ()
	return &GetGenerationByHeightParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGenerationByHeightParamsWithTimeout creates a new GetGenerationByHeightParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGenerationByHeightParamsWithTimeout(timeout time.Duration) *GetGenerationByHeightParams {
	var ()
	return &GetGenerationByHeightParams{

		timeout: timeout,
	}
}

// NewGetGenerationByHeightParamsWithContext creates a new GetGenerationByHeightParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGenerationByHeightParamsWithContext(ctx context.Context) *GetGenerationByHeightParams {
	var ()
	return &GetGenerationByHeightParams{

		Context: ctx,
	}
}

// NewGetGenerationByHeightParamsWithHTTPClient creates a new GetGenerationByHeightParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGenerationByHeightParamsWithHTTPClient(client *http.Client) *GetGenerationByHeightParams {
	var ()
	return &GetGenerationByHeightParams{
		HTTPClient: client,
	}
}

/*GetGenerationByHeightParams contains all the parameters to send to the API endpoint
for the get generation by height operation typically these are written to a http.Request
*/
type GetGenerationByHeightParams struct {

	/*Height
	  The height

	*/
	Height uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get generation by height params
func (o *GetGenerationByHeightParams) WithTimeout(timeout time.Duration) *GetGenerationByHeightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get generation by height params
func (o *GetGenerationByHeightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get generation by height params
func (o *GetGenerationByHeightParams) WithContext(ctx context.Context) *GetGenerationByHeightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get generation by height params
func (o *GetGenerationByHeightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get generation by height params
func (o *GetGenerationByHeightParams) WithHTTPClient(client *http.Client) *GetGenerationByHeightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get generation by height params
func (o *GetGenerationByHeightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the get generation by height params
func (o *GetGenerationByHeightParams) WithHeight(height uint64) *GetGenerationByHeightParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the get generation by height params
func (o *GetGenerationByHeightParams) SetHeight(height uint64) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *GetGenerationByHeightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param height
	if err := r.SetPathParam("height", swag.FormatUint64(o.Height)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
