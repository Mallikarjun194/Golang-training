package main

import "fmt"

type Printer interface {
	Prints()
}

type Persons struct {
	Name string
	Age  int
}

//func (p Persons) Add() {
//	//TODO implement me
//	panic("implement me")
//}

func (p Persons) Prints() {
	fmt.Println("Person Name is", p.Name)
	fmt.Println("Person Age is", p.Age)
}

func main() {
	person := Persons{
		Name: "Zayn",
		Age:  25,
	}
	PersonInfo(person)
}

func PersonInfo(Iobj Printer) {
	Iobj.Prints()
}
