package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lavender-snow/TweetDeleteApp/app/controllers"
)

//Routes はルート情報を設定します
func Routes(engine *gin.Engine) {
	engine.GET("/", indexHandler)
	engine.GET("/index", indexHandler)
	engine.GET("/callback", callbackHandler)
	engine.GET("/main", mainHandler)
	engine.GET("/logout", logoutHandler)
	engine.POST("/tweet/delete", tweetDeleteHandler)
}

func indexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func callbackHandler(context *gin.Context) {
	session := sessions.Default(context)
	oauthToken := context.Query("oauth_token")
	oauthVerifier := context.Query("oauth_verifier")
	requestSecret := session.Get("request_secret").(string)

	// ユーザ認証情報取得
	accessToken, accessSecret, err := controllers.GetAccessToken(requestSecret, oauthToken, oauthVerifier)

	if err != nil {
		log.Println(err)
	}
	session.Set("access_token", accessToken)
	session.Set("access_secret", accessSecret)
	session.Save()

	log.Println("コールバック処理実行")

	// メイン画面へリダイレクト
	context.Redirect(http.StatusFound, "/main")
}

func logoutHandler(context *gin.Context) {
	session := sessions.Default(context)

	// セッション情報クリア
	session.Clear()
	session.Save()
	log.Println("ユーザログアウト処理実行")

	// ログイン画面へリダイレクト
	context.Redirect(http.StatusFound, "/")
}

func mainHandler(context *gin.Context) {
	session := sessions.Default(context)
	accessToken := session.Get("access_token")
	accessSecret := session.Get("access_secret")

	// ユーザ認証情報が無い場合Twitterのアプリケーション連携許可へリダイレクトさせる
	if accessToken == nil || accessSecret == nil {
		requestToken, requestSecret := controllers.RequestToken()
		session.Set("request_secret", requestSecret)
		session.Save()
		twitterURL := fmt.Sprintf("https://api.twitter.com/oauth/authenticate?oauth_token=%s", requestToken)

		log.Println("ユーザ認証実行")

		context.Redirect(http.StatusFound, twitterURL)
	} else {
		// ユーザのツイートを取得
		Tweets := controllers.GetTweets(accessToken.(string), accessSecret.(string))
		userName := Tweets[0].ScreenName
		log.Println("ユーザツイート取得処理実行")

		context.HTML(http.StatusOK, "main.tmpl", gin.H{
			"UserName": userName,
			"count":    len(Tweets),
			"Tweets":   Tweets,
		})
	}
}

func tweetDeleteHandler(context *gin.Context) {
	session := sessions.Default(context)
	accessToken := session.Get("access_token")
	accessSecret := session.Get("access_secret")

	// ユーザ認証情報が無い場合Twitterのアプリケーション連携許可へリダイレクトさせる
	if accessToken == nil || accessSecret == nil {
		context.Redirect(http.StatusFound, "/main")
	} else {
		tweetIDs := context.PostFormArray("tweetid")
		log.Printf(tweetIDs[0])
		controllers.DeleteTweets(accessToken.(string), accessSecret.(string), tweetIDs)
		log.Println("ユーザツイート削除処理実行")

		context.Redirect(http.StatusFound, "/main")
	}

}
