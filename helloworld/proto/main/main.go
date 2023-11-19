package main

import (
	__ "OldPackageTest/helloworld/proto"
	"fmt"
	"google.golang.org/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := __.HelloRequest{
		Name:    "bobby",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	//jsonStruct := Hello{
	//	Name:    "bobby",
	//	Age:     18,
	//	Courses: []string{"go", "gin", "微服务"},
	//}
	//jsonRsp, _ := json.Marshal(jsonStruct)
	//fmt.Println(len(jsonRsp))
	rsp, _ := proto.Marshal(&req) //具体的编码如何做到的，那大家可以百度一下protobuf的原理 varint
	//fmt.Println(len(rsp))
	newReq := __.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
