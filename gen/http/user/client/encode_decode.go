// Code generated by goa v3.0.4, DO NOT EDIT.
//
// User HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	user "github.com/tonouchi510/Jeeek/gen/user"
	userviews "github.com/tonouchi510/Jeeek/gen/user/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetCurrentUserRequest instantiates a HTTP request object with method
// and path set to call the "User" service "Get current user" endpoint
func (c *Client) BuildGetCurrentUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetCurrentUserUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "Get current user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetCurrentUserRequest returns an encoder for requests sent to the User
// Get current user server.
func EncodeGetCurrentUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.SessionTokenPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "Get current user", "*user.SessionTokenPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		return nil
	}
}

// DecodeGetCurrentUserResponse returns a decoder for responses returned by the
// User Get current user endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeGetCurrentUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeGetCurrentUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetCurrentUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Get current user", err)
			}
			p := NewGetCurrentUserJeeekUserOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &userviews.JeeekUser{p, view}
			if err = userviews.ValidateJeeekUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("User", "Get current user", err)
			}
			res := user.NewJeeekUser(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetCurrentUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Get current user", err)
			}
			return nil, NewGetCurrentUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "Get current user", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateUserRequest instantiates a HTTP request object with method and
// path set to call the "User" service "Update user" endpoint
func (c *Client) BuildUpdateUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateUserUserPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "Update user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateUserRequest returns an encoder for requests sent to the User
// Update user server.
func EncodeUpdateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.UpdateUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "Update user", "*user.UpdateUserPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		body := NewUpdateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("User", "Update user", err)
		}
		return nil
	}
}

// DecodeUpdateUserResponse returns a decoder for responses returned by the
// User Update user endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeUpdateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Update user", err)
			}
			p := NewUpdateUserJeeekUserOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &userviews.JeeekUser{p, view}
			if err = userviews.ValidateJeeekUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("User", "Update user", err)
			}
			res := user.NewJeeekUser(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UpdateUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Update user", err)
			}
			return nil, NewUpdateUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "Update user", resp.StatusCode, string(body))
		}
	}
}

// BuildListUserRequest instantiates a HTTP request object with method and path
// set to call the "User" service "List user" endpoint
func (c *Client) BuildListUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListUserUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "List user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListUserRequest returns an encoder for requests sent to the User List
// user server.
func EncodeListUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.SessionTokenPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "List user", "*user.SessionTokenPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		return nil
	}
}

// DecodeListUserResponse returns a decoder for responses returned by the User
// List user endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeListUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "List user", err)
			}
			p := NewListUserJeeekUserCollectionOK(body)
			view := resp.Header.Get("goa-view")
			vres := userviews.JeeekUserCollection{p, view}
			if err = userviews.ValidateJeeekUserCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("User", "List user", err)
			}
			res := user.NewJeeekUserCollection(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ListUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "List user", err)
			}
			return nil, NewListUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "List user", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserRequest instantiates a HTTP request object with method and path
// set to call the "User" service "Get user" endpoint
func (c *Client) BuildGetUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID string
	)
	{
		p, ok := v.(*user.GetUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("User", "Get user", "*user.GetUserPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserUserPath(userID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "Get user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetUserRequest returns an encoder for requests sent to the User Get
// user server.
func EncodeGetUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.GetUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "Get user", "*user.GetUserPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		return nil
	}
}

// DecodeGetUserResponse returns a decoder for responses returned by the User
// Get user endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeGetUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Get user", err)
			}
			p := NewGetUserJeeekUserOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &userviews.JeeekUser{p, view}
			if err = userviews.ValidateJeeekUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("User", "Get user", err)
			}
			res := user.NewJeeekUser(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Get user", err)
			}
			return nil, NewGetUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "Get user", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteUserRequest instantiates a HTTP request object with method and
// path set to call the "User" service "Delete user" endpoint
func (c *Client) BuildDeleteUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteUserUserPath()}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "Delete user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteUserRequest returns an encoder for requests sent to the User
// Delete user server.
func EncodeDeleteUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.SessionTokenPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "Delete user", "*user.SessionTokenPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		return nil
	}
}

// DecodeDeleteUserResponse returns a decoder for responses returned by the
// User Delete user endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeDeleteUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusUnauthorized:
			var (
				body DeleteUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Delete user", err)
			}
			return nil, NewDeleteUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "Delete user", resp.StatusCode, string(body))
		}
	}
}