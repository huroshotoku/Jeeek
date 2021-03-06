package design

import (
	. "goa.design/goa/v3/dsl"
)

var digitPattern = `^[0-9]+$`
var phoneNumberPattern = `^\+?[\d]{10,}$`

var IDKeyDefinition = func() {
	Description("IDKey of datastore")
	Example(5644004762845184)
}

var JWT = Type("JWT Token", func() {
	Token("token", String, func() {
		Description("JWT used for authentication")
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
})

var SessionTokenPayload = Type("SessionTokenPayload", func() {
	Reference(JWT)
	Token("token")
})


var TimeStamp = Type("TimeStamp", func() {
	Attribute("created_at", String, func() {
		Example("yyyy-mm-dd:xx:yy:zz")
	})
	Attribute("updated_at", String, func() {
		Example("yyyy-mm-dd:xx:yy:zz")
	})
	Attribute("deleted_at", String, func() {
		Example("yyyy-mm-dd:xx:yy:zz")
	})
})


//Admin
var AdminSignInPayload = Type("AdminSignInPayload", func() {
	Attribute("uid", String, func() {
		Example("4WIbqiNIpIYXqrfBMVZsbKCepau1")
	})
	Attribute("password", String, func() {
		Example("b09jw04i1024")
	})
	Required("uid", "password")
})

var AdminCreateUserPayload = Type("AdminCreateUserPayload", func(){
	Reference(JWT)
	Reference(UserProfile)
	Token("token")

	Attribute("user_name")
	Attribute("email_address")
	Attribute("phone_number")
	Attribute("photo_url")
	Required("user_name", "email_address", "phone_number", "photo_url")
})

var AdminUpdateUserPayload = Type("AdminUpdateUserPayload", func() {
	Reference(JWT)
	Token("token")
	Reference(UserProfile)
	Attribute("user_id")
	Attribute("user_name")
	Attribute("email_address")
	Attribute("phone_number")
	Attribute("photo_url")
	Attribute("email_verified")
	Attribute("disabled")

	Required("user_id")
})

var AdminDeleteUserPayload = Type("AdminDeleteUserPayload", func() {
	Reference(JWT)
	Token("token")
	Attribute("user_id", String, UserIDDefinition)
	Required("user_id")
})


// User
var UserProfile = Type("UserProfile", func() {
	Attribute("user_id", String, UserIDDefinition)

	Attribute("user_name", String, func() {
		Description("ユーザーの表示名")
		MinLength(1)
		MaxLength(20)
		Example("keisuke.honda")
	})

	Attribute("email_address", String, func() {
		Description("ユーザーのプライマリ メールアドレス")
		Format("email")
		Example("keisuke.honda+testuser@ynu.jp")
	})

	Attribute("phone_number", String, func() {
		Description("ユーザのメインの電話番号")
		Pattern(phoneNumberPattern)
		Example("08079469367")
	})

	Attribute("photo_url", String, func() {
		Description("ユーザーの写真 URL")
		Example("https://imageurl.com")
	})

	Attribute("email_verified", Boolean, func() {
		Description("ユーザーのプライマリ メールアドレスが確認されているかどうか")
	})

	Attribute("disabled", Boolean, func() {
		Description("ユーザが停止状態かどうか（論理削除）")
	})

	Attribute("created_at", String, func() {
		Description("ユーザが作成された日時")
	})

	Attribute("last_signedin_at", String, func() {
		Description("最後にログインした日時")
	})

})

var UserIDDefinition = func() {
	Description("User id of firebase")
	Example("XRQ85mtXnINISH25zfM0m5RlC6L2")
	MinLength(28)
	MaxLength(28)
}

var GetUserPayload = Type("GetUserPayload", func() {
	Reference(JWT)
	Token("token")
	Attribute("user_id", String, UserIDDefinition)
	Required("user_id")
})

var UpdateUserPayload = Type("UpdateUserPayload", func() {
	Reference(JWT)
	Token("token")
	Reference(UserProfile)
	Attribute("user_name")
	Attribute("email_address")
	Attribute("phone_number")
	Attribute("photo_url")
})

var ActivityPostPayload = Type("ActivityPostPayload", func() {
	Reference(JWT)
	Token("token")
	Reference(UserProfile)
	Attribute("Activity", Activity)
})

var Activity = Type("Activity", func() {
	Attribute("id", String, func() {
		Description("投稿のID（Firestore上ではドキュメントIDになる）")
		Example("0000abcds6z57pqbpkin")
	})
	Attribute("userTiny", UserTiny)
	Attribute("category", Int, func() {
		Description("投稿のカテゴリー（0: 学習, 1: 開発, 2: 執筆, 3: 賞等）")
		Example(0)
	})
	Attribute("rank", Int, func() {
		Description("投稿のランク（0~3 -> C~S に対応してレベルを設定）")
		Example(0)
	})
	Attribute("content", Content)
	Attribute("tags", ArrayOf(String), func() {
		Description("投稿に紐づく技術タグを設定する")
		Example([]string{"Golang", "GCP"})
	})
	Attribute("favorites", ArrayOf(String), func() {
		Description("投稿に対して'いいね'したユーザのUID")
		Example([]string{"4sra3r4zibfrzp4i", "akkynv4v3v8d5evx"})
	})
	Attribute("gifts", ArrayOf(String), func() {
		Description("投稿に対して'Gifting'したユーザのUID")
		Example([]string{})
	})

	Required("id", "userTiny", "category", "rank", "content", "tags", "favorites", "gifts")
})

var UserTiny = Type("UserTiny", func() {
	Attribute("uid", String, func() {
		Description("投稿したユーザのUID")
		Example("p2qfpb2gvxrzedu2")
	})
	Attribute("name", String, func() {
		Description("投稿したユーザの名前")
		Example("トノウチ")
	})
	Attribute("photoUrl", String, func() {
		Description("投稿したユーザの写真Url")
		Example("https://storage.tenki.jp/storage/static-images/suppl/article/image/9/97/971/9711/1/large.jpg")
	})

	Required("uid", "name")
})

var Content = Type("Content", func() {
	Attribute("subject", String, func() {
		Description("投稿の主題")
		Example("PRML本の4章を読んだ。")
	})
	Attribute("url", String, func() {
		Description("投稿に関連するUrl（オプション）")
		Example("https://www.amazon.co.jp/パターン認識と機械学習-上-C-M-ビショップ/dp/4621061224")
	})
	Attribute("comment", String, func() {
		Description("投稿についての自由記述欄")
		Example("ロジスティック回帰が使われている理由がよくわかった")
	})

	Required("subject")
})
