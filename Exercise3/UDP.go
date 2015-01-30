package main

import (
  "fmt"
  "net"
  "time"
)

// constant ()

func callUDP() {
    serverAddr, err := net.ResolveUDPAddr("udp","129.241.187.255:20001" )

    buffer := make([]byte,1024)

    // conn, err := net.Dial("udp", "129.241.187.255:30000")
    if err != nil {
      fmt.Printf("Noe gikk til helvette")
    }

    socket, _ := net.ListenUDP("udp4", serverAddr)
    defer socket.Close()

  for {
  	n, _, _ := socket.ReadFromUDP(buffer)

  	fmt.Printf("%s%i \n", buffer, n)
    time.Sleep(1000 * time.Millisecond)
  }
  

}

func main() {
  callUDP()
  
}