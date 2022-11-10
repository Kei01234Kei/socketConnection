package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port := 8080
	if len(os.Args) == 1 {
		fmt.Println("Error: Enter command line argument")
		return
	}
	conn, err := net.Dial("tcp", fmt.Sprintf(":%s", strconv.Itoa(port)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sent string to server:", os.Args[1])
}
