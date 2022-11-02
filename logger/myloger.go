package logger

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

var std, _ = New("debug", os.Stderr)

type MyLog struct {
	Loglevel   level
	out        io.Writer
	fileObject *os.File
	Cut        bool
	CutSize    int64
	LogPath    string
	LogName    string
	Pid        int
}

// New 初始化NewLogger对象
func New(strLevel string, w io.Writer) (m *MyLog, err error) {
	var logLevel level
	logLevel, err = parseLoglevel(strLevel)
	if err != nil {
		fmt.Println("unknown log level")
	}
	pid := os.Getpid()
	log := &MyLog{Loglevel: logLevel, out: w, Pid: pid}

	return log, err
}

/*
   日志切割：根据文件大小切割
   1、定义日志切割大小
   2、获取文件大小状态
   3、判断日志文件大小是否符合切割条件
   4、文件切割
   	4-1、关闭旧文件
   	4-2、旧的文件重命名
	4-3、创建新建文件
	4-4、重新打开一个新文件
*/

func (m *MyLog) checkFileSize() (cut bool) {
	// 获取日志文件
	logFile := path.Join(m.LogPath, m.LogName)
	logFileInfo, err := os.Stat(logFile)

	if err != nil {
		fmt.Println("get log file info failed, error:", err)
	}
	if logFileInfo.Size() > m.CutSize {
		return true
	} else {
		return false
	}
}

func (m *MyLog) fileCut() {

	err := m.fileObject.Close()
	if err != nil {
		fmt.Println("log file close failed, error:", err)
		return
	}
	logFile := path.Join(m.LogPath, m.LogName)
	timeStamp := time.Now().Format("20060102150405")
	logSlice := strings.Split(logFile, ".")
	logName := fmt.Sprintf("%s%s.%s", logSlice[0], timeStamp, logSlice[1])
	oldFileName := path.Join(m.LogPath, logName)
	err = os.Rename(logFile, oldFileName)
	if err != nil {
		fmt.Println("log file rename failed, error:", err)
		return
	}
	// 重新打开一个日志文件
	err = m.SetLogFile()
	if err != nil {
		fmt.Println("open new log file failed, error:", err)
	}
}

func (m *MyLog) SetCut(cutSize int) error {
	if cutSize > 1024 {
		return errors.New("cut size is too big, value should less than 1024")
	}
	m.Cut = true
	m.CutSize = int64(cutSize * cutFileSize)
	return nil
}

func (m *MyLog) SetLogFile() error {
	logFile := path.Join(m.LogPath, m.LogName)
	logfile, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file faild, error info %s", err)
		return err
	}
	m.fileObject = logfile
	return nil
}
