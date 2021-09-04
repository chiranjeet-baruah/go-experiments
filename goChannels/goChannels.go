package goChannels

import (
	"fmt"
	"time"
)

// RunWaitGroupExample runs wait group example
func RunGoChannelsExample() {
	messages := make(chan string, 50) //buffered channel

	go doSomethingGoChannels(messages) // sends messages in channel

	for message := range messages {
		fmt.Println("Received: ", message)
		time.Sleep(3 * time.Second)
	}
}

func doSomethingGoChannels(messages chan string) {
	message_list := []string{"banana", "apple", "oranges", "grapes", "coconut", "watermelon"}
	for i, mesg := range message_list {
		fmt.Println("Sending message: ", i+1, " with text : ", mesg)
		messages <- mesg
		time.Sleep(2 * time.Second)
	}
	close(messages) // closing channel after sending all messages
}
