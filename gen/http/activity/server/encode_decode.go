// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Activity HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package server

import (
	"context"
	"io"
	"net/http"
	"strings"

	activity "github.com/tonouchi510/Jeeek/gen/activity"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeManualPostOfActivityResponse returns an encoder for responses returned
// by the Activity Manual post of activity endpoint.
func EncodeManualPostOfActivityResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeManualPostOfActivityRequest returns a decoder for requests sent to the
// Activity Manual post of activity endpoint.
func DecodeManualPostOfActivityRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ManualPostOfActivityRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateManualPostOfActivityRequestBody(&body)
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
		payload := NewManualPostOfActivityActivityPostPayload(&body, token)
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

// EncodeManualPostOfActivityError returns an encoder for errors returned by
// the Manual post of activity Activity endpoint.
func EncodeManualPostOfActivityError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(activity.Unauthorized)
			enc := encoder(ctx, w)
			body := NewManualPostOfActivityUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRefreshActivitiesOfAllCooperationServicesResponse returns an encoder
// for responses returned by the Activity Refresh activities of all cooperation
// services endpoint.
func EncodeRefreshActivitiesOfAllCooperationServicesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeRefreshActivitiesOfAllCooperationServicesRequest returns a decoder for
// requests sent to the Activity Refresh activities of all cooperation services
// endpoint.
func DecodeRefreshActivitiesOfAllCooperationServicesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewRefreshActivitiesOfAllCooperationServicesSessionTokenPayload(token)
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

// EncodeRefreshActivitiesOfAllCooperationServicesError returns an encoder for
// errors returned by the Refresh activities of all cooperation services
// Activity endpoint.
func EncodeRefreshActivitiesOfAllCooperationServicesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(activity.Unauthorized)
			enc := encoder(ctx, w)
			body := NewRefreshActivitiesOfAllCooperationServicesUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRefreshQiitaActivitiesResponse returns an encoder for responses
// returned by the Activity Refresh qiita activities endpoint.
func EncodeRefreshQiitaActivitiesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeRefreshQiitaActivitiesRequest returns a decoder for requests sent to
// the Activity Refresh qiita activities endpoint.
func DecodeRefreshQiitaActivitiesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewRefreshQiitaActivitiesSessionTokenPayload(token)
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

// EncodeRefreshQiitaActivitiesError returns an encoder for errors returned by
// the Refresh qiita activities Activity endpoint.
func EncodeRefreshQiitaActivitiesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(activity.Unauthorized)
			enc := encoder(ctx, w)
			body := NewRefreshQiitaActivitiesUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodePickOutAllPastActivitiesOfQiitaResponse returns an encoder for
// responses returned by the Activity Pick out all past activities of qiita
// endpoint.
func EncodePickOutAllPastActivitiesOfQiitaResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodePickOutAllPastActivitiesOfQiitaRequest returns a decoder for requests
// sent to the Activity Pick out all past activities of qiita endpoint.
func DecodePickOutAllPastActivitiesOfQiitaRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewPickOutAllPastActivitiesOfQiitaSessionTokenPayload(token)
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

// EncodePickOutAllPastActivitiesOfQiitaError returns an encoder for errors
// returned by the Pick out all past activities of qiita Activity endpoint.
func EncodePickOutAllPastActivitiesOfQiitaError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(activity.Unauthorized)
			enc := encoder(ctx, w)
			body := NewPickOutAllPastActivitiesOfQiitaUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// unmarshalActivityRequestBodyToActivityActivity builds a value of type
// *activity.Activity from a value of type *ActivityRequestBody.
func unmarshalActivityRequestBodyToActivityActivity(v *ActivityRequestBody) *activity.Activity {
	if v == nil {
		return nil
	}
	res := &activity.Activity{
		ID:       *v.ID,
		Category: *v.Category,
		Rank:     *v.Rank,
	}
	res.UserTiny = unmarshalUserTinyRequestBodyToActivityUserTiny(v.UserTiny)
	res.Content = unmarshalContentRequestBodyToActivityContent(v.Content)
	res.Tags = make([]string, len(v.Tags))
	for i, val := range v.Tags {
		res.Tags[i] = val
	}
	res.Favorites = make([]string, len(v.Favorites))
	for i, val := range v.Favorites {
		res.Favorites[i] = val
	}
	res.Gifts = make([]string, len(v.Gifts))
	for i, val := range v.Gifts {
		res.Gifts[i] = val
	}

	return res
}

// unmarshalUserTinyRequestBodyToActivityUserTiny builds a value of type
// *activity.UserTiny from a value of type *UserTinyRequestBody.
func unmarshalUserTinyRequestBodyToActivityUserTiny(v *UserTinyRequestBody) *activity.UserTiny {
	res := &activity.UserTiny{
		UID:      *v.UID,
		Name:     *v.Name,
		PhotoURL: v.PhotoURL,
	}

	return res
}

// unmarshalContentRequestBodyToActivityContent builds a value of type
// *activity.Content from a value of type *ContentRequestBody.
func unmarshalContentRequestBodyToActivityContent(v *ContentRequestBody) *activity.Content {
	res := &activity.Content{
		Subject: *v.Subject,
		URL:     v.URL,
		Comment: v.Comment,
	}

	return res
}
