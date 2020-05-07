/*
   https://appliedgo.net/concurrencyslower/
   Slow down your code with goroutines
*/

package main

import (
	"fmt"
	"time"
)

const (
	limit = 10000000000
)

/*****************************************************************************
   [1] v1 - serial version
   limit = 10000000000
   avg = 5.982986747 / 5.939953371 / 5.959049145 /  5.948986609 / 6.001934823
******************************************************************************/

//func SerialSum() int {
//	sum := 0
//	for i := 0; i < limit; i++ {
//		sum += i
//	}
//	return sum
//}
//
//func main() {
//	startTime := time.Now()
//	sum := SerialSum()
//	fmt.Println(sum)
//
//	endTime := time.Now()
//	diff := endTime.Sub(startTime)
//	fmt.Println("total time taken ", diff.Seconds(), "seconds")
//}

// end v1 - serial version

/***************************************************************************
   [2] v2 - go routine
   limit = 10000000000
   avg = 8.850305243 / 8.842119288 / 9.316972303 / 8.831026022 / 8.8568185
****************************************************************************/

//func SerialSum(ch chan int){
//	sum := 0
//	for i := 0; i < limit; i++ {
//		sum += i
//	}
//	ch <- sum
//}
//
//func main() {
//	ch := make(chan int)
//	startTime := time.Now()
//	go SerialSum(ch)
//	fmt.Println(<-ch)
//
//	endTime := time.Now()
//	diff := endTime.Sub(startTime)
//	fmt.Println("total time taken ", diff.Seconds(), "seconds")
//}

//  end v2 - go routine

/******************************************************************************
   [3] v3 - go routine cache memory fix
   limit = 10000000000
   avg = 8.850305243 / 8.842119288 / 9.316972303 / 8.831026022 / 8.8568185
********************************************************************************/

//func SerialSum(ch chan int){
//	startTime := time.Now()
//	sum := 0
//	for i := 0; i < limit; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//	endTime := time.Now()
//	diff := endTime.Sub(startTime)
//	fmt.Println("total time taken ", diff.Seconds(), "seconds")
//
//	ch <- sum
//}
//
//func main() {
//	ch := make(chan int)
//	n := runtime.GOMAXPROCS(0)
//
//	//startTime := time.Now()
//
//	//for i := 0; i < limit; i++ {
//		go SerialSum(ch)
//	    <-ch
//	//}
//
//	//sum := 0
//	//for i := 0; i < limit; i++ {
//	//	sum += <-ch
//	//}
//
//	//fmt.Println(sum)
//	fmt.Println(n)
//	//endTime := time.Now()
//	//diff := endTime.Sub(startTime)
//	//fmt.Println("total time taken ", diff.Seconds(), "seconds")
//}

//  end v3 - go routine

/******************************************************************************
   [4] v4 - go routine cache memory fix
   limit = 10000000000
   avg = 8.850305243 / 8.842119288 / 9.316972303 / 8.831026022 / 8.8568185
********************************************************************************/
//
//func main() {
//	startTime := time.Now()
//
//	chanOwner := func() <-chan int {
//		results := make(chan int, 5)
//		go func() {
//			defer close(results)
//			for i := 0; i <= 100000; i++ {
//				results <- i
//			}
//		}()
//		return results
//	}
//
//	consumer := func(results <-chan int) {
//		for result := range results {
//			fmt.Printf("Received: %d\n", result)
//		}
//		fmt.Println("Done receiving!")
//	}
//
//	results := chanOwner()
//	consumer(results)
//	endTime := time.Now()
//	diff := endTime.Sub(startTime)
//	fmt.Println("total time taken ", diff.Seconds(), "seconds")
//}

//end v4 - go routine

/******************************************************************************
   [5] v5 - go routine cache memory fix
   limit = 10000000000
   avg = 8.850305243 / 8.842119288 / 9.316972303 / 8.831026022 / 8.8568185
********************************************************************************/

func main() {
	startTime := time.Now()

	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

//  end v4 - go routine
