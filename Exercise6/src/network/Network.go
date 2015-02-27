// NETWORK MODULE //
package network

import (
	"fmt"
	"net"
	"time"
	"os"
)

const (
	OwnIP   = "129.241.187.140"
	OwnPort = "30001"
	Baddr   = "129.241.187.255"
)

func UDPBroadcast(c_broadcast chan []byte) {

	raddr, err1 := net.ResolveUDPAddr("udp", "localhost:"+OwnPort)

	if err1 != nil {
		fmt.Printf("Problemer med resolveUDPaddr")
		os.Exit(1)
	}

	socket, err2 := net.DialUDP("udp", nil, raddr)

	if err2 != nil {
		fmt.Printf("Problemer med Dial")
		os.Exit(2)
	}

	for {
		buffer := <-c_broadcast
		_, err3 := socket.Write(buffer)
		//fmt.Printf("skrev %i bytes", n)

		if err3 != nil {
			fmt.Printf("Problemer med Write")
			os.Exit(3)
		}

	}

}

func UDPListen(c_listen chan []byte, c_close chan int) {
	buffer := make([]byte, 1024)

	raddr, err1 := net.ResolveUDPAddr("udp", "localhost:"+OwnPort)

	if err1 != nil {
		fmt.Printf("Problemer med resolveUDPaddr")
		os.Exit(4)
	}

	socket, _ := net.ListenUDP("udp", raddr)
	fmt.Printf("Opened socket\n")

	for {
		select{
			case lukk := <- c_close:
				fmt.Printf("skal lukke\n")
				if (lukk == 1){
					fmt.Printf("Closing\n")
					socket.Close()

					fmt.Printf("Closed connection\n")
					return
				}
				
			case <-time.After(500 * time.Millisecond):
				_, err4 := socket.Read(buffer)

				if err4 != nil {
					fmt.Printf("Problemer med readUDP")
					os.Exit(5)
				}

				//fmt.Printf(string(buffer))
				c_listen <- buffer
				//time.Sleep(100*time.Millisecond)
		}


	}

}
