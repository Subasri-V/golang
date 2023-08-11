package main

import (
	"fmt"
)

// func main(){
// 	x:=sum(10,20) //datatype of x is func
// 	fmt.Println(x())
// }

func main(){
	x:=sum(10,20)() // datatype of x is int //chain of functions
	fmt.Println(x)
}

func sum(a,b int) func() int{
	c:=100
	return func() int{
		return a+b+c
	}

}