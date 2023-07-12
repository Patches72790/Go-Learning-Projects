package ftp

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func Serve(ftpConn *FtpConn) {
	ftpConn.status(status220)

	log.Printf("Serving connection from %s\n", ftpConn.conn.LocalAddr())

	scanner := bufio.NewScanner(ftpConn.conn)

	for scanner.Scan() {

		fmt.Println(ftpConn)

		input := strings.Fields(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {

		case "LIST":
			ftpConn.list(args)
		case "CWD":
			ftpConn.cwd(args)
		case "PORT":
			ftpConn.port()
		case "USER":
			ftpConn.user(args)
		case "QUIT":
			ftpConn.status(status221)
			return
		case "RETR":
			ftpConn.retr(args)
		case "TYPE":
			ftpConn.setDataType(args)
		default:
			ftpConn.status(status502)
		}

	}

	if scanner.Err() != nil {
		log.Print(scanner.Err())
	}
}
