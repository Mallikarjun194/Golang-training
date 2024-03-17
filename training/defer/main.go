package main

import (
	"fmt"
	"os"
)

func main() {
	//defer fmt.Println("TechBrick")
	//
	//fmt.Println("Welcome back")
	//fmt.Println("To")

	//stackOrder()
	//deferInLoop()
	//argumentsInDefer1()
	//guessWhat()
	//WrapDefer()
	callMe()
}

func stackOrder() {
	defer fmt.Println("techBrick") // bottom item
	defer fmt.Println("To")
	defer fmt.Println("Welcome back") // top

	fmt.Println("Hey!!!")
}

/*
-------------------------
fmt.Println("Welcome back") -> Top
-------------------------
fmt.Println("To")
-------------------------
fmt.Println("techBrick")
_________________________

*/

func deferInLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("main")
}

func argumentsInDefer() {
	count := 1
	defer fmt.Println("Deferrr", count)
	fmt.Println("Beforeeee ", count)
	count += 1
	fmt.Println("After ", count)
}

func argumentsInDefer1() {
	count := 1
	defer func() {
		fmt.Println("Defer", count)
	}()
	fmt.Println("Beforeeee ", count)
	count += 1
	fmt.Println("After ", count)
}

func guessWhat() {
	element := 0

	defer func(value int) {
		fmt.Println("Deferr", value)
	}(element)

	element += 1
	fmt.Println("element is", element)
}

func WrapDefer() {
	count := 0
	for i := 0; i < 5; i++ {
		count += 1
		defer func(counter int) {
			fmt.Println("Count is", counter)
		}(count)
	}
	fmt.Println("Welcome!!")
}

func InvokeDefer(msg string) {
	defer fmt.Println("Invoking defer", msg)
	fmt.Println("Invoking", msg)
}

func callMe() {
	InvokeDefer("One")
	defer InvokeDefer("Two")
	fmt.Println("Back!!!")
}

func CreateFile() {
	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("error while creating a file!!", err)
		return
	}
	defer file.Close()
	modifyFile(file)

	//file.Close()

}

func modifyFile(file *os.File) {
	// TODO
}
