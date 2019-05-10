/*
传统rpc不适合上传/下载数据量大的场景，grpc提供了stream特性

 */
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "io"
    "log"
    . "test/protobuf"
    "time"
)

func main() {
    conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client := NewHelloServiceClient(conn)

    // 先获取stream对象
    stream, err := client.Channel(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    // 模拟客户端发送数据
    go func() {
        for {
            if err := stream.Send(&String{Value: "hi"}); err != nil {
                log.Fatal(err)
            }
            time.Sleep(time.Second)
        }
    }()

    // 循环接收数据
    for {
        reply, err := stream.Recv()
        if err != nil {
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }
        fmt.Println(reply.GetValue())
    }
}