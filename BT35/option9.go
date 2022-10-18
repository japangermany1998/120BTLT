package main

import (
	"baitap/model"
	"fmt"
	"sort"
)

func option9() {
	fmt.Println("--------------------")
	var option int
	fmt.Println("1. Tăng dần theo tên")
	fmt.Println("2. Giảm dần theo tên")
	fmt.Println("3. Tăng dần theo mức lương")
	fmt.Println("4. Giảm dần theo mức lương")
	fmt.Print("Chọn thứ tự sắp xếp: ")
	scan(&option)
	var anotherListEmployee = make([]model.Employee, len(model.Employees))
	copy(anotherListEmployee, model.Employees)
	sort.SliceStable(anotherListEmployee, func(i, j int) bool {
		switch option {
		case 1:
			return anotherListEmployee[i].GetName() <= anotherListEmployee[j].GetName()
		case 2:
			return anotherListEmployee[i].GetName() >= anotherListEmployee[j].GetName()
		case 3:
			return anotherListEmployee[i].GetSalary() <= anotherListEmployee[j].GetSalary()
		}
		return anotherListEmployee[i].GetSalary() >= anotherListEmployee[j].GetSalary()
	})
	for _, v := range anotherListEmployee {
		printEmployee(v)
	}
	fmt.Print("Nhấn bất kỳ để tiếp tục: ")
	scan()
	fmt.Println("---------------------")
}
