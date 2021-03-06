// Code generated by goa v3.0.7, DO NOT EDIT.
//
// User client HTTP transport
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the User service endpoint HTTP clients.
type Client struct {
	// GetCurrentUser Doer is the HTTP client used to make requests to the Get
	// current user endpoint.
	GetCurrentUserDoer goahttp.Doer

	// UpdateUser Doer is the HTTP client used to make requests to the Update user
	// endpoint.
	UpdateUserDoer goahttp.Doer

	// ListUser Doer is the HTTP client used to make requests to the List user
	// endpoint.
	ListUserDoer goahttp.Doer

	// GetUser Doer is the HTTP client used to make requests to the Get user
	// endpoint.
	GetUserDoer goahttp.Doer

	// DeleteUser Doer is the HTTP client used to make requests to the Delete user
	// endpoint.
	DeleteUserDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the User service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetCurrentUserDoer:  doer,
		UpdateUserDoer:      doer,
		ListUserDoer:        doer,
		GetUserDoer:         doer,
		DeleteUserDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GetCurrentUser returns an endpoint that makes HTTP requests to the User
// service Get current user server.
func (c *Client) GetCurrentUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetCurrentUserRequest(c.encoder)
		decodeResponse = DecodeGetCurrentUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetCurrentUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetCurrentUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "Get current user", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateUser returns an endpoint that makes HTTP requests to the User service
// Update user server.
func (c *Client) UpdateUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateUserRequest(c.encoder)
		decodeResponse = DecodeUpdateUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "Update user", err)
		}
		return decodeResponse(resp)
	}
}

// ListUser returns an endpoint that makes HTTP requests to the User service
// List user server.
func (c *Client) ListUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeListUserRequest(c.encoder)
		decodeResponse = DecodeListUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "List user", err)
		}
		return decodeResponse(resp)
	}
}

// GetUser returns an endpoint that makes HTTP requests to the User service Get
// user server.
func (c *Client) GetUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetUserRequest(c.encoder)
		decodeResponse = DecodeGetUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "Get user", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteUser returns an endpoint that makes HTTP requests to the User service
// Delete user server.
func (c *Client) DeleteUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteUserRequest(c.encoder)
		decodeResponse = DecodeDeleteUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "Delete user", err)
		}
		return decodeResponse(resp)
	}
}
