package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("goroutine: done")
		case <-ctx.Done():
			fmt.Println("goroutine: cancelled, err:", ctx.Err())
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}
