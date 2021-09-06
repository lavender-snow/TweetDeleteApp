package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/lavender-snow/site/config"
)

// Tweet ツイートの情報を格納する構造体
type Tweet struct {
	ID            int64
	Name          string
	ScreenName    string
	Text          string
	ReplyCount    int
	FavoriteCount int
	RetweetCount  int
	CreatedAt     string
}

// Tweets 複数のツイートを格納する配列
type Tweets []Tweet

// OAuth1Config OAuth認証に必要な情報を管理
var OAuth1Config oauth1.Config

func init() {

	consumerKey := config.Config.ConsumerKey
	consumerSecret := config.Config.ConsumerSecret

	if consumerKey == "" || consumerSecret == "" {
		log.Fatal("Consumer key/secret required")
	}

	OAuth1Config = oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    config.Config.CallbackURL,
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}

}

// GetAccessToken リクエストトークンを基にアクセストークンを返却する
func GetAccessToken(requestSecret string, oauthToken string, oauthVerifier string) (string, string, error) {

	accessToken, accessSecret, err := OAuth1Config.AccessToken(oauthToken, requestSecret, oauthVerifier)
	return accessToken, accessSecret, err
}

// RequestToken リクエストトークンを作成する
func RequestToken() (string, string) {

	requestToken, requestSecret, err := OAuth1Config.RequestToken()
	if err != nil {
		log.Println(err)
	}

	return requestToken, requestSecret
}

// GetTweets 認証ユーザの直近ツイートを一括取得する
func GetTweets(accessToken string, accessSecret string) []Tweet {

	var userTweets Tweets

	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := OAuth1Config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// UserTimeLine Apiを使用しツイートを取得
	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count: config.Config.TweetCount,
	})

	if err != nil {
		log.Println(err.Error())
	}

	for _, tweet := range tweets {
		l, _ := time.LoadLocation("Asia/Tokyo")
		t, _ := tweet.CreatedAtTime()

		userTweets = append(userTweets, Tweet{
			ID:            tweet.ID,
			Name:          tweet.User.Name,
			ScreenName:    tweet.User.ScreenName,
			Text:          tweet.Text,
			ReplyCount:    tweet.ReplyCount,
			FavoriteCount: tweet.FavoriteCount,
			RetweetCount:  tweet.RetweetCount,
			CreatedAt:     t.In(l).Format("2006/01/02 15:04:05"),
		})
	}

	return userTweets
}

// DeleteTweets ツイートIDと一致するツイートを削除する
func DeleteTweets(accessToken string, accessSecret string, tweetIDstring []string) {

	var tweetID []int64
	for _, v := range tweetIDstring {
		log.Println(v)
		ID, _ := strconv.ParseInt(v, 10, 64)
		tweetID = append(tweetID, ID)
	}

	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := OAuth1Config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// Destroy Apiを使用しツイートを削除
	for _, ID := range tweetID {

		tweet, resp, err := client.Statuses.Destroy(ID, &twitter.StatusDestroyParams{
			//ID:       ID,
			TrimUser: new(bool)})
		log.Println(tweet, resp, err)
	}

}
