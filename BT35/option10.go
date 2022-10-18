package main

import (
	"baitap/model"
	"fmt"
	"sort"
)

func option10() {
	fmt.Println("--------------------")

	var anotherListEmployee = make([]model.Employee, len(model.Employees))
	copy(anotherListEmployee, model.Employees)
	sort.SliceStable(anotherListEmployee, func(i, j int) bool {
		return anotherListEmployee[i].GetSalary() >= anotherListEmployee[j].GetSalary()
	})
	for i, v := range anotherListEmployee {
		if i == 5 {
			break
		}
		printEmployee(v)
	}
	fmt.Print("Nhấn bất kỳ để tiếp tục: ")
	scan()
	fmt.Println("---------------------")
}
