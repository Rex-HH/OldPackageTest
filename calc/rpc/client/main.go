package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

//rpc 远程过程调用，做到像本地调用

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	//传输协议：http
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
	body, _ := res.Body()
	//fmt.Println(string(body))
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}
func main() {
	fmt.Println(Add(1, 2))
}
