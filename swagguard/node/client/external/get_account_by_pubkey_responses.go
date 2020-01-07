// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
)

// GetAccountByPubkeyReader is a Reader for the GetAccountByPubkey structure.
type GetAccountByPubkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountByPubkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountByPubkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAccountByPubkeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAccountByPubkeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountByPubkeyOK creates a GetAccountByPubkeyOK with default headers values
func NewGetAccountByPubkeyOK() *GetAccountByPubkeyOK {
	return &GetAccountByPubkeyOK{}
}

/*GetAccountByPubkeyOK handles this case with default header values.

Successful operation
*/
type GetAccountByPubkeyOK struct {
	Payload *models.Account
}

func (o *GetAccountByPubkeyOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}][%d] getAccountByPubkeyOK  %+v", 200, o.Payload)
}

func (o *GetAccountByPubkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountByPubkeyBadRequest creates a GetAccountByPubkeyBadRequest with default headers values
func NewGetAccountByPubkeyBadRequest() *GetAccountByPubkeyBadRequest {
	return &GetAccountByPubkeyBadRequest{}
}

/*GetAccountByPubkeyBadRequest handles this case with default header values.

Invalid public key
*/
type GetAccountByPubkeyBadRequest struct {
	Payload *models.Error
}

func (o *GetAccountByPubkeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}][%d] getAccountByPubkeyBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountByPubkeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountByPubkeyNotFound creates a GetAccountByPubkeyNotFound with default headers values
func NewGetAccountByPubkeyNotFound() *GetAccountByPubkeyNotFound {
	return &GetAccountByPubkeyNotFound{}
}

/*GetAccountByPubkeyNotFound handles this case with default header values.

Account not found
*/
type GetAccountByPubkeyNotFound struct {
	Payload *models.Error
}

func (o *GetAccountByPubkeyNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{pubkey}][%d] getAccountByPubkeyNotFound  %+v", 404, o.Payload)
}

func (o *GetAccountByPubkeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
