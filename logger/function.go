package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
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

func SetLogFile(LogPath, LogName string) error {
	std.LogPath = LogPath
	std.LogName = LogName

	logFile := path.Join(LogPath, LogName)
	logfile, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file faild, error info %s", err)
		return err
	}
	std.fileObject = logfile
	return nil
}

func SetLevel(strLevel string) error {
	level, err := parseLoglevel(strLevel)
	if err != nil {
		return fmt.Errorf("failed set level,error: %s", err)
	}

	std.Loglevel = level
	return err

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
