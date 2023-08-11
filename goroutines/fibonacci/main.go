package main

import (
	"fmt"
	"time"
)

func fibinacci()
func main(){
	
	start:=time.Now()
	n1:=0
	n2:=1
	fmt.Print(n2," ")
	for i:=0;i<50;i++{
		n3:=n1+n2
		fmt.Print(n3," ")
		n1=n2
		n2=n3
	}
	end:=time.Now().Sub(start)
	fmt.Println()
	fmt.Println(start)
	fmt.Println(end)
}