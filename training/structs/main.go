package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Age    int
	Gender string
	City   string
}

type Company struct {
	NameOfCompany string
	Emp           Employee
}

//type Abc struct {
//}
//
//type XYZ struct {
//	Cmp Company
//	Abc Abc
//}
//
//type Comp struct {
//	cmp []Car
//}
//
//type Car struct {
//	Model      string
//	Properties Properties
//}
//
//type Properties struct {
//}

func main() {
	fmt.Println("Struct in golang!!")

	emp := Employee{
		"arjun", 25, "male", "Dvg",
	}
	fmt.Println(emp)

	emp1 := Employee{
		Name:   "arjun",
		Gender: "male",
		City:   "dvg",
		Age:    25,
	}

	fmt.Printf("Emp details are %+v\n", emp1)
	fmt.Printf("Emp name %v and his age %v\n", emp1.Name, emp1.Age)

	emp1.Age = 27
	fmt.Printf("Emp details are %+v\n", emp1)

	emp2 := Employee{Name: "Zayn", Gender: "male"}
	fmt.Printf("Emp details are %+v\n", emp2)

	emp4 := Employee{Name: "Harry"}
	emp4.Age = 24
	emp4 = changeEmp(emp4)
	fmt.Printf("Emp details are using pointer %+v\n", emp4)

	cmp := Company{
		NameOfCompany: "google",
		Emp: Employee{
			Name: "harry",
		},
	}
	fmt.Println(cmp)
	fmt.Println(cmp.NameOfCompany)
	fmt.Println(cmp.Emp.Name)

	cmp1 := Comp{
		Name: "Google",
		Employee: Employee{
			Name: "Zayn",
			Age:  25,
		},
	}
	fmt.Println(cmp1)
	fmt.Println(cmp1.Name)
	fmt.Println(cmp1.Employee.Name)

	s := newServerConfig(UpdateTLS)
	fmt.Println(s)

	s = newServerConfig(UpdateTLS, updateMaxCons(1000))
	fmt.Println(s)
}

type Comp struct {
	Name     string
	Employee // struct embedding
}

func changeEmp(emp Employee) Employee {
	emp.Age = 34
	emp.Gender = "male"
	return emp
}

type serverConfig struct {
	Opts
}

type Opts struct {
	maxCons int
	tls     bool
	name    string
	id      string
}

func defaultOptions() Opts {
	return Opts{
		maxCons: 15,
		id:      "default-12345",
		tls:     false,
		name:    "default-name",
	}
}

func UpdateTLS(opts *Opts) {
	opts.tls = true
}

func updateMaxCons(noOfConn int) OptFunc {
	return func(opts *Opts) {
		opts.maxCons = noOfConn
	}
}

func newServerConfig(opts ...OptFunc) *serverConfig {
	DefOpts := defaultOptions()
	for _, function := range opts {
		function(&DefOpts)
	}
	return &serverConfig{
		Opts: DefOpts,
	}
}

type OptFunc func(*Opts)
