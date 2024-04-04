package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//n := 10000
	//startTime := time.Now()
	//
	//for i := 0; i < n; i++ {
	//	calculate(i)
	//}
	//fmt.Println("total goroutine:", runtime.NumGoroutine())
	//fmt.Println("time taken: ", time.Now().Sub(startTime))

	withGoroutine()
}

// total goroutine: 1
//time taken:  1.59357525s

func calculate(i int) {
	x := 100
	for j := 0; j < 100*i; j++ {
		x = x / (j + 1)
	}
}

func withGoroutine() {
	n := 100000
	startTime := time.Now()
	c := make(chan bool, n) // 95
	// 50

	for i := 0; i < n; i++ {
		go calculateWithGoRoutine(i, c)
	}
	fmt.Println("total goroutine:", runtime.NumGoroutine())
	for i := 0; i < n; i++ {
		<-c
	}
	fmt.Println("time taken: ", time.Now().Sub(startTime))
}

func calculateWithGoRoutine(i int, c chan bool) {
	x := 100
	for j := 0; j < 100*i; j++ {
		x = x / (j + 1)
	}
	c <- false
}

// productIDs

// inventorydetails
// priceDetails
// image and detilas

//func f1() {
//	uniqueID := fun2()
//	go inventorydetails(uniQueID)
//	go priceDetails()
//
//}
