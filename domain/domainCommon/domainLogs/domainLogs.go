package domainLogs

import (
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

//var instance *logrus

var instance *log.Logger

func init() {
	instance = log.New()
	formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
	// Enable line number logging as well
	formatter.Line = false
	formatter.File = false
	instance.SetFormatter(&formatter)
	//log.SetOutput(os.Stdout)
	instance.SetLevel(log.DebugLevel)
}

func GetLogInstance() *log.Logger {
	return instance
}

func PrintError(message interface{}) {
	log.Error(message)
}

func PrintLog(message ...interface{}) {
	log.Println(message)
}
