package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	company Company
}

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintln(employee Employee) {
	//rpc中的第二个点 传输协议， 数据编码协议
	//http 协议， 1.x ， 2.0协议
	//http 协议底层使用的是tcp http现在主流的是1.x 这种协议有性能问题，
	//一次性，一点结果返回，连接就断开
	// 1. 直接自己基于tcp/udp 协议去封装一层协议，myhttp， 没有通用性，http2.0
	//既有http的特性，也有长连接的特性，grpc
	/*
			客户端
				1. 建立连接 tcp/http
				2. 将employee 对象序列化成json字符串 - 序列化
				3. 发送json 字符串 - 调用成功后实际上你接受到的是一个二进制数据
				4.等待服务器发送结果
				5.将服务器返回的树蕨解析成PrintResult 对象 - 反序列化
			服务端
				1.监听网路端口 80
				2.读取数据 - 二进制的json数据
				3.对数据进行反序列化Employee对象
				4.开始处理业务逻辑
				5.奖处理的结果PrintResult序列化成json二进制数据 - 序列化
				6.将数据返回
		序列化和反序列化是可以选择的，不一定要采用json，xml，proto，msgpack
	*/

}

func main() {
	//现在我们想把Add函数编程一个远程的函数调用，也就意味着要把Add函数放在远程服务器上去运行
	/*
			我们原本的电商系统，这里有一段逻辑，这个逻辑是扣减库存， 但是库存服务是一个独立的系统， reduce
			那如何调用，一定会牵扯到网络，做成一个web服务（gin， beego， net/httpserver）

			1. 这个函数的调用参数如何传递-json （json是一种数据格式协议）/xml/proto/msgpack -编码协议
			   json并不是一个高效的协议
		       现在网络有两个端 - 客户端，应该干嘛？将数据传输到gin
		       gin-服务端，服务端负责解析数据

	*/
	//fmt.Println(Add(1, 2))
	//将这个打印的工作放在另一台服务器上 ，我需要将本地的内存对象 struct ，这样不行
	//可行的方法就是将struct 序列化成json对象，并进行传输

	fmt.Println(Employee{
		Name: "bobby",
		company: Company{
			Name:    "慕课网",
			Address: "北京市",
		},
	})
	//远程的服务需要将二进制对象反向解成struct对象
	//搞这么麻烦，这全部使用json去格式化不香吗？ 这种做法再浏览器和gin服务之间是可行的
	//但是如果你是一个大型的分布式系统

}
