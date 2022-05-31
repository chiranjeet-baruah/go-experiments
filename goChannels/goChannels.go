package goChannels

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"time"
)

type Order struct {
	item        string
	customer_id int64
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func waiter(order_channel *chan Order, online_waiters *int, quitShop *chan bool) {
	waiter_id := time.Now().Unix()
	fmt.Println("Waiter Online: ", waiter_id)
	*online_waiters++
	for {
		select {
		case <-*quitShop:
			return
		case order := <-*order_channel:
			fmt.Println("Received order : ", order.item, " from customer : ", order.customer_id, " by waiter : ", waiter_id, " STARTED WORKING")
			time.Sleep(time.Duration(rand.Intn(15-10)+10) * time.Second)
			fmt.Println("Received order : ", order.item, " from customer : ", order.customer_id, " by waiter : ", waiter_id, " FINISHED WORKING")
		case <-time.After(time.Duration(rand.Intn(10-1)+1) * time.Second):
			fmt.Println("Waiter Offline: ", waiter_id)
			*online_waiters--
			return
		}
	}
}

func shopOpen(order_channel *chan Order, quitShop *chan bool) {
	for {
		select {
		case <-*quitShop:
			fmt.Println("Shop is closing soon...")
			return
		default:
			item_list := []string{"banana", "apple", "oranges", "grapes", "coconut", "watermelon"}
			for {
				customer_id := time.Now().Unix()
				order := Order{item_list[rand.Intn(len(item_list))], customer_id}
				time.Sleep(time.Duration(rand.Intn(3-1)+1) * time.Second)
				*order_channel <- order
				fmt.Println("Sent order : ", order.item, " by customer : ", order.customer_id)
			}
		}
	}
}

func startWaiters(max_waiters int, order_channel *chan Order, online_waiters *int, quitShop *chan bool) {
	for {
		select {
		case <-*quitShop:
			fmt.Println("All waiters are going offline...")
			*online_waiters = 0
			return
		default:
			ticker := time.NewTicker(500 * time.Millisecond)
			for range ticker.C {
				if *online_waiters < max_waiters {
					fmt.Println("Total online waiters: ", *online_waiters)
					go waiter(order_channel, online_waiters, quitShop)
				}
			}
		}
	}
}

func RunGoChannelsExample() {
	quitShop := make(chan bool)
	order_channel := make(chan Order, 10)

	max_waiters := MaxParallelism()
	online_waiters := 0

	go shopOpen(&order_channel, &quitShop)
	go startWaiters(max_waiters, &order_channel, &online_waiters, &quitShop)

	// Wait for interrupt signal to gracefully shutdown with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	quitShop <- true
	fmt.Println("Closing shop...")
	time.Sleep(5 * time.Second)
	fmt.Println("Shop closed")
}
