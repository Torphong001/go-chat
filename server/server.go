package main

import (
	"fmt"
	"net"

)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	clientData := string(buffer[:n])

	// ตรวจสอบข้อมูล
	if clientData == fmt.Sprintf("%s:%s", "std1", "p@ssw0rd") {

		conn.Write([]byte("Hello\n"))
	} else {

		conn.Write([]byte("Invalid credentials\n"))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 5000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New connection established")

		go handleConnection(conn)
	}
}
