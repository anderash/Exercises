package main

import (
  "fmt"
  "net"
  "time"
  "os"
  "runtime"
)

func TCPsend() {
	str := "Connect to: 129.241.187.140:33546\x00"
//	p := make([]byte, 1024)
	

	serverAddr, err := net.ResolveTCPAddr("tcp","129.241.187.136:33546" )
	if err != nil {
		fmt.Printf("Addresse dritt")
		os.Exit(1)
	}

	//buffer := make([]byte,1024)
	socket, err2 := net.DialTCP("tcp", nil, serverAddr)

	if err2 != nil {
		fmt.Printf("Dial dritt")
		os.Exit(2)
	}	

	socket.Write([]byte(str))


}

func TCPconnect() {
	melding := "Martin er kuul \x00"
	// serverAddr, err := net.ResolveTCPAddr("tcp","129.241.187.136:20001" )
	// if err != nil {
	// 	fmt.Printf("Addresse dritt")
	// }
	buffer := make([]byte,1024)

	listener, err1 := net.Listen("tcp", ":33546")
	if err1 != nil {
		fmt.Printf("Socket dritt")
	}
	
	socket, err3 := listener.Accept()
	if err3 != nil {
		os.Exit(3)
	}
	fmt.Printf("kjører4?")

	for{
		socket.Write([]byte(melding))
		fmt.Printf("Skrevet")
	 	

		fmt.Printf("kjører?")
		n, err2 := socket.Read(buffer)
		fmt.Printf("lest")
		if err2 != nil {
			fmt.Printf("Read dritt")
		}	

  		fmt.Printf("%s%i \n", buffer, n)
    	time.Sleep(1000 * time.Millisecond)

	}

}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go TCPsend()
	go TCPconnect()

	c := make(chan int)
	<- c
}