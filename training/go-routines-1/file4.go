package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go firstService(c1)
	go secondServcie(c2)
	for {
		select {
		case val := <-c1:
			fmt.Println("servcie 1: ", val)

		case val := <-c2:
			fmt.Println("servcie 2: ", val)

		case <-time.After(time.Second):
			fmt.Println("No routine")
		}
	}
	fmt.Println("Done!!!")
}

func firstService(c1 chan string) {
	//time.Sleep(2 * time.Millisecond)
	c1 <- "first service called!!"
}

func secondServcie(c1 chan string) {
	//time.Sleep(3 * time.Millisecond)
	c1 <- "second service called!!"
}
