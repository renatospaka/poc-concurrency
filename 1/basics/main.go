package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 10
	select {
	case val := <-ch:
		fmt.Printf("Recebido 1 de ch: %d\n", val)
	case val := <-ch:
		fmt.Printf("Recebido 2 de ch: %d\n", val)
	case val := <-ch:
		fmt.Printf("Recebido 3 de ch: %d\n", val)
	default:
		fmt.Println("esse é o default")
	}
}
