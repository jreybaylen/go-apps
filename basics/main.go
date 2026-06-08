package main

import (
	"fmt"

	"github.com/jreybaylen/go-apps/basics/generics"
	go_routine "github.com/jreybaylen/go-apps/basics/go-routine"
	"github.com/jreybaylen/go-apps/basics/http"
	"github.com/jreybaylen/go-apps/basics/pointer"
	"github.com/jreybaylen/go-apps/basics/structs"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	person := Person{
		Name: "Jrey",
		Age:  31,
	}
	person.Greet()

	// Data Assignment
	// With var or :=
	normalArray := [2]int{1, 2}
	array(normalArray)

	slice()
	generics.GenericsMain()
	structs.StructsMain()
	pointer.PointerMain()
	go_routine.GoRoutineMain()
	http.HttpMain()
}

func array(normalArray [2]int) {
	fmt.Println("\n-------- Array -------")
	for idx, value := range normalArray {
		fmt.Printf("Index: %v, Value: %v \n", idx, value)
	}

	fmt.Println("\n-------- Multi-Dimension Array -------")
	multiDimensionArray := [2][2]int{{3, 4}, {5, 6}}
	for _, array := range multiDimensionArray {
		for idx, value := range array {
			fmt.Printf("Index: %v, Value: %v \n", idx, value)
		}
		fmt.Println("------------------")
	}

	fmt.Println("\n-------- Dynamic Range Multi-Dimension Array -------")
	dynamicMultiDimensionArray := [...]([2]int){{7, 8}, {9, 10}, {11, 12}}
	for _, array := range dynamicMultiDimensionArray {
		for idx, value := range array {
			fmt.Printf("Index: %v, Value: %v \n", idx, value)
		}
		fmt.Println("------------------")
	}
}

func slice() {
	fmt.Println("\n-------- Slice -------")

	array := [5]int{1, 2, 3, 4, 5}
	sl := array[:3]
	fmt.Println(sl, len(sl), cap(sl))

	sl = sl[:2]
	fmt.Println(sl, len(sl), cap(sl))
}
