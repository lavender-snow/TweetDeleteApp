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
	PortNo         string
	RunMode        string
	CertFile       string
	KeyFile        string
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
	secretKey := cfg.Section("session").Key("secret_key").String()

	if consumerKey == "" || consumerSecret == "" || secretKey == "" {
		log.Fatalln("iniファイルにconsumerKey,consumerSecret,secretKeyのいずれかが設定されていません")
	}

	tweetCount, err := cfg.Section("twitter").Key("tweet_count").Int()

	if err != nil {
		tweetCount = 5
		log.Println("取得ツイート数の設定が誤っているため、5件として設定します")
	}

	portNo := fmt.Sprintf(":%s", cfg.Section("gin").Key("port_no").String())

	if len(portNo) == 0 {
		portNo = ":8080"
		log.Println("ポート番号の設定がされていないため、8080として設定します")
	}

	runMode := cfg.Section("gin").Key("run_mode").String()

	if runMode != "https" && runMode != "http" {
		runMode = "http"
		log.Println("ginの実行モードが設定されていないため、httpとして設定します")
	}

	certFile := cfg.Section("gin").Key("cert_file").String()
	keyFile := cfg.Section("gin").Key("key_file").String()

	if runMode == "https" && (len(certFile) == 0 || len(keyFile) == 0) {
		log.Fatalln("実行モードがhttpsに設定されていますが、SSL証明書の情報が設定されていません")
	}

	Config = configList{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    cfg.Section("twitter").Key("callback_url").String(),
		SecretKey:      secretKey,
		PortNo:         portNo,
		RunMode:        runMode,
		CertFile:       certFile,
		KeyFile:        keyFile,
		TweetCount:     tweetCount,
	}
}
