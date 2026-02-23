package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet("hello", done)
	go greet("Yo", done)
	go slowGreet("Knock knock!", done)
	go greet("Who's der?", done)
	for range done {
	}
}

func greet(phrase string, doneChannel chan bool) {
	fmt.Println(phrase)
	doneChannel <- true
}

func slowGreet(phrase string, doneChannel chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println(phrase)
	doneChannel <- true
	close(doneChannel)
}
