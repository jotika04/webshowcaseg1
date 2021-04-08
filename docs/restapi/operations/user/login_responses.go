// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"showcase-backend/docs/models"
)

// LoginOKCode is the HTTP code returned for type LoginOK
const LoginOKCode int = 200

/*LoginOK Successful login

swagger:response loginOK
*/
type LoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginSuccess `json:"body,omitempty"`
}

// NewLoginOK creates LoginOK with default headers values
func NewLoginOK() *LoginOK {

	return &LoginOK{}
}

// WithPayload adds the payload to the login o k response
func (o *LoginOK) WithPayload(payload *models.LoginSuccess) *LoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login o k response
func (o *LoginOK) SetPayload(payload *models.LoginSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginBadRequestCode is the HTTP code returned for type LoginBadRequest
const LoginBadRequestCode int = 400

/*LoginBadRequest Bad Request

swagger:response loginBadRequest
*/
type LoginBadRequest struct {
}

// NewLoginBadRequest creates LoginBadRequest with default headers values
func NewLoginBadRequest() *LoginBadRequest {

	return &LoginBadRequest{}
}

// WriteResponse to the client
func (o *LoginBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// LoginNotFoundCode is the HTTP code returned for type LoginNotFound
const LoginNotFoundCode int = 404

/*LoginNotFound User not found

swagger:response loginNotFound
*/
type LoginNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewLoginNotFound creates LoginNotFound with default headers values
func NewLoginNotFound() *LoginNotFound {

	return &LoginNotFound{}
}

// WithPayload adds the payload to the login not found response
func (o *LoginNotFound) WithPayload(payload string) *LoginNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login not found response
func (o *LoginNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// LoginInternalServerErrorCode is the HTTP code returned for type LoginInternalServerError
const LoginInternalServerErrorCode int = 500

/*LoginInternalServerError Server error

swagger:response loginInternalServerError
*/
type LoginInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewLoginInternalServerError creates LoginInternalServerError with default headers values
func NewLoginInternalServerError() *LoginInternalServerError {

	return &LoginInternalServerError{}
}

// WithPayload adds the payload to the login internal server error response
func (o *LoginInternalServerError) WithPayload(payload string) *LoginInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login internal server error response
func (o *LoginInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
