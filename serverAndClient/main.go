package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

func server(port int) {
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
		fmt.Println("Served string from client:", string(data[:count]))
	}
}

func client(port int) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		content := scanner.Text()
		conn, err := net.Dial("tcp", fmt.Sprintf(":%s", strconv.Itoa(port)))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		_, err = conn.Write([]byte(content))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Sent string to server:", content)
	}
}

func main() {
	serverPort, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Enter valid command line argument.")
		fmt.Println(err)
		return
	}
	clientPort, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Enter valid command line argument.")
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		server(serverPort)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		client(clientPort)
	}()
	wg.Wait()
}
