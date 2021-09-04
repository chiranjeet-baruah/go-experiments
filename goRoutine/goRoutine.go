package goRoutine

import (
	"fmt"
	"time"
)

// RunGoRoutineExample runs go routine example
func RunGoRoutineExample() {
	go doSomethingGoRoutine("banana", 5)
	doSomethingGoRoutine("apple", 3)
}

func doSomethingGoRoutine(thing string, count int) {
	for i := 0; i <= count; i++ {
		fmt.Println(i+1, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
