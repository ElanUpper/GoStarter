package emp

import "fmt"

// 所有type/function 第一字母都需要大写
type Employer struct {
	name 		string;
	underling 	Employee
}

func (emp *Employer) SetEmp(lv_name string) {
	emp.name = lv_name ;
}

func (emp Employer) GetEmp() {
	fmt.Printf("the name %s, the age :%d\n", emp.name, emp.underling.Name)
}

// 添加小弟
func (employer *Employer) AddEmployee(employee Employee){
	employer.underling = employee ;
}
