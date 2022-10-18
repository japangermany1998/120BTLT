package model

type Manager struct {
	UserInfo
	ResponsibilitySalary int
}

func (manager *Manager) SetMonthSalary(month_salary int) {
	manager.MonthSalary = month_salary
}

func (manager *Manager) GetSalary() int {
	return setTaxIncome(manager.ResponsibilitySalary + manager.MonthSalary)
}
