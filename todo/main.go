package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy bread")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v", todos)
}
