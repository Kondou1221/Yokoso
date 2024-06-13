package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"yokoso_api/db"
	"yokoso_api/repository"

	"github.com/labstack/echo/v4"
)

func ApiWord(c echo.Context) error {
	// パスパラメータから"id"を取得
	id := c.Param("id")
	// 文字列から整数に変換
	idInt, err := strconv.Atoi(id)
	if err != nil {
		// エラーが発生した場合は、Bad Requestを返す
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	dbCon := db.NewDB()
	defer db.CloseDB(dbCon)
	words, err := repository.GetWords(dbCon, idInt)
	if err != nil {
		// エラーハンドリング
		fmt.Println("ワード情報の取得に失敗しました:", err)
		return err
	}

	return c.JSON(http.StatusOK, words)
}
