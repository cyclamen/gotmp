package log

import (
	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var (
	Log *logging.Logger
)

var fileLogFormat = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x} %{message}`,
)

var stderrLogFormat = logging.MustStringFormatter(
	`%{color}%{time:2006/01/02 15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func InitLog(module string, maxSize int, level string, setStderr bool) {

	ll := &lumberjack.Logger{Filename: module + ".log", MaxSize: maxSize, MaxBackups: 3, MaxAge: 15}

	fileBackend := logging.NewLogBackend(ll, "", 0)
	stderrBackend := logging.NewLogBackend(os.Stderr, "", 0)

	fileBackendFormatter := logging.NewBackendFormatter(fileBackend, fileLogFormat)
	stderrBackendFormatter := logging.NewBackendFormatter(stderrBackend, stderrLogFormat)

	fileLeveledBackend := logging.AddModuleLevel(fileBackendFormatter)
	stderrLeveledBackend := logging.AddModuleLevel(stderrBackendFormatter)

	var logLevel logging.Level

	switch strings.ToUpper(strings.TrimSpace(level)) {
	case "DEBUG":
		logLevel = logging.DEBUG
	case "INFO":
		logLevel = logging.INFO
	case "WARNING":
		logLevel = logging.WARNING
	case "ERROR":
		logLevel = logging.ERROR
	case "NOTICE":
		logLevel = logging.NOTICE
	case "CRITICAL":
		logLevel = logging.CRITICAL
	default:
		logLevel = logging.DEBUG
	}

	fileLeveledBackend.SetLevel(logLevel, "")
	stderrLeveledBackend.SetLevel(logLevel, "")

	Log = logging.MustGetLogger(module)

	if setStderr {
		Log.SetBackend(logging.SetBackend(fileLeveledBackend, stderrLeveledBackend))
	} else {
		Log.SetBackend(logging.SetBackend(fileLeveledBackend))
	}

}
