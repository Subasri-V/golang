package main

import "fmt"

func main(){
	fmt.Println("Bubble sort")

	var arr=[5]int{8,7,6,5,4}

	var i int=0
	for i=0;i<5;i++{
		fmt.Print(arr[i] )
	}
	fmt.Println()
	var n int=5
	for ; n>0; {
		for j:=0;j<4;j++{
			if arr[j]>arr[j+1]{
				var temp int=arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=temp
			}
		}
		n--
	}
	for i=0;i<5;i++{
		fmt.Print(arr[i] )
	}
	fmt.Println()
}