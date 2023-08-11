package main

import "fmt"

func main(){
	var a=100
	var p=&a
	var pp=&p

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(pp)
	fmt.Println(*pp)
	fmt.Println(**pp)
	*p=2000
	fmt.Println(a)
	fmt.Println(p)


}