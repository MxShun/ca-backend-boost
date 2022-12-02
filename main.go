package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("Traceのログ情報です")
	logrus.Debug("Debugのログ情報です")
	logrus.Info("Infoのログ情報です")

	logrus.SetLevel(logrus.InfoLevel)

	logrus.Trace("Traceのログ情報です")
	logrus.Debug("Debugのログ情報です")
	logrus.Info("Infoのログ情報です")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("JSONで構造化されたログ出力です")

	logrus.WithFields(logrus.Fields{
		"code": "400",
		"message": "エラーが発生しました",
	}).Info("フィールド付きのログ出力")
}
