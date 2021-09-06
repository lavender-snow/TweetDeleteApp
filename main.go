package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/lavender-snow/TweetDeleteApp/config"
	"github.com/lavender-snow/TweetDeleteApp/routes"
	"github.com/lavender-snow/TweetDeleteApp/utils"
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
	if config.Config.RunMode == "https" {
		engine.RunTLS(config.Config.PortNo, config.Config.CertFile, config.Config.KeyFile)
	} else {
		engine.Run(config.Config.PortNo)
	}

}
