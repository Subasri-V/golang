package main

import (
	"fmt"
	"time"
)


func hello(){
	fmt.Println("Hello world goroutine")
}
func main(){
	go hello()
	time.Sleep(10*time.Second)
	fmt.Println("Main function")
}