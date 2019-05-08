/*
Go官方rpc包是net/rpc
 */
package server

import (
    "net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
    Hello(request string, reply *string) error
}

// rpc.Register将对象类型下所有符合Go rpc规则的方法都注册为rpc方法，
// rpc方法托管在HelloServiceName命名空间下
func RegisterHelloService(svc HelloServiceInterface) error {
    return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct {}

/*
rpc方法必须满足Go rpc规则:
1. 公开方法
2. 只有两个序列化参数，第二个参数是指针类型
3. 返回值是error类型
*/
func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "hello:" + request
    return nil
}
