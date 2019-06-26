// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra/sdk/go/hydra/models"
)

// DeleteOAuth2ClientReader is a Reader for the DeleteOAuth2Client structure.
type DeleteOAuth2ClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOAuth2ClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteOAuth2ClientNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteOAuth2ClientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteOAuth2ClientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOAuth2ClientNoContent creates a DeleteOAuth2ClientNoContent with default headers values
func NewDeleteOAuth2ClientNoContent() *DeleteOAuth2ClientNoContent {
	return &DeleteOAuth2ClientNoContent{}
}

/*DeleteOAuth2ClientNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type DeleteOAuth2ClientNoContent struct {
}

func (o *DeleteOAuth2ClientNoContent) Error() string {
	return fmt.Sprintf("[DELETE /clients/{id}][%d] deleteOAuth2ClientNoContent ", 204)
}

func (o *DeleteOAuth2ClientNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOAuth2ClientNotFound creates a DeleteOAuth2ClientNotFound with default headers values
func NewDeleteOAuth2ClientNotFound() *DeleteOAuth2ClientNotFound {
	return &DeleteOAuth2ClientNotFound{}
}

/*DeleteOAuth2ClientNotFound handles this case with default header values.

genericError
*/
type DeleteOAuth2ClientNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteOAuth2ClientNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clients/{id}][%d] deleteOAuth2ClientNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOAuth2ClientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOAuth2ClientInternalServerError creates a DeleteOAuth2ClientInternalServerError with default headers values
func NewDeleteOAuth2ClientInternalServerError() *DeleteOAuth2ClientInternalServerError {
	return &DeleteOAuth2ClientInternalServerError{}
}

/*DeleteOAuth2ClientInternalServerError handles this case with default header values.

genericError
*/
type DeleteOAuth2ClientInternalServerError struct {
	Payload *models.GenericError
}

func (o *DeleteOAuth2ClientInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /clients/{id}][%d] deleteOAuth2ClientInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOAuth2ClientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
