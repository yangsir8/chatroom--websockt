package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil{
		fmt.Println("erro:", err)
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)

	for{
		input, _ := inputReader.ReadString('\n')
		inputEnd := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputEnd) == "Q"{
			return
		}
		conn.Write(([]byte(inputEnd)))

		buf := [1024]byte{}
		reader := bufio.NewReader(conn)
		n,_ := reader.Read(buf[:])
		fmt.Println("recive: ", string(buf[:n]))


	}
}