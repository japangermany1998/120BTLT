package main

import (
	"baitap/model"
	"fmt"
)

func option6() {
	fmt.Println("---------------")
	fmt.Print("Nhập mã trưởng phòng: ")
	var id string
	scan(&id)
	for _, v := range model.Employees {
		if manager, ok := v.(*model.Manager); ok && v.GetId() == id {
			fmt.Print("Nhập lương trách nhiệm trưởng phòng: ")
			scan(&manager.ResponsibilitySalary)
			break
		}
	}
	fmt.Println("-----------------------")
}
