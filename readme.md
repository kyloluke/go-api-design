## 用的第三方包
- github.com/iancoleman/strcase  是用来处理大小写
- github.com/gertd/go-pluralize  用来处理英文单复数
- embed
  - `var stubsFS embed.FS`
  - `stubsFS.ReadFile("stubs/" + stubName + ".stub")`
- os
  - `os.WriteFile(to, data, 0644)`
  - `os.Stdout`
  - `os.Args[1:]`
  - `os.SyscallError`
  - `os.IsNotExist(err)`
  - `os.Stat()`
  - `os.Exit(1)`

## string()
  - string和数字之间转换可使用标准库strconv
  - 想要转换byte数组（[]byte或 []rune）为string字符串类型，这种情况下可以用string()

## 关于时间
- 获取当前时间
```go
  currentTime := time.Now()      //获取当前时间，类型是Go的时间类型Time 类型为：time.Time, 值为:time.Date(2022, time.August, 29, 0, 32, 45, 248462100, time.Local)
  t1 := time.Now().Year()        //年
  t2 := time.Now().Month()       //月
  t3 := time.Now().Day()         //日
  t4 := time.Now().Hour()        //小时
  t5 := time.Now().Minute()      //分钟
  t6 := time.Now().Second()      //秒
  t7 := time.Now().Nanosecond()  //纳秒
  
  currentTimeData:=time.Date(t1,t2,t3,t4,t5,t6,t7,time.Local) //获取当前时间，返回当前时间Time （time.Local 指定时区）
  fmt.Println(currentTime)            //打印结果：2017-04-11 12:52:52.794351777 +0800 CST
  fmt.Println(t1,t2,t3,t4,t5,t6)      //打印结果：2017 April 11 12 52 52
  fmt.Println(currentTimeData)        //打印结果：2017-04-11 12:52:52.794411287 +0800 CST
  
  timeUnix:=time.Now().Unix()          // 获取时间戳  单位s,打印结果:1491888244
  timeUnixNano:=time.Now().UnixNano()  //获取时间戳 单位纳秒,打印结果：1491888244752784461
  
  timeStr := time.Now().Format("2006-01-02 15:04:05")  //当前时间字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
  chinaTimezone, _ := time.LoadLocation("America/New_York") // 获取指定时区的当前时间
  fmt.Println(time.Now().In(chinaTimezone))// 获取指定时区的当前时间
  
  // 时间戳转时间字符串 (int64 —> string) 
  timeUnix:=time.Now().Unix()             //已知的时间戳 int64
  formatTimeStr:=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")
  fmt.Println(formatTimeStr)              //打印结果：2017-04-11 13:30:39
  
  // 字符串转时间Time (string —> Time)
  formatTimeStr := "2017-04-11 13:33:37"
  formatTime, _ := time.Parse("2006-01-02 15:04:05", formatTimeStr) // 转到 time.Now()
```


