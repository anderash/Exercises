package main

import (
  "fmt"
  "net"
)

func main() {
  buffer byte
  conn, err := net.Dial("udp", ":30000")
  addr := conn.LocalAddr()
  fmt.Printf(addr.String())
}
  
