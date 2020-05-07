package main

import "fmt"

const (
	limit = 100
)

func SerialSum() int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += i
	}
	return sum
}

func main() {
	sum := SerialSum()
	fmt.Println(sum)
}
