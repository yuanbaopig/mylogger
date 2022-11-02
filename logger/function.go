package logger

import (
	"errors"
	"os"
	"strings"
)

// parseLoglevel 解析日志级别
func parseLoglevel(ls string) (level, error) {
	s := strings.ToUpper(ls)
	switch s {
	case "DEBUG":
		return DebugLevel, nil
	case "INFO":
		return InfoLevel, nil
	case "ERROR":
		return ErrorLevel, nil
	case "WARNING":
		return WarningLevel, nil
	case "FATAL":
		return FatalLevel, nil
	default:
		err := errors.New("unknown log level")
		return Unknown, err
	}
}

func Error(format string, a ...interface{}) {
	if std.Loglevel <= ErrorLevel {
		std.output("ERROR", format, a...)
	}

}

func Warning(format string, a ...interface{}) {
	if std.Loglevel <= WarningLevel {
		std.output("WARNING", format, a...)
	}
}

func Info(format string, a ...interface{}) {
	if std.Loglevel <= InfoLevel {
		std.output("INFO", format, a...)
	}
}

func Debug(format string, a ...interface{}) {
	if std.Loglevel <= DebugLevel {
		std.output("DEBUG", format, a...)
	}
}

func Fatal(format string, a ...interface{}) {
	if std.Loglevel <= FatalLevel {
		std.output("FATAL", format, a...)
	}
	os.Exit(1)
}
