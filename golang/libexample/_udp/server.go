package main

import (
	"log"
	"net"
	"os"
)

func main() {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	updLn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	log.Println("Starting udp server...")

	for {
		n, addr, err := updLn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		log.Println("hoge")

		go func() {
			log.Printf("Reciving data: %s from %s", string(buf[:n]), addr.String())

			log.Printf("Sending data..")
			updLn.WriteTo([]byte("Pong"), addr)
			log.Printf("Complete Sending data..")
		}()
	}
}
