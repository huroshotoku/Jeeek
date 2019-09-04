// Code generated by goa v3.0.4, DO NOT EDIT.
//
// Admin HTTP server
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package server

import (
	"context"
	"net/http"

	admin "github.com/tonouchi510/Jeeek/gen/admin"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the Admin service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	AdminHealthCheck   http.Handler
	AdminSignin        http.Handler
	AdminCreateNewUser http.Handler
	AdminUpdateUser    http.Handler
	AdminListUser      http.Handler
	AdminGetUser       http.Handler
	AdminDeleteUser    http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
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

// New instantiates HTTP handlers for all the Admin service endpoints.
func New(
	e *admin.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"AdminHealthCheck", "GET", "/v1/admin/healthcheck"},
			{"AdminSignin", "POST", "/v1/admin/signin"},
			{"AdminCreateNewUser", "POST", "/v1/admin/users"},
			{"AdminUpdateUser", "PUT", "/v1/admin/users/{user_id}"},
			{"AdminListUser", "GET", "/v1/admin/users"},
			{"AdminGetUser", "GET", "/v1/admin/users/{user_id}"},
			{"AdminDeleteUser", "DELETE", "/v1/admin/users/{user_id}"},
		},
		AdminHealthCheck:   NewAdminHealthCheckHandler(e.AdminHealthCheck, mux, dec, enc, eh),
		AdminSignin:        NewAdminSigninHandler(e.AdminSignin, mux, dec, enc, eh),
		AdminCreateNewUser: NewAdminCreateNewUserHandler(e.AdminCreateNewUser, mux, dec, enc, eh),
		AdminUpdateUser:    NewAdminUpdateUserHandler(e.AdminUpdateUser, mux, dec, enc, eh),
		AdminListUser:      NewAdminListUserHandler(e.AdminListUser, mux, dec, enc, eh),
		AdminGetUser:       NewAdminGetUserHandler(e.AdminGetUser, mux, dec, enc, eh),
		AdminDeleteUser:    NewAdminDeleteUserHandler(e.AdminDeleteUser, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Admin" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.AdminHealthCheck = m(s.AdminHealthCheck)
	s.AdminSignin = m(s.AdminSignin)
	s.AdminCreateNewUser = m(s.AdminCreateNewUser)
	s.AdminUpdateUser = m(s.AdminUpdateUser)
	s.AdminListUser = m(s.AdminListUser)
	s.AdminGetUser = m(s.AdminGetUser)
	s.AdminDeleteUser = m(s.AdminDeleteUser)
}

// Mount configures the mux to serve the Admin endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAdminHealthCheckHandler(mux, h.AdminHealthCheck)
	MountAdminSigninHandler(mux, h.AdminSignin)
	MountAdminCreateNewUserHandler(mux, h.AdminCreateNewUser)
	MountAdminUpdateUserHandler(mux, h.AdminUpdateUser)
	MountAdminListUserHandler(mux, h.AdminListUser)
	MountAdminGetUserHandler(mux, h.AdminGetUser)
	MountAdminDeleteUserHandler(mux, h.AdminDeleteUser)
}

// MountAdminHealthCheckHandler configures the mux to serve the "Admin" service
// "admin health-check" endpoint.
func MountAdminHealthCheckHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/admin/healthcheck", f)
}

// NewAdminHealthCheckHandler creates a HTTP handler which loads the HTTP
// request and calls the "Admin" service "admin health-check" endpoint.
func NewAdminHealthCheckHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminHealthCheckRequest(mux, dec)
		encodeResponse = EncodeAdminHealthCheckResponse(enc)
		encodeError    = EncodeAdminHealthCheckError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin health-check")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminSigninHandler configures the mux to serve the "Admin" service
// "admin signin" endpoint.
func MountAdminSigninHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/admin/signin", f)
}

// NewAdminSigninHandler creates a HTTP handler which loads the HTTP request
// and calls the "Admin" service "admin signin" endpoint.
func NewAdminSigninHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminSigninRequest(mux, dec)
		encodeResponse = EncodeAdminSigninResponse(enc)
		encodeError    = EncodeAdminSigninError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin signin")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminCreateNewUserHandler configures the mux to serve the "Admin"
// service "admin create new user" endpoint.
func MountAdminCreateNewUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/admin/users", f)
}

// NewAdminCreateNewUserHandler creates a HTTP handler which loads the HTTP
// request and calls the "Admin" service "admin create new user" endpoint.
func NewAdminCreateNewUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminCreateNewUserRequest(mux, dec)
		encodeResponse = EncodeAdminCreateNewUserResponse(enc)
		encodeError    = EncodeAdminCreateNewUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin create new user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminUpdateUserHandler configures the mux to serve the "Admin" service
// "admin update user" endpoint.
func MountAdminUpdateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/admin/users/{user_id}", f)
}

// NewAdminUpdateUserHandler creates a HTTP handler which loads the HTTP
// request and calls the "Admin" service "admin update user" endpoint.
func NewAdminUpdateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminUpdateUserRequest(mux, dec)
		encodeResponse = EncodeAdminUpdateUserResponse(enc)
		encodeError    = EncodeAdminUpdateUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin update user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminListUserHandler configures the mux to serve the "Admin" service
// "admin list user" endpoint.
func MountAdminListUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/admin/users", f)
}

// NewAdminListUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "Admin" service "admin list user" endpoint.
func NewAdminListUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminListUserRequest(mux, dec)
		encodeResponse = EncodeAdminListUserResponse(enc)
		encodeError    = EncodeAdminListUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin list user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminGetUserHandler configures the mux to serve the "Admin" service
// "admin get user" endpoint.
func MountAdminGetUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/admin/users/{user_id}", f)
}

// NewAdminGetUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "Admin" service "admin get user" endpoint.
func NewAdminGetUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminGetUserRequest(mux, dec)
		encodeResponse = EncodeAdminGetUserResponse(enc)
		encodeError    = EncodeAdminGetUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin get user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountAdminDeleteUserHandler configures the mux to serve the "Admin" service
// "admin delete user" endpoint.
func MountAdminDeleteUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/admin/users/{user_id}", f)
}

// NewAdminDeleteUserHandler creates a HTTP handler which loads the HTTP
// request and calls the "Admin" service "admin delete user" endpoint.
func NewAdminDeleteUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAdminDeleteUserRequest(mux, dec)
		encodeResponse = EncodeAdminDeleteUserResponse(enc)
		encodeError    = EncodeAdminDeleteUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "admin delete user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
