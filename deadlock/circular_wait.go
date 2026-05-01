package main

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		<-ch1
		ch2 <- 42
	}()

	<-ch2
	ch1 <- 100
}
