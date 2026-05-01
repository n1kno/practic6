package main

import "fmt"

func NilChannelClose() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	var ch chan int
	close(ch)
}

func NilChannelSend() {
	var ch chan int
	ch <- 42 // deadlock
}

func NilChannelReceive() {
	var ch chan int
	<-ch 
}

func main() {

	fmt.Println("раскомментируйте нужную функцию для демонстрации")
}
