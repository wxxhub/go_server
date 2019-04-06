package main

import (
	"fmt"
	"net"
)

func main()  {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("server start error")
	}

	defer listen_socket.Close()
	fmt.Println("server is waiting ...")

	for {
		conn, err := listen_socket.Accept()
		if err != nil {
			fmt.Println("connect faile")
		}

		fmt.Println("connect client successed!")

		var msg string
		for {
			msg = ""
			data := make([]byte, 255)
			n, err := conn.Read(data)

			if n == 0 || err != nil {
				break
			}

			fmt.Println("client say: ", string(data[:n]))
			fmt.Println("say to client: ")
			fmt.Scan(&msg)
			conn.Write([]byte(msg))
		}
		fmt.Println("client Close")
		conn.Close()
	}
}
