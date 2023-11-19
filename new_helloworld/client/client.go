package main

import (
	"OldPackageTest/new_helloworld/client_proxy"
	"fmt"
)

func main() {
	//1.建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	//client, err := rpc.Dial("tcp", "localhost:1234")
	//客户端
	var reply string
	err := client.Hello("bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)

	//1. 这些概念在grpc中都有对应
	//2. 这些灵魂的拷问：server_proxy 和 cliet_proxy 能否自动生成， 为多种语言生成
	//3. 都能满足， 这个就是protobuf + grpc
}
