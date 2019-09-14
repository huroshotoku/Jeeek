// Code generated by goa v3.0.4, DO NOT EDIT.
//
// Admin HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package server

import (
	"context"
	"io"
	"net/http"
	"strings"
	"unicode/utf8"

	admin "github.com/tonouchi510/Jeeek/gen/admin"
	adminviews "github.com/tonouchi510/Jeeek/gen/admin/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAdminHealthCheckResponse returns an encoder for responses returned by
// the Admin admin health-check endpoint.
func EncodeAdminHealthCheckResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*adminviews.JeeekHealthcheck)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json; charset=utf-8")
		enc := encoder(ctx, w)
		body := NewAdminHealthCheckResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAdminHealthCheckRequest returns a decoder for requests sent to the
// Admin admin health-check endpoint.
func DecodeAdminHealthCheckRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewAdminHealthCheckSessionTokenPayload(token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeAdminHealthCheckError returns an encoder for errors returned by the
// admin health-check Admin endpoint.
func EncodeAdminHealthCheckError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(admin.Unauthorized)
			enc := encoder(ctx, w)
			body := NewAdminHealthCheckUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAdminSigninResponse returns an encoder for responses returned by the
// Admin admin signin endpoint.
func EncodeAdminSigninResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*adminviews.JeeekAdminSignin)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json; charset=utf-8")
		enc := encoder(ctx, w)
		body := NewAdminSigninResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAdminSigninRequest returns a decoder for requests sent to the Admin
// admin signin endpoint.
func DecodeAdminSigninRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AdminSigninRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAdminSigninRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAdminSigninAdminSignInPayload(&body)

		return payload, nil
	}
}

// EncodeAdminSigninError returns an encoder for errors returned by the admin
// signin Admin endpoint.
func EncodeAdminSigninError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(admin.Unauthorized)
			enc := encoder(ctx, w)
			body := NewAdminSigninUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAdminCreateNewUserResponse returns an encoder for responses returned
// by the Admin admin create new user endpoint.
func EncodeAdminCreateNewUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*adminviews.JeeekUser)
		w.Header().Set("goa-view", res.View)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json; charset=utf-8")
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewAdminCreateNewUserResponseBody(res.Projected)
		case "tiny":
			body = NewAdminCreateNewUserResponseBodyTiny(res.Projected)
		case "admin":
			body = NewAdminCreateNewUserResponseBodyAdmin(res.Projected)
		}
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAdminCreateNewUserRequest returns a decoder for requests sent to the
// Admin admin create new user endpoint.
func DecodeAdminCreateNewUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AdminCreateNewUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAdminCreateNewUserRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewAdminCreateNewUserAdminCreateUserPayload(&body, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeAdminCreateNewUserError returns an encoder for errors returned by the
// admin create new user Admin endpoint.
func EncodeAdminCreateNewUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(admin.Unauthorized)
			enc := encoder(ctx, w)
			body := NewAdminCreateNewUserUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAdminUpdateUserResponse returns an encoder for responses returned by
// the Admin admin update user endpoint.
func EncodeAdminUpdateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*adminviews.JeeekUser)
		w.Header().Set("goa-view", res.View)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json; charset=utf-8")
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewAdminUpdateUserResponseBody(res.Projected)
		case "tiny":
			body = NewAdminUpdateUserResponseBodyTiny(res.Projected)
		case "admin":
			body = NewAdminUpdateUserResponseBodyAdmin(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAdminUpdateUserRequest returns a decoder for requests sent to the
// Admin admin update user endpoint.
func DecodeAdminUpdateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AdminUpdateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAdminUpdateUserRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			userID string
			token  *string

			params = mux.Vars(r)
		)
		userID = params["user_id"]
		if utf8.RuneCountInString(userID) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("userID", userID, utf8.RuneCountInString(userID), 28, true))
		}
		if utf8.RuneCountInString(userID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("userID", userID, utf8.RuneCountInString(userID), 28, false))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewAdminUpdateUserPayload(&body, userID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeAdminUpdateUserError returns an encoder for errors returned by the
// admin update user Admin endpoint.
func EncodeAdminUpdateUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(admin.Unauthorized)
			enc := encoder(ctx, w)
			body := NewAdminUpdateUserUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAdminListUserResponse returns an encoder for responses returned by the
// Admin admin list user endpoint.
func EncodeAdminListUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(adminviews.JeeekUserCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewJeeekUserResponseCollection(res.Projected)
		case "tiny":
			body = NewJeeekUserResponseTinyCollection(res.Projected)
		case "admin":
			body = NewJeeekUserResponseAdminCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAdminListUserRequest returns a decoder for requests sent to the Admin
// admin list user endpoint.
func DecodeAdminListUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewAdminListUserSessionTokenPayload(token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeAdminListUserError returns an encoder for errors returned by the admin
// list user Admin endpoint.
func EncodeAdminListUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(admin.Unauthorized)
			enc := encoder(ctx, w)
			body := NewAdminListUserUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAdminGetUserResponse returns an encoder for responses returned by the
// Admin admin get user endpoint.
func EncodeAdminGetUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*adminviews.JeeekUser)
		w.Header().Set("goa-view", res.View)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json; charset=utf-8")
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewAdminGetUserResponseBody(res.Projected)
		case "tiny":
			body = NewAdminGetUserResponseBodyTiny(res.Projected)
		case "admin":
			body = NewAdminGetUserResponseBodyAdmin(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAdminGetUserRequest returns a decoder for requests sent to the Admin
// admin get user endpoint.
func DecodeAdminGetUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID string
			token  *string
			err    error

			params = mux.Vars(r)
		)
		userID = params["user_id"]
		if utf8.RuneCountInString(userID) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("userID", userID, utf8.RuneCountInString(userID), 28, true))
		}
		if utf8.RuneCountInString(userID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("userID", userID, utf8.RuneCountInString(userID), 28, false))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewAdminGetUserGetUserPayload(userID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeAdminGetUserError returns an encoder for errors returned by the admin
// get user Admin endpoint.
func EncodeAdminGetUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(admin.Unauthorized)
			enc := encoder(ctx, w)
			body := NewAdminGetUserUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAdminDeleteUserResponse returns an encoder for responses returned by
// the Admin admin delete user endpoint.
func EncodeAdminDeleteUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeAdminDeleteUserRequest returns a decoder for requests sent to the
// Admin admin delete user endpoint.
func DecodeAdminDeleteUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID string
			token  *string
			err    error

			params = mux.Vars(r)
		)
		userID = params["user_id"]
		if utf8.RuneCountInString(userID) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("userID", userID, utf8.RuneCountInString(userID), 28, true))
		}
		if utf8.RuneCountInString(userID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("userID", userID, utf8.RuneCountInString(userID), 28, false))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewAdminDeleteUserPayload(userID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeAdminDeleteUserError returns an encoder for errors returned by the
// admin delete user Admin endpoint.
func EncodeAdminDeleteUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(admin.Unauthorized)
			enc := encoder(ctx, w)
			body := NewAdminDeleteUserUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}