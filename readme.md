## 用的第三方包
### github.com/iancoleman/strcase  是用来处理大小写
### github.com/gertd/go-pluralize  用来处理英文单复数
### embed
  - `embed.FS.ReadFile("stubs/" + stubName + ".stub")` 读文件, 
### os
  - 读取文件内容时，结果为[]byte，需要 string()去转成string
  - 写入文件内容时，格式为[]byte，需要[]byte(string_data) 转换
  - `os.WriteFile(to, []byte, 0644)` 将内容写进文件
  - `os.ReadDir(migrator.Folder)`读取指定目录中的所有的文件，返回切片数组和err
  - `os.ReadFile()` 读文件，返回[]byte 和 err
  - `os.Stdout`
  - `os.Args[1:]`
  - `os.SyscallError`
  - `os.IsNotExist(err)`
  - `os.Stat()`
  - `os.Exit(1)`
  - `os.MkdirAll(dir, os.ModePerm)` 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777 如果目录存在，则不做处理
### gorm.io/gorm
  - *gorm.DB.Table("users").Create(&users) 批量创建users，不会调用模型钩子
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
## migration command
### 思路
  - 1.如果第一次迁移，创建 migrations表
  - 2.从database/migrations 目录中获取所有的 有效的迁移文件
  - 3.从migrations表中获取所有的已经迁移的数据，每个迁移表一条数据
  - 4.如果2中有某条或者多条没有在3中出现，则将这些执行迁移操作，migrations的batch字段自增1
1. pkg/migrate/migration_file.go
  - 此文件对象对应单个的迁移文件
  - 单个的迁移文件需要具备三个属性
    - up
    - down
    - 文件名
  - 此文件中的`migrationFiles`是个切片数组，存放的是所有的迁移文件，每次添加新的迁移文件（通过命令添加），都会调用此文件中的`add`方法，append到此切片数组中
  
2. pkg/migrate/migrator.go
  - 此文件对象是最终的操作对象
  - 负责创建migration数据表
  - 拥有 up down reset fresh等动作
