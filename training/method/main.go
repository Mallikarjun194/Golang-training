package main

import "fmt"

type author struct {
	name   string
	branch string
	salary int
}

type data int

func (a author) Print() {
	fmt.Println("authors name", a.name)
}

func (d1 data) multi(d2 data) data {
	return d1 * d2
}

//func (d1 data) multiply(d2 data) data {
//	return d1 * d2
//}

//func (p *Type) method_name(...Type)Type{
//	//
//}

type authorPointer struct {
	name   string
	branch string
	salary int
}

// Update Print Method with pointer receiver
func (obj *authorPointer) Update(branchName string) {
	(*obj).branch = branchName
}

func (obj *authorPointer) Print() {
	fmt.Println("Branch name ", (*obj).branch)
}

func main() {
	res := author{"Abc", "Meh", 2500}
	res.Print()

	n1 := data(2)
	n2 := data(5)
	res1 := n2.multi(n1)
	fmt.Println(res1)

	ptr := authorPointer{name: "Harry", branch: "EC"}
	ptr.Print()

	p := &ptr
	p.Update("CS")
	p.Print()
	ptr.Print()

}
