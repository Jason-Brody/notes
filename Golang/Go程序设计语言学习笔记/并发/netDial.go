package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		fmt.Println("not blocked")
		io.Copy(os.Stdout, conn)

		log.Println("done")
		done <- struct{}{}
		fmt.Println("not blocked")
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
