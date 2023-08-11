package main

import "fmt"

//import "fmt"


type Student struct{
	name string
	rollNumber int
} 


func main(){
	var s Student
	s=Student{"suba",304}
	fmt.Println(s)

	s1:=Student{"gayatri",150}
	fmt.Println(s1.name,s1.rollNumber)

	ss1:=&s1
	fmt.Println(ss1)

	var s2 Student
	fmt.Println(s2)
	
	s2=Student{name: "abc"}
	fmt.Println(s2)


}