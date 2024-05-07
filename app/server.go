package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Response struct{
	statusLine string // HTTP/1.1 200 OK \r\n
	responseHeader string // \r\n if it's empty
	responseBody string // this can be optional
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	response := Response{}
	buff := make([]byte,1024)
	_, err := conn.Read(buff)
	if err != nil{
		fmt.Println("Error reading data: ", err.Error())
		return
	}
	data := string(buff)
	requestSections := strings.Split(data, "\r\n")
	startline := strings.Split(requestSections[0], " ")
	path := startline[1]
	if path == "/"{
		response.statusLine = "HTTP/1.1 200 OK\r\n"
		response.responseHeader = "\r\n"
		response.responseBody = "" 
	}else{
		response.statusLine = "HTTP/1.1 404 Not Found\r\n"
		response.responseHeader = "\r\n"
		response.responseBody = ""
	}
	responseString := fmt.Sprintf("%s\r\n",response.statusLine)
	conn.Write([]byte(responseString))
}
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	
	l, err := net.Listen("tcp", "127.0.0.1:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()
	for{
		conn, err := l.Accept()
		if err != nil{
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}
