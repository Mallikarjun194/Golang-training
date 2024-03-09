package main

import (
	"fmt"
)

func main() {
	var num = 25

	fmt.Println("Num is", num)
	fmt.Println("Address of num", &num)

	var ptr *int
	ptr = &num
	fmt.Println("num is", num)
	fmt.Println("Pointer value is", *ptr)
	fmt.Println("Address of num", ptr)
	fmt.Println("Address of num", &num)

	// modify the pointer

	*ptr = 250
	fmt.Println("Num is", num)
	fmt.Println("Pointer value", *ptr)

	// double pointer

	var doublePointer **int
	doublePointer = &ptr

	fmt.Println("double pointer value", **doublePointer)
	fmt.Println("ptr value is", *ptr)

	**doublePointer = 2500
	fmt.Println("double pointer value", **doublePointer)
	fmt.Println("ptr value is", *ptr)
	fmt.Println("num is", num)

	fmt.Println("address of pointer", &ptr)
	fmt.Println("value of double pointer", doublePointer)

	// call by value & call by ref
	var num1 = 45
	multiply(num1)
	fmt.Println("num inside main function")
	fmt.Println("num1 is", num1, &num1)

	multiplyRef(&num1)
	fmt.Println("num inside main function")
	fmt.Println("num1 is", num1, &num1)

	var x int
	x = 5
	var newPointer *int
	fmt.Printf("Pointer is %T and value is %v\n", newPointer, newPointer)

	newPointer = &x

	m := map[string]int{"zayn": 25, "harry": 56, "abc": 45} // Mm = &hMap
	fmt.Println("map is", m)

	myFun(m) // Mm = &hMap
	fmt.Println("map m is", m)

	myFun1(&m)
	fmt.Println("map m is after using pointer", m)

	// slice
	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1, 6)
	fmt.Println("slice is", s1, len(s1), cap(s1))
	myFunc2(s1)
	fmt.Println("slice is", s1)

	multiplySlice(s1)
	fmt.Println("slice after multiplying", s1)

	//myFunc3(&s1)
	//fmt.Println("slice after using pointer", s1)

	var i int
	Func(i)

}

func Func(n int) {
	for i := 0; i < 1000; i++ {
		go func() {
			Add(i)
		}()
	}
}

func Add(i int) {

}

func multiplySlice(s1 []int) {
	for key, value := range s1 {
		s1[key] = value * value
	}
}

func myFunc3(s *[]int) {
	*s = append(*s, 7)

}

func myFunc2(s1 []int) {
	s1 = append(s1, 7)
}

func myFun1(m1 *map[string]int) {
	(*m1)["zayn"] = 28
	(*m1)["xyz"] = 89
	*m1 = nil
}

func myFun(m1 map[string]int) {
	m1["zayn"] = 27 // Mm1 = &hMap
	m1["xyz"] = 78  // Mm1 = &hMap
	m1 = nil        // Mm1 = 0x00

}

func multiplyRef(num *int) {
	*num *= 2
	fmt.Println("Inside multiplyref fucntion")
	fmt.Println("num is", *num)
	fmt.Println("address of num", num)
}

func multiply(num int) {
	num *= 2
	fmt.Println("Inside multiply func")
	fmt.Println("num is", num)
	fmt.Println("address of num", &num)
	//num = num*2
}
