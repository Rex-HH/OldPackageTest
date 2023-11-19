package main

import (
	"net"
	"net/rpc"
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
	conn, _ := listener.Accept() //当一个新的连接进来的时候，
	rpc.ServeConn(conn)

	//一连串的代码大部分都是net的包 好像和rpc没有关系
	//不行， rpc调用中有几个问题需要解决 1. call 2. 序列化和反序列化 编码和解码
	//python 下的开发而言， 这个显得不好用
	//可以跨语言调用呢 1. go 语言的rpc的序列化·和反序列化协议是什么（gob）
	//2. 能否替换成常见的序列化
}
