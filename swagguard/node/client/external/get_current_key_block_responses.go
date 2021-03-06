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

// GetCurrentKeyBlockReader is a Reader for the GetCurrentKeyBlock structure.
type GetCurrentKeyBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentKeyBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCurrentKeyBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCurrentKeyBlockNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrentKeyBlockOK creates a GetCurrentKeyBlockOK with default headers values
func NewGetCurrentKeyBlockOK() *GetCurrentKeyBlockOK {
	return &GetCurrentKeyBlockOK{}
}

/*GetCurrentKeyBlockOK handles this case with default header values.

Successful operation
*/
type GetCurrentKeyBlockOK struct {
	Payload *models.KeyBlock
}

func (o *GetCurrentKeyBlockOK) Error() string {
	return fmt.Sprintf("[GET /key-blocks/current][%d] getCurrentKeyBlockOK  %+v", 200, o.Payload)
}

func (o *GetCurrentKeyBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyBlock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentKeyBlockNotFound creates a GetCurrentKeyBlockNotFound with default headers values
func NewGetCurrentKeyBlockNotFound() *GetCurrentKeyBlockNotFound {
	return &GetCurrentKeyBlockNotFound{}
}

/*GetCurrentKeyBlockNotFound handles this case with default header values.

Block not found
*/
type GetCurrentKeyBlockNotFound struct {
	Payload *models.Error
}

func (o *GetCurrentKeyBlockNotFound) Error() string {
	return fmt.Sprintf("[GET /key-blocks/current][%d] getCurrentKeyBlockNotFound  %+v", 404, o.Payload)
}

func (o *GetCurrentKeyBlockNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
