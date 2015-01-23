packadge main

import (
  "fmt"
  "net"
)

func main() {
  buffer := byte
  l, err := net.Dial("udp", ":30000")
  fmt.Printf(l)
}
  
