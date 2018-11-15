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

##### 1.3.1 curl 测试





