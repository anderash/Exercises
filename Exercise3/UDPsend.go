package main

import (
  "fmt"
  "net"
  "time"
)


func sendUDP() {

rport := "20001"
//lport := "30000"
serverIP := "129.241.187.255"

  serverAddr, _ := net.ResolveUDPAddr("udp", serverIP+":"+rport )
 // ownAddr, _ := net.ResolveUDPAddr("udp", serverIP+:":"+rport)
  str := "Sindre Langeveld har skylda"
  //info := make([]byte,1024)

  UDPconn, _ := net.DialUDP("udp", nil, serverAddr)
  
  
  for {
  	n, _ := UDPconn.Write([]byte(str))
  	time.Sleep(1000 * time.Millisecond)
  	fmt.Printf("%i", n)
  }

}

  func main() {
  	sendUDP()
  }