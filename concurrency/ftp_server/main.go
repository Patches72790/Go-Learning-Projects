package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/patches72790/ftp_server/ftp"
)

var port int
var rootDir string

func init() {
	flag.IntVar(&port, "port", 8080, "Port for ftp server")
	flag.StringVar(&rootDir, "rootDir", "public/", "Default directory for ftp server")
	flag.Parse()
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("FTP server listening on port %d\n", port)

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

	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	ftp.Serve(ftp.NewConn(c, absPath))
}
