package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<-  string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println(<- c)
		time.Sleep(time.Second * 1)
	}
}

func MySleep(sec int) {
	select {
		case <- time.After(time.Second * time.Duration(sec)):
			fmt.Printf("%d seconds passed.\n", sec)
	}
}

func main() {
	//var c chan string = make(chan string)
	//
	//go pinger(c)
	//go ponger(c)
	//go printer(c)

	var input string
	fmt.Scanln(&input)
	MySleep(10)
}
