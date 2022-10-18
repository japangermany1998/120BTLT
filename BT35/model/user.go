package model

type UserInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	MonthSalary int    `json:"month_salary"`
}

func (u *UserInfo) SetInfo(id, name string, monthSalary int) {
	u.Name = name
	u.Id = id
	u.MonthSalary = monthSalary
}

func (u *UserInfo) GetId() string {
	return u.Id
}

func (u *UserInfo) GetName() string {
	return u.Name
}

func (u *UserInfo) GetMonthSalary() int {
	return u.MonthSalary
}
