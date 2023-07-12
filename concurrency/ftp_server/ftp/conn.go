package ftp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type dataType int
type dataPort string

const (
	ascii dataType = iota
	binary
)

type FtpConn struct {
	conn     net.Conn
	RootDir  string
	WorkDir  string
	dataType dataType
	dataPort dataPort
}

func (c *FtpConn) String() string {
	b, _ := json.MarshalIndent(struct {
		RootDir string
		WorkDir string
		Type    dataType
		Port    string
	}{RootDir: c.RootDir, WorkDir: c.WorkDir, Type: c.dataType, Port: string(c.dataPort)}, "", " ")

	return string(b)
}

func NewConn(c net.Conn, rootDir string) *FtpConn {
	return &FtpConn{
		conn:    c,
		RootDir: rootDir,
		WorkDir: "/",
	}
}

func (c *FtpConn) status(code string) {
	log.Print(">> ", code)
	_, err := fmt.Fprintf(c.conn, "%v%v", code, c.EOL())
	if err != nil {
		log.Print(err)
	}
}

func (c *FtpConn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}

func (c *FtpConn) list(args []string) {
	var target string

	if len(args) > 0 {
		target = filepath.Join(c.RootDir, c.WorkDir, args[0])
	} else {
		target = filepath.Join(c.RootDir, c.WorkDir)
	}

	files, err := ioutil.ReadDir(target)
	if err != nil {
		log.Print(err)
		c.status(status550)
		return
	}

	c.status(status150)

	for _, file := range files {
		_, err := fmt.Fprintf(c.conn, "%s%s", file.Name(), c.EOL())
		if err != nil {
			log.Print(err)
			c.status(status426)
		}
	}

	_, err = fmt.Fprintf(c.conn, "%s", c.EOL())
	if err != nil {
		log.Print(err)
		c.status(status426)
	}

	c.status(status226)

}

func (c *FtpConn) cwd(args []string) {
	if len(args) != 1 {
		c.status(status501)
		return
	}

	workDir := filepath.Join(c.WorkDir, args[0])
	absPath := filepath.Join(c.RootDir, workDir)
	_, err := os.Stat(absPath)

	if err != nil {
		log.Print(err)
		c.status(status550)
		return
	}

	c.WorkDir = workDir
	c.status(status200)
}

func (c *FtpConn) port() {
	c.dataPort = dataPort(c.conn.RemoteAddr().String())

	c.status(status200)
}

func (c *FtpConn) user(args []string) {
	c.status(fmt.Sprintf(status230, strings.Join(args, " ")))
}

func (c *FtpConn) retr(args []string) {

}

func (c *FtpConn) setDataType(args []string) {
	if len(args) == 0 {
		c.status(status501)
		return
	}

	switch args[0] {
	case "A":
		c.dataType = ascii
	case "I":
		c.dataType = binary
	default:
		c.status(status504)
		return
	}
	c.status(status200)
}

func (c *FtpConn) dataConnect() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.conn.RemoteAddr().String())
	if err != nil {
		return nil, fmt.Errorf("Error parsing data port address: %v", err)
	}

	return conn, nil
}
