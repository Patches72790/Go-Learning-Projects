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

	fmt.Printf("ftp> ")
	for {
		write(&conn)
		go read(&conn)
	}
}
func write(conn *net.Conn) error {

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	buf := bytes.NewBuffer([]byte(s))
	_, e := io.Copy(*conn, buf)

	return e
}

func read(conn *net.Conn) {
	buffer := make([]byte, 1024)

	(*conn).Read(buffer)

	fmt.Println()
	fmt.Println(string(buffer))
	fmt.Printf("ftp> ")
}
