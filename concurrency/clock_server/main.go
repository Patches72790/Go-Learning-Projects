package clockserver

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func ClockServer() {
	p := flag.Int("port", 8000, "Port to listen on")
	flag.Parse()

	tz := os.Getenv("TZ")

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *p))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on port %d\n", *p)

	loc, err := time.LoadLocation(tz)

	fmt.Println(loc)
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
