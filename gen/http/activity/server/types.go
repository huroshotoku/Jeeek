// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Activity HTTP server types
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package server

import (
	activity "github.com/tonouchi510/Jeeek/gen/activity"
)

// FetchQiitaArticleByQiitaUserIDUnauthorizedResponseBody is the type of the
// "Activity" service "Fetch qiita article by qiita-user-id" endpoint HTTP
// response body for the "unauthorized" error.
type FetchQiitaArticleByQiitaUserIDUnauthorizedResponseBody string

// NewFetchQiitaArticleByQiitaUserIDUnauthorizedResponseBody builds the HTTP
// response body from the result of the "Fetch qiita article by qiita-user-id"
// endpoint of the "Activity" service.
func NewFetchQiitaArticleByQiitaUserIDUnauthorizedResponseBody(res activity.Unauthorized) FetchQiitaArticleByQiitaUserIDUnauthorizedResponseBody {
	body := FetchQiitaArticleByQiitaUserIDUnauthorizedResponseBody(res)
	return body
}

// NewFetchQiitaArticleByQiitaUserIDGetActivityPayload builds a Activity
// service Fetch qiita article by qiita-user-id endpoint payload.
func NewFetchQiitaArticleByQiitaUserIDGetActivityPayload(userID string, token *string) *activity.GetActivityPayload {
	return &activity.GetActivityPayload{
		UserID: userID,
		Token:  token,
	}
}
