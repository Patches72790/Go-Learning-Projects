package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	args := os.Args

	fmt.Println(args[1:])

	for _, tzInfo := range args[1:] {
		go handleTimeZone(tzInfo)
	}

	for {
	}
}

func handleTimeZone(tzInfo string) {
	split := strings.Split(tzInfo, ":")

	tz := strings.Split(split[0], "=")[0]
	port := split[1]

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Printf("Listening for timezone %s on port %s\n", tz, port)

	mustCopy(os.Stdout, conn, tz)
}

func mustCopy(dst io.Writer, src io.Reader, tz string) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
