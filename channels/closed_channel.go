package main

import "fmt"

func ClosedChannelClose() {
	ch := make(chan int)
	close(ch)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	close(ch) 
}

func ClosedChannelSend() {
	ch := make(chan int)
	close(ch)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	ch <- 42 
}

func ClosedChannelReceive() {
	ch := make(chan int, 1)
	ch <- 100
	close(ch)

	val, ok := <-ch
	fmt.Printf("1-е чтение: %v, ok=%v\n", val, ok) 
	val, ok = <-ch
	fmt.Printf("2-е чтение: %v, ok=%v\n", val, ok) 
}

func main() {
	ClosedChannelReceive()
}
