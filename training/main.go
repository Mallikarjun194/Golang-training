package main

import (
	"fmt"
	"reflect"
)

var abc string

var (
	Name string
	Age  int
)

const ab = 24

func main() {

	var a = "3"
	fmt.Println(a, reflect.TypeOf(a))

	b := "abc"
	abc = "2"
	fmt.Println(b, abc, ab)

	var x, y = 12, 13
	x, y = y, x
	fmt.Println(x, y)

	age := 13
	if age > 15 {
		fmt.Println("age is", age)
	} else if age == 30 {
		fmt.Println("age is", age)
	} else {
		fmt.Println("age is", age)
	}

	c := 2
	switch c {
	case 2, 3, 5, 7, 11, 13, 17:
		fmt.Println("case 2")
	case 4, 6, 8, 10:
		fmt.Println()
	default:
		fmt.Println()
	}

	arr := [6]int{1, 2, 3, 4, 5, 6} // Mm : []
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	//slice := []int{}
	//var sli []int
	var sli = make([]int, 0)
	fmt.Println(sli)
	abcd := []int{2, 3}
	sli = append(sli, 1, 2)
	fmt.Println("len is", len(sli))
	fmt.Println("cap is", cap(sli))
	sli = append(sli, 7)
	fmt.Println("len is", len(sli))
	fmt.Println("cap is", cap(sli))
	sli = append(sli, abcd...)
	fmt.Println(sli)
	fmt.Println("len is", len(sli))
	fmt.Println("cap is", cap(sli))

}
