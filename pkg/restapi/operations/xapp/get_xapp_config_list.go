// Code generated by go-swagger; DO NOT EDIT.

package xapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetXappConfigListHandlerFunc turns a function with the right signature into a get xapp config list handler
type GetXappConfigListHandlerFunc func(GetXappConfigListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetXappConfigListHandlerFunc) Handle(params GetXappConfigListParams) middleware.Responder {
	return fn(params)
}

// GetXappConfigListHandler interface for that can handle valid get xapp config list params
type GetXappConfigListHandler interface {
	Handle(GetXappConfigListParams) middleware.Responder
}

// NewGetXappConfigList creates a new http.Handler for the get xapp config list operation
func NewGetXappConfigList(ctx *middleware.Context, handler GetXappConfigListHandler) *GetXappConfigList {
	return &GetXappConfigList{Context: ctx, Handler: handler}
}

/*GetXappConfigList swagger:route GET /config xapp getXappConfigList

Returns the configuration of all xapps

*/
type GetXappConfigList struct {
	Context *middleware.Context
	Handler GetXappConfigListHandler
}

func (o *GetXappConfigList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetXappConfigListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
