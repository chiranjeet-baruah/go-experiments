package goWaitGroup

import (
	"fmt"
	"sync"
	"time"
)

// RunWaitGroupExample runs wait group example
func RunGoWaitGroupExample() {
	var wg sync.WaitGroup
	animals := []string{"dog", "cat", "monkey", "donkey", "elephant"}
	for i, animal := range animals {
		wg.Add(1) // creates one entry for each
		go doSomethingWaitGroup(animal, i+1, &wg)
	}
	wg.Wait() // waits for all tasks to get over
}

func doSomethingWaitGroup(thing string, count int, wg *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		fmt.Println(i+1, thing)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}
