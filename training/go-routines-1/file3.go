package main

import (
	"fmt"
	"time"
)

func main() {

	//readOnlyChannel := make(<- chan int)
	//writeOnlyChannel := make(chan <- int)
	c := make(chan string)
	go printMSg(c)
	go writeFromChannel(c)
	time.Sleep(time.Millisecond)
}

func printMSg(roc <-chan string) {
	//roc <- "Hey"
	fmt.Println("Hello", <-roc)
}

func writeFromChannel(woc chan<- string) {
	woc <- "Hey!!!"
}
