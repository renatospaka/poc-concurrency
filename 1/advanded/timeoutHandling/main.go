package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(){
		time.Sleep(2*time.Second)
		ch <- 1
	}()

	select{
	case <-ch:
		fmt.Println("dados lidos do channel")
	case <-time.After(3*time.Second):
		fmt.Println("timeout")
	}
}
