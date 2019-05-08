package main

import (
    . "../server"
    "fmt"
    "log"
    "net/rpc"
)

type HelloServiceClient struct {
    *rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
    // 与rpc server建立rpc连接
    c, err := rpc.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
    // 调用rpc方法
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