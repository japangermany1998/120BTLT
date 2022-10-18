package model

type Employee interface {
	SetMonthSalary(monthSalary int)
	GetSalary() int
	SetInfo(id, name string, monthSalary int)
	GetName() string
	GetId() string
	GetMonthSalary() int
}

func setTaxIncome(salary int) int {
	if salary < 9000000 {
		return salary
	}
	if salary > 15000000 {
		return salary - (salary * 12 / 100)
	}
	return salary - (salary * 10 / 100)
}
