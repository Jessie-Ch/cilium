package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// HTTP code for type GetEndpointOK
const GetEndpointOKCode int = 200

/*GetEndpointOK Success

swagger:response getEndpointOK
*/
type GetEndpointOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Endpoint `json:"body,omitempty"`
}

// NewGetEndpointOK creates GetEndpointOK with default headers values
func NewGetEndpointOK() *GetEndpointOK {
	return &GetEndpointOK{}
}

// WithPayload adds the payload to the get endpoint o k response
func (o *GetEndpointOK) WithPayload(payload []*models.Endpoint) *GetEndpointOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoint o k response
func (o *GetEndpointOK) SetPayload(payload []*models.Endpoint) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Endpoint, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
