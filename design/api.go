package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("Jeeek", func() {
	Title("JeeekAPI")
	Description("エンジニアの活動支援プラットフォーム")
	Version("0.1")
	Contact(func() {
		Name("tonouchi510")
		Email("tonouchi27@gmail.com")
		URL("https://github.com/tonouchi510/Jeeek/issues")
	})
	Docs(func() {
		Description("Docs")
		URL("https://github.com/tonouchi510/Jeeek/wiki")
	})

	Server("JeeekAPI", func() {
		Description("this hosts the api of Jeeek service.")

		Services("Admin", "User", "Activity")

		Host("host", func() {
			URI("http://localhost:8080")
		})
	})
})

var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the firebase.`)
})
