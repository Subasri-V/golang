package main

import "fmt"

func main(){
	slice1:=[]string{"c","c++","java"}
	slice2:=append(slice1,"Python","ruby","go")

	fmt.Printf("slice1=%v,len=%d,cap=%d\n",slice1,len(slice1),cap(slice1))
	fmt.Printf("slice2=%v,len=%d,cap=%d\n",slice2,len(slice2),cap(slice2))

	slice2[4]="xyz"
	fmt.Println("\nslice1 = ",slice1)
	fmt.Println("\nslice2 = ",slice2)

}