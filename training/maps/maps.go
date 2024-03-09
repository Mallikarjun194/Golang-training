package main

import "fmt"

func main() {
	fmt.Println("Maps!!!")

	m1 := make(map[string]int)
	// {"Zayn": 25, "Harry": 35}
	fmt.Println(m1)

	m1["Zayn"] = 1
	fmt.Println(m1)

	m2 := map[string]int{} // map literal
	m2["Harry"] = 23
	fmt.Println(m2)

	var m3 map[string]int // allocated but not initialised

	m3 = make(map[string]int)
	// "12345":{"Name": Zayn, Age: 27}
	m3["Louis"] = 34
	fmt.Println("m3 map is", m3)

	var m4 = map[string]string{
		"a": "12",
		"b": "23",
	}
	fmt.Println("m4 map is", m4)

	// Fetch
	fmt.Println("Fetch value of b ", m4["c"])

	// delete

	if _, ok := m4["c"]; ok {
		delete(m4, "c")
		fmt.Println("m4 map after deleting", m4)
	}

	v, _ := m4["a"]
	fmt.Println(v)

	// loop the map
	fmt.Println("Map items are!!!!")
	for key, value := range m4 {
		fmt.Println(key, value)
	}
	fmt.Println("Before clearing a map m4", m4)
	clear(m4)
	fmt.Println("After clearing a map m4", m4)

	// copy
	var m5 = map[string]int{
		"xyz": 13,
		"abc": 45,
	}

	copyM := m5
	copyM["xyz"] = 15
	copyM["tgj"] = 25
	fmt.Println("m5 map is", m5)
	fmt.Println("copyM map is", copyM)
	m5["yu"] = 145
	fmt.Println("m5 map is", m5)
	fmt.Println("copyM map is", copyM)

	var m6 map[string]int // Mm6 = 0x00(nil map)
	fmt.Println(m6)

	var m7 = make(map[string]int) // Mm7 = &hMap
	m7["a"] = 1
	m7["b"] = 2
	fmt.Println(m7)

	m7 = AlterFun(m7) // Mm7 = &hMap
	fmt.Println("Map after alter", m7)

	var m8 = map[string]int{}
	m8 = handleError(m8)
	if m8 == nil {

	}
	m8["x"] = 6
	fmt.Println(m8)

}

func AlterFun(map7 map[string]int) map[string]int {
	map7["a"] = 7 // map7m = &hMap
	map7["c"] = 8 // map7m = &hMap
	map7 = nil    // map7m = nil
	return map7
}

func handleError(m map[string]int) map[string]int {
	if m == nil {
		return map[string]int{}
	}
	m["a"] = 1
	m = nil
	return m
}
