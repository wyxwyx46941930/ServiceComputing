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