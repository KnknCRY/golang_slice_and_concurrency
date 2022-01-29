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
		time.Sleep(2 * time.Second)
		channel <- 10
		time.Sleep(2 * time.Second)
		channel <- 20
		time.Sleep(2 * time.Second)
		channel <- 30
		defer close(ch)
	}(ch)

	// 這個for會被卡住，等ch有數值才會動
	// go很聰明，當ch將不再有數值被丟入，就會離開這個for
	for i := range ch {
		fmt.Println(i)
	}
}
