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

---

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

---

#### 1.3 服务器测试

##### 1.3.1 运行服务器输出helloworld

截图：

![2](C:\Users\WYX\Desktop\ServiceComputing\week5\2.png)

----

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

---

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

---

### 2. 拓展任务 net/http 源码阅读与关键功能解释

总的来说，Golang的`HTTP`框架可以由下图简略表示：

![6](C:\Users\WYX\Desktop\ServiceComputing\week5\6.png)

#### 2.1 HTTP的处理流程

理解 HTTP 构建的网络应用只要关注两个端---`客户端（clinet）`和`服务端（server）`，两个端的交互来自 client 的 `request`，以及server端的`response`。所谓的http服务器，主要在于如何接受 client 的 `request`，并向client返回`response`。

接收`request`的过程中，最重要的莫过于路由（`router`），即实现一个`Multiplexer`器。Go中既可以使用内置的mutilplexer --- `DefautServeMux`，也可以自定义。`Multiplexer路由`的目的就是为了找到处理器函数（`handler`），后者将对`request`进行处理，同时构建`response`。

简单总结为如下流程：

```mathematica
Client -> Requests ->  [Multiplexer(router) -> Handler  -> Response -> Client
```

#### 2.2 Handler

理解go中的http服务，最重要就是要理解`Multiplexer`和`Handler`，Golang中的`Multiplexer`基于`ServeMux`结构，同时也实现了`Handler`接口。

`Handler` 可以有以下几种类型：

如图：

![5](C:\Users\WYX\Desktop\ServiceComputing\week5\5.png)

- `handler函数`： 具有`func(w http.ResponseWriter, r *http.Requests)`签名的函数
- `handler处理器(函数)`: 经过`HandlerFunc`结构包装的`handler函数`，它实现了ServeHTTP接口方法的函数。调用handler处理器的ServeHTTP方法时，即调用handler函数本身。
- `handler对象`：实现了Handler接口ServeHTTP方法的结构。

**注**：`handler处理器`和`handler对象`的差别在于，一个是**函数**，另外一个是**结构**，它们都有实现了ServeHTTP方法。

Golang没有继承，类多态的方式可以通过接口实现。所谓接口则是定义声明了函数签名，任何结构只要实现了与接口函数签名相同的方法，就等同于实现了接口。go的HTTP服务都是基于handler进行处理。

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

任何结构体，只要实现了ServeHTTP方法，这个结构就可以称之为handler对象。ServeMux会使用handler并调用其ServeHTTP方法处理请求并返回响应。

---

#### 2.3 ServerMux

ServeMux的源码:

```go
type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry
    hosts bool 
}

type muxEntry struct {
    explicit bool
    h        Handler
    pattern  string
}
```

ServeMux结构中最重要的字段为`m`，这是一个map，key是一些url模式，value是一个muxEntry结构，后者里定义存储了具体的url模式和handler。 

---

#### 2.4 Server

从`http.ListenAndServe`的源码可以看出，**Server**创建了一个server对象，并调用**server**对象的ListenAndServe方法：

```go
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
```

**Server** 的结构如下:

```go
type Server struct {
    Addr         string        
    Handler      Handler       
    ReadTimeout  time.Duration 
    WriteTimeout time.Duration 
    TLSConfig    *tls.Config   

    MaxHeaderBytes int

    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
    disableKeepAlives int32     nextProtoOnce     sync.Once 
    nextProtoErr      error     
}
```

server结构存储了服务器处理请求常见的字段。其中Handler字段也保留Handler接口。如果Server接口没有提供`Handler`结构对象，那么会使用`DefautServeMux`做`multiplexer`.

#### 2.5 创建HTTP服务

创建一个http服务，大致需要经历两个过程，

- 注册路由，即提供url模式和handler函数的映射
- 实例化一个server对象，并开启对客户端的监听

##### 2.5.1 注册路由

net/http包暴露的注册路由的api很简单，`http.HandleFunc`选取了`DefaultServeMux`作为`multiplexer`

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
```

`DefaultServeMux`是`ServeMux`的一个实例。`http`包也提供了`NewServeMux`方法创建一个`ServeMux`实例，默认则创建一个`DefaultServeMux`：

```go
// NewServeMux allocates and returns a new ServeMux.
func NewServeMux() *ServeMux { return new(ServeMux) }

// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux
```

`DefaultServeMux`的`HandleFunc(pattern, handler)`方法实际是定义在`ServeMux`下的。

----

```go
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	mux.Handle(pattern, HandlerFunc(handler))
}
```

上述代码中，`HandlerFunc`是一个函数类型。同时实现了`Handler`接口的`ServeHTTP`方法。使用`HandlerFunc`类型包装一下路由定义的`indexHandler`函数，其目的就是为了让这个函数也实现`ServeHTTP`方法，即转变成一个handler处理器(函数)。一旦这样做了，就意味着我们的 `indexHandler `函数也有了`ServeHTTP`方法。此外，`ServeMux`的`Handle`方法，将会对`pattern`和`handler`函数做一个`map`映射。

----

```go
func (mux *ServeMux) Handle(pattern string, handler Handler) {
    mux.mu.Lock()
    defer mux.mu.Unlock()
    if pattern == "" {
        panic("http: invalid pattern " + pattern)
    }
    if handler == nil {
        panic("http: nil handler")
    }
    if mux.m[pattern].explicit {
        panic("http: multiple registrations for " + pattern)
    }

    if mux.m == nil {
        mux.m = make(map[string]muxEntry)
    }
    mux.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern}

    if pattern[0] != '/' {
        mux.hosts = true
    }

    n := len(pattern)
    if n > 0 && pattern[n-1] == '/' && !mux.m[pattern[0:n-1]].explicit {

        path := pattern
        if pattern[0] != '/' {
            path = pattern[strings.Index(pattern, "/"):]
        }
        url := &url.URL{Path: path}
        mux.m[pattern[0:n-1]] = muxEntry{
            h: RedirectHandler(url.String(),StatusMovedPermanently), pattern: pattern
        }
	}
}    
```

由此可见，`Handle函数`的主要目的在于把`handler`和`pattern`模式绑定到`map[string]muxEntry`的`map`上，其中`muxEntry`保存了更多`pattern`和`handler`的信息，前面讨论的`Server结构的m字段`就是`map[string]muxEntry`这样一个map。  此时，pattern和handler的路由注册完成。

##### 2.5.2 开始监听

开始`server`的监听，以接收客户端的请求。

```go
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

func (srv Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(tcpKeepAliveListener{ln.(net.TCPListener)})
}
```

Server的`ListenAndServe方法`中，会`初始化监听地址Addr`，同时`调用Listen方法设置监听`。最后将监听的TCP对象传入Serve方法。

---

##### 2.5.3 处理请求

监听开启之后，一旦客户端请求到底，go就开启一个协程处理请求，主要逻辑都在`serve方法`之中。 ` serve方法`比较长，其主要职能就是，创建一个上下文对象，然后`调用Listener的Accept方法`用来获取连接数据并使用`newConn方法创建连接对象`。最后使用`goroutein协程`的方式处理连接请求。因为每一个连接都开起了一个协程，请求的上下文都不同，同时又保证了go的高并发。

`goserver`方法如下：

```go
func (c *conn) serve(ctx context.Context) {
	c.remoteAddr = c.rwc.RemoteAddr().String()
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
            buf := make([]byte, size)
            buf = buf[:runtime.Stack(buf, false)]
            c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
		if !c.hijacked() {
			c.close()
			c.setState(c.rwc, StateClosed)
		}
	}()

...

for {
    w, err := c.readRequest(ctx)
    if c.r.remain != c.server.initialReadLimitSize() {
        // If we read any bytes off the wire, we're active.
        c.setState(c.rwc, StateActive)
    }
    ...
    
    }
    
    ...
  
    serverHandler{c.server}.ServeHTTP(w, w.req)
    w.cancelCtx()
    if c.hijacked() {
        return
    }
    w.finishRequest()
    if !w.shouldReuseConnection() {
        if w.requestBodyLimitHit || w.closedRequestBodyEarly() {
            c.closeWriteAndWait()
        }
        return
    }
    c.setState(c.rwc, StateIdle)
}
```

defer定义了函数退出时，连接关闭相关的处理。然后就是读取连接的网络数据，并处理读取完毕时候的状态。接下来就是调用`serverHandler{c.server}.ServeHTTP(w, w.req)`方法处理请求了。最后就是请求处理完毕的逻辑。serverHandler是一个重要的结构，它具有一个字段，即Server结构，同时它也实现了Handler接口方法ServeHTTP，并在该接口方法中做了一个重要的事情，初始化multiplexer路由多路复用器。如果server对象没有指定Handler，则使用`默认的DefaultServeMux`作为路由Multiplexer，并调用初始化Handler的ServeHTTP方法。

---

```go
func (mux *ServeMux) (w ResponseWriter, r Request) {
    if r.RequestURI == "" {
    	if r.ProtoAtLeast(1, 1) {
    		w.Header().Set("Connection", "close")
    	}
    	w.WriteHeader(StatusBadRequest)
    	return
    }
    h, _ := mux.Handler(r)
    h.ServeHTTP(w, r)
}

func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
	if r.Method != "CONNECT" {
		if p := cleanPath(r.URL.Path); p != r.URL.Path {
            _, pattern = mux.handler(r.Host, p)
            url := *r.URL
            url.Path = p
            return RedirectHandler(url.String(), StatusMovedPermanently), pattern
		}
	}
	return mux.handler(r.Host, r.URL.Path)
}

func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
    mux.mu.RLock()
    defer mux.mu.RUnlock()

    // Host-specific pattern takes precedence over generic ones
    if mux.hosts {
        h, pattern = mux.match(host + path)
    }
    if h == nil {
        h, pattern = mux.match(path)
    }
    if h == nil {
        h, pattern = NotFoundHandler(), ""
    }
    return
    }

    func (mux *ServeMux) match(path string) (h Handler, pattern string) {
    	var n = 0
    	for k, v := range mux.m {
    		if !pathMatch(k, path) {
    			continue
    		}
    	if h == nil || len(k) > n {
    		n = len(k)
    		h = v.h
    		pattern = v.pattern
    	}
    }
    return
}
```

mux的ServeHTTP方法通过调用其Handler方法寻找注册到路由上的handler函数，并调用该函数的ServeHTTP方法，本例则是IndexHandler函数。  mux的Handler方法对URL简单的处理，然后调用handler方法，后者会创建一个锁，同时调用match方法返回一个handler和pattern。  在match方法中，mux的m字段是map[string]muxEntry图，后者存储了pattern和handler处理器函数，因此通过迭代m寻找出注册路由的patten模式与实际url匹配的handler函数并返回。  返回的结构一直传递到mux的ServeHTTP方法，接下来调用handler函数的ServeHTTP方法，即IndexHandler函数，然后把response写到http.RequestWirter对象返回给客户端。  上述函数运行结束即`serverHandler{c.server}.ServeHTTP(w, w.req)`运行结束。接下来就是对请求处理完毕之后上希望和连接断开的相关逻辑。  

----

### 3. 源代码与Readme

[Github](https://github.com/wyxwyx46941930/ServiceComputing/tree/master/week5)

----

### 4. 参考文献

- [在windows上安装测试ab](https://www.cnblogs.com/wxinyu/p/8929992.html)
- [Martini框架的使用](http://www.cnblogs.com/tanghui/p/4846156.html)
- [http/net的源码解读](https://blog.csdn.net/HOMERUNIT/article/details/78518430)