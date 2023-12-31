package main

import (
	"fmt"
	"sync"
)

func routine(i int, wg*sync.WaitGroup){
	fmt.Println("Started Routine",i)
	fmt.Printf("Routine %d ended \n",i)
	wg.Done()
}

func main(){
	noOfRoutines:=3
	var wg sync.WaitGroup
	for i:=0;i<noOfRoutines;i++{
		wg.Add(1)
		go routine(i,&wg)
	}
	wg.Wait();
	fmt.Println("All routines are finished")
}