package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("операция успешна")
	case <-ctx.Done():
		fmt.Println("таймаут:", ctx.Err())
	}
}
