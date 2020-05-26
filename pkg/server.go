package pkg

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func StartServer(config *ServerConfig, port int) (err error) {

	var l net.Listener
	l, err = net.Listen("tcp", "localhost:" + strconv.Itoa(port))
	if err != nil {
		log.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		log.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}

	return
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}