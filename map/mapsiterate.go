package main

import "fmt"

func main(){
	var personAge = map[string]int{
		"Suba":21,
		"gayatri":19,
		"mohana":53,
	}
	personAge["abc"]=60
	personAge["xyz"]=35

	x,isExisting:=personAge["abc"]

	fmt.Println(x)
	fmt.Println(isExisting)

	delete(personAge,"abc")
	fmt.Println(personAge)

	// for name,age:=range personAge{
	// 	fmt.Println(name,age)
	// }

	fmt.Println(x)
}