package main

import "fmt"

func main(){
	s:= []int{10,20,30,40,50,60,70,80,90,100}
	fmt.Println("Original Array")
	fmt.Printf("s=%v , len=%d , cap=%d\n",s,len(s),cap(s))

	s=s[1:5]
	fmt.Printf("s=%v , len=%d , cap=%d\n",s,len(s),cap(s))

	s=s[ :8]
	fmt.Printf("s=%v , len=%d , cap=%d\n",s,len(s),cap(s))

	s=s[2:]
	fmt.Printf("s=%v , len=%d , cap=%d\n",s,len(s),cap(s))
}