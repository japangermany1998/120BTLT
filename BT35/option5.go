package main

import (
	"baitap/model"
	"fmt"
)

func getOrderInput(userId string) (order *model.Order) {
	order = new(model.Order)
	order.UserId = userId
	fmt.Println("--------")
	fmt.Println("Nhập thông tin đơn hàng")
	fmt.Print("Nhập tên item: ")
	scan(&order.ItemsName)
	fmt.Print("Nhập doanh thu: ")
	scan(&order.Amount)
	return
}

func option5() {
	fmt.Println("---------------")
	fmt.Print("Nhập mã nhân viên tiếp thị: ")
	var id string
	scan(&id)
	for _, v := range model.Employees {
		if marketing, ok := v.(*model.MarketingStaff); ok && v.GetId() == id {
			if marketing.Orders == nil {
				marketing.Orders = []*model.Order{}
			}
			for {
				marketing.Orders = append(marketing.Orders, getOrderInput(id))
				fmt.Print("Bạn có muốn nhập đơn hàng nữa? (y/n): ")
				var check string
				scan(&check)
				if check != "y" {
					break
				}
			}
			break
		}
	}
	fmt.Println("-----------------------")
}
