package main

import "fmt"

type Read1 interface {
	Read1(p []byte) (n int, err error)
}

type Write1 interface {
	Write1(p []byte) (n int, err error)
}

type Read1Write1 interface {
	Read1
	Write1
}

type readAndWrite interface {
	Read1Write1
	Abc()
}

func PrintSomething(data interface{}) {
	fmt.Println("data is", data)
}

func main() {

	PrintSomething(2)
	PrintSomething("Hello")
	PrintSomething(false)
	a := abc{}

	PrintSomething(a)

}

type abc struct {
	Name string
	Age  int
}

//{
//slice : []
//map : {}
//struct
//string
//nil
//}
