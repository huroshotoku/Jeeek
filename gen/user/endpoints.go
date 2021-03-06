// Code generated by goa v3.0.7, DO NOT EDIT.
//
// User endpoints
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "User" service endpoints.
type Endpoints struct {
	GetCurrentUser goa.Endpoint
	UpdateUser     goa.Endpoint
	ListUser       goa.Endpoint
	GetUser        goa.Endpoint
	DeleteUser     goa.Endpoint
}

// NewEndpoints wraps the methods of the "User" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		GetCurrentUser: NewGetCurrentUserEndpoint(s, a.JWTAuth),
		UpdateUser:     NewUpdateUserEndpoint(s, a.JWTAuth),
		ListUser:       NewListUserEndpoint(s, a.JWTAuth),
		GetUser:        NewGetUserEndpoint(s, a.JWTAuth),
		DeleteUser:     NewDeleteUserEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "User" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetCurrentUser = m(e.GetCurrentUser)
	e.UpdateUser = m(e.UpdateUser)
	e.ListUser = m(e.ListUser)
	e.GetUser = m(e.GetUser)
	e.DeleteUser = m(e.DeleteUser)
}

// NewGetCurrentUserEndpoint returns an endpoint function that calls the method
// "Get current user" of service "User".
func NewGetCurrentUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SessionTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.GetCurrentUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekUser(res, view)
		return vres, nil
	}
}

// NewUpdateUserEndpoint returns an endpoint function that calls the method
// "Update user" of service "User".
func NewUpdateUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.UpdateUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekUser(res, view)
		return vres, nil
	}
}

// NewListUserEndpoint returns an endpoint function that calls the method "List
// user" of service "User".
func NewListUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SessionTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.ListUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekUserCollection(res, view)
		return vres, nil
	}
}

// NewGetUserEndpoint returns an endpoint function that calls the method "Get
// user" of service "User".
func NewGetUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.GetUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekUser(res, view)
		return vres, nil
	}
}

// NewDeleteUserEndpoint returns an endpoint function that calls the method
// "Delete user" of service "User".
func NewDeleteUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SessionTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteUser(ctx, p)
	}
}
