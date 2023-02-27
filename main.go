package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	const (
		tcpPort = 9090
		udpPort = 9091
	)

	errChan := make(chan error)

	go func() {
		errChan <- startTcp(tcpPort)
	}()

	go func() {
		errChan <- startUdp(udpPort)
	}()

	log.Fatalf("server error: %s", <-errChan)
}

func startUdp(port int) error {
	const network = "udp"

	listener, err := net.ListenUDP(network, &net.UDPAddr{Port: port})
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Println(network, "server ready and listening")

	for {
		buff := make([]byte, 1<<7)
		_, remote, err := listener.ReadFromUDP(buff)
		if err != nil {
			log.Printf("%s: error handling request: %s", network, err)
		}

		log.Printf(
			"%s: incoming connection from remote: %s, message: %s",
			network, remote, buff,
		)
	}
}

func startTcp(port int) error {
	const network = "tcp"

	listener, err := net.Listen(network, fmt.Sprint(":", port))
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Println(network, "server ready and listening")

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handle(network, conn)
	}
}

func handle(network string, conn net.Conn) {
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	buff := make([]byte, 1<<7)
	_, err := conn.Read(buff)
	if err != nil {
		log.Printf("%s: error handling request: %s", network, err)
		return
	}

	log.Printf(
		"%s: incoming connection from remote: %s, message: %s",
		network, conn.RemoteAddr(), buff,
	)

	if network != "tcp" {
		return
	}

	_, err = conn.Write([]byte(fmt.Sprint("hello from the other side. time: ", time.Now().Format(time.ANSIC))))
	if err != nil {
		log.Printf("%s: error writing response: %s", network, err)
	}
}
