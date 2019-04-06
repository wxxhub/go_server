package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"example"
	"os"
)

func write() {
	fmt.Println("test")
	test := &example.Test{
		Id: 1110,
		Name: "wxx",
	}

	fmt.Println(test)

	data, _ := proto.Marshal(test);
	fmt.Println("data: ", data)
	ioutil.WriteFile("test.txt", data, os.ModePerm)
}

func read() {
	data, _ := ioutil.ReadFile("test.txt")
	test := &example.Test{}
	proto.Unmarshal(data, test);
	fmt.Println("read: ", test)
}

func main()  {
	write()
	read()
}
