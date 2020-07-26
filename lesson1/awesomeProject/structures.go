package main

import "fmt"

type Worker struct {
	age           int64
	name          string
	salary        string
	hobbies       [3]string
	hourlyRate    float64
	workingShifts float64
}
type Programmer struct{
	name string
	age int
	hourlyRate float64


}
type rectangle struct{
	color string
	length float64
	geometry struct{
		area float64
		perimeter float64
	}
}
//Nested structures
type Salary struct{
	Basic,HRA,TA float64
}
type Employee struct{
	FirstName,LastName,Email string
	Age int
	MonthlySalary []Salary
}
type JSEmployee struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	City string `json:"city"`
}
func (e *Employee) EmpInfo() string{
	fmt.Println(e.FirstName,e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info:=range e.MonthlySalary{
		fmt.Println("===============")
		fmt.Println(info.Basic)
		fmt.Println(info.TA)
		fmt.Println(info.HRA)
	}
	return "-----------------------"
}
//a kind of constructor with default values
func (obj *Employee) Info() {
	if obj.FirstName==""||obj.LastName==""{
		obj.FirstName="Jack"
		obj.LastName="Black"
	}
	if obj.Age==0{
		obj.Age=25
	}
}

