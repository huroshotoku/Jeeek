// Code generated by goa v3.0.7, DO NOT EDIT.
//
// User client
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "User" service client.
type Client struct {
	GetCurrentUserEndpoint goa.Endpoint
	UpdateUserEndpoint     goa.Endpoint
	ListUserEndpoint       goa.Endpoint
	GetUserEndpoint        goa.Endpoint
	DeleteUserEndpoint     goa.Endpoint
}

// NewClient initializes a "User" service client given the endpoints.
func NewClient(getCurrentUser, updateUser, listUser, getUser, deleteUser goa.Endpoint) *Client {
	return &Client{
		GetCurrentUserEndpoint: getCurrentUser,
		UpdateUserEndpoint:     updateUser,
		ListUserEndpoint:       listUser,
		GetUserEndpoint:        getUser,
		DeleteUserEndpoint:     deleteUser,
	}
}

// GetCurrentUser calls the "Get current user" endpoint of the "User" service.
func (c *Client) GetCurrentUser(ctx context.Context, p *SessionTokenPayload) (res *JeeekUser, err error) {
	var ires interface{}
	ires, err = c.GetCurrentUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*JeeekUser), nil
}

// UpdateUser calls the "Update user" endpoint of the "User" service.
func (c *Client) UpdateUser(ctx context.Context, p *UpdateUserPayload) (res *JeeekUser, err error) {
	var ires interface{}
	ires, err = c.UpdateUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*JeeekUser), nil
}

// ListUser calls the "List user" endpoint of the "User" service.
func (c *Client) ListUser(ctx context.Context, p *SessionTokenPayload) (res JeeekUserCollection, err error) {
	var ires interface{}
	ires, err = c.ListUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(JeeekUserCollection), nil
}

// GetUser calls the "Get user" endpoint of the "User" service.
func (c *Client) GetUser(ctx context.Context, p *GetUserPayload) (res *JeeekUser, err error) {
	var ires interface{}
	ires, err = c.GetUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*JeeekUser), nil
}

// DeleteUser calls the "Delete user" endpoint of the "User" service.
func (c *Client) DeleteUser(ctx context.Context, p *SessionTokenPayload) (err error) {
	_, err = c.DeleteUserEndpoint(ctx, p)
	return
}
