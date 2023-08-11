package main

import "fmt"

func main(){
	var a=[6]string{"Apple","Banana","carrot","dog","Egg","abcdef"}
	var s []string = a[1:4]

	fmt.Println("Array a = ",a)
	fmt.Println("Slice s = ", s)

	fmt.Println(cap(s)) //capacity -upto max size it can expand
	fmt.Println(len(s)) //no of elements in the slice
}