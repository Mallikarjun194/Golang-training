package main

import "fmt"

func main() {

	sliceA := []string{"Mango", "Jack", "Apple", "orange", "Banana", "pineapple"}
	fmt.Println(sliceA)

	for i := 0; i < len(sliceA); i++ {
		fmt.Println(sliceA[i])
	}
	//
	//for index, fruit := range sliceA {
	//	fmt.Println(index, fruit)
	//}

	for i := range sliceA {
		fmt.Println(sliceA[i])
	}

	for i := 0; i < 10; i++ {
		if i == 6 {
			goto letsGo
		}
		fmt.Println("num is", i)
	}
letsGo:
	fmt.Println("Welcome back to go pgm!!!")

}
