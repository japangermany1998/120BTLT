package main

import (
	"baitap/model"
	"fmt"
)

func option7() {
	fmt.Println("-------------")
	fmt.Print("Nhập phần trăm doanh thu: ")
	scan(&model.PercentFinance)
	fmt.Println("----------------")
}
