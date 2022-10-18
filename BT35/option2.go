package main

import (
	"baitap/model"
	"fmt"
	"strconv"
)

func printEmployee(e model.Employee) {
	fmt.Println("Nhân viên mã số " + e.GetId() + " {")
	fmt.Println("\t Họ tên: " + e.GetName())
	fmt.Print("\t Loại nhân viên: ")
	switch e.(type) {
	case *model.Manager:
		fmt.Println("Trưởng phòng")
		fmt.Println("\t Mức lương trách nhiệm hiện tại: " + strconv.Itoa(e.(*model.Manager).ResponsibilitySalary))
		break
	case *model.AdministrativeStaff:
		fmt.Println("Nhân viên hành chính")
		break
	case *model.MarketingStaff:
		fmt.Println("Nhân viên tiếp thị")
		fmt.Println("\t Các đơn hàng của nhân viên: ")
		for i, v := range e.(*model.MarketingStaff).Orders {
			fmt.Println("\t \t Đơn hàng " + strconv.Itoa(i+1) + " {")
			fmt.Println("\t \t \tTên mặt hàng: " + v.ItemsName)
			fmt.Println("\t \t \tDoanh thu: " + strconv.Itoa(v.Amount))
			fmt.Println("\t \t}")
		}
	}
	fmt.Println("\t Mức lương cố định hàng tháng: " + strconv.Itoa(e.GetMonthSalary()) + " đồng")
	fmt.Println("\t Mức lương hiện tại: " + strconv.Itoa(e.GetSalary()) + " đồng")
	fmt.Println("}")
}

func option2() {
	fmt.Println("--------------------------------\nDanh sách nhân viên:")
	for _, val := range model.Employees {
		printEmployee(val)
	}
	fmt.Println("---------------")
	fmt.Print("Nhấn enter để tiếp tục ")
	scan()
	fmt.Println("---------------------------")
}
