package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"

	"git.in.chaitin.net/patronus/t1k-sdk/sdk/go/pkg/gosnserver"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sReq := "POST /form.php HTTP/1.1\r\n" +
		"Host: a.com\r\n" +
		"Content-Length: 40\r\n" +
		"Content-Type: application/json\r\n\r\n" +
		"{\"name\": \"youcai\", \"password\": \"******\"}"
	req, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer([]byte(sReq))))
	panicIf(err)

	conn, err := net.Dial("tcp", "169.254.0.5:8000")
	panicIf(err)
	defer conn.Close()
	result, err := gosnserver.DetectHttpRequest(conn, req)
	panicIf(err)

	if result.Passed() {
		fmt.Println("Passed")
	}
	if result.Blocked() {
		fmt.Println("Blocked")
	}
}
