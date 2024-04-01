package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	concurrency := 3
	semaphore := make(chan struct{}, concurrency)
	tasks := []string{"tarefa 1", "tarefa 2", "tarefa 3"}

	for _, task := range tasks {
		wg.Add(1)

		go func(t string) {
			defer wg.Done()

			semaphore <- struct{}{}
			fmt.Println("executando", t)
			
			//sinalizando que a tarefa foi concluída
			<-semaphore
		}(task)
	}
	wg.Wait()
	fmt.Println("todas as tarefas foram concluídas")
}
