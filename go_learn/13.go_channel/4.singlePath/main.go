package main

func main()  {
	stringChan := make(chan<- string, 3)
	intChan := make(<-chan int, 3)
	stringChan <- "hello world"
	<- intChan
	// <- stringChan
	// intChan <- 2
}