package main
import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
// This is probably not a good idea for large projects...
	"runtime"
	"time"
)



func someGoroutine1(chan c int) {
	for j := 0; j<1000000; j++ {
		i:= <- c
		i++
		c <- i
	}
}
func someGoroutine2(chan c int) {
	for j := 0; j<1000000; j++ {
		i := <- c
		i--
		c <- i
	}
}
func main() {
	c := make (chan int, 1)
	c <- 0
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
										// Try doing the exercise both with and without it!
	go someGoroutine1(c) // This spawns someGoroutine() as a goroutine
	go someGoroutine2(c)
// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
// We'll come back to using channels in Exercise 2. For now: Sleep.
	time.Sleep(100*time.Millisecond)
	Println(i)
	Println("Hello from main!")
}
