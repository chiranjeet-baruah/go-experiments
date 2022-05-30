package goWaitGroup

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RunWaitGroupExample runs wait group example
func RunGoWaitGroupExample() {
	var wg sync.WaitGroup
	worker_count := 10
	for i := 1; i <= worker_count; i++ {
		wg.Add(1)
		go doSomethingWaitGroup(i, &wg)
		time.Sleep(time.Duration(rand.Intn(5-1)+1) * time.Second)
	}
	wg.Wait()
	fmt.Println("All workers have completed")
}

func doSomethingWaitGroup(worker_id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Started worker : ", worker_id)
	time.Sleep(time.Duration(rand.Intn(10-5)+5) * time.Second)
	fmt.Println("Finished worker : ", worker_id)
}
