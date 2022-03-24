package main

import "fmt"

func chDemo1() {
	ch := make(chan int, 1)

	ch <- 1
	// for {
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	default:
		fmt.Println("no")
	}
	// }
}
