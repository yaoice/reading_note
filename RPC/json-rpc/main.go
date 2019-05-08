/*
采用官方net/rpc/jsonrpc扩展，只要是同样的json结构，就可以进行跨语言rpc

参考链接：https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-01-rpc-intro.html
*/
package main

import (
    . "./server"
    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
)

func main() {
    RegisterHelloService(new(HelloService))

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("Accept error:", err)
        }

        // json编解码器包装
        go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
    }
}