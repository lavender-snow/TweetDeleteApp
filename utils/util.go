package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// LoggingSettings ログファイルを保存するための設定を実行
func LoggingSettings(logFile string) {
	l, _ := time.LoadLocation("Asia/Tokyo")
	logFileName := fmt.Sprintf("%s_%s.log", logFile, time.Now().In(l).Format("20060102150405"))

	logfile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 066)
	if err != nil {
		log.Fatalf("file=logfile err=%s", err.Error())
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
