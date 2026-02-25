package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	fmt.Println(os.Getenv("BLAH"))
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
