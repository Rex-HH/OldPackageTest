package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer struct{}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}
func main() {
	//1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	//2.注册处理逻辑handler
	_ = rpc.RegisterName("HelloServer", &HelloServer{})
	//3.启动服务
	for {
		conn, _ := listener.Accept() //当一个新的连接进来的时候，
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
