package main

import "fmt"

func main() {
	c := make(chan bool)
	go printValue(c)
	c <- true
	fmt.Println("main true")
	c <- false
	fmt.Println("main false")
}

func printValue(c chan bool) {
	fmt.Println(<-c)
	fmt.Println("print true")
	fmt.Println(<-c)
	fmt.Println("print false")

}
