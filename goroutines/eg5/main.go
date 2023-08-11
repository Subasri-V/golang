package main

import "fmt"

func myfunc(channel chan string){
	for i:=1;i<=10;i++{
		channel<-"Go language is awesome"
	}
	close(channel)
}



func main(){
	c:=make(chan string,8)// 8 is a buffer size
	go myfunc(c)

	counter:=0
	for {
		res,ok:=<-c
		counter++
		if !ok{
			fmt.Println("channel is closed ",ok)
			break
		}
		fmt.Println(counter)
		fmt.Println("Channel is open and data is ",res,ok)
	}
}