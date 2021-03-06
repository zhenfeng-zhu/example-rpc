package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial error", err)
	}

	var reply service.String
	err = client.Call("HelloService.Hello", pb.String{Value: "你好，我是"}, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
