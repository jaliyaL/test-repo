package main

import "fmt"
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
	emps := make([]employee, 10000)

	for _, em := range emps {
		otAmount := calOTAmount(em.otRate, em.otHrs)
		fullSalary := calfullSalary(em.basic, otAmount)
		fmt.Println(em.name, "Full Salary: ", fullSalary)
		//time.Sleep(1 * time.Second)
	}
	endTime := time.Now()
	fmt.Println("endTime time  ", endTime)
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
