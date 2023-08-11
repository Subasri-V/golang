package main

import (
	"fmt"
	"time"
)

func main(){
	for i:=0;i<5;i++{
		//time.Sleep(1*time.Second)
		go func (){
			//time.Sleep(1*time.Second)
			fmt.Println(i)
		}()
		time.Sleep(1*time.Second)
	}
	//time.Sleep(1*time.Second)
}