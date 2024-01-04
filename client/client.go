package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// ป้อน username และ password
	fmt.Print("Connecting to server...\n")
	fmt.Print("Enter username: ")
	username, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Print("Enter password: ")
	password, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// ตัดช่องว่างและขึ้นบรรทัดใหม่ที่เพิ่มเข้ามาจากการใช้ ReadString
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	// สร้างข้อมูลที่จะส่งไปยัง Server
	data := fmt.Sprintf("%s:%s", username, password)

	// เชื่อมต่อกับ Server
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// ส่งข้อมูลไปยัง Server
	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Error sending data to server:", err)
		return
	}

	// รับข้อมูลจาก Server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving data from server:", err)
		return
	}

	fmt.Println("Server :", string(buffer[:n]))
}