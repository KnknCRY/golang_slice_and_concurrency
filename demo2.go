package main

import (
	"fmt"
	"time"
)

func concurrencyDemo2() {
	ch := make(chan int)

	go func(channel chan<- int) {
		// for i:=0;i<10000;i++ {
		// 	channel <- 10
		// }
		channel <- 10
		time.Sleep(2 * time.Second)
		channel <- 20
		time.Sleep(2 * time.Second)
		channel <- 30
		defer close(ch)
	}(ch)

	for i := range ch {
		fmt.Println(i)
	}

}
