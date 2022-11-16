package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

type level uint8

// 日志级别
const (
	Unknown level = iota
	DebugLevel
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel

	cutFileSize = 1024 * 1024
)

var rwMutex sync.RWMutex

func (m *MyLog) Output(format interface{}, fileObject *os.File) {
	// 检查日志级别，如果是日志级别小于Info级别，则在标准输出

	fmt.Print(format)
	go func(format interface{}) {
		rwMutex.Lock()
		_, err := fmt.Fprint(fileObject, format)
		if err != nil {
			fmt.Println("write log failed, error:", err)
		}
		rwMutex.Unlock()
	}(format)

}

// output 日志输出
func (m *MyLog) output(level string, format string, a ...interface{}) {
	// 检查日志是否需要切割
	if m.Cut == true {

		// 检查日志文件大小
		if m.checkFileSize() == true {
			// 日志切割
			rwMutex.Lock()
			m.fileCut()
			rwMutex.Unlock()
		}
	}
	// 获取go runtime
	funcName, file, line, ok := m.getInfo(3)

	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	// 日志信息
	msg := fmt.Sprintf(format, a...)

	fmt.Printf("%s,%s,%d,%s,%s,%d,%s\n", time.Now().Format("2006-01-02 15:04:05"), level, m.Pid, file, funcName, line, msg)

	// 写入日志文件
	if m.fileObject != nil {
		go func(msg *string) {
			rwMutex.Lock()
			_, err := fmt.Fprintf(m.fileObject, "%s,%s,%d,%s,%s,%d,%s\n", time.Now().Format("2006-01-02 15:04:05"), level, m.Pid, file, funcName, line, *msg)
			if err != nil {
				fmt.Println("write log failed, error:", err)
			}
			rwMutex.Unlock()
		}(&msg)
	}
}

// getInfo n 代表获取runtime函数调用层级的信息，最底层0代表自身函数的数据
func (m *MyLog) getInfo(n int) (funcName, fileName string, line int, ok bool) {
	/*
		runtime 包为记录运行时状态的功能，例如堆栈信息，执行的函数名，调用的文件具柄等
		pc 代表运行时对象（执行函数）
			- 但是想要获取函数名称，还需要再调用FuncForPC()特殊方法来获得
		file 代表正在执行的文件（调用Caller该函数的文件）
		line 代表正在执行的行号
		ok 如果可以取到对象信息则为True
	*/
	pc, file, line, ok := runtime.Caller(n)
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	//s := strings.Split(bagFuncName, ".")
	//funcName = s[1]
	return

}

func (m *MyLog) Error(format string, a ...interface{}) {
	if m.Loglevel <= ErrorLevel {
		m.output("ERROR", format, a...)
	}
}

func (m *MyLog) Warning(format string, a ...interface{}) {
	if m.Loglevel <= WarningLevel {
		m.output("WARNING", format, a...)
	}
}

func (m *MyLog) Info(format string, a ...interface{}) {
	if m.Loglevel <= InfoLevel {
		m.output("INFO", format, a...)
	}
}

func (m *MyLog) Debug(format string, a ...interface{}) {
	if m.Loglevel <= DebugLevel {
		m.output("DEBUG", format, a...)
	}
}

func (m *MyLog) Fatal(format string, a ...interface{}) {
	if m.Loglevel <= FatalLevel {
		m.output("FATAL", format, a...)
	}
	os.Exit(1)
}
