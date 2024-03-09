package main

import (
	"fmt"
)

func main() {

	// A = {1, 2, 3, 4, 5, 6, 7, 8}
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8} //MSliceA: [ Ptr to A, 8, 8]

	fmt.Println(sliceA)
	fmt.Println("Len of sliceA", len(sliceA))
	fmt.Println("Cap of sliceA", cap(sliceA))

	fmt.Printf("addrs of ptr to A %p\n", sliceA)
	fmt.Println(&sliceA[0])

	sliceB := sliceA[:5] // [Ptr to A, 5, 8]
	fmt.Println(sliceB)
	fmt.Println("Len of sliceB", len(sliceB))
	fmt.Println("Cap of sliceB", cap(sliceB))

	sliceB = append(sliceB, 9, 10, 11)
	fmt.Println("----", sliceB) // [Ptr to A, 8, 8]
	fmt.Println("Len of sliceB", len(sliceB))
	fmt.Println("Cap of sliceB", cap(sliceB))
	fmt.Printf("addrs of ptr to A %p\n", sliceB)

	// B := [1 2 3 4 5 9 10 11 12     ]
	sliceB = append(sliceB, 12)
	fmt.Println("----", sliceB) // [Ptr to B, 9, 16]
	fmt.Println("Len of sliceB", len(sliceB))
	fmt.Println("Cap of sliceB", cap(sliceB))
	fmt.Printf("addrs of ptr to B %p\n", sliceB)

	// len, cap, make, copy, append
	sA := make([]int, 2)
	sA[0] = 1
	sA[1] = 2
	fmt.Println(sA)

	sB := make([]int, 5)

	copiedItemCount := copy(sB, sA)
	fmt.Println("No of items copied", copiedItemCount)
	fmt.Println(sB)

	sliceAppend := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceAppend, len(sliceAppend), cap(sliceAppend))

	sliceAppend = append(sliceAppend, 6)
	fmt.Println("slice with append", sliceAppend)

	//FunAppend(sliceAppend)
	//fmt.Println("slice with append after calling a Fun", sliceAppend)
	multiply(sliceAppend)
	fmt.Println("Slice items after multiply", sliceAppend)
	sliceAppend = append(sliceAppend, 100)
	sliceAppend = FunAppend(sliceAppend)
	fmt.Println("slice with append after calling a Fun", sliceAppend)
	fmt.Printf("addrees withing main %p\n", &sliceAppend)

	slice, status := multipleReturnItems(sliceAppend)
	fmt.Println(slice, status)

	slice1 := []int{2, 1, 20, 30, 4, 2, 1, 6}
	//sort.Ints(slice1)
	//fmt.Println("sorted items are", slice1)

	slice1 = deleteItemFromSlice(slice1, 2)
	fmt.Println(slice1)
	slice1 = deleteItemFromSlice(slice1, 2)
	fmt.Println(slice1)

}

func deleteItemFromSlice(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func FunAppend(slice []int) []int {
	slice = append(slice, 100, 120)
	fmt.Printf("addrees withing Fun %p\n", &slice)
	return slice
}

func multiply(sliceAppend []int) {
	for key, value := range sliceAppend {
		sliceAppend[key] = value * 2
	}
}

func multipleReturnItems(slice []int) ([]int, bool) {
	if len(slice) == 0 {
		return []int{}, false
	}
	slice = append(slice, 122, 134)
	return slice, true
}
