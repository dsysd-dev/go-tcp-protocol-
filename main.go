package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	l, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalln("couldn't listen to network")
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("err while accept", err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	// we keep reading all the messages from the connection
	// until connection is closed
	fmt.Println("connected to: ", conn.RemoteAddr())
	for {
		var buffer [1024]byte // 1KB
		_, err := conn.Read(buffer[:])
		if err != nil {
			log.Printf("err while reading from conn: %v, exiting ...", err)
			return
		}
		fmt.Println("message read: ", string(buffer[:]))
	}
}
