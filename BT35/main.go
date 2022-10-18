package main

import (
	"baitap/model"
	"fmt"
)

func insertEmployee(employee model.Employee) {
	uniqueId := map[string]bool{}
	for _, v := range model.Employees {
		uniqueId[v.GetId()] = true
	}
	if !uniqueId[employee.GetId()] {
		model.Employees = append(model.Employees, employee)
	} else {
		fmt.Println("Đã tồn tại mã số, mời nhập lại")
		option1()
	}
}

func main() {
out:
	for {
		var choose int
		printGuidelines()
		fmt.Print("Chọn: ")
		scan(&choose)
		switch choose {
		case 1:
			option1()
			break
		case 2:
			option2()
			break
		case 3:
			option3()
			break
		case 4:
			option4()
			break
		case 5:
			option5()
			break
		case 6:
			option6()
			break
		case 7:
			option7()
			break
		case 8:
			option8()
			break
		case 9:
			option9()
			break
		case 10:
			option10()
			break
		default:
			break out
		}
	}
}

func printGuidelines() {
	fmt.Println("1. Nhập nhân viên")
	fmt.Println("2. Xuất danh sách nhân viên")
	fmt.Println("3. Xóa nhân viên")
	fmt.Println("4. Cập nhật thông tin nhân viên")
	fmt.Println("5. Nhập đơn hàng cho nhân viên tiếp thị")
	fmt.Println("6. Nhập lương trách nhiệm cho trưởng phòng")
	fmt.Println("7. Nhập phần trăm doanh thu với nhân viên tiếp thị (mặc định 10%)")
	fmt.Println("8. Tìm kiếm nhân viên theo lương")
	fmt.Println("9. Sắp xếp nhân viên theo họ tên và thu nhập")
	fmt.Println("10. Xuất 5 nhân viên có thu nhập cao nhất công ty")
	fmt.Println("11. Thoát chương trình")
	fmt.Println("-----------------------")
}
