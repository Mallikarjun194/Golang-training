package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main routine called!!!")
	ch := make(chan struct{}) // Unbuffered channel
	// 1. Read is blocked, until some other go-routine is going to write data to it
	// 2. Write is blocked, until some other go-routine is going to read data to it
	//ch <- true // write data to channel
	//<- ch // read data from channel
	go firstRoutine(ch)
	//ch <- true
	ch <- struct {
	}{}
	fmt.Println("End of main")
	//fmt.Println(<-ch) // read block
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//go secondRoutine()
	//time.Sleep(100 * time.Millisecond)
	//fmt.Println("end!!")
}

func secondRoutine() {
	for i := 0; i < 3; i++ {

		fmt.Println("second Routine called!!!", i)
		time.Sleep(5 * time.Millisecond)
	}
} // 35ms

func firstRoutine(ch chan struct{}) {
	//for i := 0; i < 4; i++ {
	//
	//	fmt.Println("firstRoutine called!!!", i)
	//	time.Sleep(5 * time.Millisecond)
	//}
	<-ch
	fmt.Println("firstRoutine called!!!")
	fmt.Println("firstRoutine called!!!")
	fmt.Println("firstRoutine called!!!")
	fmt.Println("firstRoutine called!!!")
	fmt.Println("firstRoutine called!!!")
	//ch <- struct{}{}
	//<-ch
	//c <- false

} // 25ms
