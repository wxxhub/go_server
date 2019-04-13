package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"JarvisMessage"
	// "os"
	"io"
	"net"
)

func returnResult(conn net.Conn, result bool) {
	result_msg := &JarvisMessage.LoginResult{
		Result: true,
	}

	return_msg,_ := proto.Marshal(result_msg)
	conn.Write([]byte(return_msg))
}

func verifyLogin(conn net.Conn) bool {
	for i := 0; i < 3; i++{
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
	
		if err != nil {
			fmt.Println("read data failed...")
			continue
		}
	
		msg_buf := buf[0:size]
	
		recive_msg := &JarvisMessage.Login{}
	
		err = proto.Unmarshal(msg_buf, recive_msg)
	
		if err != nil {
			fmt.Println("unmarshaling error: ", recive_msg)
			continue
		}
		
		if recive_msg.Username == "中国移动" && recive_msg.Password == "1008611" {
			return true
		}
	}

	return false
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// login verify
	verify_result := verifyLogin(conn)
	returnResult(conn, verify_result)
	if verify_result == false {
		conn.Close()
		return
	}

	// revive msg
	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
	
		if err != nil {
			fmt.Println("read data failed...")
			if err == io.EOF {
				fmt.Println("client closed!")
				return
			}
		}
	
		msg_buf := buf[0:size]
	
		recive_msg := &JarvisMessage.Message{}
	
		err = proto.Unmarshal(msg_buf, recive_msg)
	
		if err != nil {
			// fmt.Println("unmarshaling error: ", recive_msg)
		}
		fmt.Println("I recive: ", recive_msg.Message)
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