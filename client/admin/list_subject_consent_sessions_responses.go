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

// ListSubjectConsentSessionsReader is a Reader for the ListSubjectConsentSessions structure.
type ListSubjectConsentSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSubjectConsentSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSubjectConsentSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSubjectConsentSessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListSubjectConsentSessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListSubjectConsentSessionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSubjectConsentSessionsOK creates a ListSubjectConsentSessionsOK with default headers values
func NewListSubjectConsentSessionsOK() *ListSubjectConsentSessionsOK {
	return &ListSubjectConsentSessionsOK{}
}

/*ListSubjectConsentSessionsOK handles this case with default header values.

A list of used consent requests.
*/
type ListSubjectConsentSessionsOK struct {
	Payload []*models.PreviousConsentSession
}

func (o *ListSubjectConsentSessionsOK) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/sessions/consent][%d] listSubjectConsentSessionsOK  %+v", 200, o.Payload)
}

func (o *ListSubjectConsentSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSubjectConsentSessionsBadRequest creates a ListSubjectConsentSessionsBadRequest with default headers values
func NewListSubjectConsentSessionsBadRequest() *ListSubjectConsentSessionsBadRequest {
	return &ListSubjectConsentSessionsBadRequest{}
}

/*ListSubjectConsentSessionsBadRequest handles this case with default header values.

genericError
*/
type ListSubjectConsentSessionsBadRequest struct {
	Payload *models.GenericError
}

func (o *ListSubjectConsentSessionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/sessions/consent][%d] listSubjectConsentSessionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListSubjectConsentSessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSubjectConsentSessionsNotFound creates a ListSubjectConsentSessionsNotFound with default headers values
func NewListSubjectConsentSessionsNotFound() *ListSubjectConsentSessionsNotFound {
	return &ListSubjectConsentSessionsNotFound{}
}

/*ListSubjectConsentSessionsNotFound handles this case with default header values.

genericError
*/
type ListSubjectConsentSessionsNotFound struct {
	Payload *models.GenericError
}

func (o *ListSubjectConsentSessionsNotFound) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/sessions/consent][%d] listSubjectConsentSessionsNotFound  %+v", 404, o.Payload)
}

func (o *ListSubjectConsentSessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSubjectConsentSessionsInternalServerError creates a ListSubjectConsentSessionsInternalServerError with default headers values
func NewListSubjectConsentSessionsInternalServerError() *ListSubjectConsentSessionsInternalServerError {
	return &ListSubjectConsentSessionsInternalServerError{}
}

/*ListSubjectConsentSessionsInternalServerError handles this case with default header values.

genericError
*/
type ListSubjectConsentSessionsInternalServerError struct {
	Payload *models.GenericError
}

func (o *ListSubjectConsentSessionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/sessions/consent][%d] listSubjectConsentSessionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSubjectConsentSessionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
