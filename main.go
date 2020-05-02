package main

import (
	"fmt"
	"math/rand"
)
import "time"

type employee struct {
	name   string
	basic  int
	otRate int
	otHrs  int
}

func calfullSalary(basic int, otAmount int) int {
	return basic + otAmount
}

func calOTAmount(otRate int, otHrs int) int {
	//time.Sleep(1 * time.Second)
	return otRate * otHrs
}

func main() {
	startTime := time.Now()
	fmt.Println("startTime time  ", startTime)
	//emps := []employee{}
	emps := make([]employee, 0)

	for i := 0; i < 1000; i++ {
		emp := employee{
			name:   "wewewe",
			basic:  rand.Intn(100000),
			otRate: rand.Intn(1000),
			otHrs:  rand.Intn(100),
		}
		emps = append(emps, emp)
	}

	for _, em := range emps {
		otAmount := calOTAmount(em.otRate, em.otHrs)
		fullSalary := calfullSalary(em.basic, otAmount)
		fmt.Println(em.name, "| basic:", em.basic, "| otRate:", em.otRate, "| otHrs:", em.otHrs, "| Full Salary:", fullSalary)
		//time.Sleep(1 * time.Second)
	}
	endTime := time.Now()
	fmt.Println("endTime time  ", endTime)
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
