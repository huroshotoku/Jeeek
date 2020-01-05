// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Activity endpoints
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package activity

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "Activity" service endpoints.
type Endpoints struct {
	ReflectionActivity goa.Endpoint
}

// NewEndpoints wraps the methods of the "Activity" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		ReflectionActivity: NewReflectionActivityEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "Activity" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ReflectionActivity = m(e.ReflectionActivity)
}

// NewReflectionActivityEndpoint returns an endpoint function that calls the
// method "Reflection activity" of service "Activity".
func NewReflectionActivityEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ActivityPostPayload)
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
		return nil, s.ReflectionActivity(ctx, p)
	}
}
