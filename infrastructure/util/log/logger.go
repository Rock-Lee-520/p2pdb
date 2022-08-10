package log

import (
	"os"

	debug "github.com/favframework/debug"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	// log.SetFormatter(&log.TextFormatter{
	// 	DisableColors: true,
	// 	FullTimestamp: true,
	// })

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.ErrorLevel)
}

func Trace(args ...interface{}) {
	log.Trace(args)
}

func Debug(args ...interface{}) {
	debug.Dump(args)
	//log.Debug(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Panic(args ...interface{}) {
	log.Panic(args)
}
