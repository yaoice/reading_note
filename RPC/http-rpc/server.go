/*
在http协议上提供jsonrpc服务

调用：
curl localhost:1234/hello -X POST \
    --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
 */
package main

import (
    "io"
    "net/http"
    "net/rpc"
    "net/rpc/jsonrpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "hello:" + request
    return nil
}

func main() {
    rpc.RegisterName("HelloService", new(HelloService))

    // 在处理函数中基于http.ResponseWriter和http.Request类型的参数
    // 构造一个io.ReadWriteCloser类型的conn通道
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        var conn io.ReadWriteCloser = struct {
            io.Writer
            io.ReadCloser
        }{
            ReadCloser: r.Body,
            Writer:     w,
        }

        rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
    })

    http.ListenAndServe(":1234", nil)
}