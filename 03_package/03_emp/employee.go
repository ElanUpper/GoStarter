package emp

import "fmt"

// type第一个字母必须大写Emp
type Employee struct {
	Name  string ;
	Level int ;
	age   int
}

func (emp *Employee) SetEmp(lv_name string, lv_age int) {
	emp.Name = lv_name ;
	emp.age  = lv_age ;
}

func (emp Employee) GetEmp() {
	fmt.Printf("the name %s, the age :%d\n", emp.Name, emp.age)
}



