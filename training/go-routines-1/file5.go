package main

import (
	"fmt"
	"sync"
)

func main() {

	//var wg sync.WaitGroup
	//// c1 - > 123 -> 45
	//// c2 - > 123 -> 46
	//// counter : 0
	//for i := 0; i < 4; i++ {
	//	wg.Add(2) // counter 1, 2, 3, 4
	//	go firstRoutine(i, &wg)
	//	go secondRouinte(i, &wg)
	//}
	//wg.Wait()
	//fmt.Println("DONE!!!!")
	guessWhat()
}

func secondRouinte(i int, s *sync.WaitGroup) {
	fmt.Printf("servcie %v executeed\n", i+10)
	s.Done()
}

func firstRoutine(i int, s *sync.WaitGroup) {
	fmt.Printf("servcie %v executeed\n", i)
	s.Done() // 3, 2, 1, 0
}

func guessWhat() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("item i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done!!!")
}
