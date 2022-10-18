package main

import (
	"baitap/model"
	"fmt"
)

func option8() {
	fmt.Println("---------------")
	var minSalary, maxSalary int
	fmt.Print("Nhập mức lương tối thiểu: ")
	scan(&minSalary)
	fmt.Print("Nhập mức lương tối đa: ")
	scan(&maxSalary)
	if minSalary >= maxSalary {
		fmt.Print("Lương tối thiểu ko thể >= lương tối đa. Nhấn bất kỳ để tiếp tục: ")
		scan()
	}

	for _, v := range model.Employees {
		if v.GetSalary() >= minSalary && v.GetSalary() <= maxSalary {
			printEmployee(v)
		}
	}
	fmt.Print("Nhấn enter để tiếp tục")
	scan()
}
