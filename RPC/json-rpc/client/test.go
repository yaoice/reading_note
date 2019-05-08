package main

import (
    . "../server"
    "fmt"
    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
)

type HelloServiceClient struct {
    *rpc.Client
}

//var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
    // rpc连接建立在io.ReadWriteCloser接口之上
    conn, err := net.Dial(network, address)
    if err != nil {
        return nil, err
    }
    // json编解码器包装
    client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

    return &HelloServiceClient{Client: client}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
    return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
    client, err := DialHelloService("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    var reply string
    err = client.Hello("world", &reply)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(reply)
}