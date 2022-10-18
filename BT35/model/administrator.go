package model

type AdministrativeStaff struct {
	UserInfo
}

func (administrator *AdministrativeStaff) SetMonthSalary(month_salary int) {
	administrator.MonthSalary = month_salary
}

func (administrator *AdministrativeStaff) GetSalary() int {
	return setTaxIncome(administrator.MonthSalary)
}
