package main

import (
	"fmt"
	"time"
)

type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	// 工作者：幫忙消化掉in這個chan裡面的工作
	for w := range in {
		w.z = w.x * w.y
		fmt.Println("do work:", w)
		time.Sleep(time.Duration(w.z) * time.Second)
		out <- w
	}
}

func sendLotsOfWork(in chan *Work) {
	// 派發者：派發工作進入in這個chan
	for i := 0; i < 10; i++ {
		fmt.Println("work is sent:", i)
		in <- &Work{x: 1, y: i}
	}
}

func getLotsOfWork(out chan *Work) {
	// 結案者：做完的會放在out裡面，結案者把它拿出來
	for result := range out {
		// <-out
		fmt.Println("work was done:", result)
	}
	defer close(out)
}

func concurrencyDemo1() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < 5; i++ {
		go worker(in, out)
		fmt.Println("work")
	}
	go sendLotsOfWork(in)
	getLotsOfWork(out)
}
