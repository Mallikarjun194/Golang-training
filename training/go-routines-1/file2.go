package main

import "fmt"

func main() {
	//n := 4
	//c := make(chan int, n)
	//
	//go printSquare(c, n)
	//for i := 0; i <= n; i++ {
	//	c <- i
	//}
	//fmt.Println("DONE!!!!")
	//lenAndCap()
	closedBufferedChannel()
}

func printSquare(c chan int, n int) {
	for i := 0; i < n; i++ {
		val := <-c
		fmt.Printf("Square of num %v is %v\n", val, val*val)
	}
}

// Len and Cap

func lenAndCap() {
	n := 4
	c := make(chan int, n)
	fmt.Println(len(c), cap(c))
	c <- 2
	c <- 3
	c <- 4
	fmt.Println(len(c), cap(c))
	fmt.Println(<-c)
	fmt.Println(len(c), cap(c))

	//c1 := make(chan int)
	//c1 <- 1
	//<-c1

}

func closedBufferedChannel() {
	n := 4
	c := make(chan int, n)
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	//c <- 6
	fmt.Println(<-c)
	// c -> 3, 4, 5
	c <- 6
	// c -> 3, 4, 5, 6
	close(c)
	//close(c)
	//c <- 7
	for val := range c {
		fmt.Println("val", val)
	}
	fmt.Println("DONE!!")

	n1 := 4
	c1 := make(chan int, n1)
	c1 <- 1
	c1 <- 2
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	//	fmt.Println(<-c1)

}
