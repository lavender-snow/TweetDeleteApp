package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/lavender-snow/site/config"
	"github.com/lavender-snow/site/routes"
	"github.com/lavender-snow/site/utils"
)

func main() {
	// ログ設定
	utils.LoggingSettings("TweetDelete")

	// gin初期化
	engine := gin.Default()
	store := sessions.NewCookieStore([]byte(config.Config.SecretKey))
	store.Options(sessions.Options{MaxAge: 3600})
	engine.Use(sessions.Sessions("tweetdelete", store))
	engine.Static("/assets", "./assets")
	engine.LoadHTMLGlob("app/views/templates/*.tmpl")
	routes.Routes(engine)

	// 開始
	engine.Run(":8080")
}
