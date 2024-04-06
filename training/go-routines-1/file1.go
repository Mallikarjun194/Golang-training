package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://google.com",
	"https://flipkart.com",
	"https://linkedin.com",
	"https://pkg.go.dev",
}

func main() {
	//startTime := time.Now()
	//
	//for _, url := range urls {
	//	hitUrl(url)
	//}
	//fmt.Println("Time taken ", time.Now().Sub(startTime))
	//withRoutine()
	//funcSquareOfNum()
	twoGoroutines()
}

func hitUrl(url string) {
	if result, err := http.Get(url); err != nil {
		fmt.Println("err", url, err)
	} else {
		fmt.Println(result.StatusCode, url)
	}
}

// Time taken  12.86110775s

func withRoutine() {
	c := make(chan bool)
	startTime := time.Now()

	for _, url := range urls {
		go hitUrlc(url, c)
	}
	for i := 0; i < len(urls); i++ {
		<-c
	}
	fmt.Println("Time taken ", time.Now().Sub(startTime))

}

func hitUrlc(url string, c chan bool) {
	if result, err := http.Get(url); err != nil {
		fmt.Println("err", url, err)
	} else {
		fmt.Println(result.StatusCode, url)
	}
	c <- true
}

func funcSquareOfNum() {
	c := make(chan int)
	go square(c)

	for {
		value, ok := <-c
		fmt.Println(value, ok)
		if !ok {
			break
		}
	}
	fmt.Println("DONE!!!!")
}

func square(c chan int) {
	for i := 0; i < 3; i++ {
		c <- i * i
	}
	close(c)
}

func twoGoroutines() {
	fmt.Println("Main routine!!!")
	c1, c2 := make(chan int), make(chan int)

	go Routine1(c1)
	go Routine2(c2)

	for val := range c1 {
		fmt.Println(val)
	}

	for val := range c2 {
		fmt.Println(val)
	}

	//for {
	//	v1, ok1 := <-c1
	//	fmt.Println(v1, ok1)
	//	v2, ok2 := <-c2
	//	fmt.Println(v2, ok2)
	//
	//	if !ok1 && !ok2 {
	//		break
	//	}
	//}
}

func Routine2(c2 chan int) {
	for i := 0; i < 3; i++ {
		fmt.Println("second routine", i)
		c2 <- i
	}
	close(c2)
}

func Routine1(c1 chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println("first routine", i)
		c1 <- i
	}
	close(c1)
}
