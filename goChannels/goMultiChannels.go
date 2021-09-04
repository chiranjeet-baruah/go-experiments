package goChannels

import (
	"fmt"
	"math/rand"
	"time"
)

// GetRandomSleep gets a random value of sleep everytime
func GetRandomSleep(max int) (sleep int) {
	rand.Seed(time.Now().UnixNano())
	sleep = rand.Intn(max)
	return sleep
}

// RunWaitGroupExample runs wait group example
func RunGoMultiChannelsExample() {
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	// Push to channel 1
	go func() {
		message_counter := 1
		for {
			msg := fmt.Sprint("CH1 Message no. ", message_counter)
			channel_1 <- msg
			message_counter = message_counter + 1

			time.Sleep(time.Duration(GetRandomSleep(7)) * time.Second)
		}
	}()

	// Push to channel 2
	go func() {
		message_counter := 1
		for {
			msg := fmt.Sprint("CH2 Message no. ", message_counter)
			channel_2 <- msg
			message_counter = message_counter + 1

			time.Sleep(time.Duration(GetRandomSleep(7)) * time.Second)
		}
	}()

	// Get messages from channel
	for {
		select {
		case msg1 := <-channel_1:
			fmt.Println("From channel 1: ", msg1)
		case msg2 := <-channel_2:
			fmt.Println("From channel 2: ", msg2)
		}
	}
}

// func doSomethingGoMultiChannels(messages chan string) {

// }
