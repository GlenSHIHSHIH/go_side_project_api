package log

import (
	"componentmod/internal/utils/file"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/golang-module/carbon"
	"github.com/sirupsen/logrus"
)

var (
	L = logrus.New()
)

func init() {
	// logger.SetReportCaller(true)
	dateTime := carbon.Now(carbon.Taipei).Format("Y-m-d")
	mydir, _ := os.Getwd()
	f := file.CreateFile(mydir+"/log", dateTime+".log", 0)
	L.SetOutput(f)
	L.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		HideKeys:        true,
		NoColors:        true,
		ShowFullLevel:   true,
	})

	initLog()
}

type LoggerFunc func(args ...interface{})

var (
	Debug LoggerFunc
	Info  LoggerFunc
	Warn  LoggerFunc
	Error LoggerFunc
	Panic LoggerFunc
	Fatal LoggerFunc
)

func initLog() {
	Debug = L.Debug
	Info = L.Info
	Warn = L.Warn
	Error = L.Error
	Panic = L.Panic
	Fatal = L.Fatal
}
