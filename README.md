# TweetDeleteApp
TwitterAPIを使用して認証ユーザのツイートを一覧表示し、選択して一括削除を行うWebアプリ。

使用言語
Go

使用フレームワーク
gin https://github.com/gin-gonic/gin
bootstrap 5.0


# 事前準備

1. TwitterAPIを実行するためにTwitterアプリを作成し、ConsumerKeyおよびConsumerSecretKeyを取得します。

1. TwitterアプリのAppDetails画面にてCallbackURLを設定します。ローカル環境にて実行したい場合は「http://127.0.0.1:8080/callback」としてください。

1. TwitterアプリのPermissions画面にてAccess permissionを「Read and write」に設定します。

1. TemplateConfig.iniを基にconfig.iniを作成します。

```
[twitter]
consumer_key = <取得したConsumerKey>
consumer_secret = <取得したConsumerSecretKey>
callback_url = <設定したCallbackURL>
tweet_count = <ログイン後一度に取得するツイートの最大件数、200件が上限>

[common]
secret_key = <session管理で使用するシークレットキー>
```

# 実行方法
cloneしたディレクトリへ移動し

```
go run main.go
```