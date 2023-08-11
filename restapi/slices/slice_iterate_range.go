package main

import "fmt"

func main(){
	// primeNumbers := []int{2,3,5,7,11,13,17,19,23,29}

	// for index,number:=range primeNumbers{
	// 	fmt.Printf("PrimeNumber(%d)=%d\n",index+1,number)
	// }


	numbers:=[]float64{3.5,7.4,9.2,5.4}

	sum:=0.0
	for _,x:=range numbers{
		sum+=x
	}
	fmt.Printf("%.2f\n",sum)
}