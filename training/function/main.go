package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Welcome to functions in golang!!")

	sayHello()
	var num1, num2 int
	num1 = 10
	num2 = 20
	fmt.Println(doSomething(num1, num2))

	fmt.Println(addNumbers(1, 2, 3, 4))

	fmt.Println(addAndSub(num1, num2))

	div, err := divide(4, 2)
	if err != nil {
		if err.Error() == "divide by zero error" {
			fmt.Println(err.Error(), http.StatusBadRequest)
		} else {
			fmt.Println(err.Error(), http.StatusInternalServerError)
		}
		return
	}
	fmt.Println(div)
}

func divide(i int, i2 int) (float32, error) {
	if i > 4 {
		return 0.0, errors.New("i more than 4")
	}
	if i2 == 0 {
		return 0.0, errors.New("divide by zero error")
	}
	return float32(i / i2), nil
}

func addAndSub(num1, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func addNumbers(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func doSomething(num1 int, num2 int) int {
	return num1 + num2
}

func sayHello() {
	fmt.Println("Hello, How are you??")
}
