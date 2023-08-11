package main

import (
	"fmt"
	"errors"
)

func fact(input int){
	if input==0 {
		return 1
	} else {
		return input*fact(input-1) 
	}
	
}

func main(){
	var input=2

	//fmt.Scanln(&input)
	//fmt.Println(input)
	if input<=0 {
		err:=errors.New("Invalid input")
		fmt.Println(err)
	} else { 
		//fact(input)
		fmt.Println(fact(input))
	}
}

// package main

// import "fmt"

// func fact(n int) int {
//     if n == 0 {
//         return 1
//     }
//     return n * fact(n-1)
// }

// func main() {
//     fmt.Println(fact(7))

   
// }