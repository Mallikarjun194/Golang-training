package main

import "fmt"

func main() {
	//defer recoverFromPanic()
	var ptr *int
	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("Nil pointer")
	}

	sli := []int{4, 5, 6, 8}
	fetchIndex := 3
	if fetchIndex >= 0 && fetchIndex < len(sli) {
		fmt.Println("array item", sli[fetchIndex])
	} else {
		fmt.Println("Out of range!!!")
	}

	var inter interface{} = "Hello!"
	//number := inter.(int)
	//fmt.Println(number)
	if number, ok := inter.(string); ok {
		fmt.Println(number)
	} else {
		fmt.Println("wrong assertion")
	}
	var input int
	input = 4
	if input >= 1 && input <= 10 {
		// initialised db
		fmt.Println("Pgm is fine")
	} else {
		panic("input is not in the range")
	}

	var abc *struct{ Value string }
	//if abc == nil {
	//	fmt.Println("Nil Pointer")
	//	return
	//}
	fmt.Println(abc.Value)
	fmt.Println("Hey!!!")

}

func recoverFromPanic() {
	if err := recover(); err != nil {
		fmt.Println("Recovered!!!")
	}
	fmt.Println("Heyyyy!!!!!")
}
