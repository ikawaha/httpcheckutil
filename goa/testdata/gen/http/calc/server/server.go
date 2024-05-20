// Code generated by goa v3.16.2, DO NOT EDIT.
//
// calc HTTP server
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/gen/calc"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the calc service endpoint HTTP handlers.
type Server struct {
	Mounts       []*MountPoint
	Multiply     http.Handler
	Openapi3JSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the calc service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *calc.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemOpenapi3JSON http.FileSystem,
) *Server {
	if fileSystemOpenapi3JSON == nil {
		fileSystemOpenapi3JSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Multiply", "GET", "/multiply/{a}/{b}"},
			{"openapi3.json", "GET", "/openapi.json"},
		},
		Multiply:     NewMultiplyHandler(e.Multiply, mux, decoder, encoder, errhandler, formatter),
		Openapi3JSON: http.FileServer(fileSystemOpenapi3JSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "calc" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Multiply = m(s.Multiply)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return calc.MethodNames[:] }

// Mount configures the mux to serve the calc endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountMultiplyHandler(mux, h.Multiply)
	MountOpenapi3JSON(mux, goahttp.Replace("", "/openapi3.json", h.Openapi3JSON))
}

// Mount configures the mux to serve the calc endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountMultiplyHandler configures the mux to serve the "calc" service
// "multiply" endpoint.
func MountMultiplyHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/multiply/{a}/{b}", f)
}

// NewMultiplyHandler creates a HTTP handler which loads the HTTP request and
// calls the "calc" service "multiply" endpoint.
func NewMultiplyHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMultiplyRequest(mux, decoder)
		encodeResponse = EncodeMultiplyResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "multiply")
		ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountOpenapi3JSON configures the mux to serve GET request made to
// "/openapi.json".
func MountOpenapi3JSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}
