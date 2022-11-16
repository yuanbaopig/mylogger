package main

import (
	"fmt"
	"github.com/yuanbaopig/mylogger/logger"
	"os"
)

var info string

func main() {

	//logPrint()
	logOut()
}

func logOut() {
	// 自定义logger对象
	myLogger, err := logger.New("DEBUG", os.Stderr)
	if err != nil {
		fmt.Printf("logger initialization failed")
	}

	// 指定日志文件输出
	err = myLogger.SetLogFile("./", "test.log")
	if err != nil {
		panic(err)
	}

	// 开启日志切割
	//err = logger.SetCut(1)
	//if err != nil {
	//	panic(err)
	//}

	for {

		info = "error"
		myLogger.Error("this is %s log", info)
		info = "debug"
		myLogger.Debug("this is %s log", info)
		info = "info"
		myLogger.Info("this is %s log", info)
		info = "warning"
		myLogger.Warning("this is warning log")
		//info = "fatal"
		//myLogger.Fatal("this is fatal log")

	}

}

func logPrint() {
	for {

		info = "error"
		logger.Error("this is %s log", info)
		info = "debug"
		logger.Debug("this is %s log", info)
		info = "info"
		logger.Info("this is %s log", info)
		info = "warning"
		logger.Warning("this is warning log")
		info = "fatal"
		logger.Fatal("this is fatal log")

	}

}
