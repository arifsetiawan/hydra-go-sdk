// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIntrospectOAuth2TokenParams creates a new IntrospectOAuth2TokenParams object
// with the default values initialized.
func NewIntrospectOAuth2TokenParams() *IntrospectOAuth2TokenParams {
	var ()
	return &IntrospectOAuth2TokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIntrospectOAuth2TokenParamsWithTimeout creates a new IntrospectOAuth2TokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIntrospectOAuth2TokenParamsWithTimeout(timeout time.Duration) *IntrospectOAuth2TokenParams {
	var ()
	return &IntrospectOAuth2TokenParams{

		timeout: timeout,
	}
}

// NewIntrospectOAuth2TokenParamsWithContext creates a new IntrospectOAuth2TokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewIntrospectOAuth2TokenParamsWithContext(ctx context.Context) *IntrospectOAuth2TokenParams {
	var ()
	return &IntrospectOAuth2TokenParams{

		Context: ctx,
	}
}

// NewIntrospectOAuth2TokenParamsWithHTTPClient creates a new IntrospectOAuth2TokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIntrospectOAuth2TokenParamsWithHTTPClient(client *http.Client) *IntrospectOAuth2TokenParams {
	var ()
	return &IntrospectOAuth2TokenParams{
		HTTPClient: client,
	}
}

/*IntrospectOAuth2TokenParams contains all the parameters to send to the API endpoint
for the introspect o auth2 token operation typically these are written to a http.Request
*/
type IntrospectOAuth2TokenParams struct {

	/*Scope
	  An optional, space separated list of required scopes. If the access token was not granted one of the
	scopes, the result of active will be false.

	*/
	Scope *string
	/*Token
	  The string value of the token. For access tokens, this
	is the "access_token" value returned from the token endpoint
	defined in OAuth 2.0. For refresh tokens, this is the "refresh_token"
	value returned.

	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) WithTimeout(timeout time.Duration) *IntrospectOAuth2TokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) WithContext(ctx context.Context) *IntrospectOAuth2TokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) WithHTTPClient(client *http.Client) *IntrospectOAuth2TokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) WithScope(scope *string) *IntrospectOAuth2TokenParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithToken adds the token to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) WithToken(token string) *IntrospectOAuth2TokenParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the introspect o auth2 token params
func (o *IntrospectOAuth2TokenParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *IntrospectOAuth2TokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Scope != nil {

		// form param scope
		var frScope string
		if o.Scope != nil {
			frScope = *o.Scope
		}
		fScope := frScope
		if fScope != "" {
			if err := r.SetFormParam("scope", fScope); err != nil {
				return err
			}
		}

	}

	// form param token
	frToken := o.Token
	fToken := frToken
	if fToken != "" {
		if err := r.SetFormParam("token", fToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
