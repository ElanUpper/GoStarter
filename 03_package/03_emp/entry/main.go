package main

import (
	emp "03_package/03_emp"
	"fmt"
)

// 扩展原有type（组合）
type exEmp struct {
	emp *emp.Employee
}

// 扩展一个方法
func (boss *exEmp) promote_staff(level int) {
	if boss != nil && boss.emp != nil {
		if boss.emp.Name == "Elan" {
			boss.emp.Level += level
		}
	}
}

func (boss *exEmp) print_staff_level() {
	fmt.Printf("staff level %d\n", boss.emp.Level)
}

func main() {
	// 员工
	var employee emp.Employee
	employee = emp.Employee{Name: "staff01"} //, age: 20};  如果属性为小写那么就会出现错误 unknown field 'age'
	employee.SetEmp("Elan", 30)
	employee.GetEmp()
	// 股东
	//Elaner := emp.Employer{: "elan", emp: &employee}
	//Elaner.print_staff_level()
	//Elaner.promote_staff(1)
	//Elaner.print_staff_level()

}
