package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"JarvisMessage"
	// "os"
	"net"
)

// func verifyLogin(recive_msg ) {

// }

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
	
		if err != nil {
			fmt.Println("read data failed...")
		}
	
		msg_buf := buf[0:size]
	
		recive_msg := &JarvisMessage.LoginIn{}
	
		err = proto.Unmarshal(msg_buf, recive_msg)
	
		if err != nil {
			// fmt.Println("unmarshaling error: ", recive_msg)
		}
		fmt.Println("recive: ", recive_msg)
		fmt.Println("username: ", recive_msg.Username)
	}
}

func main()  {
	listen_socket, err := net.Listen("tcp", "0.0.0.0:8000")

	if err != nil {
		fmt.Println("server start failed!")
	}

	defer listen_socket.Close()

	fmt.Println("server is waiting ...")

	for {
		connect, err := listen_socket.Accept()
		if err != nil {
			fmt.Println("connect failed!")
		}

		go handleConn(connect)
	}
}