// main
package main

import (
	"./network"
	"./phoenix"
	"runtime"
	"time"
	"strconv"
)

func main() {

	c_listen := make(chan []byte)
	//c_broadcast := make(chan []byte)

	cnt := 0

	runtime.GOMAXPROCS(runtime.NumCPU())

	go network.UDPListen(c_listen)

	select {
	case melding := <-c_listen: //Skal bli backup
		buffer := string(melding)
		cnt, _ = strconv.Atoi(buffer)
		go phoenix.Backup(cnt, c_listen)

	case <-time.After(3000 * time.Millisecond):
		go phoenix.Primary(cnt)

	}

	c := make(chan int)
	<- c

	//fmt.Printf("Antall bytes sendt: %i", nrBsendt)

}
