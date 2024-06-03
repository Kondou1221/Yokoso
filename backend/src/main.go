package main

import (
    "yokoso_api/controller"
    "github.com/labstack/echo/v4"
)

func main() {

     // Echoのインスタンスを作成
     e := echo.New()

     // ルーターを設定
     e.GET("/category", controller.ApiCategory)
 
     // サーバーを開始
     e.Start(":8080")
}