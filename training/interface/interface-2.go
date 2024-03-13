package main

import (
	"fmt"
	"math"
)

type calculations interface {
	area() float32
	perim() float32
}

type rect struct {
	Width  float32
	Length float32
}

type Circle struct {
	radius float32
}

func (r rect) area() float32 {
	return r.Width * r.Length
}

func (r Circle) area() float32 {
	return math.Pi * (r.radius * r.radius)
}

func (r rect) perim() float32 {
	return 2 * (r.Width + r.Length)
}

func (r Circle) perim() float32 {
	return 2 * math.Pi * r.radius
}

func doSomething(IObj calculations) {
	fmt.Println("Geeometry is", IObj)
	fmt.Println("Area is", IObj.area())
	fmt.Println("Perim is", IObj.perim())
}

func main() {
	rect := rect{
		Width:  5.0,
		Length: 4.0,
	}
	doSomething(rect)

	printScan := PrintScanStruct{}
	printScanFun(printScan)

	// Gin, gorilla mux, chi, fiber,
}

func printScanFun(ps PrintScan) {
	ps.Scanner1()
	ps.Printers1()
}

type Printers interface {
	Printers1()
}

type Scan interface {
	Scanner1()
}

type PrintScan interface {
	Printers
	Scan
}

type PrintScanStruct struct {
}

func (p PrintScanStruct) Printers1() {
	fmt.Println("Hey I am a printer")
}

func (p PrintScanStruct) Scanner1() {
	fmt.Println("Hey I am a scanner")
}
