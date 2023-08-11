package main

import "fmt"

// type Point struct{
// 	X,Y float64
// }

// func(p Point) IsAbove(y float64) bool{
// 	return p.X>y
// }
// func main(){

// 	p:=Point{2.0,4.0}
// 	fmt.Println("Point:",p)
// 	fmt.Println(p.IsAbove(5))

// }

type Point struct{
	X,Y int
}

func(p *Point) Translate(dx,dy int) {
	p.X=p.X-dx
	p.Y=p.Y-dy
}
func main(){

	p:=Point{2,4}
	fmt.Println("Point:",p)

	p.Translate(7,9)
	fmt.Println(p)

}