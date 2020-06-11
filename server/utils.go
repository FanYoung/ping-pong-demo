package server

import (
	"log"
	"net"
	"os"
	"strings"
)

func GetOutboundIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Printf("Can't detect OutboundIP %v\n", err)
		os.Exit(1)
	}

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
