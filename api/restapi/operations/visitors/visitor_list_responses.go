// Code generated by go-swagger; DO NOT EDIT.

package visitors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/quintsys/ga-exporter/api/models"
)

// VisitorListOKCode is the HTTP code returned for type VisitorListOK
const VisitorListOKCode int = 200

/*
VisitorListOK Returns data about visitors

swagger:response visitorListOK
*/
type VisitorListOK struct {

	/*
	  Max Items: 5000
	  In: Body
	*/
	Payload []*models.Visitor `json:"body,omitempty"`
}

// NewVisitorListOK creates VisitorListOK with default headers values
func NewVisitorListOK() *VisitorListOK {

	return &VisitorListOK{}
}

// WithPayload adds the payload to the visitor list o k response
func (o *VisitorListOK) WithPayload(payload []*models.Visitor) *VisitorListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the visitor list o k response
func (o *VisitorListOK) SetPayload(payload []*models.Visitor) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VisitorListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Visitor, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// VisitorListMethodNotAllowedCode is the HTTP code returned for type VisitorListMethodNotAllowed
const VisitorListMethodNotAllowedCode int = 405

/*
VisitorListMethodNotAllowed Method Not Allowed

swagger:response visitorListMethodNotAllowed
*/
type VisitorListMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMethodNotAllowed `json:"body,omitempty"`
}

// NewVisitorListMethodNotAllowed creates VisitorListMethodNotAllowed with default headers values
func NewVisitorListMethodNotAllowed() *VisitorListMethodNotAllowed {

	return &VisitorListMethodNotAllowed{}
}

// WithPayload adds the payload to the visitor list method not allowed response
func (o *VisitorListMethodNotAllowed) WithPayload(payload *models.ErrorMethodNotAllowed) *VisitorListMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the visitor list method not allowed response
func (o *VisitorListMethodNotAllowed) SetPayload(payload *models.ErrorMethodNotAllowed) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VisitorListMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VisitorListNotAcceptableCode is the HTTP code returned for type VisitorListNotAcceptable
const VisitorListNotAcceptableCode int = 406

/*
VisitorListNotAcceptable Not Acceptable

swagger:response visitorListNotAcceptable
*/
type VisitorListNotAcceptable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorNotAcceptable `json:"body,omitempty"`
}

// NewVisitorListNotAcceptable creates VisitorListNotAcceptable with default headers values
func NewVisitorListNotAcceptable() *VisitorListNotAcceptable {

	return &VisitorListNotAcceptable{}
}

// WithPayload adds the payload to the visitor list not acceptable response
func (o *VisitorListNotAcceptable) WithPayload(payload *models.ErrorNotAcceptable) *VisitorListNotAcceptable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the visitor list not acceptable response
func (o *VisitorListNotAcceptable) SetPayload(payload *models.ErrorNotAcceptable) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VisitorListNotAcceptable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(406)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VisitorListTooManyRequestsCode is the HTTP code returned for type VisitorListTooManyRequests
const VisitorListTooManyRequestsCode int = 429

/*
VisitorListTooManyRequests Too many request

swagger:response visitorListTooManyRequests
*/
type VisitorListTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorTooManyRequests `json:"body,omitempty"`
}

// NewVisitorListTooManyRequests creates VisitorListTooManyRequests with default headers values
func NewVisitorListTooManyRequests() *VisitorListTooManyRequests {

	return &VisitorListTooManyRequests{}
}

// WithPayload adds the payload to the visitor list too many requests response
func (o *VisitorListTooManyRequests) WithPayload(payload *models.ErrorTooManyRequests) *VisitorListTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the visitor list too many requests response
func (o *VisitorListTooManyRequests) SetPayload(payload *models.ErrorTooManyRequests) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VisitorListTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VisitorListInternalServerErrorCode is the HTTP code returned for type VisitorListInternalServerError
const VisitorListInternalServerErrorCode int = 500

/*
VisitorListInternalServerError Internal Server Error

swagger:response visitorListInternalServerError
*/
type VisitorListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorInternalServerError `json:"body,omitempty"`
}

// NewVisitorListInternalServerError creates VisitorListInternalServerError with default headers values
func NewVisitorListInternalServerError() *VisitorListInternalServerError {

	return &VisitorListInternalServerError{}
}

// WithPayload adds the payload to the visitor list internal server error response
func (o *VisitorListInternalServerError) WithPayload(payload *models.ErrorInternalServerError) *VisitorListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the visitor list internal server error response
func (o *VisitorListInternalServerError) SetPayload(payload *models.ErrorInternalServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VisitorListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VisitorListServiceUnavailableCode is the HTTP code returned for type VisitorListServiceUnavailable
const VisitorListServiceUnavailableCode int = 503

/*
VisitorListServiceUnavailable Service Unavailable

swagger:response visitorListServiceUnavailable
*/
type VisitorListServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorServiceUnavailable `json:"body,omitempty"`
}

// NewVisitorListServiceUnavailable creates VisitorListServiceUnavailable with default headers values
func NewVisitorListServiceUnavailable() *VisitorListServiceUnavailable {

	return &VisitorListServiceUnavailable{}
}

// WithPayload adds the payload to the visitor list service unavailable response
func (o *VisitorListServiceUnavailable) WithPayload(payload *models.ErrorServiceUnavailable) *VisitorListServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the visitor list service unavailable response
func (o *VisitorListServiceUnavailable) SetPayload(payload *models.ErrorServiceUnavailable) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VisitorListServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
VisitorListDefault Unexpected error

swagger:response visitorListDefault
*/
type VisitorListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVisitorListDefault creates VisitorListDefault with default headers values
func NewVisitorListDefault(code int) *VisitorListDefault {
	if code <= 0 {
		code = 500
	}

	return &VisitorListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the visitor list default response
func (o *VisitorListDefault) WithStatusCode(code int) *VisitorListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the visitor list default response
func (o *VisitorListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the visitor list default response
func (o *VisitorListDefault) WithPayload(payload *models.Error) *VisitorListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the visitor list default response
func (o *VisitorListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VisitorListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
