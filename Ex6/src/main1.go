// main
package main

import (
	"./network"
	"./phoenix"
	"runtime"
	"time"
	"strconv"
	"fmt"
)

func main() {

	c_listen := make(chan []byte, 1024)
	c_close := make(chan int, 8)
	//c_broadcast := make(chan []byte)

	cnt := 0

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("For UDP\n")
	go network.UDPListen(c_listen, c_close)
	fmt.Printf("Select state\n")
	select {
	case melding := <-c_listen: //Skal bli backup
		buffer := string(melding)
		cnt, _ = strconv.Atoi(buffer)
		fmt.Printf("For backup")
		go phoenix.Backup(cnt, c_listen, c_close)


	case <-time.After(3000 * time.Millisecond):
		
		fmt.Printf("Skal sende\n")
		c_close <- 1
		fmt.Printf("Sendt\n")
		go phoenix.Primary(cnt)




	}

	c := make(chan int)
	<- c

	//fmt.Printf("Antall bytes sendt: %i", nrBsendt)

}
