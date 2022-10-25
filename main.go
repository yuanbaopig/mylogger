package main

import (
	"logger/logger"
)

func main() {

	//logger, err := logger.New("DEBUG", "./", "proc.log", false)
	//if err != nil {
	//	fmt.Printf("loger initialization failed")
	//}

	// 指定日志文件输出
	//err := logger.SetLogFile("./", "xxx.log")
	//if err != nil {
	//	panic(err)
	//}
	// 开启日志切割
	//err = logger.SetCut(1)
	//if err != nil {
	//	panic(err)
	//}
	var info string
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
