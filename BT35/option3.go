package main

import (
	"baitap/model"
	"fmt"
)

func option3() {
	fmt.Println("---------------")
	fmt.Print("Nhập mã nhân viên cần xóa: ")
	var id string
	scan(&id)
	for i, v := range model.Employees {
		if v.GetId() == id {
			model.Employees = append(model.Employees[:i], model.Employees[i+1:]...)
			break
		}
	}
	fmt.Println("---------------------------")
}
