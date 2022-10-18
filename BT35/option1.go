package main

import (
	"baitap/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scan(i ...interface{}) {
	scanner.Scan()
	text := scanner.Text()
	if i == nil {
		return
	}
	if v, ok := i[0].(*string); ok {
		*v = text
	}
	if v, ok := i[0].(*int); ok {
		num, _ := strconv.Atoi(text)
		*v = num
	}
}

func getEmployeeInput(types int) (employee model.Employee) {
	switch types {
	case 1:
		employee = &model.AdministrativeStaff{}
		break
	case 2:
		employee = &model.MarketingStaff{}
		break
	case 3:
		employee = &model.Manager{}
	}
	var name, id string
	var monthSalary int

	fmt.Print("Nhập tên: ")
	scan(&name)
	fmt.Print("Nhập mã: ")
	scan(&id)
	fmt.Print("Nhập lương tháng: ")
	scan(&monthSalary)

	employee.SetInfo(id, name, monthSalary)
	return employee
}

func option1() {
	fmt.Println("----------------------------")
	var types int
	fmt.Println("Chọn loại nhân viên: ")
	fmt.Println("1. Nhân viên hành chính")
	fmt.Println("2. Nhân viên tiếp thị")
	fmt.Println("3. Trưởng phòng")
	fmt.Println("------------")
	fmt.Print("Chọn: ")
	scan(&types)

	insertEmployee(getEmployeeInput(types))
	fmt.Println("-------------------------------------")
}
