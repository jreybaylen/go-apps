package todoapp

import (
	"fmt"
)

func main() {
	todoList := []string{"Buy groceries", "Clean the house", "Finish the report"}

	displayTodoList(todoList)
	updateTodoList(&todoList, "Go for a run")
	fmt.Println("--Updated--")
	displayTodoList(todoList)
}

func displayTodoList(todoList []string) {
	fmt.Println("Todo List:")

	for i, item := range todoList {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func updateTodoList(todoList *[]string, newItem string) {
	*todoList = append(*todoList, newItem)
}
