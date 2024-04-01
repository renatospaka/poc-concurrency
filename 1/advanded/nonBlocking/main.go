package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	fmt.Println("dados escritos no channel - 1")

	select {
	case ch <- 2:
		fmt.Println("dados esctiros no channel - 2")
	default:
		fmt.Println("channel está cheio/vazio, impossível escrever dados")
	}

	// tentar ler dados do channel quando já contém dados não bloqueará a operação
	dados, ok := <-ch
	if ok {
		fmt.Println("dados lidos do channel:", dados)
	}

	select {
		// teentativa de ler dados de um canal vazio resultará no retorno imediato
		// porque o canal está vazio e não é possível completar a operação de leitura
		case dados, ok := <-ch:
			if ok {
				fmt.Println("dados lidos do channel:", dados)
			} else {
				fmt.Println("channel está vazio, impossível ler dados")
			}
		default:
			fmt.Println("channel está vazio, impossível ler dados")
	}
}
