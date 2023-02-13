package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

func keyString(k ssh.PublicKey) string {
	return k.Type() + " " + base64.StdEncoding.EncodeToString(k.Marshal())
}

func trustedHostKeyCallback(trustedKey string) ssh.HostKeyCallback {
	return func(_ string, _ net.Addr, k ssh.PublicKey) error {
		//ks := keyString(k)
		//		if trustedKey != ks {
		//			fmt.Println(trustedKey)
		//			fmt.Println(ks)
		//			return fmt.Errorf("Warning: trusted key does not much public key!")
		//		}
		return nil
	}
}

type SCPWriter struct {
	stuff []byte
	total int
}

type SCPReader struct {
	payload []byte
}

func (s *SCPWriter) Write(b []byte) (int, error) {
	s.stuff = b[:]
	return len(b), nil
}

func (r *SCPReader) Read(b []byte) (int, error) {
	b = r.payload[:]
	fmt.Println(b)
	return len(r.payload), nil
}

func main() {
	keyBytes, _ := ioutil.ReadFile("/Users/plharvey/.ssh/lsstaging_key.pub")

	clientConfig, err := auth.PrivateKey("plharvey", "/Users/plharvey/.ssh/lsstaging_key", trustedHostKeyCallback(string(keyBytes)))
	if err != nil {
		log.Fatalln(err)
	}

	client := scp.NewClient("lsstaging.doit.wisc.edu:22", &clientConfig)

	err = client.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	//	scpWriter := SCPWriter{stuff: make([]byte, 128)}
	//
	//	// copy from remote to writer in memory
	//	err = client.CopyFromRemotePassThru(context.Background(), &scpWriter, "/home/plharvey/testfile", func(r io.Reader, n int64) io.Reader {
	//		buffer := make([]byte, 0)
	//		r.Read(buffer)
	//		fmt.Println("Read from remote file: ", string(buffer))
	//		fmt.Println("Bytes read: ", n)
	//		return r
	//	})

	scpReader := SCPReader{payload: make([]byte, 0)}
	scpReader.payload = []byte("hello world test!")
	err = client.Copy(context.Background(), &scpReader, "/home/plharvey/testfile2", "555", int64(len(scpReader.payload)))

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println("Written to writer: ", string(scpWriter.stuff))
}
