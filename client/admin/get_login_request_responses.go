// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/arifsetiawan/hydra-go-sdk/models"
)

// GetLoginRequestReader is a Reader for the GetLoginRequest structure.
type GetLoginRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoginRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoginRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetLoginRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLoginRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetLoginRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetLoginRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoginRequestOK creates a GetLoginRequestOK with default headers values
func NewGetLoginRequestOK() *GetLoginRequestOK {
	return &GetLoginRequestOK{}
}

/*GetLoginRequestOK handles this case with default header values.

loginRequest
*/
type GetLoginRequestOK struct {
	Payload *models.LoginRequest
}

func (o *GetLoginRequestOK) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/login][%d] getLoginRequestOK  %+v", 200, o.Payload)
}

func (o *GetLoginRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoginRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoginRequestBadRequest creates a GetLoginRequestBadRequest with default headers values
func NewGetLoginRequestBadRequest() *GetLoginRequestBadRequest {
	return &GetLoginRequestBadRequest{}
}

/*GetLoginRequestBadRequest handles this case with default header values.

genericError
*/
type GetLoginRequestBadRequest struct {
	Payload *models.GenericError
}

func (o *GetLoginRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/login][%d] getLoginRequestBadRequest  %+v", 400, o.Payload)
}

func (o *GetLoginRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoginRequestNotFound creates a GetLoginRequestNotFound with default headers values
func NewGetLoginRequestNotFound() *GetLoginRequestNotFound {
	return &GetLoginRequestNotFound{}
}

/*GetLoginRequestNotFound handles this case with default header values.

genericError
*/
type GetLoginRequestNotFound struct {
	Payload *models.GenericError
}

func (o *GetLoginRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/login][%d] getLoginRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetLoginRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoginRequestConflict creates a GetLoginRequestConflict with default headers values
func NewGetLoginRequestConflict() *GetLoginRequestConflict {
	return &GetLoginRequestConflict{}
}

/*GetLoginRequestConflict handles this case with default header values.

genericError
*/
type GetLoginRequestConflict struct {
	Payload *models.GenericError
}

func (o *GetLoginRequestConflict) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/login][%d] getLoginRequestConflict  %+v", 409, o.Payload)
}

func (o *GetLoginRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoginRequestInternalServerError creates a GetLoginRequestInternalServerError with default headers values
func NewGetLoginRequestInternalServerError() *GetLoginRequestInternalServerError {
	return &GetLoginRequestInternalServerError{}
}

/*GetLoginRequestInternalServerError handles this case with default header values.

genericError
*/
type GetLoginRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetLoginRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/login][%d] getLoginRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLoginRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
