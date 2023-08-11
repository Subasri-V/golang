package main

import "fmt"

// func main(){
// 	fmt.Println("I am first")
// 	defer fmt.Println("Predict me2")
// 	defer fmt.Println("Predict me1")
// 	defer fmt.Println("Predict me3") //last in first out 
// 	fmt.Println("I am end")

// 	//why defer statements are designed in the stacked(LIFO) order
// }

func one(){
	defer two()
	defer three()
	fmt.Println("four")

	
}
func two(){
	defer three()
	fmt.Println("Two")
}
func three(){
	fmt.Println("three")
}

func main(){
	defer one()
	fmt.Println("One")
}


//defer panic recover - graceful session handling

//recover statement - the system will start execution from the recover statement