package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

type Fields = logrus.Fields

func init() {
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
	//Log.SetLevel(logrus.DebugLevel)
	//Log.SetLevel(logrus.ErrorLevel)
	Log.SetFormatter(&logrus.JSONFormatter{})
}
