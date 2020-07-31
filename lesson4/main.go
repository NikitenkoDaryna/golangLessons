package main

import "fmt"

func main() {
	a := &american{}
	u := &ukrainian{}
	sayHello(a, "Emily\n")
	sayHello(u, "Darya")
	fmt.Println()
	var cat = &cat{}
	cat.run()
	cat.walk()

	var bird = &eagle{}
	bird.fly()
	walk(bird)
	walk(&dog{})

	m := map[string]interface{}{}
	m["one"] = 1
	m["two"] = 2.0
	m["true"] = true
	for k, v := range m {
		switch v.(type) {
		case int:
			fmt.Println(k)
		case float64:
			fmt.Println(k)
		case bool:
			fmt.Println(k)
		}
	}

	student := &student{person: person{name: "Daryna", surname: "Nikitenko", age: 18}}
	nonStudent := &nonStudent{person: person{
		name:    "Harry",
		surname: "Johns",
		age:     24,
	}}
	student.attendClasses()
	student.doHomework()
	student.work()
	student.hangOutWithSb("Mya")
	nonStudent.cook()
	nonStudent.work()
	nonStudent.hangOutWithSb("Isabella")


}
