package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"JarvisMessage"
	"net"
	"time"
)

func main()  {
	clint, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("clint start failed")
		return
	}

	send_msg := &JarvisMessage.LoginIn{
		Username: "中国移动",
		Password: "1008611",
	}

	msg, _ := proto.Marshal(send_msg)

	for {
		clint.Write([]byte(msg))
		time.Sleep(2)
	}
}