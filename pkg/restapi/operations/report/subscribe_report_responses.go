// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/models"
)

// SubscribeReportCreatedCode is the HTTP code returned for type SubscribeReportCreated
const SubscribeReportCreatedCode int = 201

/*SubscribeReportCreated Subscription successfully created

swagger:response subscribeReportCreated
*/
type SubscribeReportCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SubscriptionResponse `json:"body,omitempty"`
}

// NewSubscribeReportCreated creates SubscribeReportCreated with default headers values
func NewSubscribeReportCreated() *SubscribeReportCreated {

	return &SubscribeReportCreated{}
}

// WithPayload adds the payload to the subscribe report created response
func (o *SubscribeReportCreated) WithPayload(payload *models.SubscriptionResponse) *SubscribeReportCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the subscribe report created response
func (o *SubscribeReportCreated) SetPayload(payload *models.SubscriptionResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubscribeReportCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SubscribeReportBadRequestCode is the HTTP code returned for type SubscribeReportBadRequest
const SubscribeReportBadRequestCode int = 400

/*SubscribeReportBadRequest Invalid input

swagger:response subscribeReportBadRequest
*/
type SubscribeReportBadRequest struct {
}

// NewSubscribeReportBadRequest creates SubscribeReportBadRequest with default headers values
func NewSubscribeReportBadRequest() *SubscribeReportBadRequest {

	return &SubscribeReportBadRequest{}
}

// WriteResponse to the client
func (o *SubscribeReportBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// SubscribeReportInternalServerErrorCode is the HTTP code returned for type SubscribeReportInternalServerError
const SubscribeReportInternalServerErrorCode int = 500

/*SubscribeReportInternalServerError Internal error

swagger:response subscribeReportInternalServerError
*/
type SubscribeReportInternalServerError struct {
}

// NewSubscribeReportInternalServerError creates SubscribeReportInternalServerError with default headers values
func NewSubscribeReportInternalServerError() *SubscribeReportInternalServerError {

	return &SubscribeReportInternalServerError{}
}

// WriteResponse to the client
func (o *SubscribeReportInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
