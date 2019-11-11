// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v7/swagguard/compiler/models"
)

// APIVersionReader is a Reader for the APIVersion structure.
type APIVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAPIVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewAPIVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAPIVersionOK creates a APIVersionOK with default headers values
func NewAPIVersionOK() *APIVersionOK {
	return &APIVersionOK{}
}

/*APIVersionOK handles this case with default header values.

Sophia compiler version
*/
type APIVersionOK struct {
	Payload *models.APIVersion
}

func (o *APIVersionOK) Error() string {
	return fmt.Sprintf("[GET /api-version][%d] apiVersionOK  %+v", 200, o.Payload)
}

func (o *APIVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIVersionInternalServerError creates a APIVersionInternalServerError with default headers values
func NewAPIVersionInternalServerError() *APIVersionInternalServerError {
	return &APIVersionInternalServerError{}
}

/*APIVersionInternalServerError handles this case with default header values.

Error
*/
type APIVersionInternalServerError struct {
	Payload *models.Error
}

func (o *APIVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-version][%d] apiVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *APIVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
