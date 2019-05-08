package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "net"
    . "test/protobuf"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
    reply := &String{Value: "hello:" + args.GetValue()}
    return reply, nil
}

func main() {
    grpcServer := grpc.NewServer()
    // protobuf生成注册函数
    RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

    lis, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal(err)
    }
    grpcServer.Serve(lis)
}