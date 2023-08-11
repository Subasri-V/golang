//insertion deletion retrive
package main

import "fmt"

func main(){
	m:=make(map[string] int)
	m["salary"]=2000
	m["age"]=20
	m["employeeId"]=1
	//m["name"]="Subasri_v"
	if m==nil{
		fmt.Println("m is nil")
	} else {
		fmt.Println(m)
	}
}