package main

import (
  "fmt"
  "net"
)

func main() {
 // var buffer byte = 0 
  conn, err := net.Dial("udp", ":30000")
  if err != nil {
    fmt.Printf("Noe gikk til helvette")
  }
  
  addr := conn.LocalAddr()
  fmt.Printf(addr.Network())
}
  
