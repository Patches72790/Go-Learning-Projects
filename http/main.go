package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, "https://google.com", &bytes.Buffer{})
	req.ContentLength = 0
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	length := resp.ContentLength
	fmt.Println(length)
	var buffer = make([]byte, length)
	resp.Body.Read(buffer)

	fmt.Println(string(buffer))
	fmt.Println(resp.Status)
}
