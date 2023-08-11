package main

import (
	"fmt"
	"math"
)

// func getStockPriceChange(prevPrice,currPrice float64 )( float64,float64,error){

// 	if prevPrice==0{
// 		err:=errors.New("Previous")
// 	}
// 	change:=currPrice-prevPrice
// 	percentChange:=(change/prevPrice)*100
// 	return __ ,percentChange
// }



// func getStockPriceChange(prevPrice,currPrice float64 )( float64,float64){
// 	change:=currPrice-prevPrice
// 	percentChange:=(change/prevPrice)*100
// 	return change,percentChange
// }

//func main(){
	// var i int=10
	// name:="Subasri"
	// fmt.Println("Welcome to golang",i,name)
	// fmt.Printf("Welcome to golang %t %t",i,name)
	// var names="abc"
	// fmt.Printf("Variable 'name is of type %t",names)

	// var(
	// 	firstname,lastname  string
	// 	age 				int
	// 	salary 				float64
	// 	isComfirmed 		bool
	// )
	// firstname="suba"
	// lastname="sri"
	// age=21
	// salary=15000
	// isConfirmed=true

	// fmt.Println("firstname: %s,lastname:%s,age: %d,salary: %f,isConfirmed: %t",firstname,lastname,age,salary,isConfirmed)

	// var i=10
	// if i<10{
	// 	Printf("True")
	// }
	// else {
	// 	Printf("False")
	// }
	
	//var dayOfWeek=6
	// switch dayOfWeek:=6 ;dayOfWeek{
	// case 1: fmt.Println("Monday")
	// case 2: fmt.Println("Tuesday")
	// case 6: { 
	// 	fmt.Println("Saturday")
	// 	fmt.Println("Weekend. Yaay!")
	// }
	// default: fmt.Println("Invalid day")
	// }

	// for i:=1;i<=10;i++ {
	// 	fmt.Printf("%d ",i)
	// }
	// var j=1
	// for ;j<=10;j++ {
	// 	fmt.Printf("%d ",j)
	// }

	// for{
	// 	fmt.Printf("j")
	// }

	//break statement exist in golang
	//continue should end with semicolon (continue;)

	// prevStockPrice:=0.0
	// currStockPrice:=100000.0

	// change,percentChange:=getStockPriceChange(prevStockPrice,currStockPrice)

	// if change<0{
	// 	fmt.Printf("the stock price decreased by $%.2f which is %.2f%% of the prev price\n",math.Abs(change),math.Abs(percentChange))
	// } else {
	// 	fmt.Printf("the stock price increased by $%.2f which is %.2f%% of the prev price\n",change,percentChange)
	// }
//}
func blankstatements(prevPrice,currPrice float64)(change, percentChange float64){
	change =currPrice-prevPrice
	percentChange=(change/prevPrice)*100
	return change, percentChange
}
func main(){
		prevStockPrice:=100000.0
		CurrStockPrice:=90000.0
		
		change,_:=blankstatements(prevStockPrice,CurrStockPrice)
	
		if change<0{
			fmt.Printf("the stock price decreased by $%.2f which is %.2f%% of the prev price\n",math.Abs(change))
			} else {
				fmt.Printf("the stock price increased by $%.2f which is %.2f%% of the prev price\n",change)
			}
				
	}





