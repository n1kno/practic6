package main

import (
	"fmt"
	"time"
)

func NormalChannelClose() {
	ch := make(chan int)
	close(ch)
	fmt.Println("канал закрыт")
}

func NormalChannelSendBlock() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		<-ch
		fmt.Println("данные получены")
	}()
	fmt.Println("отправка (блокировка на 2 сек)...")
	ch <- 42
	fmt.Println("отправка завершена")
}

func NormalChannelReceiveBlock() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 100
	}()
	fmt.Println("чтение (блокировка на 2 сек)...")
	val := <-ch
	fmt.Println("получено:", val)
}

func main() {
	NormalChannelSendBlock()
}
