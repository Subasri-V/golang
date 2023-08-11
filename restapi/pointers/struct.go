package main

import "fmt"

type Student struct{
	RollNumber int
	Name string
} 

func main(){
	s:=Student{304,"Subasri"}
	ps:=&s
	fmt.Println(ps)

	fmt.Println((*ps).Name)
	fmt.Println(ps.Name)
	ps.RollNumber=403
	fmt.Println(ps)

}