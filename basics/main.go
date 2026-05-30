package main

import "fmt"

func main() {
	fmt.Println("Hello Go")

	// Data Assignment
	// With var or :=
	normalArray := [2]int{1, 2}
	array(normalArray)
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
