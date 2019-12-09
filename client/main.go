package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connection success")
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("enter something: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if text == "\n" {
			continue
		}
		if text == "exit\n" {
			return
		}
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("message was sent")

		response := make([]byte, 1024)
		_, err = conn.Read(response)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("response: ")
		fmt.Println(string(response))
	}
}