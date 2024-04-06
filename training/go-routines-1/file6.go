package main

import "fmt"

func main() {
	chanTask := make(chan int, 6)
	chanResult := make(chan int, 6)

	workers := 3
	for i := 1; i <= workers; i++ {
		go squareOfNum(chanTask, chanResult, i)
	}
	for i := 1; i <= 6; i++ {
		chanTask <- i
	}
	for i := 1; i <= 6; i++ {
		result := <-chanResult
		fmt.Println("Main result ", i, result)
	}

}

func squareOfNum(task <-chan int, result chan<- int, workerID int) {
	for val := range task {
		result <- val * val
		fmt.Printf("worked %v did a square of %v is %v\n", workerID, val, val*val)
	}
}

// go-routine, channels, types of a channel, deadlock, waitgorup, select, workerpool,
