# Loger 自定义日志库

**日志库需求**

1. 区分日志级别
   - DEBUG
   - INFO
   - WARNING
   - ERROR
   - FATAL
2. 级别控制开关
3. 指定输出到文件或标准输出
   - 如果不指定文件，默认标准输出
4. 日志格式化
   - 标准化格式：时间、函数、行号、日志级别、进程ID
   - 自定义格式：
5. 日志切割开关
   - 文件大小切割
   - 日期时间切割
6. 当日志级别小于INFO时，将在标准输出中打印日志
7. ~~日志写入使用异步写~~



日志样式

```shell
2022-10-25 15:25:43,ERROR,12059,main.go,main,27,this is error log
2022-10-25 15:25:43,DEBUG,12059,main.go,main,29,this is debug log
2022-10-25 15:25:43,INFO,12059,main.go,main,31,this is info log
2022-10-25 15:25:43,WARNING,12059,main.go,main,33,this is warning log
2022-10-25 15:25:43,FATAL,12059,main.go,main,35,this is fatal log
```

