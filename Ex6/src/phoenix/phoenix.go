// Primary and backup module//
package phoenix

import (
	"../network"
	"fmt"
	"os/exec"
	"time"
	//"encoding/binary"
	"strconv"
)

func Primary(cnt int) {

	c_broadcast := make(chan []byte)
	buffer := ""

	go network.UDPBroadcast(c_broadcast)
	fmt.Printf("Starting broadcast\n")

	newHo := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run main1.go")
	//returnerer et sett datainstrukser. err = newHo.Run() //Starter ny terminal
	err := newHo.Run()
	if err != nil {
		fmt.Printf("Feil med exec")
	}
	for {

		fmt.Println(cnt)
		//binary.Write()
		buffer = strconv.Itoa(cnt)

		c_broadcast <- []byte(buffer)
		cnt += 1
		time.Sleep(1 * time.Second)

	}

}

func Backup(cnt int, c_listen chan []byte, c_close chan int) {
	buffer := ""

	for {
		select {

		case Melding := <-c_listen:
			buffer = string(Melding)
			fmt.Printf("Backupread:"+buffer+"\n")
			cnt, _ = strconv.Atoi(buffer)

		case <-time.After(2000 * time.Millisecond):
			c_close <- 1
			fmt.Println(cnt)
			go Primary(cnt)
			return
		}
	}

}
