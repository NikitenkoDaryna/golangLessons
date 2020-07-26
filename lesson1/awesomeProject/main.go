package main

import (
	"encoding/json"
	"fmt"
)

var p string

func main() {
	fmt.Println("Hello my name is " + "Daryna")
	//using var in creating structures
	var worker0 Worker
	worker0.salary = "1200"
	hobbies := [3]string{"coding", "drawing", "reading"}
	//creating a struct using a struct literal
	worker1 := Worker{20, "Hyram", "3000", hobbies, 24.30, 8}
	//using new keyword
	worker2 := new(Worker)
	worker2.age = 23
	worker2.name = "Harry"
	worker2.salary = "5000"
	worker2.hobbies = [3]string{"eSport", "cooking", "coding"}
	fmt.Println("worker1", worker1)
	fmt.Println("worker2", *worker2)
	fmt.Println("money1-", worker1.getMoneyPerDay())
	//using pointer address operator
	var programmer1 = &Programmer{"Kenny", 24, 35.6} //can't skip any value
	fmt.Println(*programmer1)

	var programmer2 = &Programmer{}
	programmer2.hourlyRate = 16.4
	programmer2.name = "Jason"
	fmt.Println(*programmer2)

	var programmer3 = &Programmer{}
	(*programmer3).name = "Diana"
	(*programmer3).hourlyRate = 23.5

	fmt.Println(programmer3)
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{Basic: 15000.00,
				HRA: 5000.00,
				TA:  2000.00},
			{Basic: 16000.00,
				HRA: 5000.00,
				TA:  2100.00},
			{Basic: 17000.00,
				HRA: 5000.00,
				TA:  2200.00},
		},
	}
	//fmt.Println(e.FirstName, e.LastName)
	//fmt.Println(e.Age)
	//fmt.Println(e.Email)
	//fmt.Println(e.MonthlySalary[1])
	fmt.Println("using PrintInfo()")
	fmt.Println(e.EmpInfo())
	json_string := `
{
"firstname":"Ed",
"lastname":"Stellary",
"city":"London"
}
`
	emp := new(JSEmployee)
	json.Unmarshal([]byte(json_string), emp)
	fmt.Println(emp)

	emp2 := new(JSEmployee)
	emp2.FirstName = "Roland"
	emp2.LastName = "Rollerda"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)

	myEmplyee := new(Employee)
	myEmplyee.Info()
	fmt.Println(*myEmplyee)
	var calc = Calculator(2)
	fmt.Println("res of cacl:", calc(7))
	div := Division(1243.3213)
	fmt.Println(div(234))
	//User Defined fuctions


}

func (w *Worker) getMoneyPerDay() float64 {
	return w.hourlyRate * w.workingShifts
}

func inputNum() {
	fmt.Print("Input variable:")
	var input float64
	fmt.Scan(&input)
	output := input * 2
	fmt.Println(output)
}

