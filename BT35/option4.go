package main

import (
	"baitap/model"
	"fmt"
)

func updateEmployee(e model.Employee, id string) {
	var name string
	var monthSalary int

	fmt.Print("Nhập tên: ")
	scan(&name)
	fmt.Print("Nhập lương tháng: ")
	scan(&monthSalary)
	e.SetInfo(id, name, monthSalary)
}

func option4() {
	fmt.Println("---------------")
	fmt.Print("Nhập mã nhân viên cần sửa: ")
	var id string
	scan(&id)
	for _, v := range model.Employees {
		if v.GetId() == id {
			updateEmployee(v, id)
			break
		}
	}
}
