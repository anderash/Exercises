package main

import (
  "fmt"
  "net"
)

// constant ()
func main() {

	serverAddr, err := net.ResolveUDPAddr("udp","129.241.187.255:30000" )

 	buffer := make([]byte,1024)
  // conn, err := net.Dial("udp", "129.241.187.255:30000")
  if err != nil {
    fmt.Printf("Noe gikk til helvette")
  }

	connUDP, _ := net.ListenUDP("udp4", serverAddr)  

	n, _, _ := connUDP.ReadFromUDP(buffer)

	fmt.Printf("%s%i \n", buffer, n)
}
  
