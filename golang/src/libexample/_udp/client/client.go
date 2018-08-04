package main

import (
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			send()
			wg.Done()
		}()
		wg.Add(1)
	}
	wg.Wait()
}

func send() {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()

	n, err := conn.Write([]byte("Ping"))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	recvBuf := make([]byte, 1024)

	n, err = conn.Read(recvBuf)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	log.Printf("Received data: %s", string(recvBuf[:n]))
}
