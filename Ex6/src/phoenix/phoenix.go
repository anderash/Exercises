// Primary and backup module//
package phoenix

import (
	"../network"
	"fmt"
	"os/exec"
	"time"
	//"encoding/binary"
	"strconv"
	"strings"
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
	netCnt := cnt
	for {
		select {

		case Melding := <-c_listen:
			buffer = string(Melding[0:2])
			buffer = strings.Trim(buffer, "")
			for i := 0; i < 2; i++ {
				fmt.Println(buffer[i])
			}			
			fmt.Printf("Backupread:"+buffer+"\n")
			netCnt, _ = strconv.Atoi(buffer)
			fmt.Println(netCnt)

		case <-time.After(2000 * time.Millisecond):
			c_close <- 1
			fmt.Printf("cnt nr primary: ")
			fmt.Println(netCnt)
			go Primary(netCnt)
			return
		}
	}

}
