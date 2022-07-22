package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("I got your UDP message"), addr)
	if err != nil {
		fmt.Printf("Error with response: %v", err)
	}
}

func server() {
	p := make([]byte, 1024)
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}

	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

	defer server.Close()

	for {
		_, remoteAddr, err := server.ReadFromUDP(p)
		fmt.Printf("Read message from %v : %s \n", remoteAddr, p)
		if err != nil {
			fmt.Printf("Error reading %v\n", err)
		}
		go sendResponse(server, remoteAddr)
	}
}

func client() {
	host := "127.0.0.1:1234"
	buffer := make([]byte, 1024)
	conn, err := net.Dial("udp", host)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		return
	}

	defer conn.Close()
	fmt.Fprintf(conn, "Hello World")

	_, err = bufio.NewReader(conn).Read(buffer)
	if err == nil {
		fmt.Printf("%s\n", buffer)
	} else {
		fmt.Printf("Error %s\n", err)
	}
}
func main() {
	go server()
	for {
		go client()
		time.Sleep(time.Millisecond * 1000)
	}
}
