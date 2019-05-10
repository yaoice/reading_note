package main

import (
    "context"
    "google.golang.org/grpc"
    "io"
    "log"
    "net"
    . "test/protobuf"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
    reply := &String{Value: "hello:" + args.GetValue()}
    return reply, nil
}

// 循环接收客户端的数据，数据重新组装后，通过stream又发给客户端；双向流数据的发送和接收是独立的
func (p *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
    for {
        args, err := stream.Recv()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }

        reply := &String{Value: "hello:" + args.GetValue()}

        err = stream.Send(reply)
        if err != nil {
            return err
        }
    }
}

func main() {
    grpcServer := grpc.NewServer()
    RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

    lis, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal(err)
    }
    grpcServer.Serve(lis)
}
