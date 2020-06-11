package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

type Server struct {
	Config          *ServerConfig
	Port            int
	PeerConnections []*net.Conn
}

// StartServer good
func (svc *Server) Start() (err error) {

	var l net.Listener
	l, err = net.Listen("tcp", "localhost:"+strconv.Itoa(svc.Port))
	if err != nil {
		log.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	go func() {
		for _, host := range svc.Config.ClusterHosts {
			conn, ok := svc.PeerConnections[host]
			if ok {

			} else {
				conn, err := net.Dial("tcp", *host+":"+*port)
				if err != nil {
					fmt.Println("Error connecting:", err)
					continue
				}
			}
		}
	}()

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
	// 使用 bufio 标准库提供的缓冲区功能
	reader := bufio.NewReader(conn)
	for {
		// ReadString 会一直阻塞直到遇到分隔符 '\n'
		// 遇到分隔符后会返回上次遇到分隔符或连接建立后收到的所有数据, 包括分隔符本身
		// 若在遇到分隔符之前遇到异常, ReadString 会返回已收到的数据和错误信息
		msg, err := reader.ReadString('\n')
		if err != nil {
			// 通常遇到的错误是连接中断或被关闭，用io.EOF表示
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}
		b := []byte("resp " + msg)
		// 将收到的信息发送给客户端
		conn.Write(b)
	}
}
