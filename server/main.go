package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

func processInput(input string) string {
	number, err := strconv.Atoi(input)
	if err == nil {
		return strconv.Itoa(number * 2)
	}
	return strings.ToUpper(input)
}

func handleConnection(conn net.Conn) {
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println("client left..")
		conn.Close()
		return
	}
	message := strings.TrimSuffix(string(bufferBytes), "\n")
	conn.Write([]byte(processInput(message)))
	handleConnection(conn)
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("tcp server listener error:", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error", err)
			return
		}
		go handleConnection(conn)
	}
}