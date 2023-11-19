package client_proxy

import (
	"OldPackageTest/new_helloworld/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

// 在go中没有类、对象， 就意味着没有初始话方法
func NewHelloServiceClient(protcol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protcol, address)
	if err != nil {
		panic("connect err !")
	}
	return HelloServiceStub{conn}
}
func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
