// Code generated by goa v3.0.7, DO NOT EDIT.
//
// User service
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package user

import (
	"context"

	userviews "github.com/tonouchi510/Jeeek/gen/user/views"
	"goa.design/goa/v3/security"
)

// ユーザー/セッションに関するエンドポイントです。
type Service interface {
	// 現在のセッションに紐付くユーザの情報を返します。
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	//	- "admin"
	GetCurrentUser(context.Context, *SessionTokenPayload) (res *JeeekUser, view string, err error)
	// 現在のセッションに紐付くユーザー情報を更新します。
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	//	- "admin"
	UpdateUser(context.Context, *UpdateUserPayload) (res *JeeekUser, view string, err error)
	// ユーザーの一覧を返します。
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	//	- "admin"
	ListUser(context.Context, *SessionTokenPayload) (res JeeekUserCollection, view string, err error)
	// 指定したIDのユーザーの情報を返します。
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	//	- "admin"
	GetUser(context.Context, *GetUserPayload) (res *JeeekUser, view string, err error)
	// 現在のセッションに紐づくユーザーを削除します。
	DeleteUser(context.Context, *SessionTokenPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "User"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"Get current user", "Update user", "List user", "Get user", "Delete user"}

// SessionTokenPayload is the payload type of the User service Get current user
// method.
type SessionTokenPayload struct {
	// JWT used for authentication
	Token *string
}

// JeeekUser is the result type of the User service Get current user method.
type JeeekUser struct {
	// User id of firebase
	UserID string
	// ユーザーの表示名
	UserName string
	// ユーザーのプライマリ メールアドレス
	EmailAddress string
	// ユーザのメインの電話番号
	PhoneNumber string
	// ユーザーの写真 URL
	PhotoURL string
	// ユーザーのプライマリ メールアドレスが確認されているかどうか
	EmailVerified bool
	// ユーザが停止状態かどうか（論理削除）
	Disabled *bool
	// ユーザが作成された日時
	CreatedAt *string
	// 最後にログインした日時
	LastSignedinAt *string
}

// UpdateUserPayload is the payload type of the User service Update user method.
type UpdateUserPayload struct {
	// JWT used for authentication
	Token *string
	// ユーザーの表示名
	UserName *string
	// ユーザーのプライマリ メールアドレス
	EmailAddress *string
	// ユーザのメインの電話番号
	PhoneNumber *string
	// ユーザーの写真 URL
	PhotoURL *string
}

// JeeekUserCollection is the result type of the User service List user method.
type JeeekUserCollection []*JeeekUser

// GetUserPayload is the payload type of the User service Get user method.
type GetUserPayload struct {
	// JWT used for authentication
	Token *string
	// User id of firebase
	UserID string
}

// Credentials are invalid
type Unauthorized string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// NewJeeekUser initializes result type JeeekUser from viewed result type
// JeeekUser.
func NewJeeekUser(vres *userviews.JeeekUser) *JeeekUser {
	var res *JeeekUser
	switch vres.View {
	case "default", "":
		res = newJeeekUser(vres.Projected)
	case "tiny":
		res = newJeeekUserTiny(vres.Projected)
	case "admin":
		res = newJeeekUserAdmin(vres.Projected)
	}
	return res
}

// NewViewedJeeekUser initializes viewed result type JeeekUser from result type
// JeeekUser using the given view.
func NewViewedJeeekUser(res *JeeekUser, view string) *userviews.JeeekUser {
	var vres *userviews.JeeekUser
	switch view {
	case "default", "":
		p := newJeeekUserView(res)
		vres = &userviews.JeeekUser{p, "default"}
	case "tiny":
		p := newJeeekUserViewTiny(res)
		vres = &userviews.JeeekUser{p, "tiny"}
	case "admin":
		p := newJeeekUserViewAdmin(res)
		vres = &userviews.JeeekUser{p, "admin"}
	}
	return vres
}

// NewJeeekUserCollection initializes result type JeeekUserCollection from
// viewed result type JeeekUserCollection.
func NewJeeekUserCollection(vres userviews.JeeekUserCollection) JeeekUserCollection {
	var res JeeekUserCollection
	switch vres.View {
	case "default", "":
		res = newJeeekUserCollection(vres.Projected)
	case "tiny":
		res = newJeeekUserCollectionTiny(vres.Projected)
	case "admin":
		res = newJeeekUserCollectionAdmin(vres.Projected)
	}
	return res
}

// NewViewedJeeekUserCollection initializes viewed result type
// JeeekUserCollection from result type JeeekUserCollection using the given
// view.
func NewViewedJeeekUserCollection(res JeeekUserCollection, view string) userviews.JeeekUserCollection {
	var vres userviews.JeeekUserCollection
	switch view {
	case "default", "":
		p := newJeeekUserCollectionView(res)
		vres = userviews.JeeekUserCollection{p, "default"}
	case "tiny":
		p := newJeeekUserCollectionViewTiny(res)
		vres = userviews.JeeekUserCollection{p, "tiny"}
	case "admin":
		p := newJeeekUserCollectionViewAdmin(res)
		vres = userviews.JeeekUserCollection{p, "admin"}
	}
	return vres
}

// newJeeekUser converts projected type JeeekUser to service type JeeekUser.
func newJeeekUser(vres *userviews.JeeekUserView) *JeeekUser {
	res := &JeeekUser{}
	if vres.UserID != nil {
		res.UserID = *vres.UserID
	}
	if vres.UserName != nil {
		res.UserName = *vres.UserName
	}
	if vres.EmailAddress != nil {
		res.EmailAddress = *vres.EmailAddress
	}
	if vres.PhoneNumber != nil {
		res.PhoneNumber = *vres.PhoneNumber
	}
	if vres.PhotoURL != nil {
		res.PhotoURL = *vres.PhotoURL
	}
	if vres.EmailVerified != nil {
		res.EmailVerified = *vres.EmailVerified
	}
	return res
}

// newJeeekUserTiny converts projected type JeeekUser to service type JeeekUser.
func newJeeekUserTiny(vres *userviews.JeeekUserView) *JeeekUser {
	res := &JeeekUser{}
	if vres.UserID != nil {
		res.UserID = *vres.UserID
	}
	if vres.UserName != nil {
		res.UserName = *vres.UserName
	}
	if vres.EmailAddress != nil {
		res.EmailAddress = *vres.EmailAddress
	}
	return res
}

// newJeeekUserAdmin converts projected type JeeekUser to service type
// JeeekUser.
func newJeeekUserAdmin(vres *userviews.JeeekUserView) *JeeekUser {
	res := &JeeekUser{
		Disabled:       vres.Disabled,
		CreatedAt:      vres.CreatedAt,
		LastSignedinAt: vres.LastSignedinAt,
	}
	if vres.UserID != nil {
		res.UserID = *vres.UserID
	}
	if vres.UserName != nil {
		res.UserName = *vres.UserName
	}
	if vres.EmailAddress != nil {
		res.EmailAddress = *vres.EmailAddress
	}
	if vres.PhoneNumber != nil {
		res.PhoneNumber = *vres.PhoneNumber
	}
	if vres.PhotoURL != nil {
		res.PhotoURL = *vres.PhotoURL
	}
	if vres.EmailVerified != nil {
		res.EmailVerified = *vres.EmailVerified
	}
	return res
}

// newJeeekUserView projects result type JeeekUser to projected type
// JeeekUserView using the "default" view.
func newJeeekUserView(res *JeeekUser) *userviews.JeeekUserView {
	vres := &userviews.JeeekUserView{
		UserID:        &res.UserID,
		UserName:      &res.UserName,
		EmailAddress:  &res.EmailAddress,
		PhoneNumber:   &res.PhoneNumber,
		PhotoURL:      &res.PhotoURL,
		EmailVerified: &res.EmailVerified,
	}
	return vres
}

// newJeeekUserViewTiny projects result type JeeekUser to projected type
// JeeekUserView using the "tiny" view.
func newJeeekUserViewTiny(res *JeeekUser) *userviews.JeeekUserView {
	vres := &userviews.JeeekUserView{
		UserID:       &res.UserID,
		UserName:     &res.UserName,
		EmailAddress: &res.EmailAddress,
	}
	return vres
}

// newJeeekUserViewAdmin projects result type JeeekUser to projected type
// JeeekUserView using the "admin" view.
func newJeeekUserViewAdmin(res *JeeekUser) *userviews.JeeekUserView {
	vres := &userviews.JeeekUserView{
		UserID:         &res.UserID,
		UserName:       &res.UserName,
		EmailAddress:   &res.EmailAddress,
		PhoneNumber:    &res.PhoneNumber,
		PhotoURL:       &res.PhotoURL,
		EmailVerified:  &res.EmailVerified,
		Disabled:       res.Disabled,
		CreatedAt:      res.CreatedAt,
		LastSignedinAt: res.LastSignedinAt,
	}
	return vres
}

// newJeeekUserCollection converts projected type JeeekUserCollection to
// service type JeeekUserCollection.
func newJeeekUserCollection(vres userviews.JeeekUserCollectionView) JeeekUserCollection {
	res := make(JeeekUserCollection, len(vres))
	for i, n := range vres {
		res[i] = newJeeekUser(n)
	}
	return res
}

// newJeeekUserCollectionTiny converts projected type JeeekUserCollection to
// service type JeeekUserCollection.
func newJeeekUserCollectionTiny(vres userviews.JeeekUserCollectionView) JeeekUserCollection {
	res := make(JeeekUserCollection, len(vres))
	for i, n := range vres {
		res[i] = newJeeekUserTiny(n)
	}
	return res
}

// newJeeekUserCollectionAdmin converts projected type JeeekUserCollection to
// service type JeeekUserCollection.
func newJeeekUserCollectionAdmin(vres userviews.JeeekUserCollectionView) JeeekUserCollection {
	res := make(JeeekUserCollection, len(vres))
	for i, n := range vres {
		res[i] = newJeeekUserAdmin(n)
	}
	return res
}

// newJeeekUserCollectionView projects result type JeeekUserCollection to
// projected type JeeekUserCollectionView using the "default" view.
func newJeeekUserCollectionView(res JeeekUserCollection) userviews.JeeekUserCollectionView {
	vres := make(userviews.JeeekUserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newJeeekUserView(n)
	}
	return vres
}

// newJeeekUserCollectionViewTiny projects result type JeeekUserCollection to
// projected type JeeekUserCollectionView using the "tiny" view.
func newJeeekUserCollectionViewTiny(res JeeekUserCollection) userviews.JeeekUserCollectionView {
	vres := make(userviews.JeeekUserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newJeeekUserViewTiny(n)
	}
	return vres
}

// newJeeekUserCollectionViewAdmin projects result type JeeekUserCollection to
// projected type JeeekUserCollectionView using the "admin" view.
func newJeeekUserCollectionViewAdmin(res JeeekUserCollection) userviews.JeeekUserCollectionView {
	vres := make(userviews.JeeekUserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newJeeekUserViewAdmin(n)
	}
	return vres
}
