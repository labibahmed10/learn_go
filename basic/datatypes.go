package main

import "fmt"


func dataTypes() {
	// Integer
	var a int = 10
	fmt.Println("Integer:", a)

	// Float
	var b float64 = 20.5
	fmt.Println("Float:", b)

	// String
	var c string = "Hello, Go!"
	fmt.Println("String:", c)

	// Boolean
	var d bool = true
	fmt.Println("Boolean:", d)

	// Array
	var e [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", e)

	// Slice
	var f []int = []int{4, 5, 6}
	fmt.Println("Slice:", f)

	// Map
	var g map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Println("Map:", g)

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	var h Person = Person{Name: "Alice", Age: 30}
	fmt.Println("Struct:", h)
}

func main() {
    dataTypes()
}