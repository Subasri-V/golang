package main

import "fmt"

func main(){
	a:=[7]int{1,2,3,4,5,6,7}
	//slices of reference datatypes
	slice1:=a[1:]
	slice2:=a[3:]
	
	fmt.Println("a = ",a)
	fmt.Println("slice1 = ",slice1)
	fmt.Println("slice2 = ",slice2)

	slice1[0]=20
	slice1[1]=30
	slice1[5]=70

	slice2[2]=60
	fmt.Println("a = ",a)
	fmt.Println("slice1 = ",slice1)
	fmt.Println("slice2 = ",slice2)
}