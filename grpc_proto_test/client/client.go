package main

import (
	"OldPackageTest/grpc_proto_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "bobby",
		//Url:  "imooc.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
