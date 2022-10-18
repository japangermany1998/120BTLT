package model

type MarketingStaff struct {
	UserInfo
	Orders []*Order
}

func (marketing *MarketingStaff) SetMonthSalary(month_salary int) {
	marketing.MonthSalary = month_salary
}

func (marketing *MarketingStaff) GetSalary() int {
	var total_finance int
	for _, value := range marketing.Orders {
		total_finance += value.Amount
	}
	return setTaxIncome(percentFinance(total_finance) + marketing.MonthSalary)
}

func (marketing *MarketingStaff) appendOrder(order *Order) {
	marketing.Orders = append(marketing.Orders, order)
}

func percentFinance(totalFinance int) int {
	return PercentFinance * totalFinance / 100
}

type Order struct {
	UserId    string
	ItemsName string
	Amount    int
}
