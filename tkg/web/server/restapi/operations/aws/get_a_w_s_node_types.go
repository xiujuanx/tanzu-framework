// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAWSNodeTypesHandlerFunc turns a function with the right signature into a get a w s node types handler
type GetAWSNodeTypesHandlerFunc func(GetAWSNodeTypesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAWSNodeTypesHandlerFunc) Handle(params GetAWSNodeTypesParams) middleware.Responder {
	return fn(params)
}

// GetAWSNodeTypesHandler interface for that can handle valid get a w s node types params
type GetAWSNodeTypesHandler interface {
	Handle(GetAWSNodeTypesParams) middleware.Responder
}

// NewGetAWSNodeTypes creates a new http.Handler for the get a w s node types operation
func NewGetAWSNodeTypes(ctx *middleware.Context, handler GetAWSNodeTypesHandler) *GetAWSNodeTypes {
	return &GetAWSNodeTypes{Context: ctx, Handler: handler}
}

/*
GetAWSNodeTypes swagger:route GET /api/providers/aws/nodetypes aws getAWSNodeTypes

Retrieve AWS supported node types
*/
type GetAWSNodeTypes struct {
	Context *middleware.Context
	Handler GetAWSNodeTypesHandler
}

func (o *GetAWSNodeTypes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAWSNodeTypesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
