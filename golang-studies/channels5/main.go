package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker: ", workerId, "Msg: ", res)
		time.Sleep(time.Second)
	}
}

func main() {

	msg := make(chan int) 
	go worker(1, msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}
}