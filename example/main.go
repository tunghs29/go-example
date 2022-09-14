package main

import (
	"fmt"
	"time"
)

func startConsumer(queue chan int, name string) {
	for {
		time.Sleep(time.Second)
		fmt.Println(name, <-queue)
	}
}

func main() {
	n := 1000000
	queue := make(chan int, n)

	for i := 0; i < n; i++ {
		queue <- i
	}

	go startConsumer(queue, "1")
	go startConsumer(queue, "2")

	time.Sleep(time.Second * 10)
}
