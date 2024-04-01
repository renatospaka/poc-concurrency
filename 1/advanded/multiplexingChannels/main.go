package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(){
		time.Sleep(2*time.Second)
		ch1 <- 1
	}()

	go func(){
		time.Sleep(1*time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("recebido de ch1")
	case <-ch2:
		fmt.Println("recebido de ch2")
	case <-time.After(3*time.Second):
		fmt.Println("timeout")
	}
}
