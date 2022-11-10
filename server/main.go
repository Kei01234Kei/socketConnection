package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	port := 8080
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", strconv.Itoa(port)))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening at localhost:", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		data := make([]byte, 1024)
		count, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:count]))
	}
}
