// Code generated by go-swagger; DO NOT EDIT.

package logins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// GetRsaPubKeyOKCode is the HTTP code returned for type GetRsaPubKeyOK
const GetRsaPubKeyOKCode int = 200

/*GetRsaPubKeyOK OK

swagger:response getRsaPubKeyOK
*/
type GetRsaPubKeyOK struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewGetRsaPubKeyOK creates GetRsaPubKeyOK with default headers values
func NewGetRsaPubKeyOK() *GetRsaPubKeyOK {

	return &GetRsaPubKeyOK{}
}

// WithPayload adds the payload to the get rsa pub key o k response
func (o *GetRsaPubKeyOK) WithPayload(payload *models.Result) *GetRsaPubKeyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rsa pub key o k response
func (o *GetRsaPubKeyOK) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRsaPubKeyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRsaPubKeyNotFoundCode is the HTTP code returned for type GetRsaPubKeyNotFound
const GetRsaPubKeyNotFoundCode int = 404

/*GetRsaPubKeyNotFound the url is not found.

swagger:response getRsaPubKeyNotFound
*/
type GetRsaPubKeyNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRsaPubKeyNotFound creates GetRsaPubKeyNotFound with default headers values
func NewGetRsaPubKeyNotFound() *GetRsaPubKeyNotFound {

	return &GetRsaPubKeyNotFound{}
}

// WithPayload adds the payload to the get rsa pub key not found response
func (o *GetRsaPubKeyNotFound) WithPayload(payload *models.Error) *GetRsaPubKeyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rsa pub key not found response
func (o *GetRsaPubKeyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRsaPubKeyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
