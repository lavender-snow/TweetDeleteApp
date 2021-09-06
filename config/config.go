package config

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

type configList struct {
	ConsumerKey    string
	ConsumerSecret string
	CallbackURL    string
	SecretKey      string
	TweetCount     int
}

const (
	configIni = "config.ini"
)

// Config は設定情報を管理します
var Config configList

func init() {
	cfg, err := ini.Load(configIni)
	if err != nil {
		log.Fatalf("iniファイルの読込に失敗しました: %v\n", err.Error())
	}

	consumerKey := cfg.Section("twitter").Key("consumer_key").String()
	consumerSecret := cfg.Section("twitter").Key("consumer_secret").String()
	secretKey := cfg.Section("common").Key("secret_key").String()

	if consumerKey == "" || consumerSecret == "" || secretKey == "" {
		log.Fatalln("iniファイルにconsumerKey,consumerSecret,secretKeyのいずれかが設定されていません")
	}

	tweetCount, err := cfg.Section("twitter").Key("tweet_count").Int()

	if err != nil {
		fmt.Println("取得ツイート数の設定が誤っているため、5件として設定します")
		tweetCount = 5
	}

	Config = configList{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    cfg.Section("twitter").Key("callback_url").String(),
		SecretKey:      secretKey,
		TweetCount:     tweetCount,
	}
}
