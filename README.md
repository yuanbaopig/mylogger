# Loger 自定义日志库

**日志库需求**

1. 区分日志级别
   - DEBUG
   - INFO
   - WARNING
   - ERROR
   - FATAL
2. 级别控制开关
   - 提供指定日志级别（默认debug）
     - 默认情况下会将所有日志信息
     - 如果指定了日志文件，则可以设置日志级别，控制日志信息输出到屏幕（`info`以上级别将不再输出`Print`）
3. 指定输出到文件或标准输出
   - 如果不指定文件，默认标准输出
4. 日志格式化
   - 标准化格式：时间、函数、行号、日志级别、进程ID
   - 自定义格式（Output）
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



## 注意事项

在`import` 导入包的同时，会定义一个全局的`std`对象，所以在项目中使用时建议定义一个对象的`logger`，避免在多个包中导入出现定义的属性冲突。如果只是单纯的当作日志输出则没有限制。

```go
var std, _ = New("debug", os.Stderr)
```

