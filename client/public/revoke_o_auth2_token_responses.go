// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/arifsetiawan/hydra-go-sdk/models"
)

// RevokeOAuth2TokenReader is a Reader for the RevokeOAuth2Token structure.
type RevokeOAuth2TokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeOAuth2TokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRevokeOAuth2TokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRevokeOAuth2TokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRevokeOAuth2TokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRevokeOAuth2TokenOK creates a RevokeOAuth2TokenOK with default headers values
func NewRevokeOAuth2TokenOK() *RevokeOAuth2TokenOK {
	return &RevokeOAuth2TokenOK{}
}

/*RevokeOAuth2TokenOK handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type RevokeOAuth2TokenOK struct {
}

func (o *RevokeOAuth2TokenOK) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] revokeOAuth2TokenOK ", 200)
}

func (o *RevokeOAuth2TokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeOAuth2TokenUnauthorized creates a RevokeOAuth2TokenUnauthorized with default headers values
func NewRevokeOAuth2TokenUnauthorized() *RevokeOAuth2TokenUnauthorized {
	return &RevokeOAuth2TokenUnauthorized{}
}

/*RevokeOAuth2TokenUnauthorized handles this case with default header values.

genericError
*/
type RevokeOAuth2TokenUnauthorized struct {
	Payload *models.GenericError
}

func (o *RevokeOAuth2TokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] revokeOAuth2TokenUnauthorized  %+v", 401, o.Payload)
}

func (o *RevokeOAuth2TokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeOAuth2TokenInternalServerError creates a RevokeOAuth2TokenInternalServerError with default headers values
func NewRevokeOAuth2TokenInternalServerError() *RevokeOAuth2TokenInternalServerError {
	return &RevokeOAuth2TokenInternalServerError{}
}

/*RevokeOAuth2TokenInternalServerError handles this case with default header values.

genericError
*/
type RevokeOAuth2TokenInternalServerError struct {
	Payload *models.GenericError
}

func (o *RevokeOAuth2TokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] revokeOAuth2TokenInternalServerError  %+v", 500, o.Payload)
}

func (o *RevokeOAuth2TokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
