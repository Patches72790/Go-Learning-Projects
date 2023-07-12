package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fmt.Printf("Dialed to %v from %v\n", conn.RemoteAddr(), conn.LocalAddr())

	for {
		fmt.Printf("ftp> ")

		reader := bufio.NewReader(os.Stdin)
		s, _ := reader.ReadString('\n')
		buf := bytes.NewBuffer([]byte(s))
		io.Copy(conn, buf)

		buffer := make([]byte, 256)
		_, e := conn.Read(buffer)
		if e != nil {
			fmt.Println("Error: %v", e)
		}

		fmt.Println(string(buffer))
	}
}
