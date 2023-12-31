package design

import (
	. "goa.design/goa/v3/dsl"
)

// API 定義
var _ = API("diary", func() {
	// API の説明（タイトルと説明）
	Title("Diary Service")
	Description("Service for Diary")

	// サーバ定義
	Server("diary", func() {
		Host("localhost", func() {
			URI("http://localhost:8000") // HTTP REST API
			URI("grpc://localhost:8080") // gRPC
		})
	})
})

// サービス定義
// var _ = Service("calc", func() {
// 	// 説明
// 	Description("The calc service performs operations on numbers.")

// 	// メソッド (HTTPでいうところのエンドポントに相当)
// 	Method("add", func() {
// 		// ペイロード定義
// 		Payload(func() {
// 			// 整数型の属性 `a` これは左の被演算子
// 			Attribute("a", Int, func() {
// 				Description("Left operand") // 説明
// 				Meta("rpc:tag", "1")        // gRPC 用のメタ情報。タグ定義
// 			})
// 			// 整数型の属性 `b` これは右の被演算子
// 			Attribute("b", Int, func() {
// 				Description("Right operand") // 説明
// 				Meta("rpc:tag", "2")         // gRPC 用のメタ情報。タグ定義
// 			})
// 			Required("a", "b") // a と b は required な属性であることの指定
// 		})

// 		Result(Int) // メソッドの返値（整数を返す）

// 		// HTTP トランスポート用の定義
// 		HTTP(func() {
// 			GET("/add/{a}/{b}") // GET エンドポイント
// 			Response(StatusOK)  // レスポンスのステータスは Status OK = 200 を返す
// 		})

// 		// GRPC トランスポート用の定義
// 		GRPC(func() {
// 			Response(CodeOK) // レスポンスのステータスは CodeOK を返す
// 		})
// 	})
// })

var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("Secures endpoint by requiring an API key.")
})

var _ = Service("diary", func() {
	Error("bad_request")
	Error("unauthorized")

	Error("user_validation_error")
	Error("user_duplication_error")

	Method("UserSignup", func() {
		Payload(func() {
			Attribute("name", String, func() {
				Description("User name")
			})
			Attribute("email", String, func() {
				Description("User email")
			})
			Required("name", "email")
		})

		Result(Empty)

		HTTP(func() {
			POST("/signup")
			Response(StatusCreated)
		})
	})

	Method("Signin", func() {
		Description("Creates a valid API token")

		Payload(func() {
			Attribute("email", String, func() {
				Description("User email")
			})
			Required("email")
		})

		Result(String)

		HTTP(func() {
			POST("/signin")
			Response(StatusOK)
		})
	})

	Method("CreateDiary", func() {
		Description("Creates a diary")

		Security(APIKeyAuth)

		Payload(func() {
			APIKey("api_key", "key", String, "API key used to perform authorization")
			Attribute("user_name", String, func() {
				Description("User name")
			})
			Attribute("title", String, func() {
				Description("Diary title")
			})
			Required("title")
		})

		Result(Empty)

		HTTP(func() {
			// URL is /users/:user_name/diaries
			POST("/users/{user_name}/diaries")
			Response(StatusCreated)
			Header("key:Authorization")
		})
	})
})
