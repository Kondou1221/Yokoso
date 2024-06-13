package router

import(
	"yokoso_api/controller"

)

func router(){
	// Echoのインスタンスを作成
	e := echo.New()

	// ルーティング
	e.GET("/", hello)
	e.GET("/category", controller.apiCategory)
}