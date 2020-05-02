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

	emps := []employee{}

	emps = append(emps, employee{
		name:   "Saman",
		basic:  100000,
		otRate: 500,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "Gasalu",
		basic:  50000,
		otRate: 300,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Raju",
		basic:  45000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "ko",
		basic:  500000,
		otRate: 700,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "roja",
		basic:  60000,
		otRate: 600,
		otHrs:  30,
	})
	emps = append(emps, employee{
		name:   "konara",
		basic:  25000,
		otRate: 700,
		otHrs:  40,
	})

	emps = append(emps, employee{
		name:   "Saman",
		basic:  100000,
		otRate: 500,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "Gasalu",
		basic:  50000,
		otRate: 300,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Raju",
		basic:  45000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "ko",
		basic:  500000,
		otRate: 700,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "roja",
		basic:  60000,
		otRate: 600,
		otHrs:  30,
	})
	emps = append(emps, employee{
		name:   "konara",
		basic:  25000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Saman",
		basic:  100000,
		otRate: 500,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "Gasalu",
		basic:  50000,
		otRate: 300,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Raju",
		basic:  45000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "ko",
		basic:  500000,
		otRate: 700,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "roja",
		basic:  60000,
		otRate: 600,
		otHrs:  30,
	})
	emps = append(emps, employee{
		name:   "konara",
		basic:  25000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Saman",
		basic:  100000,
		otRate: 500,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "Gasalu",
		basic:  50000,
		otRate: 300,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Raju",
		basic:  45000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "ko",
		basic:  500000,
		otRate: 700,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "roja",
		basic:  60000,
		otRate: 600,
		otHrs:  30,
	})
	emps = append(emps, employee{
		name:   "konara",
		basic:  25000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Saman",
		basic:  100000,
		otRate: 500,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "Gasalu",
		basic:  50000,
		otRate: 300,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Raju",
		basic:  45000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "ko",
		basic:  500000,
		otRate: 700,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "roja",
		basic:  60000,
		otRate: 600,
		otHrs:  30,
	})
	emps = append(emps, employee{
		name:   "konara",
		basic:  25000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Saman",
		basic:  100000,
		otRate: 500,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "Gasalu",
		basic:  50000,
		otRate: 300,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Raju",
		basic:  45000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "ko",
		basic:  500000,
		otRate: 700,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "roja",
		basic:  60000,
		otRate: 600,
		otHrs:  30,
	})
	emps = append(emps, employee{
		name:   "konara",
		basic:  25000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Saman",
		basic:  100000,
		otRate: 500,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "Gasalu",
		basic:  50000,
		otRate: 300,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "Raju",
		basic:  45000,
		otRate: 700,
		otHrs:  40,
	})
	emps = append(emps, employee{
		name:   "ko",
		basic:  500000,
		otRate: 700,
		otHrs:  20,
	})
	emps = append(emps, employee{
		name:   "roja",
		basic:  60000,
		otRate: 600,
		otHrs:  30,
	})
	emps = append(emps, employee{
		name:   "konara",
		basic:  25000,
		otRate: 700,
		otHrs:  40,
	})

	for _, em := range emps {
		otAmount := calOTAmount(em.otRate, em.otHrs)
		fullSalary := calfullSalary(em.basic, otAmount)
		fmt.Println(em.name, "Full Salary: ", fullSalary)

		//time.Sleep(1 * time.Second)

	}
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Nanoseconds(), "Nanoseconds")
	fmt.Println("======================================")
	fmt.Println("Average time taken '130000 - 200000' Nanoseconds")

}
