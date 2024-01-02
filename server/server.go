package main

import( "fmt"
	 "io"
	 "net"
	 "os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create Buffer
	buffer := make([]byte, 1024)

	//receive filename from client
	fileNameBuffer := make([]byte, 64)

	n , err := conn.Read(fileNameBuffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := string(fileNameBuffer[:n])
	fmt.Println("Receive File Name", fileName)

	//create file to store data
	file, err := os.Create(fileName)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close() //Close file before exit

	// Receive and weire data to file
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF{
				fmt.Println("Transfer Complete")
			} else {
				fmt.Println(err)
			}
			return
		}
		//write data from buffer to file
		file.Write(buffer[:n])
	}
}

func main() {
	//Create Listenner
	listenner, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listenner.Close() // Close listenner before exit
	fmt.Println("Server is listening on port 5000")

	//Accept connection from client
	for{
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Print Client address
		fmt.Println("Client Connected:", conn.RemoteAddr())

		//Handle connection
		go handleConnection(conn)
	}
}