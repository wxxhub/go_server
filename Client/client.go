package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"JarvisMessage"
	"net"
	"time"
)

func reciveMsg(conn net.Conn) string {

	return ""
}

func LoginResult(conn net.Conn) bool {
	
	for i := 0; i < 3; i++ {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			continue
		} else {
			msg := buf[0:size]
			result_msg := &JarvisMessage.LoginResult{}

			unmar_err := proto.Unmarshal(msg, result_msg)
			 if unmar_err != nil {
				 continue
			 }
			 return result_msg.Result
		}
	}
	return false
}

func Login(conn net.Conn, user_name string, pass_word string) bool {

	login_msg := &JarvisMessage.Login{
		Username: user_name,
		Password: pass_word,
	}

	msg, _ := proto.Marshal(login_msg)
	_, write_err := conn.Write([]byte(msg))

	if write_err != nil {
		fmt.Println("send failed!")
		return false
	}
	time.Sleep(100*time.Millisecond)

	if LoginResult(conn) == false {
		fmt.Println("login failed!")
		return false
	} else {
		fmt.Println("login success!")
		return true
	}

	return false;
}

func main()  {
	clint, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("clint start failed")
		return
	}
	defer clint.Close()

	if Login(clint, "中国移动", "1008611") == false {
		return
	}

	send_msg := &JarvisMessage.Message{
		Message: "Hello, I'm Client",
	}

	msg, _ := proto.Marshal(send_msg)

	for {
		_, write_err := clint.Write([]byte(msg))
		if write_err != nil {
			fmt.Println("connect failed")
			break
		}
		time.Sleep(1*time.Second)
	}
}