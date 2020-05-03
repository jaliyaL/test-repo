package main

import (
	"fmt"
	"time"
)

func test(chtest chan int) {

	for i := 0; i < 50; i++ {
		fmt.Println("test", i)
	}
	fmt.Println("test finished")
	chtest <- 10

}

func test2(chtest2 chan int) {

	for i := 0; i < 10; i++ {
		fmt.Println("test2", i)
	}
	fmt.Println("test2 finished")
	chtest2 <- 10

}

//func test(){
//
//	for i:=0; i<50; i++ {
//		fmt.Println("Rat", i)
//	}
//
//}
//
//func test2(){
//
//	for i:=0; i<10; i++ {
//		fmt.Println("test2", i)
//	}
//}

func main() {
	startTime := time.Now()

	//with gorutine
	chtest := make(chan int)
	chtest2 := make(chan int)
	go test(chtest)
	go test2(chtest2)
	<-chtest
	<-chtest2

	//without routine
	//test()
	//test2()

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Nanoseconds(), "Nanoseconds")
}
