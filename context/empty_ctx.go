package main

import (
	"context"
	"fmt"
)

func main() {
	bg := context.Background()
	todo := context.TODO()

	fmt.Printf("Background: %T, Done() = %v\n", bg, bg.Done())
	fmt.Printf("TODO: %T, Done() = %v\n", todo, todo.Done())
}
