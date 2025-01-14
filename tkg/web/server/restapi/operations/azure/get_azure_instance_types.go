// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAzureInstanceTypesHandlerFunc turns a function with the right signature into a get azure instance types handler
type GetAzureInstanceTypesHandlerFunc func(GetAzureInstanceTypesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAzureInstanceTypesHandlerFunc) Handle(params GetAzureInstanceTypesParams) middleware.Responder {
	return fn(params)
}

// GetAzureInstanceTypesHandler interface for that can handle valid get azure instance types params
type GetAzureInstanceTypesHandler interface {
	Handle(GetAzureInstanceTypesParams) middleware.Responder
}

// NewGetAzureInstanceTypes creates a new http.Handler for the get azure instance types operation
func NewGetAzureInstanceTypes(ctx *middleware.Context, handler GetAzureInstanceTypesHandler) *GetAzureInstanceTypes {
	return &GetAzureInstanceTypes{Context: ctx, Handler: handler}
}

/*
GetAzureInstanceTypes swagger:route GET /api/providers/azure/regions/{location}/instanceTypes azure getAzureInstanceTypes

Retrieve list of supported Azure instance types for a region
*/
type GetAzureInstanceTypes struct {
	Context *middleware.Context
	Handler GetAzureInstanceTypesHandler
}

func (o *GetAzureInstanceTypes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAzureInstanceTypesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
