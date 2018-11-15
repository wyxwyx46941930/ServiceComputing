package main

import (
    "os"
    "github.com/web/service"
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
