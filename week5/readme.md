[TOC]

## 服务计算——cloudgo

### 1. 基本任务——搭建简单web服务器

#### 1.1 框架选择

本次使用的web开发框架是Martini，Martini 是一个非常新的 Go 语言的 Web 框架，使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，也可使用自己的 DB 层、会话管理和模板。这个框架在GitHub上都有中文的解释以及用法，比较容易上手。

其特性如下：

- 使用非常简单
- 无侵入设计
- 可与其他 Go 的包配合工作
- 超棒的路径匹配和路由
- 模块化设计，可轻松添加工具
- 大量很好的处理器和中间件
- 很棒的开箱即用特性
- 完全兼容 http.HandlerFunc 接口

简单例子：

```go
package main
 
import "github.com/codegangsta/martini"
 
func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
```

请求处理器：

```go
m.Get("/", func() {
  println("hello world")
})
 
m.Get("/", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
  res.WriteHeader(200) // HTTP 200
})
```

#### 1.2 代码

##### 1.2.1 main.go

`main.go`文件使用了老师博客中给出的代码，完成`绑定端口`、`解析端口`、`启动server`完成操作的任务

```go
package main
import (
    "os"
    "web/service"
    flag "github.com/spf13/pflag"
)
const (
    //默认8080端口
    PORT string = "8080" 
)
func main() {
    //默认8080端口
    port := os.Getenv("PORT") 
    if len(port) == 0 {
        port = PORT
    }
    //端口号的解析
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
    //启动server
    service.NewServer(port)
}
```

##### 1.2.2 server.go

`server.go`文件则是使用martini框架中的函数格式具体定义`main.go`文件中启动server后要具体进行的操作

```go
package service
import (
    "github.com/go-martini/martini" 
)
func NewServer(port string) {   
    m := martini.Classic()

    m.Get("/", func(params martini.Params) string {
        return "hello world"
    })

    m.RunOnAddr(":"+port)   
}
```

#### 1.3 服务器测试

##### 1.3.1 运行服务器输出helloworld

截图：

![2](C:\Users\WYX\Desktop\ServiceComputing\week5\2.png)

##### 1.3.2 curl 测试

截图：

![1](C:\Users\WYX\Desktop\ServiceComputing\week5\1.png)

提示信息：

```mathematica
$ curl -v http://localhost:9090/  
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*  Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET / HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.61.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Thu, 15 Nov 2018 12:07:26 GMT
< Content-Length: 11
< Content-Type: text/plain; charset=utf-8
<
{ [11 bytes data]
100    11  100    11    0     0    354      0 --:--:-- --:--:-- --:--:--   354hello world
* Connection #0 to host localhost left intact

```

##### 1.3.3 ab测试

截图：

![4](C:\Users\WYX\Desktop\ServiceComputing\week5\4.png)

提示信息：

```mathematica
$ ./ab -n 1000 -c 100 http://localhost:9090/
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9090

Document Path:          /
Document Length:        11 bytes

Concurrency Level:      100
Time taken for tests:   0.271 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      128000 bytes
HTML transferred:       11000 bytes
Requests per second:    3692.76 [#/sec] (mean)
Time per request:       27.080 [ms] (mean)
Time per request:       0.271 [ms] (mean, across all concurrent requests)
Transfer rate:          461.60 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       1
Processing:     4   25   3.9     26      31
Waiting:        2   15   6.7     14      27
Total:          4   25   3.9     26      31

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     26
  75%     26
  80%     27
  90%     27
  95%     27
  98%     28
  99%     28
 100%     31 (longest request)
```

