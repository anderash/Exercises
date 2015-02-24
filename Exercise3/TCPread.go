package main

import (
  "fmt"
  "net"
  "time"
)

func TCPconnect() {

	// serverAddr, err := net.ResolveTCPAddr("tcp","129.241.187.136:20001" )
	// if err != nil {
	// 	fmt.Printf("Addresse dritt")
	// }

	buffer := make([]byte,1024)

	listener, err1 := net.Listen("tcp", ":34933")
	if err1 != nil {
		fmt.Printf("Socket dritt")
	}

	for{
		socket, _ := listener.Accept()

		n, err2 := socket.Read(buffer)

		if err2 != nil {
		fmt.Printf("Read dritt")
		}	

  		fmt.Printf("%s%i \n", buffer, n)
    	time.Sleep(1000 * time.Millisecond)

	}

}

func main() {
	TCPconnect()
}